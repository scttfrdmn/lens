# AWS IDE Documentation

Welcome to AWS IDE - a suite of command-line tools for launching secure, cloud-based development environments for academic research on AWS.

## What is AWS IDE?

AWS IDE provides three powerful applications for researchers:

- **[aws-jupyter](user-guides/jupyter.md)**: Launch Jupyter Lab instances for data science and machine learning
- **[aws-rstudio](user-guides/rstudio.md)**: Launch RStudio Server for statistical computing and R development  
- **[aws-vscode](user-guides/vscode.md)**: Launch VSCode Server for general-purpose development

All tools share robust infrastructure for AWS integration, security, cost optimization, and ease of use.

## Quick Start

Get started in under 5 minutes:

```bash
# Install (choose your platform)
# macOS
brew install scttfrdmn/tap/aws-jupyter

# Launch with interactive wizard
aws-jupyter
```

## Key Features

### 🎯 **Beginner-Friendly**
- Interactive wizard for guided setup
- Plain-English error messages
- No AWS expertise required

### 💰 **Cost-Optimized**
- Automatic instance shutdown when idle
- Cost estimates before launch
- Spot instance support for up to 90% savings

### 🔒 **Secure by Default**
- AWS Session Manager (no SSH keys needed)
- Private subnet support with NAT Gateway
- IAM role-based access

### 🚀 **Production-Ready**
- Pre-configured research environments
- GPU support for ML workloads
- S3 data sync for persistence

## Who Is This For?

AWS IDE is designed for:

- **Solo Researchers**: Quick access to cloud compute for data analysis
- **Research Labs**: Reproducible environments for team collaboration
- **Data Scientists**: GPU-accelerated ML training and experimentation
- **Graduate Students**: Cost-effective compute for thesis work
- **Course Instructors**: Consistent environments for teaching

## Documentation

<div class="grid cards" markdown>

-   :material-rocket-launch: **[Getting Started](getting-started/installation.md)**

    ---
    
    Install AWS IDE and launch your first instance

-   :material-book-open-variant: **[User Guides](user-guides/jupyter.md)**

    ---
    
    Detailed guides for each application

-   :material-puzzle: **[Architecture](architecture/overview.md)**

    ---
    
    Technical architecture and design

-   :material-code-braces: **[Development](development/contributing.md)**

    ---
    
    Contributing to AWS IDE

</div>

## Get Help

- **Questions?** Check out [GitHub Discussions](https://github.com/scttfrdmn/aws-ide/discussions)
- **Found a bug?** [Report an issue](https://github.com/scttfrdmn/aws-ide/issues/new/choose)
- **Need a feature?** [Request it here](https://github.com/scttfrdmn/aws-ide/issues/new/choose)

## Version

Current version: **v0.7.2** (Platform: v1.0.0)

See the [CHANGELOG](https://github.com/scttfrdmn/aws-ide/blob/main/CHANGELOG.md) for release notes.
