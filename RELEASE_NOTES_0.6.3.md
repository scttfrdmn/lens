# Release Notes: v0.6.3

**Release Date:** 2025-10-18

## Overview

v0.6.3 is a documentation and polish release that updates all version references, adds comprehensive documentation for the config and costs commands introduced in v0.6.1 and v0.6.2, and modernizes the project roadmap.

## What's Changed

### Documentation Improvements

#### README.md Updates
- **Version References**: Updated all installation examples from v0.5.1 to v0.6.2
- **Status Updates**: Changed aws-vscode from "beta" to "production" status
- **New Feature Documentation**: Added comprehensive sections for Configuration Management and Cost Tracking commands
- **Updated Roadmap**: Reflects completed work through v0.6.2 and clarifies future direction

#### Configuration Management Section
Added complete documentation for the `config` command:

```bash
# Initialize config file with defaults
aws-jupyter config init

# View current configuration
aws-jupyter config show

# Set defaults
aws-jupyter config set default_region us-west-2
aws-jupyter config set default_instance_type t4g.large

# Enable cost tracking
aws-vscode config set enable_cost_tracking true
aws-vscode config set cost_alert_threshold 100

# App-specific settings
aws-vscode config set vscode.port 8080
```

#### Cost Tracking Section
Added complete documentation for the `costs` command with example output:

```bash
# View costs for all instances
aws-jupyter costs

# Detailed breakdown for specific instance
aws-jupyter costs i-1234567890abcdef0

# Show detailed breakdowns for all instances
aws-rstudio costs --details
```

**Key Metrics Documented:**
- Effective Cost (total cost / elapsed hours)
- Utilization percentage
- Monthly cost estimates
- 24/7 cost comparison

#### Key Features Section Enhancements
Added two new feature categories:

**💰 Cost Optimization (Enhanced)**:
- Cost Tracking with effective cost per hour
- Monthly cost estimates based on usage patterns

**🔧 Configuration & Management (New)**:
- Unified config file (`~/.aws-ide/config.yaml`)
- Flexible settings for all tools
- Cost alerts with monthly thresholds
- Per-tool overrides

#### Versioning Section
Updated current release information:
- **Current Release**: v0.5.1 → v0.6.2
- **aws-vscode status**: beta → production
- All three tools now at v0.6.2 (production)

#### Roadmap Section
Complete rewrite to reflect current state:

**Completed**:
- v0.1.0-v0.5.0: Core features, Session Manager, cost optimization, monorepo
- v0.6.0: Testing & quality, VSCode feature parity
- v0.6.1: Unified config system, cost tracking
- v0.6.2: Full feature parity across all tools

**Up Next**:
- v0.7.0: Security hardening, audit logging
- v0.8.0: Package manager integration
- v0.9.0: Advanced networking
- v1.0.0: Production-ready release

### Quality Assurance

- ✅ All three applications build successfully
- ✅ All tests pass (vscode, jupyter, rstudio)
- ✅ No TODO/FIXME markers in codebase
- ✅ Documentation is current and accurate

## Installation

### Updated Package Installation Examples

**APT (Debian/Ubuntu)**
```bash
wget https://github.com/scttfrdmn/aws-ide/releases/latest/download/aws-jupyter_0.6.2_linux_amd64.deb
wget https://github.com/scttfrdmn/aws-ide/releases/latest/download/aws-rstudio_0.6.2_linux_amd64.deb
wget https://github.com/scttfrdmn/aws-ide/releases/latest/download/aws-vscode_0.6.2_linux_amd64.deb

sudo dpkg -i aws-jupyter_0.6.2_linux_amd64.deb
sudo dpkg -i aws-rstudio_0.6.2_linux_amd64.deb
sudo dpkg -i aws-vscode_0.6.2_linux_amd64.deb
```

**YUM/DNF (RedHat/Fedora/Amazon Linux)**
```bash
wget https://github.com/scttfrdmn/aws-ide/releases/latest/download/aws-jupyter_0.6.2_linux_amd64.rpm
wget https://github.com/scttfrdmn/aws-ide/releases/latest/download/aws-rstudio_0.6.2_linux_amd64.rpm
wget https://github.com/scttfrdmn/aws-ide/releases/latest/download/aws-vscode_0.6.2_linux_amd64.rpm

sudo rpm -i aws-jupyter_0.6.2_linux_amd64.rpm
sudo rpm -i aws-rstudio_0.6.2_linux_amd64.rpm
sudo rpm -i aws-vscode_0.6.2_linux_amd64.rpm
```

## Files Changed

- `README.md`: Major documentation updates (58 lines changed)
  - Version references updated throughout
  - Added Configuration Management section
  - Added Cost Tracking section with examples
  - Enhanced Key Features section
  - Modernized Roadmap section
  - Updated Versioning section

## Benefits

### For New Users
- **Clearer Documentation**: Updated README accurately reflects current capabilities
- **Better Onboarding**: Comprehensive examples for config and costs commands
- **Current Information**: All version references and installation examples are up-to-date

### For Existing Users
- **Feature Discovery**: May discover config and costs commands they didn't know about
- **Better Understanding**: See the full roadmap and where the project is heading
- **Accurate Examples**: Can copy-paste current installation commands

### For Contributors
- **Current Roadmap**: Clear understanding of what's completed and what's next
- **Documentation Standards**: See how new features should be documented

## Migration Guide

No code changes - this is purely a documentation update. No migration needed.

## Upgrade Instructions

No binary changes in v0.6.3. This release is documentation-only. Continue using v0.6.2 binaries.

If rebuilding from source:
```bash
git pull origin main
cd apps/vscode && go build ./cmd/aws-vscode
cd ../jupyter && go build ./cmd/aws-jupyter
cd ../rstudio && go build ./cmd/aws-rstudio
```

## What's Next

With documentation current and comprehensive, v0.7.0 will focus on:
- Security hardening and audit logging
- IAM role improvements
- Compliance reporting
- Session recording capabilities

## Compatibility

- **Go Version**: 1.22+
- **AWS CLI**: Required for credentials
- **Platforms**: macOS (ARM64/x86_64), Linux (ARM64/x86_64)
- **Config File**: `~/.aws-ide/config.yaml` (shared across all tools)
- **State File**: `~/.aws-ide/state.json` (shared across all tools)

## Contributors

- Scott Friedman (@scttfrdmn)

## Links

- [GitHub Repository](https://github.com/scttfrdmn/aws-ide)
- [v0.6.3 Release](https://github.com/scttfrdmn/aws-ide/releases/tag/v0.6.3)
- [Documentation](https://github.com/scttfrdmn/aws-ide/blob/main/README.md)
- [Roadmap](https://github.com/scttfrdmn/aws-ide/blob/main/ROADMAP.md)
