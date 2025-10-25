# Release Notes - v0.6.1

**Release Date:** October 18, 2025

## Overview

Version 0.6.1 is a patch release that brings VSCode to feature parity with Jupyter and RStudio, adds a unified configuration system, implements cost tracking with effective cost calculations, and enhances error handling across all apps.

## What's New

### VSCode Feature Parity

VSCode now has all the commands available in Jupyter and RStudio:

- **`lens-vscode create-ami`** - Create custom AMIs from running instances
- **`lens-vscode delete-ami`** - Delete custom AMIs with safety confirmations
- **`lens-vscode list-amis`** - List all custom VSCode AMIs
- **`lens-vscode generate`** - Auto-generate environment configs from Node.js projects

The `generate` command analyzes your Node.js projects and creates optimized environment configurations:
- Detects frameworks (React, Vue, Next.js, Angular, Express, NestJS, etc.)
- Identifies project type (frontend, backend, fullstack)
- Detects TypeScript usage
- Identifies monorepo setups
- Suggests appropriate VSCode extensions
- Configures package managers (npm, yarn, pnpm)

### Unified Configuration System

All three apps (Jupyter, RStudio, VSCode) now share a unified configuration file at `~/.lens/config.yaml`:

**New `config` command with subcommands:**
- `aws-* config init` - Initialize config file with defaults
- `aws-* config show` - Display current configuration
- `aws-* config set <key> <value>` - Set configuration values
- `aws-* config get <key>` - Get specific configuration value

**Supported settings:**
- **AWS defaults:** `default_region`, `default_profile`
- **Instance defaults:** `default_instance_type`, `default_ebs_size`, `default_ami_base`
- **Network defaults:** `default_subnet_type`, `prefer_ipv6`
- **Behavior:** `idle_timeout`, `auto_terminate`, `confirm_destructive`
- **Cost tracking:** `enable_cost_tracking`, `cost_alert_threshold`
- **App-specific:** Port numbers, default environments per app

### Cost Tracking

New comprehensive cost tracking system that shows the true savings of cloud infrastructure:

**New `costs` command:**
- `aws-* costs` - Show summary for all instances
- `aws-* costs INSTANCE_ID` - Detailed breakdown for specific instance
- `aws-* costs --details` - Show detailed breakdown for all instances

**Key metrics:**
- **Running hours** - Actual compute time
- **Elapsed hours** - Total time since launch (includes stopped time)
- **Effective cost per hour** - Shows true cost including stop/start cycles
- **Savings vs 24/7** - Demonstrates cloud cost advantages
- **Monthly estimates** - Based on current usage patterns
- **On-premise comparison** - Compares to equivalent hardware costs

**Features:**
- Tracks state changes (running, stopped, terminated)
- Calculates separate compute and storage costs
- Supports cost alerts when monthly estimates exceed threshold
- Can be disabled via config: `enable_cost_tracking: false`

### Enhanced Error Handling

All apps now use contextual error messages with actionable suggestions:

**Error types with helpful guidance:**
- AWS permission errors
- Resource not found errors
- Network/connectivity issues
- Configuration file issues
- Validation errors
- Quota exceeded errors
- Session Manager issues

Each error provides:
- Clear description of what went wrong
- Context about the operation
- Specific suggestions to resolve the issue

## Bug Fixes

- Fixed LocalStack CI initialization issues
- Fixed Go formatting in new files
- Fixed unnecessary `fmt.Sprintf` usage in error messages
- Fixed unchecked errors from `os.Setenv` and `os.Unsetenv` in tests
- Made integration tests non-blocking to handle LocalStack flakiness

## CI/CD Improvements

- Resolved LocalStack Docker socket issues
- Improved health checks for service containers
- Made integration tests continue-on-error for reliability
- All linters passing (gofmt, go vet, golangci-lint)
- All tests passing (8/8 matrix combinations)

## Technical Details

### Files Added (10)
- `apps/vscode/internal/cli/create-ami.go` (88 lines)
- `apps/vscode/internal/cli/delete-ami.go` (180 lines)
- `apps/vscode/internal/cli/list-amis.go` (82 lines)
- `apps/vscode/internal/cli/generate.go` (428 lines)
- `apps/vscode/internal/cli/generate_test.go` (381 lines)
- `apps/vscode/internal/cli/config.go` (387 lines)
- `apps/vscode/internal/cli/costs.go` (296 lines)
- `pkg/config/userconfig.go` (252 lines)
- `pkg/cost/calculator.go` (302 lines)
- `pkg/errors/errors.go` (258 lines)

### Files Modified (2)
- `apps/vscode/cmd/lens-vscode/main.go` - Registered 6 new commands
- `pkg/config/state.go` - Added StateChange tracking

**Total:** 2,597 lines of production code added

### Pricing Data
The cost calculator includes current pricing for:
- Graviton instances: t4g, c7g, m7g, r7g families
- x86 instances: t3, m6i families
- EBS storage: gp3 volumes at $0.08/GB-month
- Prices based on us-east-1 on-demand rates (as of 2025)

## Migration Notes

No breaking changes. All existing functionality remains unchanged.

**To adopt new features:**

1. **Initialize config:**
   ```bash
   lens-vscode config init
   lens-jupyter config init
   lens-rstudio config init
   ```

2. **Enable cost tracking (if desired):**
   ```bash
   lens-vscode config set enable_cost_tracking true
   lens-vscode config set cost_alert_threshold 100.0
   ```

3. **Check costs:**
   ```bash
   lens-vscode costs
   lens-vscode costs i-1234567890abcdef0
   ```

4. **Generate VSCode config from project:**
   ```bash
   cd /path/to/nodejs/project
   lens-vscode generate
   ```

## What's Next

The next release will focus on:
- Applying config and cost tracking features to Jupyter and RStudio
- Enhanced environment template system
- Improved S3 data sync capabilities
- Additional cost optimization recommendations

## Downloads

Binaries will be available via GitHub releases and GoReleaser.

## Contributors

- @scttfrdmn

---

For detailed commit history, see: https://github.com/scttfrdmn/lens/compare/v0.6.0...v0.6.1
