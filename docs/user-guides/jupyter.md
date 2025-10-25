# lens-jupyter User Guide

Launch secure Jupyter Lab instances on AWS for data science and machine learning.

## Quick Start

The simplest way to launch Jupyter Lab is using the interactive wizard:

```bash
lens-jupyter
```

This launches the wizard which guides you through:
1. Selecting your use case (data science, ML, etc.)
2. Choosing instance type and cost options
3. Configuring auto-stop to save money
4. Launching and connecting

## Common Use Cases

### Data Science with Python

```bash
# Interactive wizard (recommended for beginners)
lens-jupyter

# Quick launch with defaults
lens-jupyter quickstart

# Launch with specific environment
lens-jupyter launch --env data-science-python
```

### Machine Learning with GPU

```bash
# Launch GPU instance for training
lens-jupyter launch --env deep-learning-gpu --instance-type g4dn.xlarge

# Use spot instance for cost savings
lens-jupyter launch --env ml-gpu --instance-type g4dn.xlarge --spot
```

### Bioinformatics

```bash
lens-jupyter launch --env bioinformatics
```

## Available Commands

### Launch

Start a new Jupyter Lab instance:

```bash
lens-jupyter launch [flags]
```

**Common flags:**
- `--env`: Environment name (see `lens-jupyter env list`)
- `--instance-type`: EC2 instance type (default: t3.medium)
- `--region`: AWS region (default: us-east-1)
- `--spot`: Use spot instance for cost savings
- `--auto-stop`: Minutes of inactivity before auto-stop
- `--name`: Custom name for the instance

### List

View all running instances:

```bash
lens-jupyter list

# Filter by state
lens-jupyter list --state running
lens-jupyter list --state stopped
```

### Connect

Connect to an existing instance:

```bash
lens-jupyter connect [instance-id]

# Connect via Session Manager (no SSH keys needed)
lens-jupyter connect i-abc123 --session-manager
```

### Stop

Stop a running instance (data persists):

```bash
lens-jupyter stop i-abc123
```

### Start

Start a stopped instance:

```bash
lens-jupyter start i-abc123
```

### Terminate

Permanently delete an instance:

```bash
lens-jupyter terminate i-abc123
```

## Environments

List available environments:

```bash
lens-jupyter env list
```

Show environment details:

```bash
lens-jupyter env show data-science-python
```

## Tips and Best Practices

1. **Use auto-stop**: Save money by automatically stopping idle instances
2. **Start with t3.medium**: Good balance of cost and performance for most work
3. **Use spot instances**: Save up to 90% for non-critical workloads
4. **Session Manager**: No SSH keys needed, more secure
5. **Name your instances**: Use `--name` for easy identification

## Troubleshooting

See the [troubleshooting guide](troubleshooting.md) for common issues and solutions.
