# AWS Authentication Guide

This guide covers all the ways to authenticate lens-jupyter with AWS. The tool uses the standard AWS SDK credential chain, so any method that works with the AWS CLI will work with lens-jupyter.

## 🚀 Quick Setup

### Option 1: AWS Profiles (Recommended)

```bash
# Configure a new profile
aws configure --profile jupyter
# Enter: Access Key ID, Secret Access Key, Region, Output format

# Use with lens-jupyter
lens-jupyter launch --profile jupyter

# Or set as environment variable
export AWS_PROFILE=jupyter
lens-jupyter launch
```

### Option 2: Environment Variables

```bash
export AWS_ACCESS_KEY_ID="your-access-key-id"
export AWS_SECRET_ACCESS_KEY="your-secret-access-key"
export AWS_REGION="us-west-2"

lens-jupyter launch
```

### Option 3: AWS SSO (Single Sign-On)

```bash
# Configure SSO profile
aws configure sso --profile jupyter-sso

# Login and use
aws sso login --profile jupyter-sso
lens-jupyter launch --profile jupyter-sso
```

## 📋 Authentication Methods (Priority Order)

The AWS SDK checks credentials in this order:

1. **Command Line Parameters** (`--profile`, `--region`)
2. **Environment Variables** (`AWS_ACCESS_KEY_ID`, etc.)
3. **AWS Credentials File** (`~/.aws/credentials`)
4. **AWS Config File** (`~/.aws/config`)
5. **IAM Roles** (when running on EC2/ECS/Lambda)
6. **Container Credentials** (ECS tasks)
7. **Instance Metadata** (EC2 instances)

## 🔐 Detailed Authentication Methods

### 1. AWS Profiles

**Best for:** Local development, multiple AWS accounts

```bash
# Create profile interactively
aws configure --profile myprofile

# Or edit files directly
cat >> ~/.aws/credentials << EOF
[myprofile]
aws_access_key_id = AKIA...
aws_secret_access_key = abc123...
EOF

cat >> ~/.aws/config << EOF
[profile myprofile]
region = us-west-2
output = json
EOF
```

**Using with lens-jupyter:**
```bash
lens-jupyter launch --profile myprofile
lens-jupyter list --profile myprofile --region us-east-1
```

### 2. Environment Variables

**Best for:** CI/CD, containerized environments

```bash
# Required
export AWS_ACCESS_KEY_ID="AKIA..."
export AWS_SECRET_ACCESS_KEY="abc123..."
export AWS_REGION="us-west-2"

# Optional
export AWS_SESSION_TOKEN="token..."  # For temporary credentials
export AWS_PROFILE="myprofile"       # Use specific profile
export AWS_DEFAULT_REGION="us-west-2" # Alternative to AWS_REGION
```

**Using with lens-jupyter:**
```bash
# Credentials automatically detected
lens-jupyter launch

# Override region if needed
lens-jupyter launch --region eu-west-1
```

### 3. AWS SSO (Single Sign-On)

**Best for:** Enterprise environments with centralized auth

#### Initial Setup
```bash
# Configure SSO profile
aws configure sso
# Follow prompts:
# - SSO start URL (e.g., https://mycompany.awsapps.com/start)
# - SSO region (e.g., us-east-1)
# - Account ID and role name
# - CLI profile name (e.g., jupyter-sso)
# - Default region and output format
```

#### Daily Usage
```bash
# Login (opens browser)
aws sso login --profile jupyter-sso

# Verify login
aws sts get-caller-identity --profile jupyter-sso

# Use with lens-jupyter
lens-jupyter launch --profile jupyter-sso
```

#### Session Management
```bash
# Check if logged in
aws sts get-caller-identity --profile jupyter-sso

# Login if session expired
aws sso login --profile jupyter-sso

# Logout
aws sso logout
```

### 4. IAM Roles (EC2/ECS/Lambda)

**Best for:** Running lens-jupyter on AWS infrastructure

#### EC2 Instance Roles
```bash
# No configuration needed!
# lens-jupyter automatically uses the instance role

# Just run commands
lens-jupyter launch
lens-jupyter list
```

#### ECS Task Roles
```bash
# Configure in ECS task definition
{
  "taskRoleArn": "arn:aws:iam::123456789012:role/MyTaskRole",
  "executionRoleArn": "arn:aws:iam::123456789012:role/MyExecutionRole"
}

# lens-jupyter automatically uses the task role
lens-jupyter launch
```

### 5. Temporary Credentials

**Best for:** Assumed roles, MFA requirements

```bash
# Assume role with AWS CLI
aws sts assume-role \
  --role-arn arn:aws:iam::123456789012:role/MyRole \
  --role-session-name jupyter-session \
  --profile myprofile

# Extract credentials from output and set as env vars
export AWS_ACCESS_KEY_ID="ASIA..."
export AWS_SECRET_ACCESS_KEY="abc123..."
export AWS_SESSION_TOKEN="token123..."

# Use with lens-jupyter
lens-jupyter launch
```

### 6. Cross-Account Access

**Best for:** Managing resources across AWS accounts

```bash
# Configure role assumption in ~/.aws/config
cat >> ~/.aws/config << EOF
[profile cross-account]
role_arn = arn:aws:iam::999999999999:role/CrossAccountRole
source_profile = myprofile
region = us-west-2
EOF

# Use the cross-account profile
lens-jupyter launch --profile cross-account
```

## 🛡️ Required AWS Permissions

lens-jupyter needs these IAM permissions:

### Minimum Required Permissions
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:DescribeInstances",
        "ec2:DescribeImages",
        "ec2:DescribeKeyPairs",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeVpcs",
        "ec2:DescribeSubnets"
      ],
      "Resource": "*"
    }
  ]
}
```

### Full Functionality Permissions
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:RunInstances",
        "ec2:TerminateInstances",
        "ec2:StartInstances",
        "ec2:StopInstances",
        "ec2:DescribeInstances",
        "ec2:DescribeImages",
        "ec2:DescribeKeyPairs",
        "ec2:CreateKeyPair",
        "ec2:DeleteKeyPair",
        "ec2:DescribeSecurityGroups",
        "ec2:CreateSecurityGroup",
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:RevokeSecurityGroupIngress",
        "ec2:DescribeVpcs",
        "ec2:DescribeSubnets",
        "ec2:CreateTags",
        "ec2:DescribeTags"
      ],
      "Resource": "*"
    }
  ]
}
```

### Region-Specific Permissions (Optional)
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "ec2:*",
      "Resource": "*",
      "Condition": {
        "StringEquals": {
          "aws:RequestedRegion": ["us-west-2", "us-east-1"]
        }
      }
    }
  ]
}
```

## 🧪 Testing Your Configuration

### Verify AWS Credentials
```bash
# Test basic access
aws sts get-caller-identity

# Test with specific profile
aws sts get-caller-identity --profile myprofile

# Test EC2 access
aws ec2 describe-regions

# Test in specific region
aws ec2 describe-vpcs --region us-west-2
```

### Test with lens-jupyter
```bash
# Dry run to test configuration
lens-jupyter launch --dry-run
lens-jupyter launch --dry-run --profile myprofile --region us-west-2

# List existing instances (read-only)
lens-jupyter list

# Test environment validation
lens-jupyter env validate data-science
```

## 🔧 Troubleshooting

### Common Issues

#### "Unable to locate credentials"
```bash
# Check if AWS CLI is configured
aws configure list

# Check environment variables
env | grep AWS

# Verify profile exists
cat ~/.aws/credentials
```

#### "Access Denied" / "UnauthorizedOperation"
```bash
# Check current identity
aws sts get-caller-identity

# Test minimal EC2 permissions
aws ec2 describe-regions

# Review IAM permissions in AWS Console
```

#### "No default VPC" / "No subnets found"
```bash
# Check VPCs in your region
aws ec2 describe-vpcs

# Check subnets
aws ec2 describe-subnets

# Create default VPC if needed
aws ec2 create-default-vpc
```

#### SSO Session Expired
```bash
# Re-login to SSO
aws sso login --profile myprofile

# Check session status
aws sts get-caller-identity --profile myprofile
```

### Debug Commands

```bash
# Enable AWS SDK debug logging
export AWS_SDK_LOAD_CONFIG=1
export AWS_DEBUG=1

# Run with verbose output
lens-jupyter launch --dry-run

# Check credential resolution
aws configure list --profile myprofile
aws configure get region --profile myprofile
```

## 🏢 Enterprise Setup Examples

### Multi-Account Setup
```bash
# ~/.aws/config
[default]
region = us-west-2

[profile dev]
role_arn = arn:aws:iam::111111111111:role/DeveloperRole
source_profile = default

[profile staging]
role_arn = arn:aws:iam::222222222222:role/DeveloperRole
source_profile = default

[profile prod]
role_arn = arn:aws:iam::333333333333:role/ProductionRole
source_profile = default
mfa_serial = arn:aws:iam::123456789012:mfa/myusername
```

### CI/CD Pipeline Setup
```bash
# GitHub Actions / GitLab CI
export AWS_ACCESS_KEY_ID=${{ secrets.AWS_ACCESS_KEY_ID }}
export AWS_SECRET_ACCESS_KEY=${{ secrets.AWS_SECRET_ACCESS_KEY }}
export AWS_REGION=us-west-2

# Use in pipeline
lens-jupyter launch --env minimal --dry-run
```

### Docker Container Setup
```dockerfile
# Pass credentials as environment variables
docker run -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY -e AWS_REGION \
  myorg/lens-jupyter:latest launch --env data-science

# Or mount credential files
docker run -v ~/.aws:/root/.aws:ro \
  myorg/lens-jupyter:latest launch --profile myprofile
```

## 🔐 Security Best Practices

### 1. Use IAM Roles When Possible
- ✅ EC2 instance roles
- ✅ ECS task roles
- ✅ Lambda execution roles
- ❌ Hard-coded access keys

### 2. Principle of Least Privilege
- Grant only necessary EC2 permissions
- Use resource-based restrictions when possible
- Regularly audit and rotate credentials

### 3. Temporary Credentials
- Use AWS SSO for human access
- Use temporary credentials for automation
- Set appropriate session duration limits

### 4. Secure Storage
- Never commit credentials to source control
- Use environment variables or secrets managers
- Encrypt credential files at rest

### 5. Multi-Factor Authentication
- Enable MFA for AWS root account
- Require MFA for sensitive operations
- Use hardware tokens for production access

---

## 📞 Need Help?

- **AWS Documentation**: [AWS CLI Configuration](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)
- **GitHub Issues**: [Report authentication problems](https://github.com/scttfrdmn/lens-jupyter/issues/new/choose)
- **AWS Support**: For AWS account and IAM issues

**Still having trouble?** Open a [GitHub issue](https://github.com/scttfrdmn/lens-jupyter/issues/new/choose) with:
- Your authentication method
- Commands you've tried
- Error messages (sanitized)
- Output of `aws sts get-caller-identity`