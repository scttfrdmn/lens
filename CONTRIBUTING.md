# Contributing to AWS IDE

Thank you for your interest in contributing to AWS IDE! This document provides guidelines and instructions for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Development Workflow](#development-workflow)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Documentation](#documentation)
- [Community](#community)

## Code of Conduct

This project follows a professional and inclusive code of conduct. We expect all contributors to:

- Be respectful and welcoming to all participants
- Focus on constructive feedback and collaboration
- Respect differing viewpoints and experiences
- Accept responsibility and learn from mistakes
- Focus on what is best for the community and research users

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21 or later** - [Install Go](https://golang.org/doc/install)
- **Make** - For build automation
- **Git** - For version control
- **AWS CLI** - For AWS integration testing (optional)
- **golangci-lint** - For code linting: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/aws-ide.git
   cd aws-ide
   ```
3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/scttfrdmn/aws-ide.git
   ```

## Development Setup

### Building the Project

The project uses a monorepo structure with three apps and shared infrastructure:

```bash
# Build all apps
make build

# Build specific app
cd apps/jupyter && go build -o ../../bin/aws-jupyter cmd/aws-jupyter/main.go
cd apps/rstudio && go build -o ../../bin/aws-rstudio cmd/aws-rstudio/main.go
cd apps/vscode && go build -o ../../bin/aws-vscode cmd/aws-vscode/main.go
```

### Running Tests

```bash
# Run all tests
make test

# Run unit tests only
make test-unit

# Run integration tests (requires LocalStack)
make test-integration

# Run linting
make lint
```

### Project Structure

```
aws-ide/
├── apps/
│   ├── jupyter/       # Jupyter Lab launcher
│   ├── rstudio/       # RStudio Server launcher
│   └── vscode/        # VSCode Server launcher
├── pkg/               # Shared infrastructure library
│   ├── aws/          # AWS integrations
│   ├── cli/          # CLI utilities
│   ├── config/       # Configuration management
│   ├── errors/       # Error handling
│   ├── output/       # Terminal output formatting
│   ├── readiness/    # Service health checks
│   └── version.go    # Platform version
├── docs/              # Documentation
└── .github/           # GitHub workflows and templates
```

## Development Workflow

### Creating a Branch

1. Ensure your main branch is up-to-date:
   ```bash
   git checkout main
   git pull upstream main
   ```

2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix-name
   ```

### Making Changes

1. Make your changes in logical, focused commits
2. Write clear commit messages following [Conventional Commits](https://www.conventionalcommits.org/):
   ```
   feat: add support for custom AMIs
   fix: resolve session manager connection timeout
   docs: update launch command documentation
   test: add unit tests for config module
   refactor: simplify error handling in pkg/aws
   chore: bump version to 0.7.3
   ```

3. Keep commits small and focused on a single change
4. Test your changes locally before pushing

### Commit Message Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `chore`: Build process or auxiliary tool changes

**Examples:**
```
feat(jupyter): add GPU instance type selection

Add support for selecting GPU instance types (p3, g4dn, etc.)
in the interactive wizard for Jupyter notebooks.

Closes #123
```

## Pull Request Process

### Before Submitting

1. **Run tests and linting:**
   ```bash
   make test
   make lint
   ```

2. **Update documentation:**
   - Add/update relevant documentation in `docs/`
   - Update CHANGELOG.md with your changes
   - Update README if adding new features

3. **Self-review your code:**
   - Remove debug statements
   - Check for proper error handling
   - Ensure code follows project conventions
   - Add comments for complex logic

### Submitting a Pull Request

1. Push your changes to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

2. Open a Pull Request on GitHub

3. Fill out the PR template completely:
   - Describe your changes
   - Link related issues
   - Check all applicable boxes
   - Add screenshots/demos if relevant

4. Wait for review and address feedback

### PR Review Process

- Maintainers will review your PR within 3-5 business days
- Address review feedback promptly
- Keep the PR focused and avoid scope creep
- Be responsive to questions and suggestions
- Once approved, a maintainer will merge your PR

## Coding Standards

### Go Style Guidelines

- Follow standard Go conventions and idioms
- Run `gofmt` and `goimports` on your code
- Use meaningful variable and function names
- Write clear, self-documenting code
- Add comments for exported functions and complex logic
- Keep functions small and focused
- Handle errors explicitly, don't ignore them

### Error Handling

Use the `pkg/errors` package for user-friendly error messages:

```go
import "github.com/scttfrdmn/aws-ide/pkg/errors"

// Create friendly error
err := errors.NewFriendlyError(
    "Cannot connect to AWS",
    "Your AWS credentials are not configured or have expired.",
    []string{
        "Run 'aws configure' to set up your credentials",
        "Check that your AWS credentials are valid",
        "Ensure your IAM user has the necessary permissions",
    },
)
```

### Testing

- Write unit tests for new functions
- Aim for meaningful test coverage, not arbitrary percentages
- Test both happy paths and error conditions
- Use table-driven tests for multiple test cases
- Mock external dependencies (AWS API calls)

Example:
```go
func TestFormatDuration(t *testing.T) {
    tests := []struct {
        name     string
        duration time.Duration
        want     string
    }{
        {"seconds", 45 * time.Second, "45s"},
        {"minutes", 3 * time.Minute, "3m0s"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := formatDuration(tt.duration)
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Documentation

### Code Documentation

- Add godoc comments for exported functions
- Include usage examples in documentation
- Document parameters and return values
- Explain non-obvious behavior

### User Documentation

- Update user guides for new features
- Add examples and tutorials
- Include screenshots for UI changes
- Keep language clear and accessible for researchers

### Documentation Structure

```
docs/
├── getting-started/
├── user-guides/
├── architecture/
├── development/
└── examples/
```

## Community

### Getting Help

- **GitHub Discussions:** Ask questions and share ideas
- **GitHub Issues:** Report bugs and request features
- **Documentation:** Check docs for answers

### Communication

- Be patient and respectful
- Provide context and details
- Search before asking duplicate questions
- Help others when you can

## Recognition

Contributors will be recognized in:
- CHANGELOG.md for significant contributions
- GitHub contributors page
- Release notes for major features

## Questions?

If you have questions about contributing, please:
1. Check this guide and project documentation
2. Search existing GitHub Discussions
3. Open a new Discussion in Q&A category

Thank you for contributing to AWS IDE! Your efforts help make cloud-based research tools more accessible to the academic community.
