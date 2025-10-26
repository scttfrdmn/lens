package dcv

// Config holds DCV-specific configuration
type Config struct {
	Port            int    // DCV server port (default 8443)
	SessionName     string // DCV session name
	SessionType     string // virtual or console
	Quality         string // high, medium, low
	EnableGPU       bool   // Enable GPU acceleration
	EnableUSB       bool   // Enable USB redirection
	EnableClipboard bool   // Enable clipboard sharing
	Owner           string // Session owner (default: ubuntu)
}

// DefaultConfig returns a sensible default DCV configuration
func DefaultConfig() *Config {
	return &Config{
		Port:            8443,
		SessionName:     "lens-session",
		SessionType:     "virtual",
		Quality:         "high",
		EnableGPU:       false,
		EnableUSB:       false,
		EnableClipboard: true,
		Owner:           "ubuntu",
	}
}

// DefaultGPUConfig returns a DCV configuration optimized for GPU workloads
func DefaultGPUConfig() *Config {
	cfg := DefaultConfig()
	cfg.EnableGPU = true
	return cfg
}
