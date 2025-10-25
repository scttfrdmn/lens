# Lens Documentation

Welcome to Lens - a suite of command-line tools for launching secure, cloud-based development environments for academic research on AWS.

## What is Lens?

Lens provides three powerful applications for researchers:

- **[lens-jupyter](user-guides/jupyter.md)**: Launch Jupyter Lab instances for data science and machine learning
- **[lens-rstudio](user-guides/rstudio.md)**: Launch RStudio Server for statistical computing and R development  
- **[lens-vscode](user-guides/vscode.md)**: Launch VSCode Server for general-purpose development

All tools share robust infrastructure for AWS integration, security, cost optimization, and ease of use.

## Quick Start

Get started in under 5 minutes:

```bash
# Install (choose your platform)
# macOS
brew install scttfrdmn/tap/lens-jupyter

# Launch with interactive wizard
lens-jupyter
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

Lens is designed for academic researchers at all levels:

<div class="grid cards" markdown>

-   :material-account: **[Solo Researcher](USER_SCENARIOS/01_SOLO_RESEARCHER_WALKTHROUGH.md)**

    ---

    Individual researchers with limited budgets needing simple, reliable cloud access

    **Key benefits**: 96% faster setup, 98% cost reduction, zero AWS expertise needed

-   :material-school: **[Graduate Student](USER_SCENARIOS/02_GRADUATE_STUDENT_WALKTHROUGH.md)**

    ---

    PhD students needing GPU access with strict budget constraints and advisor oversight

    **Key benefits**: 3x research productivity, 67% under budget, advisor trust restored

-   :material-account-group: **[Lab PI](USER_SCENARIOS/03_LAB_PI_WALKTHROUGH.md)**

    ---

    Principal investigators managing 5-10 researchers with limited time and budget

    **Key benefits**: 85% budget savings, 247 hours/year saved, 100% reproducibility

-   :material-teach: **[Course Instructor](USER_SCENARIOS/04_COURSE_INSTRUCTOR_WALKTHROUGH.md)**

    ---

    Professors teaching 30+ students needing consistent, reproducible environments

    **Key benefits**: 96% setup time reduction, 87% under budget, zero "works on my machine"

-   :material-server: **[Research Computing Manager](USER_SCENARIOS/05_RESEARCH_COMPUTING_MANAGER_WALKTHROUGH.md)**

    ---

    IT managers supporting 100+ researchers across multiple disciplines

    **Key benefits**: 83% support reduction, 100% cost visibility, 97% security compliance

</div>

## Documentation

<div class="grid cards" markdown>

-   :material-rocket-launch: **[Getting Started](getting-started/installation.md)**

    ---
    
    Install Lens and launch your first instance

-   :material-book-open-variant: **[User Guides](user-guides/jupyter.md)**

    ---
    
    Detailed guides for each application

-   :material-puzzle: **[Architecture](architecture/overview.md)**

    ---
    
    Technical architecture and design

-   :material-code-braces: **[Development](development/contributing.md)**

    ---
    
    Contributing to Lens

</div>

## Get Help

- **Questions?** Check out [GitHub Discussions](https://github.com/scttfrdmn/lens/discussions)
- **Found a bug?** [Report an issue](https://github.com/scttfrdmn/lens/issues/new/choose)
- **Need a feature?** [Request it here](https://github.com/scttfrdmn/lens/issues/new/choose)

## Version

Current version: **v0.7.2** (Platform: v1.0.0)

See the [CHANGELOG](https://github.com/scttfrdmn/lens/blob/main/CHANGELOG.md) for release notes.
