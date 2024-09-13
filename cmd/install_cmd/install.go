package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// Function to execute command and handle error
func executeCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Confirmation prompt for the user
func confirmAction(message string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/n): ", message)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	return strings.ToLower(response) == "y"
}

// Function to check system requirements
func check() error {
	// Check for at least 4 CPU cores
	cpuCores := runtime.NumCPU()
	if cpuCores < 4 {
		return fmt.Errorf("Error: The system must have at least 4 CPU cores. Current cores: %d", cpuCores)
	}
	fmt.Printf("CPU Cores: %d - OK\n", cpuCores)

	// Check for at least 8GB of memory
	var sysInfo syscall.Sysinfo_t
	err := syscall.Sysinfo(&sysInfo)
	if err != nil {
		return fmt.Errorf("Error getting system memory information: %v", err)
	}

	memoryGB := sysInfo.Totalram / (1024 * 1024 * 1024) // Convert from bytes to GB
	if memoryGB < 8 {
		return fmt.Errorf("Error: The system must have at least 8GB of memory. Current memory: %dGB", memoryGB)
	}
	fmt.Printf("Memory: %dGB - OK\n", memoryGB)

	// Check for at least 20GB of free disk space
	var stat syscall.Statfs_t
	err = syscall.Statfs("/", &stat)
	if err != nil {
		return fmt.Errorf("Error getting disk space information: %v", err)
	}

	diskSpaceGB := (stat.Bavail * uint64(stat.Bsize)) / (1024 * 1024 * 1024) // Convert from bytes to GB
	if diskSpaceGB < 20 {
		return fmt.Errorf("Error: The system must have at least 20GB of free disk space. Current space: %dGB", diskSpaceGB)
	}
	fmt.Printf("Disk Space: %dGB - OK\n", diskSpaceGB)

	fmt.Println("All preconditions are met.")
	return nil
}

// Handle installation process with subcommands
func install() {
	if len(os.Args) > 2 {
		subcommand := os.Args[2]
		switch subcommand {
		case "openiiot":
			if confirmAction("Are you sure you want to install OpenIIOT?") {
				err := installOpeniiot()
				if err != nil {
					fmt.Println("Error installing OpenIIOT:", err)
				} else {
					fmt.Println("OpenIIOT installed successfully.")
				}
			} else {
				fmt.Println("Installation of OpenIIOT aborted.")
			}
		case "supabase":
			if confirmAction("Are you sure you want to install Supabase?") {
				err := installSupabase()
				if err != nil {
					fmt.Println("Error installing Supabase:", err)
				} else {
					fmt.Println("Supabase installed successfully.")
				}
			} else {
				fmt.Println("Installation of Supabase aborted.")
			}
		default:
			fmt.Println("Unknown install subcommand:", subcommand)
			os.Exit(1)
		}
	} else {
		if confirmAction("Are you sure you want to proceed with the installation of FreeFlow?") {
			err := installK8s()
			if err != nil {
				fmt.Println("Error installing Kubernetes:", err)
				return
			}
			err = installOpeniiot()
			if err != nil {
				fmt.Println("Error installing OpenIIOT:", err)
				return
			}
			err = installDocker()
			if err != nil {
				fmt.Println("Error installing Docker:", err)
				return
			}
			err = installSupabase()
			if err != nil {
				fmt.Println("Error installing Supabase:", err)
				return
			}
			err = installNginx()
			if err != nil {
				fmt.Println("Error installing Nginx:", err)
				return
			}
			fmt.Println("FreeFlow installed successfully.")
		} else {
			fmt.Println("Installation aborted.")
		}
	}
}

func uninstallNginx() error {
	fmt.Println("Uninstalling Nginx...")
	err := executeCommand("sudo", "docker-compose", "-f", "nginx/docker-compose.yml", "down", "--volumes", "--rmi", "all", "--remove-orphans")
	if err != nil {
		return fmt.Errorf("failed to uninstall Nginx: %v", err)
	}
	fmt.Println("Nginx uninstalled successfully.")
	return nil
}

func uninstallSupabase() error {
	fmt.Println("Uninstalling Supabase...")
	err := executeCommand("sudo", "docker-compose", "-f", "supabase/docker-compose.yml", "down", "--volumes", "--rmi", "all", "--remove-orphans")
	if err != nil {
		return fmt.Errorf("failed to uninstall Supabase: %v", err)
	}
	fmt.Println("Supabase uninstalled successfully.")
	return nil
}

func uninstallDocker() error {
	fmt.Println("Uninstalling Docker...")

	// Remove containerd config for Docker registry
	fmt.Println("Removing Docker registry from containerd config...")
	err := executeCommand("sudo", "rm", "-rf", "/etc/containerd/certs.d/registry:5000")
	if err != nil {
		fmt.Printf("Error removing Docker registry config: %v\n", err)
	}

	err = executeCommand("sudo", "systemctl", "restart", "containerd")
	if err != nil {
		fmt.Printf("Error restarting containerd: %v\n", err)
	}

	// Stopping and removing Docker Compose managed resources
	fmt.Println("Stopping and removing Docker Compose managed resources...")
	err = executeCommand("sudo", "docker-compose", "down")
	if err != nil {
		fmt.Printf("Error stopping Docker Compose managed resources: %v\n", err)
	}

	// Uninstall Docker Compose
	fmt.Println("Uninstalling Docker Compose...")
	err = executeCommand("sudo", "rm", "/usr/local/bin/docker-compose")
	if err != nil {
		fmt.Printf("Error uninstalling Docker Compose: %v\n", err)
	}

	// Uninstall Docker registry
	fmt.Println("Uninstalling Docker registry...")
	err = executeCommand("sudo", "kubectl", "delete", "-f", "registry/registry-deployment.yaml")
	if err != nil {
		fmt.Printf("Error uninstalling Docker registry deployment: %v\n", err)
	}
	err = executeCommand("sudo", "kubectl", "delete", "-f", "registry/registry-pv-pvc.yaml")
	if err != nil {
		fmt.Printf("Error uninstalling Docker registry PV/PVC: %v\n", err)
	}

	err = executeCommand("sudo", "sed", "-i", "/registry/d", "/etc/hosts")
	if err != nil {
		fmt.Println("No entry for 'registry' found in /etc/hosts.")
	}
	err = executeCommand("sudo", "rm", "-rf", "/volumes/data/default_registry")
	if err != nil {
		fmt.Printf("Error removing default registry data: %v\n", err)
	}

	// Uninstall Docker
	fmt.Println("Uninstalling Docker...")
	err = executeCommand("sudo", "systemctl", "stop", "docker")
	if err != nil {
		fmt.Printf("Error stopping Docker: %v\n", err)
	}

	err = executeCommand("sudo", "apt-get", "remove", "-y", "docker.io")
	if err != nil {
		fmt.Printf("Error removing Docker package: %v\n", err)
	}

	err = executeCommand("sudo", "rm", "-rf", "/etc/docker")
	if err != nil {
		fmt.Printf("Error removing Docker configuration: %v\n", err)
	}

	err = executeCommand("sudo", "rm", "-rf", "/var/lib/docker")
	if err != nil {
		fmt.Printf("Error removing Docker library: %v\n", err)
	}

	err = executeCommand("sudo", "rm", "/etc/systemd/system/docker.service")
	if err != nil {
		fmt.Printf("Error removing Docker service file: %v\n", err)
	}

	err = executeCommand("sudo", "rm", "/etc/systemd/system/docker.socket")
	if err != nil {
		fmt.Printf("Error removing Docker socket file: %v\n", err)
	}

	err = executeCommand("sudo", "rm", "-rf", "/volumes/docker")
	if err != nil {
		fmt.Printf("Error removing Docker volumes: %v\n", err)
	}

	err = executeCommand("sudo", "systemctl", "daemon-reload")
	if err != nil {
		fmt.Printf("Error reloading daemon: %v\n", err)
	}

	fmt.Println("Docker uninstalled successfully.")
	return nil
}

func uninstallK8s() error {

	fmt.Println("Uninstalling Kubernetes...")
	err := executeCommand("sh", "-c", "echo y | sudo sealos reset 2>/dev/null || true")
	if err != nil {
		fmt.Printf("Error resetting Kubernetes using Sealos: %v\n", err)
	}
	time.Sleep(1 * time.Second)

	/*fmt.Println("Unmask containerd...")
	err = executeCommand("sh", "-c", "sudo systemctl unmask containerd && echo 'Unmask successful' || echo 'Unmask failed'")
	if err != nil {
		fmt.Printf("Error unmasking containerd: %v\n", err)
	}
	time.Sleep(1 * time.Second)*/

	fmt.Println("Uninstall sealos")
	err = executeCommand("sudo", "dpkg", "--purge", "sealos")
	if err != nil {
		fmt.Printf("Error purging Sealos: %v\n", err)
	}

	fmt.Println("Kubernetes uninstalled successfully.")
	return nil
}

func uninstallOpeniiot() error {
	fmt.Println("Uninstalling OpenIIOT...")

	// Delete all namespaces starting with 'openiiot'
	namespacesOutput, err := exec.Command("sudo", "kubectl", "get", "namespaces", "--no-headers", "-o", "custom-columns=:metadata.name").Output()
	if err != nil {
		fmt.Printf("Error getting namespaces: %v\n", err)
	} else {
		namespaces := strings.Split(string(namespacesOutput), "\n")
		for _, ns := range namespaces {
			if strings.HasPrefix(ns, "openiiot") {
				fmt.Printf("Deleting namespace: %s\n", ns)
				err = executeCommand("sudo", "kubectl", "delete", "namespace", ns)
				if err != nil {
					fmt.Printf("Error deleting namespace %s: %v\n", ns, err)
				}
			}
		}
	}

	// Delete all PersistentVolumes starting with 'openiiot'
	pvsOutput, err := exec.Command("sudo", "kubectl", "get", "pv", "--no-headers", "-o", "custom-columns=:metadata.name").Output()
	if err != nil {
		fmt.Printf("Error getting persistent volumes: %v\n", err)
	} else {
		pvs := strings.Split(string(pvsOutput), "\n")
		for _, pv := range pvs {
			if strings.HasPrefix(pv, "openiiot") {
				fmt.Printf("Deleting PersistentVolume: %s\n", pv)
				err = executeCommand("sudo", "kubectl", "delete", "pv", pv)
				if err != nil {
					fmt.Printf("Error deleting PersistentVolume %s: %v\n", pv, err)
				}
			}
		}
	}

	// Delete volume data
	fmt.Println("Deleting volume data...")
	err = executeCommand("sudo", "rm", "-rf", "/volumes/openiiot")
	if err != nil {
		fmt.Printf("Error deleting volume data: %v\n", err)
	}

	// Unload images
	fmt.Println("Unloading OpenIIOT images...")
	images := []string{
		"docker.io/library/openiiot-server:1.0.0",
		"docker.io/library/openiiot-consolemanager:1.0.0",
		"docker.io/library/openiiot-nodered:1.0.0",
		"docker.io/library/openiiot-grafana:1.0.0",
		"docker.io/library/openiiot-emqx:1.0.0",
	}
	for _, image := range images {
		err = executeCommand("sudo", "ctr", "-n", "k8s.io", "images", "delete", image)
		if err != nil {
			fmt.Printf("Error deleting image %s: %v\n", image, err)
		}
	}

	fmt.Println("OpenIIOT uninstalled successfully.")
	return nil
}

func createSystemdService() error {
	serviceContent := `[Unit]
Description=Unmask containerd after reboot
After=multi-user.target

[Service]
Type=oneshot
ExecStart=/bin/sh -c 'sudo systemctl unmask containerd && sudo systemctl disable unmask-containerd.service'
RemainAfterExit=true

[Install]
WantedBy=multi-user.target
`

	// Write the service file to /etc/systemd/system/unmask-containerd.service
	err := os.WriteFile("/etc/systemd/system/unmask-containerd.service", []byte(serviceContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write systemd service file: %v", err)
	}

	// Enable the service to run at the next boot
	err = executeCommand("sudo", "systemctl", "enable", "unmask-containerd.service")
	if err != nil {
		return fmt.Errorf("failed to enable systemd service: %v", err)
	}

	return nil
}

// Handle uninstallation process with subcommands
func uninstall() {
	if len(os.Args) > 2 {
		subcommand := os.Args[2]
		switch subcommand {
		case "openiiot":
			if confirmAction("Are you sure you want to uninstall OpenIIOT? This action cannot be undone.") {
				err := uninstallOpeniiot()
				if err != nil {
					fmt.Println("Error uninstalling OpenIIOT:", err)
				} else {
					fmt.Println("OpenIIOT uninstalled successfully.")
				}
			} else {
				fmt.Println("Uninstallation of OpenIIOT aborted.")
			}
		case "supabase":
			if confirmAction("Are you sure you want to uninstall Supabase? This action cannot be undone.") {
				err := uninstallSupabase()
				if err != nil {
					fmt.Println("Error uninstalling Supabase:", err)
				} else {
					fmt.Println("Supabase uninstalled successfully.")
				}
			} else {
				fmt.Println("Uninstallation of Supabase aborted.")
			}
		default:
			fmt.Println("Unknown uninstall subcommand:", subcommand)
			os.Exit(1)
		}
	} else {
		if confirmAction("Are you sure you want to uninstall FreeFlow? This action cannot be undone.") {
			fmt.Println("Proceeding with uninstallation...")

			err := uninstallNginx()
			if err != nil {
				fmt.Printf("Error uninstalling Nginx: %v\n", err)
			}

			err = uninstallSupabase()
			if err != nil {
				fmt.Printf("Error uninstalling Supabase: %v\n", err)
			}

			err = uninstallDocker()
			if err != nil {
				fmt.Printf("Error uninstalling Docker: %v\n", err)
			}

			err = uninstallOpeniiot()
			if err != nil {
				fmt.Printf("Error uninstalling OpenIIOT: %v\n", err)
			}

			err = uninstallK8s()
			if err != nil {
				fmt.Printf("Error uninstalling Kubernetes: %v\n", err)
			}

			// Cleanup remaining configurations
			fmt.Println("Cleaning up remaining configurations...")
			err = executeCommand("sudo", "apt-get", "autoremove", "-y")
			if err != nil {
				fmt.Printf("Error during cleanup: %v\n", err)
			}
			err = executeCommand("sudo", "apt-get", "clean")
			if err != nil {
				fmt.Printf("Error during cleanup: %v\n", err)
			}

			// Create systemd service to mask containerd after reboot
			err = createSystemdService()
			if err != nil {
				fmt.Printf("Error creating systemd service: %v\n", err)
			} else {
				fmt.Println("Systemd service created to unmask containerd after reboot.")
			}

			// Reboot system if confirmed by the user
			if confirmAction("FreeFlow uninstallation complete. It's recommended to reboot the system to fully clean up. Would you like to reboot now?") {
				fmt.Println("Rebooting the system...")
				err = executeCommand("sudo", "reboot")
				if err != nil {
					fmt.Printf("Error rebooting the system: %v\n", err)
				}
			} else {
				fmt.Println("Reboot skipped. Please reboot the system manually if needed.")
			}
		} else {
			fmt.Println("Uninstallation aborted.")
		}
	}
}

// Handle version output
func version() {
	fmt.Println("FreeFlow Version: v1.0.0")
}

// Helper function to wait for a Kubernetes job to complete with a timeout
func waitForJobCompletion(jobName, namespace string, timeoutSeconds int) error {
	startTime := time.Now()
	fmt.Printf("Waiting for job %s to complete in namespace %s", jobName, namespace)
	for {
		// Check if the job is complete
		output, err := exec.Command("sudo", "kubectl", "get", "job", jobName, "-n", namespace, "-o", "jsonpath={.status.succeeded}").Output()
		if err != nil {
			return fmt.Errorf("failed to get status of job %s: %v", jobName, err)
		}

		if string(output) == "1" {
			fmt.Println("\nJob completed successfully.")
			break
		}

		// Check if the timeout has been reached
		if time.Since(startTime).Seconds() > float64(timeoutSeconds) {
			return fmt.Errorf("timeout reached waiting for job %s to complete", jobName)
		}

		// Wait for a short interval before checking again
		time.Sleep(10 * time.Second)
		fmt.Print(".")
	}
	return nil
}

// Helper function to wait for a Kubernetes deployment to be ready
func waitForDeploymentCompletion(deploymentName, namespace string, timeoutSeconds int) error {
	startTime := time.Now()
	fmt.Printf("Waiting for deployment %s in namespace %s", deploymentName, namespace)
	for {
		// Check the deployment status
		availableReplicas, err := exec.Command("sudo", "kubectl", "get", "deployment", deploymentName, "-n", namespace, "-o", "jsonpath={.status.availableReplicas}").Output()
		if err != nil {
			return fmt.Errorf("failed to get status of deployment %s: %v", deploymentName, err)
		}
		desiredReplicas, err := exec.Command("sudo", "kubectl", "get", "deployment", deploymentName, "-n", namespace, "-o", "jsonpath={.spec.replicas}").Output()
		if err != nil {
			return fmt.Errorf("failed to get spec of deployment %s: %v", deploymentName, err)
		}

		if string(availableReplicas) == string(desiredReplicas) && len(availableReplicas) > 0 {
			fmt.Println("\nDeployment completed successfully.")
			break
		}

		// Check if the timeout has been reached
		if time.Since(startTime).Seconds() > float64(timeoutSeconds) {
			return fmt.Errorf("timeout reached waiting for deployment %s to complete", deploymentName)
		}

		// Wait for a short interval before checking again
		time.Sleep(10 * time.Second)
		fmt.Print(".")
	}
	return nil
}

func installOpeniiot() error {
	fmt.Println("Installing OpenIIOT components...")

	// Check for existing PersistentVolumes
	output, err := exec.Command("sudo", "kubectl", "get", "pv").Output()
	if err == nil && strings.Contains(string(output), "openiiot") {
		return fmt.Errorf("there are PersistentVolumes for 'openiiot'. Please delete them before proceeding")
	}

	// Check for existing 'openiiot' namespace
	_, err = exec.Command("sudo", "kubectl", "get", "namespace", "openiiot").Output()
	if err == nil {
		return fmt.Errorf("namespace 'openiiot' already exists. Please delete it before proceeding")
	}

	// Create the 'openiiot' namespace
	err = executeCommand("sudo", "kubectl", "create", "namespace", "openiiot")
	if err != nil {
		return fmt.Errorf("failed to create namespace 'openiiot': %v", err)
	}

	// Load container images
	fmt.Println("Loading container images...")
	images := []string{
		"bin/openiiot-server.tar",
		"bin/openiiot-consolemanager.tar",
		"bin/openiiot-nodered.tar",
		"bin/openiiot-grafana.tar",
		"bin/openiiot-emqx.tar",
	}
	for _, image := range images {
		err = executeCommand("sudo", "ctr", "-n", "k8s.io", "images", "import", image)
		if err != nil {
			return fmt.Errorf("failed to load image %s: %v", image, err)
		}
	}

	// Install MySQL
	fmt.Println("Installing MySQL...")
	err = executeCommand("sudo", "mkdir", "-p", "/volumes/openiiot/mysql")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "chmod", "777", "/volumes/openiiot/mysql")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "chown", "nobody:nogroup", "/volumes/openiiot/mysql")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-mysql-pv.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-mysql-pvc.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-mysql-deployment.yaml")
	if err != nil {
		return err
	}
	err = waitForDeploymentCompletion("openiiot-mysql", "openiiot", 300)
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-mysql-service.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-mysql-cm.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-mysql-job.yaml")
	if err != nil {
		return err
	}
	err = waitForJobCompletion("openiiot-mysql-job", "openiiot", 300)
	if err != nil {
		return err
	}

	// Install database tools
	fmt.Println("Installing database tools...")
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-pma-deployment.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-pma-service.yaml")
	if err != nil {
		return err
	}

	// Install the server
	fmt.Println("Installing OpenIIOT server...")
	err = executeCommand("sudo", "mkdir", "-p", "/volumes/openiiot")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "chmod", "777", "/volumes/openiiot")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "chown", "-R", "nobody:nogroup", "/volumes/openiiot")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-server-pv.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-server-pvc.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-server-rbac.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-server-deployment.yaml")
	if err != nil {
		return err
	}
	err = waitForDeploymentCompletion("openiiot-server", "openiiot", 300)
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-server-service.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-server-ingress.yaml")
	if err != nil {
		return err
	}

	// Prepare data
	fmt.Println("Preparing server data...")
	err = executeCommand("sudo", "mkdir", "-p", "/volumes/openiiot/server")
	if err != nil {
		return err
	}
	err = executeCommand("sh", "-c", "sudo cp -f volumes/* /volumes/openiiot/server")
	if err != nil {
		return err
	}

	// Create default tenant
	fmt.Println("Creating default tenant...")
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-server-job.yaml")
	if err != nil {
		return err
	}
	err = waitForJobCompletion("openiiot-server-job", "openiiot", 300)
	if err != nil {
		return err
	}

	// Install Console Manager
	fmt.Println("Installing Console Manager...")
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-consolemanager-deployment.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-consolemanager-service.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "manifest/openiiot-consolemanager-ingress.yaml")
	if err != nil {
		return err
	}

	fmt.Println("OpenIIOT installed successfully.")
	return nil
}

func updateApt() error {
	fmt.Println("Updating package list...")
	err := executeCommand("sudo", "apt-get", "update")
	if err != nil {
		fmt.Printf("Error updating package list: %v\n", err)
		return err
	}
	fmt.Println("Package list updated successfully.")
	return nil
}

func installK8s() error {
	// Run the environment check first
	err := check()
	if err != nil {
		return err
	}

	// Update apt package list
	err = updateApt()
	if err != nil {
		return err
	}

	fmt.Println("Installing Kubernetes...")

	// Install Sealos
	err = executeCommand("sudo", "dpkg", "-i", "./middleware/sealos_4.3.7_linux_amd64.deb")
	if err != nil {
		return err
	}

	// Run Sealos to set up Kubernetes, Helm, Cilium, and Ingress-Nginx
	err = executeCommand("sudo", "sealos", "run", "labring/kubernetes:v1.29.7", "labring/helm:v3.14.1", "labring/cilium:v1.14.13", "labring/ingress-nginx:v1.11.2")
	return err
}

// Function to install Docker and Docker Compose (completed with containerd config modification)
func installDocker() error {
	fmt.Println("Installing Docker...")

	// Install Docker
	err := executeCommand("sudo", "DEBIAN_FRONTEND=noninteractive", "apt-get", "install", "-y", "docker.io")
	if err != nil {
		return err
	}

	// Setup Docker directory
	err = executeCommand("sudo", "mkdir", "-p", "/volumes/docker")
	if err != nil {
		return err
	}

	// Copy daemon.json
	err = executeCommand("sudo", "cp", "./config/daemon.json", "/etc/docker/daemon.json")
	if err != nil {
		return err
	}

	// Restart Docker
	err = executeCommand("sudo", "systemctl", "restart", "docker")
	if err != nil {
		return err
	}

	// Install Docker Compose
	err = executeCommand("sudo", "cp", "./middleware/docker-compose-linux-x86_64", "/usr/local/bin/docker-compose")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "chmod", "+x", "/usr/local/bin/docker-compose")
	if err != nil {
		return err
	}

	// Install Docker registry
	err = executeCommand("sudo", "mkdir", "-p", "/volumes/data/default_registry")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "registry/registry-pv-pvc.yaml")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "kubectl", "apply", "-f", "registry/registry-deployment.yaml")
	if err != nil {
		return err
	}

	// Retrieve the ClusterIP for the registry service
	clusterIP, err := exec.Command("sudo", "kubectl", "get", "svc", "registry", "-o", "jsonpath={.spec.clusterIP}").Output()
	if err != nil || len(clusterIP) == 0 {
		return fmt.Errorf("failed to retrieve ClusterIP for the 'registry' service")
	}

	// Update /etc/hosts with the registry’s ClusterIP
	fmt.Println("Updating /etc/hosts with Docker registry ClusterIP...")
	hostsEntry := fmt.Sprintf("%s registry", string(clusterIP))
	if strings.Contains(string(clusterIP), "registry") {
		err = executeCommand("sudo", "sed", "-i", fmt.Sprintf("/registry/c\\%s", hostsEntry), "/etc/hosts")
	} else {
		err = executeCommand("sudo", "sh", "-c", fmt.Sprintf("echo '%s' >> /etc/hosts", hostsEntry))
	}
	if err != nil {
		return err
	}

	// Add Docker registry to containerd config
	fmt.Println("Modifying containerd configuration for the Docker registry...")
	err = executeCommand("sudo", "mkdir", "-p", "/etc/containerd/certs.d/registry:5000")
	if err != nil {
		return err
	}
	err = executeCommand("sudo", "cp", "./config/hosts.toml", "/etc/containerd/certs.d/registry:5000/hosts.toml")
	if err != nil {
		return err
	}

	// Restart containerd to apply changes
	err = executeCommand("sudo", "systemctl", "restart", "containerd")
	if err != nil {
		return err
	}

	fmt.Println("Docker and Docker Compose installed and configured successfully.")
	return nil
}

func installSupabase() error {
	fmt.Println("Installing Supabase...")

	// Load the Supabase Docker image
	err := executeCommand("sudo", "docker", "load", "-i", "bin/openiiot-supabase-studio.tar")
	if err != nil {
		return fmt.Errorf("failed to load Supabase Docker image: %v", err)
	}

	// Start the Supabase containers using Docker Compose
	err = executeCommand("sudo", "docker-compose", "-f", "supabase/docker-compose.yml", "up", "-d")
	if err != nil {
		return fmt.Errorf("failed to start Supabase with Docker Compose: %v", err)
	}

	fmt.Println("Supabase installed successfully.")
	return nil
}

// Function to install and configure Nginx (converted from install_nginx.sh)
func installNginx() error {
	fmt.Println("Installing and configuring Nginx...")

	// Start the Nginx service using Docker Compose
	err := executeCommand("sudo", "docker-compose", "-f", "nginx/docker-compose.yml", "up", "-d")
	if err != nil {
		return fmt.Errorf("failed to start Nginx with Docker Compose: %v", err)
	}

	// Retrieve the current Ingress Controller's HTTP NodePort
	nodePort, err := exec.Command("sudo", "kubectl", "get", "svc", "-n", "ingress-nginx", "ingress-nginx-controller", "-o", "jsonpath={.spec.ports[?(@.port==80)].nodePort}").Output()
	if err != nil || len(nodePort) == 0 {
		return fmt.Errorf("failed to retrieve NodePort for Ingress Controller: %v", err)
	}
	fmt.Printf("Current HTTP NodePort is %s\n", string(nodePort))

	// Generate new content for the default.conf file
	newConf := fmt.Sprintf(`server {
    listen 80 default_server;
    listen [::]:80 default_server;

    server_name _;

    location / {
        proxy_pass http://127.0.0.1:%s;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

		proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "Upgrade";
    }

	location /supabase/ {
        proxy_pass http://127.0.0.1:8000/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    error_page 502 503 504 /50x.html;
    location = /50x.html {
        root /usr/share/nginx/html;
    }
}`, string(nodePort))

	// Write the new default.conf content into the Nginx container's configuration file
	containerName := "nginx-server"
	err = exec.Command("sudo", "docker", "exec", containerName, "bash", "-c", fmt.Sprintf("echo '%s' > /etc/nginx/conf.d/default.conf", newConf)).Run()
	if err != nil {
		return fmt.Errorf("failed to update Nginx configuration: %v", err)
	}

	// Reload Nginx to apply the new configuration
	err = executeCommand("sudo", "docker", "exec", containerName, "nginx", "-s", "reload")
	if err != nil {
		return fmt.Errorf("failed to reload Nginx: %v", err)
	}

	fmt.Println("Nginx default.conf updated and reloaded successfully.")
	return nil
}

// Main function to parse commands
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: freeflow <command>")
		fmt.Println("Commands: install, uninstall, version")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "install":
		install()
	case "uninstall":
		uninstall()
	case "version":
		version()
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Usage: freeflow <install|uninstall|version>")
		os.Exit(1)
	}
}
