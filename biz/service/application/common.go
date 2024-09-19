package application

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
	"freezonex/openiiot/biz/service/k8s"
	"freezonex/openiiot/biz/service/tenant"
)

type ApplicationService struct {
	db     *mysql.MySQL
	c      *config.GeneralConfig
	k8s    *k8s.K8sService
	tenant *tenant.TenantService
}

var (
	service *ApplicationService
	once    sync.Once
)

func NewApplicationService(db *mysql.MySQL, c *config.GeneralConfig, k8s *k8s.K8sService, tenant *tenant.TenantService) *ApplicationService {
	once.Do(func() {
		service = &ApplicationService{
			db:     db,
			c:      c,
			k8s:    k8s,
			tenant: tenant,
		}
	})
	return service
}

func DefaultApplicationService() *ApplicationService {
	return service
}

// BuildDockerImage builds a Docker image using the provided appName, componentType, and files.
func (a *ApplicationService) BuildDockerImage(appName string, componentType string, fileHeaders []*multipart.FileHeader) (string, error) {
	// Create a temporary directory as the Docker build context
	tmpDir, err := os.MkdirTemp("", "docker-build-context")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmpDir) // Clean up the temporary directory after use

	// Create the Dockerfile content
	dockerfileContent := `FROM nginx:alpine
COPY . /usr/share/nginx/html/
EXPOSE 80
`

	// Write the Dockerfile to the temporary directory
	err = os.WriteFile(filepath.Join(tmpDir, "Dockerfile"), []byte(dockerfileContent), 0644)
	if err != nil {
		return "", err
	}

	// Iterate over the files and write each file to the temporary directory
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			return "", fmt.Errorf("failed to open file %s: %v", fileHeader.Filename, err)
		}
		defer file.Close()

		// Create a file in the temporary directory with the original name
		outPath := filepath.Join(tmpDir, fileHeader.Filename)
		outFile, err := os.Create(outPath)
		if err != nil {
			return "", fmt.Errorf("failed to create file %s: %v", outPath, err)
		}

		// Copy the file content to the newly created file
		if _, err := io.Copy(outFile, file); err != nil {
			outFile.Close()
			return "", fmt.Errorf("failed to copy content to file %s: %v", outPath, err)
		}
		outFile.Close()
	}

	// Build the Docker image using the Docker CLI
	imageName := fmt.Sprintf("openiiot-app-%s-%s:1.0.0", appName, componentType)
	cmd := exec.Command("docker", "build", "-t", imageName, tmpDir)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Run the Docker build command and capture output
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to build Docker image: %v, output: %s", err, out.String())
	}

	fmt.Printf("Build output: %s\n", out.String())
	return imageName, nil
}

// TagAndPushDockerImage tags the built image and pushes it to the registry
func (a *ApplicationService) TagAndPushDockerImage(imageName string, registry string) error {
	// Tag the Docker image with the registry
	taggedImageName := fmt.Sprintf("%s/%s", registry, imageName)
	cmdTag := exec.Command("docker", "tag", imageName, taggedImageName)
	var outTag bytes.Buffer
	cmdTag.Stdout = &outTag
	cmdTag.Stderr = &outTag

	// Run the Docker tag command
	if err := cmdTag.Run(); err != nil {
		return fmt.Errorf("failed to tag Docker image: %v, output: %s", err, outTag.String())
	}

	fmt.Printf("Tag output: %s\n", outTag.String())

	// Push the tagged Docker image to the registry
	cmdPush := exec.Command("docker", "push", taggedImageName)
	var outPush bytes.Buffer
	cmdPush.Stdout = &outPush
	cmdPush.Stderr = &outPush

	// Run the Docker push command
	if err := cmdPush.Run(); err != nil {
		return fmt.Errorf("failed to push Docker image: %v, output: %s", err, outPush.String())
	}

	fmt.Printf("Push output: %s\n", outPush.String())
	return nil
}

// checkApplicationName checks if the application_name is a valid Kubernetes deployment name
func (a *ApplicationService) CheckApplicationName(applicationName string) error {
	// Rule: Maximum length of 240 characters
	if len(applicationName) > 240 {
		return fmt.Errorf("application name exceeds the maximum length of 240 characters")
	}

	// List of internal component names that the application name cannot start with
	internalComponents := []string{"nodered", "grafana", "emqx"}

	// Check if applicationName starts with any of the internal components
	for _, component := range internalComponents {
		if strings.HasPrefix(applicationName, component) {
			return fmt.Errorf("application name cannot start with internal component name: %s", component)
		}
	}

	// Rule: Must consist of lowercase letters, numbers, and hyphens
	// Rule: Cannot contain uppercase letters, spaces, underscores, or special characters
	// Rule: Cannot start with a number or end with a hyphen
	// Regex breakdown:
	// ^[a-z]       : must start with a lowercase letter
	// [a-z0-9-]*   : can contain lowercase letters, numbers, and hyphens
	// [a-z0-9]$    : must end with a lowercase letter or number
	validNamePattern := `^[a-z]([a-z0-9-]*[a-z0-9])?$`
	matched, err := regexp.MatchString(validNamePattern, applicationName)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid application name: it must consist of lowercase letters, numbers, and hyphens, and cannot start with a number or end with a hyphen")
	}

	return nil
}
