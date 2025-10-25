# Release Notes: v0.6.2

**Release Date:** 2025-10-18

## Overview

v0.6.2 extends the config and costs commands from lens-vscode to lens-jupyter and lens-rstudio, achieving full feature parity across all three CLI tools in the lens suite.

## What's New

### Feature Parity: Jupyter & RStudio

All three CLI tools (lens-vscode, lens-jupyter, lens-rstudio) now support:

#### Config Command (`config`)
Unified configuration management across all tools:

```bash
# Initialize config with defaults
lens-jupyter config init
lens-rstudio config init

# View current configuration
lens-jupyter config show
lens-rstudio config show

# Set configuration values
lens-jupyter config set default_region us-west-2
lens-rstudio config set default_instance_type t4g.large

# Get specific values
lens-jupyter config get default_region
lens-rstudio config get cost_alert_threshold
```

**Configuration Options:**
- AWS settings: `default_region`, `default_profile`
- Instance defaults: `default_instance_type`, `default_ebs_size`, `default_ami_base`
- Networking: `default_subnet_type`, `prefer_ipv6`
- Behavior: `idle_timeout`, `auto_terminate`, `confirm_destructive`
- Cost tracking: `enable_cost_tracking`, `cost_alert_threshold`
- App-specific settings: `jupyter.*`, `rstudio.*`, `vscode.*`

All tools share the same config file at `~/.lens/config.yaml`.

#### Costs Command (`costs`)
Full cost tracking and analysis:

```bash
# View costs for all instances
lens-jupyter costs
lens-rstudio costs

# Detailed cost breakdown for specific instance
lens-jupyter costs i-1234567890abcdef0
lens-rstudio costs i-1234567890abcdef0

# Show detailed breakdowns
lens-jupyter costs --details
lens-rstudio costs --details
```

**Cost Metrics:**
- **Total Cost**: Compute + storage costs
- **Running Hours**: Actual time instance was running
- **Elapsed Hours**: Total time since launch
- **Utilization**: Running hours / elapsed hours
- **Effective Rate**: True cost per hour (shows stop/start savings)
- **Monthly Estimates**: Based on current usage patterns
- **Savings vs 24/7**: How much you save by stopping instances

**Example Output:**
```
Instance: i-abc123 (ml-gpu)
  Type: t4g.large
  Running: 12.5h / 48.0h (26% utilization)
  Total Cost: $1.23
  Effective Rate: $0.026/hour

  Savings vs 24/7: $0.073/hour (74%)
```

## Implementation Details

### Commands Added

**lens-jupyter:**
- `lens-jupyter config` with subcommands: `init`, `show`, `set`, `get`
- `lens-jupyter costs [INSTANCE_ID]` with `--details` flag

**lens-rstudio:**
- `lens-rstudio config` with subcommands: `init`, `show`, `set`, `get`
- `lens-rstudio costs [INSTANCE_ID]` with `--details` flag

### Files Changed

**Jupyter:**
- `apps/jupyter/cmd/lens-jupyter/main.go` - Added config and costs command registration
- `apps/jupyter/internal/cli/config.go` - Config command implementation (387 lines)
- `apps/jupyter/internal/cli/costs.go` - Costs command implementation (296 lines)

**RStudio:**
- `apps/rstudio/cmd/lens-rstudio/main.go` - Added config and costs command registration
- `apps/rstudio/internal/cli/config.go` - Config command implementation (387 lines)
- `apps/rstudio/internal/cli/costs.go` - Costs command implementation (296 lines)

**Version Updates:**
- `apps/vscode/cmd/lens-vscode/main.go` - Version: 0.6.1 → 0.6.2
- `apps/jupyter/cmd/lens-jupyter/main.go` - Version: 0.6.1 → 0.6.2
- `apps/rstudio/cmd/lens-rstudio/main.go` - Version: 0.6.1 → 0.6.2

### Shared Infrastructure

All commands use the shared `pkg/` infrastructure:
- `pkg/config/userconfig.go` - Config file management
- `pkg/cost/calculator.go` - Cost calculation engine
- `pkg/errors/errors.go` - Contextual error handling

## Benefits

### For Users

1. **Consistent Experience**: Same commands work across all three tools
2. **Unified Configuration**: One config file for all lens tools
3. **Cost Visibility**: Track spending across Jupyter, RStudio, and VSCode instances
4. **Cost Optimization**: See exactly how much you save with stop/start cycles

### For Developers

1. **Code Reuse**: CLI wrappers share the same `pkg/` implementation
2. **Easy Maintenance**: Changes to shared code benefit all tools
3. **Feature Parity**: No gaps between tool capabilities

## Migration Guide

No migration needed! If you're already using v0.6.1 config or costs features with lens-vscode, they'll continue to work. The same config file and state tracking work for all tools.

## Upgrade Instructions

### From Source

```bash
# Pull latest changes
git pull origin main

# Build all tools
cd apps/vscode && go build ./cmd/lens-vscode && cd ../..
cd apps/jupyter && go build ./cmd/lens-jupyter && cd ../..
cd apps/rstudio && go build ./cmd/lens-rstudio && cd ../..
```

### Using Go Install

```bash
go install github.com/scttfrdmn/lens/apps/vscode/cmd/lens-vscode@v0.6.2
go install github.com/scttfrdmn/lens/apps/jupyter/cmd/lens-jupyter@v0.6.2
go install github.com/scttfrdmn/lens/apps/rstudio/cmd/lens-rstudio@v0.6.2
```

## Compatibility

- **Go Version**: 1.22+
- **AWS CLI**: Required for credentials
- **Platforms**: macOS (ARM64/x86_64), Linux (ARM64/x86_64)
- **Config File**: `~/.lens/config.yaml` (shared across all tools)
- **State File**: `~/.lens/state.json` (shared across all tools)

## What's Next

With full feature parity achieved, upcoming releases will focus on:
- Security hardening and audit logging (v0.7.0)
- Package manager support (v0.8.0)
- Advanced networking features (v0.9.0)
- Production readiness for v1.0.0

## Contributors

- Scott Friedman (@scttfrdmn)

## Links

- [GitHub Repository](https://github.com/scttfrdmn/lens)
- [v0.6.2 Release](https://github.com/scttfrdmn/lens/releases/tag/v0.6.2)
- [Documentation](https://github.com/scttfrdmn/lens/blob/main/README.md)
- [Roadmap](https://github.com/scttfrdmn/lens/blob/main/ROADMAP.md)
