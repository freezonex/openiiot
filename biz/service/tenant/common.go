package tenant

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/dal/mysql"
	"freezonex/openiiot/biz/service/k8s"
	"freezonex/openiiot/biz/service/utils/common"
)

type TenantService struct {
	db  *mysql.MySQL
	c   *config.GeneralConfig
	k8s *k8s.K8sService
}

var (
	service *TenantService
	once    sync.Once
)

func NewTenantService(db *mysql.MySQL, c *config.GeneralConfig, k8s *k8s.K8sService) *TenantService {
	once.Do(func() {
		service = &TenantService{
			db:  db,
			c:   c,
			k8s: k8s,
		}
	})
	return service
}

func DefaultTenantService() *TenantService {
	return service
}

// check tenant input
func (a *TenantService) CheckTenant(ctx context.Context, id string, name string) (tenant_id int64, tenant_name string, err error) {

	var tenants []*model_openiiot.Tenant

	// Check if id is provided
	if id != "" {
		tenants, err = a.GetTenantDB(ctx, common.StringToInt64(id), "")
		if err != nil {
			return 0, "", fmt.Errorf("failed to retrieve tenant from db using tenant id %s: %w", id, err)
		}
	}

	// Check if name is provided and id is not
	if name != "" && id == "" {
		tenants, err = a.GetTenantDB(ctx, 0, name)
		if err != nil {
			return 0, "", fmt.Errorf("failed to retrieve tenant from db using tenant name %s: %w", name, err)
		}
	}

	// Check if no tenants were found
	if len(tenants) == 0 {
		return 0, "", errors.New("cannot find tenant in db with the provided criteria")
	}

	return tenants[0].ID, tenants[0].Name, nil
}

// SplitComponentName splits a component name into its English word and numeric parts.
// It returns an error if the format is not "english+number".
func (a *TenantService) SplitComponentNameStrict(componentName string) (string, string, error) {
	// Define a regex pattern to match the component name (English word) and number
	pattern := regexp.MustCompile(`^([a-zA-Z]+)(\d+)$`)

	// Find the matches
	matches := pattern.FindStringSubmatch(componentName)

	// If matches are found, return the two parts
	if len(matches) == 3 {
		return matches[1], matches[2], nil
	}

	// Return an error if the format is incorrect
	return "", "", errors.New("invalid format: expected component name format is 'english+number'")
}

// "nodered1" => "nodered","1",  "nodered" => "nodered", "", "" => "", ""
func (a *TenantService) SplitComponentNameLoose(componentName string) (string, string) {
	// Define the regex pattern to match the name part and the number part
	re := regexp.MustCompile(`^([a-zA-Z]*)(\d*)$`)

	// Find the matches
	matches := re.FindStringSubmatch(componentName)

	// The first submatch is the entire string, the second is the name, and the third is the number
	if len(matches) > 2 {
		return matches[1], matches[2]
	}

	// If no match, return the original string and an empty number part
	return componentName, ""
}

// FindComponentName checks if the componentName is found in the list of deploymentNames after trimming the prefix.
func (a *TenantService) FindComponentName(deploymentNames []string, componentName string) bool {
	// Define the prefix to trim
	prefix := "openiiot-"

	// Iterate over the list of deployment names
	for _, deployment := range deploymentNames {
		// Trim the prefix
		trimmedName := strings.TrimPrefix(deployment, prefix)

		// Check if the trimmed name matches the componentName
		if trimmedName == componentName {
			return true // Found the componentName
		}
	}

	return false // Not found
}
