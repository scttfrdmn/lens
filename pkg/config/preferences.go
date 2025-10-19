package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// WizardPreferences stores user's previous wizard choices
type WizardPreferences struct {
	LastEnvironment  string `json:"last_environment,omitempty"`
	LastInstanceType string `json:"last_instance_type,omitempty"`
	LastEBSSize      int    `json:"last_ebs_size,omitempty"`
	LastIdleTimeout  string `json:"last_idle_timeout,omitempty"`
	LastRegion       string `json:"last_region,omitempty"`
}

var (
	prefsCache = make(map[string]*WizardPreferences)
	prefsMutex sync.RWMutex
)

// GetWizardPreferences loads saved preferences for a specific app
func GetWizardPreferences(appName string) (*WizardPreferences, error) {
	prefsMutex.RLock()
	if prefs, ok := prefsCache[appName]; ok {
		prefsMutex.RUnlock()
		return prefs, nil
	}
	prefsMutex.RUnlock()

	prefsPath := getPreferencesPath(appName)

	// If file doesn't exist, return empty preferences
	if _, err := os.Stat(prefsPath); os.IsNotExist(err) {
		return &WizardPreferences{}, nil
	}

	data, err := os.ReadFile(prefsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read preferences: %w", err)
	}

	var prefs WizardPreferences
	if err := json.Unmarshal(data, &prefs); err != nil {
		return nil, fmt.Errorf("failed to parse preferences: %w", err)
	}

	// Cache it
	prefsMutex.Lock()
	prefsCache[appName] = &prefs
	prefsMutex.Unlock()

	return &prefs, nil
}

// SaveWizardPreferences saves user's wizard choices for future use
func SaveWizardPreferences(appName string, prefs *WizardPreferences) error {
	prefsPath := getPreferencesPath(appName)

	// Ensure directory exists
	if err := EnsureConfigDir(); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(prefs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal preferences: %w", err)
	}

	if err := os.WriteFile(prefsPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write preferences: %w", err)
	}

	// Update cache
	prefsMutex.Lock()
	prefsCache[appName] = prefs
	prefsMutex.Unlock()

	return nil
}

// ClearWizardPreferences removes saved preferences for an app
func ClearWizardPreferences(appName string) error {
	prefsPath := getPreferencesPath(appName)

	// Clear from cache
	prefsMutex.Lock()
	delete(prefsCache, appName)
	prefsMutex.Unlock()

	// Remove file
	if err := os.Remove(prefsPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove preferences: %w", err)
	}

	return nil
}

// getPreferencesPath returns the path to the preferences file for an app
func getPreferencesPath(appName string) string {
	configDir := GetConfigDir()
	return filepath.Join(configDir, fmt.Sprintf("%s-preferences.json", appName))
}

// UpdateWizardPreference updates a single preference field
func (p *WizardPreferences) UpdateEnvironment(env string) {
	p.LastEnvironment = env
}

func (p *WizardPreferences) UpdateInstanceType(instanceType string) {
	p.LastInstanceType = instanceType
}

func (p *WizardPreferences) UpdateEBSSize(size int) {
	p.LastEBSSize = size
}

func (p *WizardPreferences) UpdateIdleTimeout(timeout string) {
	p.LastIdleTimeout = timeout
}

func (p *WizardPreferences) UpdateRegion(region string) {
	p.LastRegion = region
}

// HasPreferences returns true if any preferences have been saved
func (p *WizardPreferences) HasPreferences() bool {
	return p.LastEnvironment != "" ||
		p.LastInstanceType != "" ||
		p.LastEBSSize > 0 ||
		p.LastIdleTimeout != "" ||
		p.LastRegion != ""
}

// GetDefault returns a default value or the saved preference
func (p *WizardPreferences) GetEnvironmentOrDefault(defaultVal string) string {
	if p.LastEnvironment != "" {
		return p.LastEnvironment
	}
	return defaultVal
}

func (p *WizardPreferences) GetInstanceTypeOrDefault(defaultVal string) string {
	if p.LastInstanceType != "" {
		return p.LastInstanceType
	}
	return defaultVal
}

func (p *WizardPreferences) GetEBSSizeOrDefault(defaultVal int) int {
	if p.LastEBSSize > 0 {
		return p.LastEBSSize
	}
	return defaultVal
}

func (p *WizardPreferences) GetIdleTimeoutOrDefault(defaultVal string) string {
	if p.LastIdleTimeout != "" {
		return p.LastIdleTimeout
	}
	return defaultVal
}

func (p *WizardPreferences) GetRegionOrDefault(defaultVal string) string {
	if p.LastRegion != "" {
		return p.LastRegion
	}
	return defaultVal
}
