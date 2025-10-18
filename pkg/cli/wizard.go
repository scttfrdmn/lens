package cli

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// WizardConfig holds the configuration options for the wizard
type WizardConfig struct {
	AppName      string
	AppType      string // "jupyter", "rstudio", "vscode"
	Environments map[string]string
}

// WizardResult holds the user's choices from the wizard
type WizardResult struct {
	Environment  string
	InstanceType string
	EBSSize      int
	IdleTimeout  string
	Name         string
}

// RunLaunchWizard runs an interactive wizard to help users launch their first instance
func RunLaunchWizard(config WizardConfig) (*WizardResult, error) {
	result := &WizardResult{}

	fmt.Println()
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Printf("  Welcome to %s! Let's set up your cloud environment.\n", config.AppName)
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println()

	// Question 1: What type of work do you want to do?
	envPrompt, envOptions := buildEnvironmentPrompt(config)
	var selectedEnv string
	err := survey.AskOne(envPrompt, &selectedEnv, survey.WithValidator(survey.Required))
	if err != nil {
		return nil, err
	}
	result.Environment = envOptions[selectedEnv]

	fmt.Println()

	// Question 2: How powerful a computer do you need?
	var powerChoice string
	powerPrompt := &survey.Select{
		Message: "How powerful a computer do you need?",
		Options: []string{
			"Light work (browsing data, small datasets) - $0.0042/hour",
			"Normal work (typical analysis) - $0.0084/hour [Recommended]",
			"Heavy work (large datasets, complex models) - $0.0168/hour",
			"Very heavy work (big data, deep learning) - $0.0336/hour",
		},
		Default: "Normal work (typical analysis) - $0.0084/hour [Recommended]",
	}
	err = survey.AskOne(powerPrompt, &powerChoice)
	if err != nil {
		return nil, err
	}
	result.InstanceType = mapPowerChoiceToInstanceType(powerChoice)

	fmt.Println()

	// Question 3: How much storage space do you need?
	var storageChoice string
	storagePrompt := &survey.Select{
		Message: "How much storage space do you need?",
		Options: []string{
			"Small (20 GB) - Good for trying things out",
			"Normal (50 GB) - Sufficient for most projects [Recommended]",
			"Large (100 GB) - For big datasets",
			"Very large (200 GB) - For very large datasets",
		},
		Default: "Normal (50 GB) - Sufficient for most projects [Recommended]",
	}
	err = survey.AskOne(storagePrompt, &storageChoice)
	if err != nil {
		return nil, err
	}
	result.EBSSize = mapStorageChoiceToGB(storageChoice)

	fmt.Println()

	// Question 4: Auto-stop behavior
	var idleChoice string
	idlePrompt := &survey.Select{
		Message: "Should your instance automatically stop when idle to save money?",
		Options: []string{
			"Yes, stop after 1 hour of inactivity [Recommended]",
			"Yes, stop after 2 hours of inactivity",
			"Yes, stop after 4 hours of inactivity",
			"No, keep it running (you'll need to stop it manually)",
		},
		Default: "Yes, stop after 1 hour of inactivity [Recommended]",
		Help:    "Instances are charged by the hour when running. Auto-stop helps prevent unexpected costs.",
	}
	err = survey.AskOne(idlePrompt, &idleChoice)
	if err != nil {
		return nil, err
	}
	result.IdleTimeout = mapIdleChoiceToTimeout(idleChoice)

	fmt.Println()

	// Question 5: Give it a name (optional)
	var instanceName string
	namePrompt := &survey.Input{
		Message: "Give your instance a name (optional, press Enter to skip):",
		Help:    "This helps you identify your instance later. For example: 'research-project' or 'thesis-analysis'",
	}
	err = survey.AskOne(namePrompt, &instanceName)
	if err != nil {
		return nil, err
	}
	result.Name = strings.TrimSpace(instanceName)

	// Show summary
	fmt.Println()
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("  Summary of your choices:")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Printf("  Environment:     %s\n", result.Environment)
	fmt.Printf("  Computer power:  %s\n", result.InstanceType)
	fmt.Printf("  Storage:         %d GB\n", result.EBSSize)
	if result.IdleTimeout != "" {
		fmt.Printf("  Auto-stop:       After %s of inactivity\n", result.IdleTimeout)
	} else {
		fmt.Printf("  Auto-stop:       Disabled (manual stop required)\n")
	}
	if result.Name != "" {
		fmt.Printf("  Name:            %s\n", result.Name)
	}
	fmt.Println()

	// Calculate estimated cost
	hourlyCost := getHourlyCost(result.InstanceType)
	storageCost := float64(result.EBSSize) * 0.0001 // Approximate EBS cost per GB-hour
	totalHourlyCost := hourlyCost + storageCost
	monthlyEstimate := totalHourlyCost * 730 // Approximate hours per month

	fmt.Printf("  Estimated cost:\n")
	fmt.Printf("    Per hour:      $%.4f\n", totalHourlyCost)
	fmt.Printf("    Per month:     $%.2f (if running 24/7)\n", monthlyEstimate)
	if result.IdleTimeout != "" {
		fmt.Printf("    Note: With auto-stop enabled, your actual cost will be much lower!\n")
	}
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println()

	// Final confirmation
	var confirm bool
	confirmPrompt := &survey.Confirm{
		Message: "Ready to launch your instance?",
		Default: true,
	}
	err = survey.AskOne(confirmPrompt, &confirm)
	if err != nil {
		return nil, err
	}

	if !confirm {
		return nil, fmt.Errorf("launch cancelled by user")
	}

	return result, nil
}

func buildEnvironmentPrompt(config WizardConfig) (*survey.Select, map[string]string) {
	var message string
	var options []string
	envMap := make(map[string]string)

	switch config.AppType {
	case "jupyter":
		message = "What type of analysis do you want to do?"
		options = []string{
			"Data science (Python + R, pandas, tidyverse)",
			"Machine learning (PyTorch, scikit-learn)",
			"Deep learning (PyTorch, TensorFlow, GPU support)",
			"Statistical analysis (R-focused, statistical packages)",
			"Just start me with the basics",
		}
		envMap[options[0]] = "data-science"
		envMap[options[1]] = "ml-pytorch"
		envMap[options[2]] = "deep-learning"
		envMap[options[3]] = "statistics"
		envMap[options[4]] = "base"

	case "rstudio":
		message = "What type of analysis do you want to do?"
		options = []string{
			"Statistical analysis (R with common packages)",
			"Data science (tidyverse, data wrangling)",
			"Bioinformatics (genomics, RNA-seq analysis)",
			"Just start me with the basics",
		}
		envMap[options[0]] = "statistics"
		envMap[options[1]] = "data-science"
		envMap[options[2]] = "bioinformatics"
		envMap[options[3]] = "base"

	case "vscode":
		message = "What type of development do you want to do?"
		options = []string{
			"Web development (HTML, CSS, JavaScript, Node.js)",
			"Python development (data analysis, scripting)",
			"Go development (systems programming)",
			"Full-stack development (Python + JavaScript)",
		}
		envMap[options[0]] = "web-dev"
		envMap[options[1]] = "python-dev"
		envMap[options[2]] = "go-dev"
		envMap[options[3]] = "fullstack"
	}

	return &survey.Select{
		Message: message,
		Options: options,
	}, envMap
}

func mapPowerChoiceToInstanceType(choice string) string {
	if strings.Contains(choice, "Light work") {
		return "t4g.small"
	} else if strings.Contains(choice, "Normal work") {
		return "t4g.medium"
	} else if strings.Contains(choice, "Heavy work") {
		return "t4g.large"
	} else {
		return "t4g.xlarge"
	}
}

func mapStorageChoiceToGB(choice string) int {
	if strings.Contains(choice, "Small (20 GB)") {
		return 20
	} else if strings.Contains(choice, "Normal (50 GB)") {
		return 50
	} else if strings.Contains(choice, "Large (100 GB)") {
		return 100
	} else {
		return 200
	}
}

func mapIdleChoiceToTimeout(choice string) string {
	if strings.Contains(choice, "1 hour") {
		return "1h"
	} else if strings.Contains(choice, "2 hours") {
		return "2h"
	} else if strings.Contains(choice, "4 hours") {
		return "4h"
	} else {
		return "" // No timeout
	}
}

func getHourlyCost(instanceType string) float64 {
	// Approximate costs for t4g instances in us-east-1
	costs := map[string]float64{
		"t4g.small":  0.0042,
		"t4g.medium": 0.0084,
		"t4g.large":  0.0168,
		"t4g.xlarge": 0.0336,
	}
	if cost, ok := costs[instanceType]; ok {
		return cost
	}
	return 0.01 // Default fallback
}
