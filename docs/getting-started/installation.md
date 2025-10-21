# Installation

## Prerequisites

Before installing AWS IDE, ensure you have:

- **AWS Account**: An active AWS account with appropriate permissions
- **AWS CLI**: Configured with valid credentials (`aws configure`)
- **macOS, Linux, or Windows**: Supported on all major platforms

## Installation Methods

### macOS (Homebrew)

```bash
# Add tap (first time only)
brew tap scttfrdmn/tap

# Install
brew install aws-jupyter  # or aws-rstudio, aws-vscode
```

### Linux

#### Debian/Ubuntu (.deb)

```bash
# Download from releases page
wget https://github.com/scttfrdmn/aws-ide/releases/download/v0.7.2/aws-jupyter_0.7.2_linux_amd64.deb

# Install
sudo dpkg -i aws-jupyter_0.7.2_linux_amd64.deb
```

#### RHEL/CentOS/Fedora (.rpm)

```bash
# Download from releases page
wget https://github.com/scttfrdmn/aws-ide/releases/download/v0.7.2/aws-jupyter_0.7.2_linux_amd64.rpm

# Install
sudo rpm -i aws-jupyter_0.7.2_linux_amd64.rpm
```

#### From Binary

```bash
# Download
wget https://github.com/scttfrdmn/aws-ide/releases/download/v0.7.2/aws-jupyter_Linux_x86_64.tar.gz

# Extract
tar -xzf aws-jupyter_Linux_x86_64.tar.gz

# Move to PATH
sudo mv aws-jupyter /usr/local/bin/
```

### Windows

#### Download Binary

1. Download from [releases page](https://github.com/scttfrdmn/aws-ide/releases)
2. Extract the `.zip` file
3. Add the directory to your PATH

#### Windows Subsystem for Linux (WSL)

Use the Linux installation instructions within WSL.

## Verify Installation

```bash
aws-jupyter --version
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
