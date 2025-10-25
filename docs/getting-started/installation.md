# Installation

## Prerequisites

Before installing Lens, ensure you have:

- **AWS Account**: An active AWS account with appropriate permissions
- **AWS CLI**: Configured with valid credentials (`aws configure`)
- **macOS, Linux, or Windows**: Supported on all major platforms

## Installation Methods

### macOS (Homebrew)

```bash
# Add tap (first time only)
brew tap scttfrdmn/tap

# Install
brew install lens-jupyter  # or lens-rstudio, lens-vscode
```

### Linux

#### Debian/Ubuntu (.deb)

```bash
# Download from releases page
wget https://github.com/scttfrdmn/lens/releases/download/v0.7.2/lens-jupyter_0.7.2_linux_amd64.deb

# Install
sudo dpkg -i lens-jupyter_0.7.2_linux_amd64.deb
```

#### RHEL/CentOS/Fedora (.rpm)

```bash
# Download from releases page
wget https://github.com/scttfrdmn/lens/releases/download/v0.7.2/lens-jupyter_0.7.2_linux_amd64.rpm

# Install
sudo rpm -i lens-jupyter_0.7.2_linux_amd64.rpm
```

#### From Binary

```bash
# Download
wget https://github.com/scttfrdmn/lens/releases/download/v0.7.2/lens-jupyter_Linux_x86_64.tar.gz

# Extract
tar -xzf lens-jupyter_Linux_x86_64.tar.gz

# Move to PATH
sudo mv lens-jupyter /usr/local/bin/
```

### Windows

#### Download Binary

1. Download from [releases page](https://github.com/scttfrdmn/lens/releases)
2. Extract the `.zip` file
3. Add the directory to your PATH

#### Windows Subsystem for Linux (WSL)

Use the Linux installation instructions within WSL.

## Verify Installation

```bash
lens-jupyter --version
# Output: v0.7.2 (platform: v1.0.0, commit: abc1234, date: 2025-10-19)
```

## AWS Configuration

Configure AWS credentials if you haven't already:

```bash
aws configure
# AWS Access Key ID [None]: YOUR_ACCESS_KEY
# AWS Secret Access Key [None]: YOUR_SECRET_KEY
# Default region name [None]: us-east-1
# Default output format [None]: json
```

## Next Steps

- [Launch your first instance](../user-guides/jupyter.md)
- [Configure environments](../user-guides/environments.md)
- [Understand costs](../user-guides/cost-management.md)
