package environments

import (
	"fmt"
)

// DesktopEnvironment defines a desktop configuration for DCV
type DesktopEnvironment struct {
	Name            string            `yaml:"name"`
	Description     string            `yaml:"description"`
	InstanceType    string            `yaml:"instance_type"`
	RequiresGPU     bool              `yaml:"requires_gpu"`
	EBSVolumeSize   int               `yaml:"ebs_volume_size"`
	DesktopType     string            `yaml:"desktop_type"` // gnome, xfce, or minimal
	PreInstalledApps []string         `yaml:"pre_installed_apps"`
	EnvironmentVars map[string]string `yaml:"environment_vars"`
	DCVSettings     DCVConfig         `yaml:"dcv_settings"`
}

// DCVConfig holds DCV-specific configuration
type DCVConfig struct {
	Port            int    `yaml:"port"`
	SessionType     string `yaml:"session_type"` // virtual or console
	Quality         string `yaml:"quality"`      // high, medium, low
	EnableGPU       bool   `yaml:"enable_gpu"`
	EnableUSB       bool   `yaml:"enable_usb"`
	EnableClipboard bool   `yaml:"enable_clipboard"`
}

// GetDefaultEnvironments returns all built-in desktop environments
func GetDefaultEnvironments() map[string]*DesktopEnvironment {
	return map[string]*DesktopEnvironment{
		"general-desktop": {
			Name:         "general-desktop",
			Description:  "Ubuntu desktop with common research tools",
			InstanceType: "t3.xlarge",
			RequiresGPU:  false,
			EBSVolumeSize: 50,
			DesktopType:  "xfce",
			PreInstalledApps: []string{
				"firefox",
				"code",
				"git",
				"python3",
				"r-base",
			},
			DCVSettings: DCVConfig{
				Port:            8443,
				SessionType:     "virtual",
				Quality:         "high",
				EnableGPU:       false,
				EnableUSB:       false,
				EnableClipboard: true,
			},
		},
		"gpu-workstation": {
			Name:         "gpu-workstation",
			Description:  "CUDA-enabled desktop for GPU computing",
			InstanceType: "g4dn.xlarge",
			RequiresGPU:  true,
			EBSVolumeSize: 100,
			DesktopType:  "xfce",
			PreInstalledApps: []string{
				"nvidia-cuda-toolkit",
				"nvidia-driver",
				"firefox",
				"code",
				"python3-tensorflow",
				"python3-pytorch",
			},
			DCVSettings: DCVConfig{
				Port:            8443,
				SessionType:     "virtual",
				Quality:         "high",
				EnableGPU:       true,
				EnableUSB:       false,
				EnableClipboard: true,
			},
		},
		"matlab-desktop": {
			Name:         "matlab-desktop",
			Description:  "Desktop configured for MATLAB (user provides license)",
			InstanceType: "g4dn.xlarge",
			RequiresGPU:  true,
			EBSVolumeSize: 100,
			DesktopType:  "xfce",
			PreInstalledApps: []string{
				"firefox",
				// MATLAB will be installed by user or via AMI
			},
			EnvironmentVars: map[string]string{
				"MATLAB_PREFDIR": "/home/ubuntu/.matlab",
			},
			DCVSettings: DCVConfig{
				Port:            8443,
				SessionType:     "virtual",
				Quality:         "high",
				EnableGPU:       true,
				EnableUSB:       true, // USB for hardware-in-loop
				EnableClipboard: true,
			},
		},
		"data-viz-desktop": {
			Name:         "data-viz-desktop",
			Description:  "ParaView and visualization tools",
			InstanceType: "g4dn.xlarge",
			RequiresGPU:  true,
			EBSVolumeSize: 100,
			DesktopType:  "xfce",
			PreInstalledApps: []string{
				"paraview",
				"visit",
				"firefox",
				"code",
			},
			DCVSettings: DCVConfig{
				Port:            8443,
				SessionType:     "virtual",
				Quality:         "high",
				EnableGPU:       true,
				EnableUSB:       false,
				EnableClipboard: true,
			},
		},
		"image-analysis": {
			Name:         "image-analysis",
			Description:  "ImageJ, Fiji, QuPath, CellProfiler",
			InstanceType: "t3.xlarge",
			RequiresGPU:  false,
			EBSVolumeSize: 50,
			DesktopType:  "xfce",
			PreInstalledApps: []string{
				"imagej",
				"fiji",
				"cellprofiler",
				"firefox",
			},
			DCVSettings: DCVConfig{
				Port:            8443,
				SessionType:     "virtual",
				Quality:         "high",
				EnableGPU:       false,
				EnableUSB:       false,
				EnableClipboard: true,
			},
		},
		"bioinformatics-gui": {
			Name:         "bioinformatics-gui",
			Description:  "Geneious, UGENE, bioinformatics tools",
			InstanceType: "t3.xlarge",
			RequiresGPU:  false,
			EBSVolumeSize: 100,
			DesktopType:  "xfce",
			PreInstalledApps: []string{
				"ugene",
				"jalview",
				"firefox",
				"code",
			},
			DCVSettings: DCVConfig{
				Port:            8443,
				SessionType:     "virtual",
				Quality:         "high",
				EnableGPU:       false,
				EnableUSB:       false,
				EnableClipboard: true,
			},
		},
	}
}

// Get returns a desktop environment by name
func Get(name string) (*DesktopEnvironment, error) {
	envs := GetDefaultEnvironments()
	env, ok := envs[name]
	if !ok {
		return nil, fmt.Errorf("desktop environment %q not found", name)
	}
	return env, nil
}

// List returns all available desktop environment names
func List() []string {
	envs := GetDefaultEnvironments()
	names := make([]string, 0, len(envs))
	for name := range envs {
		names = append(names, name)
	}
	return names
}

// GetRecommendedInstanceTypes returns GPU and non-GPU recommended instance types
func GetRecommendedInstanceTypes(requiresGPU bool) []string {
	if requiresGPU {
		return []string{
			"g4dn.xlarge",  // 4 vCPU, 16GB, NVIDIA T4
			"g4dn.2xlarge", // 8 vCPU, 32GB, NVIDIA T4
			"g5.xlarge",    // 4 vCPU, 16GB, NVIDIA A10G
			"g5.2xlarge",   // 8 vCPU, 32GB, NVIDIA A10G
		}
	}
	return []string{
		"t3.xlarge",  // 4 vCPU, 16GB
		"t3.2xlarge", // 8 vCPU, 32GB
		"m5.xlarge",  // 4 vCPU, 16GB
		"m5.2xlarge", // 8 vCPU, 32GB
	}
}
