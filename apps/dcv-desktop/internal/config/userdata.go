package config

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/scttfrdmn/lens/apps/dcv-desktop/internal/environments"
)

// GenerateUserData creates a cloud-init user data script for DCV Desktop
func GenerateUserData(env *environments.DesktopEnvironment, idleTimeoutSeconds int) (string, error) {
	script := generateUserDataScript(env, idleTimeoutSeconds)
	// AWS expects user data to be base64 encoded
	encoded := base64.StdEncoding.EncodeToString([]byte(script))
	return encoded, nil
}

// generateUserDataScript creates the actual bash script for DCV Desktop setup
func generateUserDataScript(env *environments.DesktopEnvironment, idleTimeoutSeconds int) string {
	var sb strings.Builder

	// Start with bash shebang and error handling
	sb.WriteString("#!/bin/bash\n")
	sb.WriteString("set -e\n")
	sb.WriteString("set -x\n\n")

	// Log file for debugging
	sb.WriteString("exec > >(tee /var/log/user-data.log)\n")
	sb.WriteString("exec 2>&1\n\n")

	// Create progress log file
	sb.WriteString("# Setup progress tracking\n")
	sb.WriteString("PROGRESS_LOG=\"/var/log/setup-progress.log\"\n")
	sb.WriteString("touch $PROGRESS_LOG\n")
	sb.WriteString("chmod 644 $PROGRESS_LOG\n\n")

	sb.WriteString("log_progress() {\n")
	sb.WriteString("  echo \"STEP:$1\" | tee -a $PROGRESS_LOG\n")
	sb.WriteString("}\n\n")

	sb.WriteString(fmt.Sprintf("log_progress 'Starting lens-dcv-desktop environment setup (%s)'\n", env.Name))
	sb.WriteString(fmt.Sprintf("echo 'Environment: %s'\n", env.Description))
	sb.WriteString(fmt.Sprintf("echo 'Desktop Type: %s'\n\n", env.DesktopType))

	// Update system
	sb.WriteString("# Update system packages\n")
	sb.WriteString("log_progress 'Updating system packages'\n")
	sb.WriteString("apt-get update -y\n")
	sb.WriteString("DEBIAN_FRONTEND=noninteractive apt-get upgrade -y\n\n")

	// Install SSM Agent
	sb.WriteString("# Install AWS Systems Manager Agent\n")
	sb.WriteString("log_progress 'Installing SSM Agent'\n")
	sb.WriteString("snap install amazon-ssm-agent --classic\n")
	sb.WriteString("systemctl enable snap.amazon-ssm-agent.amazon-ssm-agent.service\n")
	sb.WriteString("systemctl start snap.amazon-ssm-agent.amazon-ssm-agent.service\n\n")

	// GPU setup if required
	if env.RequiresGPU {
		sb.WriteString(generateGPUSetupScript())
	}

	// Install desktop environment
	sb.WriteString(generateDesktopInstallScript(env.DesktopType))

	// Install NICE DCV Server
	sb.WriteString(generateDCVInstallScript(env))

	// Install pre-installed applications
	if len(env.PreInstalledApps) > 0 {
		sb.WriteString("# Install pre-configured applications\n")
		sb.WriteString("log_progress 'Installing desktop applications'\n")
		sb.WriteString("DEBIAN_FRONTEND=noninteractive apt-get install -y \\\n")
		for i, app := range env.PreInstalledApps {
			if i == len(env.PreInstalledApps)-1 {
				sb.WriteString("  " + app + "\n\n")
			} else {
				sb.WriteString("  " + app + " \\\n")
			}
		}
	}

	// Configure DCV session
	sb.WriteString(generateDCVSessionScript(env))

	// Setup idle timeout monitoring
	if idleTimeoutSeconds > 0 {
		sb.WriteString(generateIdleMonitorScript(idleTimeoutSeconds))
	}

	// Set environment variables
	if len(env.EnvironmentVars) > 0 {
		sb.WriteString("# Configure environment variables\n")
		sb.WriteString("cat >> /etc/environment << 'EOF'\n")
		for key, value := range env.EnvironmentVars {
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		}
		sb.WriteString("EOF\n\n")
	}

	// Final setup steps
	sb.WriteString("# Finalize setup\n")
	sb.WriteString("log_progress 'DCV Desktop setup complete'\n")
	sb.WriteString("echo 'SETUP_COMPLETE' > /var/log/dcv-ready\n\n")

	// Reboot to ensure all changes take effect
	sb.WriteString("# Reboot to apply all changes\n")
	sb.WriteString("log_progress 'Rebooting system'\n")
	sb.WriteString("reboot\n")

	return sb.String()
}

// generateGPUSetupScript creates the GPU driver installation script
func generateGPUSetupScript() string {
	var sb strings.Builder

	sb.WriteString("# GPU Setup\n")
	sb.WriteString("log_progress 'Detecting GPU'\n")
	sb.WriteString("if lspci | grep -i nvidia; then\n")
	sb.WriteString("  log_progress 'NVIDIA GPU detected, installing drivers'\n")
	sb.WriteString("  \n")
	sb.WriteString("  # Install NVIDIA drivers\n")
	sb.WriteString("  apt-get install -y linux-headers-$(uname -r)\n")
	sb.WriteString("  apt-get install -y nvidia-driver-535\n")
	sb.WriteString("  \n")
	sb.WriteString("  # Install CUDA toolkit\n")
	sb.WriteString("  wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2204/x86_64/cuda-keyring_1.0-1_all.deb\n")
	sb.WriteString("  dpkg -i cuda-keyring_1.0-1_all.deb\n")
	sb.WriteString("  apt-get update\n")
	sb.WriteString("  apt-get install -y cuda-toolkit-12-2\n")
	sb.WriteString("  \n")
	sb.WriteString("  # Add CUDA to PATH\n")
	sb.WriteString("  echo 'export PATH=/usr/local/cuda/bin:$PATH' >> /etc/profile.d/cuda.sh\n")
	sb.WriteString("  echo 'export LD_LIBRARY_PATH=/usr/local/cuda/lib64:$LD_LIBRARY_PATH' >> /etc/profile.d/cuda.sh\n")
	sb.WriteString("  \n")
	sb.WriteString("  log_progress 'GPU drivers installed'\n")
	sb.WriteString("else\n")
	sb.WriteString("  log_progress 'No GPU detected, skipping GPU setup'\n")
	sb.WriteString("fi\n\n")

	return sb.String()
}

// generateDesktopInstallScript creates the desktop environment installation script
func generateDesktopInstallScript(desktopType string) string {
	var sb strings.Builder

	sb.WriteString("# Install Desktop Environment\n")
	sb.WriteString(fmt.Sprintf("log_progress 'Installing %s desktop environment'\n", desktopType))

	switch desktopType {
	case "gnome":
		sb.WriteString("DEBIAN_FRONTEND=noninteractive apt-get install -y \\\n")
		sb.WriteString("  ubuntu-desktop \\\n")
		sb.WriteString("  gnome-session \\\n")
		sb.WriteString("  gdm3\n\n")
	case "xfce":
		sb.WriteString("DEBIAN_FRONTEND=noninteractive apt-get install -y \\\n")
		sb.WriteString("  xubuntu-desktop \\\n")
		sb.WriteString("  xfce4 \\\n")
		sb.WriteString("  xfce4-goodies \\\n")
		sb.WriteString("  lightdm\n\n")
	case "minimal":
		sb.WriteString("DEBIAN_FRONTEND=noninteractive apt-get install -y \\\n")
		sb.WriteString("  xserver-xorg \\\n")
		sb.WriteString("  openbox \\\n")
		sb.WriteString("  lightdm\n\n")
	}

	sb.WriteString("systemctl set-default graphical.target\n\n")

	return sb.String()
}

// generateDCVInstallScript creates the NICE DCV installation script
func generateDCVInstallScript(env *environments.DesktopEnvironment) string {
	var sb strings.Builder

	sb.WriteString("# Install NICE DCV Server\n")
	sb.WriteString("log_progress 'Installing NICE DCV Server'\n\n")

	// Download and install DCV
	sb.WriteString("cd /tmp\n")
	sb.WriteString("wget https://d1uj6qtbmh3dt5.cloudfront.net/nice-dcv-ubuntu2204-x86_64.tgz\n")
	sb.WriteString("tar -xvzf nice-dcv-ubuntu2204-x86_64.tgz\n")
	sb.WriteString("cd nice-dcv-*-ubuntu2204-x86_64\n\n")

	sb.WriteString("# Install DCV packages\n")
	sb.WriteString("apt-get install -y ./nice-dcv-server_*.deb\n")
	sb.WriteString("apt-get install -y ./nice-dcv-web-viewer_*.deb\n")
	sb.WriteString("apt-get install -y ./nice-xdcv_*.deb\n\n")

	// Configure DCV
	sb.WriteString("# Configure DCV Server\n")
	sb.WriteString("cat > /etc/dcv/dcv.conf << 'EOF'\n")
	sb.WriteString("[connectivity]\n")
	sb.WriteString(fmt.Sprintf("web-port=%d\n", env.DCVSettings.Port))
	sb.WriteString("web-use-https=true\n")
	sb.WriteString("\n")
	sb.WriteString("[security]\n")
	sb.WriteString("authentication=\"system\"\n")
	sb.WriteString("\n")
	sb.WriteString("[session-management/automatic-console-session]\n")
	sb.WriteString("owner=\"ubuntu\"\n")
	sb.WriteString("\n")
	sb.WriteString("[display]\n")
	sb.WriteString("enable-cu-desktops=true\n")

	if env.DCVSettings.EnableGPU {
		sb.WriteString("\n")
		sb.WriteString("[gpu]\n")
		sb.WriteString("enable-gpu=true\n")
	}

	sb.WriteString("EOF\n\n")

	// Enable and start DCV
	sb.WriteString("# Enable and start DCV service\n")
	sb.WriteString("systemctl enable dcvserver\n")
	sb.WriteString("systemctl start dcvserver\n\n")

	// Set ubuntu user password for DCV login
	sb.WriteString("# Set ubuntu user password for DCV login\n")
	sb.WriteString("echo 'ubuntu:$(openssl rand -base64 12)' | chpasswd\n\n")

	return sb.String()
}

// generateDCVSessionScript creates the DCV session creation script
func generateDCVSessionScript(env *environments.DesktopEnvironment) string {
	var sb strings.Builder

	sb.WriteString("# Create DCV Session\n")
	sb.WriteString("log_progress 'Creating DCV session'\n\n")

	sb.WriteString("# Wait for DCV server to be ready\n")
	sb.WriteString("for i in {1..30}; do\n")
	sb.WriteString("  if systemctl is-active --quiet dcvserver; then\n")
	sb.WriteString("    break\n")
	sb.WriteString("  fi\n")
	sb.WriteString("  sleep 2\n")
	sb.WriteString("done\n\n")

	sb.WriteString("# Create virtual session\n")
	sb.WriteString("dcv create-session --type=virtual --owner ubuntu lens-desktop\n\n")

	return sb.String()
}

// generateIdleMonitorScript creates the idle timeout monitoring script
func generateIdleMonitorScript(idleTimeoutSeconds int) string {
	var sb strings.Builder

	sb.WriteString("# Setup idle timeout monitoring\n")
	sb.WriteString(fmt.Sprintf("log_progress 'Configuring auto-stop after %d seconds of idle'\n\n", idleTimeoutSeconds))

	sb.WriteString("# Create idle monitor script\n")
	sb.WriteString("cat > /usr/local/bin/dcv-idle-monitor.sh << 'EOF'\n")
	sb.WriteString("#!/bin/bash\n")
	sb.WriteString(fmt.Sprintf("IDLE_TIMEOUT=%d\n", idleTimeoutSeconds))
	sb.WriteString("\n")
	sb.WriteString("while true; do\n")
	sb.WriteString("  # Check if any DCV sessions are active\n")
	sb.WriteString("  if dcv list-sessions | grep -q 'lens-desktop'; then\n")
	sb.WriteString("    # Check for active connections\n")
	sb.WriteString("    CONNECTIONS=$(dcv list-connections -session lens-desktop | wc -l)\n")
	sb.WriteString("    if [ $CONNECTIONS -eq 0 ]; then\n")
	sb.WriteString("      # No active connections, shutdown\n")
	sb.WriteString("      logger 'DCV idle timeout reached, shutting down'\n")
	sb.WriteString("      shutdown -h now\n")
	sb.WriteString("    fi\n")
	sb.WriteString("  fi\n")
	sb.WriteString("  sleep 60\n")
	sb.WriteString("done\n")
	sb.WriteString("EOF\n\n")

	sb.WriteString("chmod +x /usr/local/bin/dcv-idle-monitor.sh\n\n")

	sb.WriteString("# Create systemd service for idle monitor\n")
	sb.WriteString("cat > /etc/systemd/system/dcv-idle-monitor.service << 'EOF'\n")
	sb.WriteString("[Unit]\n")
	sb.WriteString("Description=DCV Idle Monitor\n")
	sb.WriteString("After=dcvserver.service\n")
	sb.WriteString("\n")
	sb.WriteString("[Service]\n")
	sb.WriteString("Type=simple\n")
	sb.WriteString("ExecStart=/usr/local/bin/dcv-idle-monitor.sh\n")
	sb.WriteString("Restart=on-failure\n")
	sb.WriteString("\n")
	sb.WriteString("[Install]\n")
	sb.WriteString("WantedBy=multi-user.target\n")
	sb.WriteString("EOF\n\n")

	sb.WriteString("systemctl daemon-reload\n")
	sb.WriteString("systemctl enable dcv-idle-monitor.service\n")
	sb.WriteString("systemctl start dcv-idle-monitor.service\n\n")

	return sb.String()
}
