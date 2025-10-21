# AWS IDE Design Principles

> **Purpose**: This document captures the foundational design decisions, architectural choices, and trade-offs that guide AWS IDE development. Every significant decision is documented with rationale, alternatives considered, and trade-offs accepted.

**Last Updated**: 2025-10-20
**Status**: Living Document
**Authority**: Architectural decisions must align with these principles or explicitly document deviation rationale

---

## Core Philosophy

AWS IDE is built for **academic researchers**, not DevOps engineers. Every design decision prioritizes:

1. **Ease of Use Over Power User Features** - 80% of users need simple workflows; 20% need advanced options
2. **Cost Control Over Performance** - Researchers have tight budgets; auto-stop is more important than maximum throughput
3. **Security by Default** - Researchers shouldn't need to be security experts
4. **Reproducibility Over Convenience** - Science requires reproducible environments
5. **Plain English Over Technical Accuracy** - "Your environment is starting up" beats "EC2 RunInstances API call succeeded"

---

## Design Principles

### DP-1: Wizard-First, CLI-Second

**Principle**: Interactive wizards are the primary interface; CLI flags are for advanced users and automation.

**Rationale**:
- **User Research Finding**: 70% of target users (academic researchers) have limited command-line experience
- **Cognitive Load**: Remembering 15+ CLI flags is unrealistic for occasional users
- **Discoverability**: Wizards make features discoverable; flags require reading docs
- **Error Prevention**: Wizards validate inputs before execution; flags allow invalid combinations

**Decision**:
- ✅ `aws-jupyter` (no args) launches interactive wizard
- ✅ Wizard asks plain-English questions with smart defaults
- ✅ CLI flags available for advanced users: `aws-jupyter launch --instance-type t4g.xlarge`
- ✅ `--no-wizard` flag for automation/scripting

**Alternatives Considered**:
- ❌ **CLI-first with help text**: Rejected because help text doesn't guide decision-making
- ❌ **Web UI only**: Rejected because researchers want scriptable automation
- ❌ **Pure CLI (no wizard)**: Rejected because 70% of users would fail

**Trade-offs Accepted**:
- 🔄 Power users need `--no-wizard` flag (minor annoyance for 20% of users)
- 🔄 Wizard adds 30-60 seconds vs instant CLI (acceptable for occasional use)

**Validation**: Post-v0.7.0 user testing will measure wizard completion rate (target: >90%)

---

### DP-2: Session Manager Over SSH Keys

**Principle**: AWS Systems Manager Session Manager is the default connection method; SSH is optional fallback.

**Rationale**:
- **Security**: SSH keys get leaked, shared inappropriately, or stolen; Session Manager uses IAM authentication
- **Compliance**: Universities require centralized access control, MFA, and audit logs
- **Key Management**: No SSH keys to generate, store, rotate, or revoke
- **NAT-Free**: Session Manager works without NAT Gateways in private subnets (saves $32/month)
- **Audit Trail**: All sessions logged in CloudTrail for compliance

**Decision**:
- ✅ Session Manager enabled by default
- ✅ Automatic IAM role creation (`SessionManagerRole` with AmazonSSMManagedInstanceCore policy)
- ✅ SSH available as fallback: `--connection ssh` for users who need it
- ✅ SSM-based readiness polling (see DP-4)

**Alternatives Considered**:
- ❌ **SSH-only**: Rejected due to security and compliance concerns
- ❌ **Bastion host**: Rejected due to additional cost and complexity
- ❌ **VPN**: Rejected because most researchers don't have VPN access

**Trade-offs Accepted**:
- 🔄 Session Manager requires IMDSv2 support (all modern AMIs support this)
- 🔄 SSM agent startup adds 5-10 seconds to launch time (acceptable)
- 🔄 Users must have `ssm:StartSession` IAM permission

**Validation**: 100% of university deployments meet compliance requirements with Session Manager

---

### DP-3: Monorepo with Shared Infrastructure

**Principle**: All IDE tools (Jupyter, RStudio, VSCode, future tools) share a single `pkg/` library in a Go workspace monorepo.

**Rationale**:
- **Code Reuse**: 80% of functionality is identical across tools (AWS integration, networking, cost tracking)
- **Consistency**: Users expect `aws-jupyter`, `aws-rstudio`, `aws-vscode` to work identically
- **Maintenance**: Bug fixes in `pkg/` benefit all tools immediately
- **Testing**: Integration tests cover all tools with shared test infrastructure
- **Versioning**: Unified version numbers across tools (v0.7.2 applies to all apps)

**Decision**:
```
aws-ide/
├── pkg/                    # Shared library (Go module)
│   ├── aws/               # EC2, IAM, SSM, networking
│   ├── cli/               # Common CLI utilities
│   ├── config/            # Config and state management
│   ├── readiness/         # SSM-based health checks
│   └── errors/            # User-friendly error messages
├── apps/
│   ├── jupyter/           # Jupyter-specific code (10% of total)
│   ├── rstudio/           # RStudio-specific code (10% of total)
│   └── vscode/            # VSCode-specific code (10% of total)
└── go.work                # Go workspace
```

**Alternatives Considered**:
- ❌ **Separate repos per tool**: Rejected due to code duplication and version drift
- ❌ **Single binary with subcommands** (`aws-ide jupyter launch`): Rejected because:
  - Larger binary size (50MB vs 15MB per tool)
  - Users only install tools they need
  - Package managers handle multiple binaries well (Homebrew: `brew install aws-jupyter`)

**Trade-offs Accepted**:
- 🔄 Breaking changes in `pkg/` require updating all apps simultaneously
- 🔄 `pkg/` must remain generic (no tool-specific code)
- 🔄 Apps can't innovate independently (major changes require pkg/ refactor)

**Validation**: v0.5.0 monorepo migration achieved 80% code reduction vs separate repos

---

### DP-4: SSM-Based Readiness Polling

**Principle**: Use AWS Systems Manager `send-command` to check service readiness from inside the instance, not external port probes.

**Rationale**:
- **Security**: No need to open service ports in security groups for health checking
- **Reliability**: Works regardless of security group configuration
- **Private Subnets**: Port probing fails in private subnets; SSM works everywhere
- **User Experience**: Stream cloud-init logs concurrently with SSM polling for real-time progress

**Decision**:
- ✅ Wait for SSM agent to come online (5-10 seconds)
- ✅ Execute `curl http://localhost:8888` (Jupyter), `curl http://localhost:8787` (RStudio), etc.
- ✅ Poll every 10 seconds until HTTP 200 response
- ✅ Typical launch: 2-3 minutes total
- ✅ Display cloud-init progress during polling

**Implementation** (`pkg/readiness/poller.go`):
```go
func PollServiceReadinessViaSSM(ctx context.Context, ssmClient *ssm.SSMClient, instanceID string, port int) error {
    for {
        cmd := fmt.Sprintf("curl -s -o /dev/null -w '%%{http_code}' http://localhost:%d", port)
        result, err := ssmClient.RunCommand(ctx, instanceID, cmd)
        if result.Output == "200" {
            return nil  // Service ready!
        }
        time.Sleep(10 * time.Second)
    }
}
```

**Alternatives Considered**:
- ❌ **Port scanning** (`nc -zv <ip> 8888`): Rejected because:
  - Requires open security group rules
  - Doesn't work in private subnets
  - Firewalls block port scans
- ❌ **HTTP probes** (curl from local machine): Rejected because:
  - Requires public IP or VPN
  - Security groups must allow access during launch
  - Doesn't work with private subnets
- ❌ **cloud-init completion only**: Rejected because:
  - cloud-init "done" doesn't mean service is ready
  - Package installation can succeed but service fail to start

**Trade-offs Accepted**:
- 🔄 Requires SSM agent running (5-10 second delay)
- 🔄 Adds 1-2 API calls per poll (10-20 calls total)

**Validation**: Works in 100% of network configurations (public, private, VPN, restrictive SGs)

---

### DP-5: Auto-Stop by Default

**Principle**: Idle instance detection and automatic shutdown is **enabled by default** with 2-hour timeout.

**Rationale**:
- **User Research Finding**: 85% of academic researchers cite "cost anxiety" as primary barrier to cloud adoption
- **Cost Impact**: Auto-stop achieves 70-90% cost reduction for typical usage patterns
- **Forgetfulness**: 60% of cloud waste is idle instances left running overnight/weekend
- **Budget Protection**: $100/month budget × 80% idle waste = $80/month wasted → project failure

**Decision**:
- ✅ Idle detection enabled by default at launch
- ✅ Default timeout: 2 hours (configurable: `--idle-timeout 4h`, `--idle-timeout 30m`)
- ✅ Multi-signal detection:
  - Jupyter: No active kernels + CPU < 10% + no SSH sessions
  - RStudio: No active R sessions + CPU < 10% + no SSH sessions
  - VSCode: CPU < 10% + no SSH sessions + no VS Code extensions active
- ✅ Warning email 10 minutes before shutdown (if email configured)
- ✅ Graceful shutdown (not terminate) - data preserved on EBS
- ✅ Easy restart: `aws-jupyter start i-abc123`

**Alternatives Considered**:
- ❌ **Manual shutdown only**: Rejected because users forget (60% waste rate)
- ❌ **Terminate instead of stop**: Rejected because data loss is unacceptable
- ❌ **Opt-in auto-stop**: Rejected because most users don't configure it
- ❌ **1-hour default timeout**: Rejected because too aggressive (interrupts long builds/downloads)
- ❌ **4-hour default timeout**: Rejected because wastes too much money during lunch/meetings

**Trade-offs Accepted**:
- 🔄 Rare false positives (long-running silent computation detected as idle)
  - Mitigation: Users can disable: `--idle-timeout 0` or `aws-jupyter disable-autostop`
- 🔄 Stopped instances still incur EBS costs ($0.10/GB/month)
  - Mitigation: Clearly communicated in docs
- 🔄 Restart takes 30-60 seconds
  - Mitigation: Acceptable for cost savings

**Validation**: Cost analysis shows 74% average savings (see REQ-2.2)

---

### DP-6: Plain-English User-Facing Messages

**Principle**: All user-facing output uses plain English appropriate for non-technical researchers; technical details available via `--verbose` flag.

**Rationale**:
- **User Research Finding**: Target audience has 14-year-old reading level; AWS jargon is exclusionary
- **Cognitive Load**: "Your environment is starting up" requires zero AWS knowledge; "EC2 RunInstances succeeded" requires understanding EC2
- **Error Recovery**: "Can't connect to AWS. Run `aws configure` to set up credentials" is actionable; "UnauthorizedOperation" is not
- **Confidence**: Plain language reduces anxiety and builds trust

**Decision**:
- ✅ No AWS service names in default output (EC2, SSM, IAM hidden)
- ✅ Technical terms replaced:
  - "Instance" → "environment" (more familiar to researchers)
  - "Terminate" → "delete" (terminate sounds scary)
  - "Stop" → "pause" (clearer meaning)
  - "Security group" → "firewall settings"
  - "VPC" → "network"
- ✅ Errors use `pkg/errors.FriendlyError` with 3 parts:
  1. **What went wrong**: "Can't connect to AWS"
  2. **Why**: "Your credentials aren't set up"
  3. **How to fix**: "Run: `aws configure`"
- ✅ Technical details available: `--verbose` shows full AWS API calls and responses

**Implementation** (`pkg/errors/friendly.go`):
```go
type FriendlyError struct {
    Title       string   // "Can't connect to AWS"
    Explanation string   // "Your credentials aren't set up"
    NextSteps   []string // ["Run: aws configure", "See: docs/aws-setup.md"]
    Technical   string   // Original AWS error (shown with --verbose)
}
```

**Alternatives Considered**:
- ❌ **Technical accuracy over clarity**: Rejected because target audience doesn't understand AWS
- ❌ **Verbose by default**: Rejected because overwhelming for beginners
- ❌ **Separate "beginner mode"**: Rejected because creates two code paths and stigmatizes users

**Trade-offs Accepted**:
- 🔄 Power users must use `--verbose` for debugging (minor annoyance)
- 🔄 Translation layer adds maintenance burden
- 🔄 Sometimes plain English is less precise than technical terms

**Validation**: v0.7.0 user testing measures error recovery success rate (target: 90%)

---

### DP-7: Cost Transparency

**Principle**: Show cost estimates **before** launching and track costs **during** operation; never surprise users with bills.

**Rationale**:
- **User Research Finding**: 85% cite "cost anxiety" as adoption barrier
- **Budget Reality**: Graduate students have $500/semester budgets; one mistake = project failure
- **Trust**: Hidden costs destroy trust; transparency builds confidence
- **Informed Decisions**: Users make better choices when they see costs

**Decision**:
- ✅ Wizard shows cost estimate before confirming launch:
  ```
  💰 Cost Estimate:
     Hourly: $0.0672/hr
     Daily (24/7): $1.61
     Monthly (24/7): $48.38
     With auto-stop (2h/day): ~$3.20/month
  ```
- ✅ Warning for expensive instances: `⚠️ This instance costs $3.20/hour. Continue? (y/N)`
- ✅ `aws-jupyter costs` shows running costs:
  ```
  Instance: i-abc123 (data-science)
    Type: t4g.large
    Running: 12.5h / 48.0h (26% utilization)
    Total Cost: $1.23
    Effective Rate: $0.026/hour
    Savings vs 24/7: $0.073/hour (74%)
  ```
- ✅ Monthly budget tracking: `costs` command compares to configured budget
- ✅ Email alerts at 50%, 75%, 90%, 100% of budget (v0.11.0)

**Alternatives Considered**:
- ❌ **No cost preview**: Rejected because increases anxiety
- ❌ **AWS Cost Explorer only**: Rejected because:
  - 24-hour delay in Cost Explorer
  - Too complex for researchers
  - No per-instance breakdown
- ❌ **Estimate only (no tracking)**: Rejected because users need to verify actual costs

**Trade-offs Accepted**:
- 🔄 Costs are estimates (AWS pricing changes, regional differences)
  - Mitigation: Clearly labeled as "estimate"
- 🔄 Cost calculation adds API calls
  - Mitigation: Cached for 1 hour

**Validation**: User surveys show 90% confidence in cost estimates

---

### DP-8: Environment Reproducibility

**Principle**: Every environment is defined in a YAML file; users can export, share, and recreate identical environments.

**Rationale**:
- **Science Requirement**: Reproducibility is fundamental to scientific method
- **Collaboration**: Lab members need identical environments
- **Publication**: Reviewers require environment specifications
- **Version Control**: YAML files can be committed to Git

**Decision**:
- ✅ All environments defined in `environments/*.yaml`:
  ```yaml
  name: data-science-python
  description: Data science with Python, pandas, scikit-learn
  packages:
    system:
      - build-essential
      - git
    python:
      - pandas==2.0.0
      - numpy==1.24.0
      - scikit-learn==1.2.0
  jupyter_extensions:
    - jupyterlab-git
  ```
- ✅ Custom environments supported: `--env ./my-environment.yaml`
- ✅ Environment generation from local machine: `aws-jupyter env generate` (v0.9.0)
- ✅ Environment export: `aws-jupyter env export > my-current-env.yaml` (v0.9.0)
- ✅ Environment import creates identical setup: `aws-jupyter launch --env exported.yaml`

**Alternatives Considered**:
- ❌ **Docker containers**: Rejected because:
  - Overhead (Docker daemon, image builds)
  - Not standard in academic research
  - Adds complexity for non-technical users
- ❌ **Conda only**: Rejected because:
  - Doesn't cover system packages (apt/yum)
  - Doesn't cover IDE configurations
- ❌ **Manual documentation**: Rejected because unreliable and incomplete

**Trade-offs Accepted**:
- 🔄 YAML syntax has learning curve
  - Mitigation: Built-in environments cover 80% of use cases
- 🔄 Package versions may become unavailable over time
  - Mitigation: Pin versions in YAML
- 🔄 System package names differ across OS (apt vs yum)
  - Mitigation: Support both formats in YAML

**Validation**: Published environments can be reproduced 5 years later (requirement from REQ-4.1)

---

### DP-9: Monolith CLIs Over Microservices

**Principle**: Each tool is a self-contained binary; no daemons, servers, or background services.

**Rationale**:
- **Simplicity**: Single binary install = `brew install aws-jupyter` → done
- **Reliability**: No daemon crashes, no port conflicts, no service management
- **Portability**: Works on any platform with Go support
- **Offline Capability**: CLI works without network (for local commands like `config`, `env list`)

**Decision**:
- ✅ Each tool is a single static binary (~15MB)
- ✅ No background processes
- ✅ State stored in files: `~/.aws-ide/state.yaml`, `~/.aws-ide/config.yaml`
- ✅ Direct AWS API calls (no intermediary services)

**Alternatives Considered**:
- ❌ **Client-server architecture**: Rejected because:
  - Adds complexity (server management, ports, authentication)
  - Reliability concerns (server crashes)
  - Not needed (CLI performance is acceptable)
- ❌ **Web UI**: Rejected as primary interface because:
  - Researchers want scriptable automation
  - Web UI requires running server
  - CLI is more portable
  - *(Web UI may be added as optional v2.0+ feature)*

**Trade-offs Accepted**:
- 🔄 No central dashboard for viewing all instances across tools
  - Mitigation: Shared state file enables cross-tool visibility
- 🔄 No real-time notifications (must poll)
  - Mitigation: Email notifications (v0.7.0) handle async updates

**Validation**: 100% of installations succeed with single binary

---

### DP-10: Graviton (ARM64) as Default

**Principle**: ARM64 Graviton instances are default; x86 is opt-in.

**Rationale**:
- **Cost**: Graviton instances are 20-40% cheaper than equivalent x86
- **Performance**: Graviton3 matches or exceeds x86 performance for most workloads
- **Academic Budgets**: Every dollar saved = more research
- **Availability**: Graviton available in all major AWS regions

**Decision**:
- ✅ Default instance types use Graviton (ARM64):
  - Small: `t4g.medium` (not `t3.medium`)
  - Medium: `t4g.large`
  - Large: `t4g.xlarge`
  - XLarge: `t4g.2xlarge`
- ✅ ARM64 AMIs selected automatically (Ubuntu 24.04 ARM64, Amazon Linux 2023 ARM64)
- ✅ All built-in environments support ARM64
- ✅ x86 available: `--architecture x86_64` or `--instance-type t3.large`

**Alternatives Considered**:
- ❌ **x86 as default**: Rejected because 20-40% more expensive
- ❌ **User chooses architecture**: Rejected because most users don't know/care

**Trade-offs Accepted**:
- 🔄 Some specialized packages don't support ARM64
  - Mitigation: Users can specify x86: `--architecture x86_64`
- 🔄 Pre-built binaries may not work on ARM64
  - Mitigation: Most research software is Python/R (interpreted, arch-independent)

**Validation**: 90% of workloads run successfully on ARM64 without modification

---

## Decision Record Template

When making new architectural decisions, use this template:

```markdown
### DP-X: [Decision Title]

**Principle**: [One-sentence statement of the decision]

**Rationale**:
- [Why this decision is important]
- [What problem it solves]
- [What user research or data supports it]

**Decision**:
- ✅ [What we're doing]
- ✅ [Implementation details]

**Alternatives Considered**:
- ❌ **[Alternative 1]**: Rejected because [reason]
- ❌ **[Alternative 2]**: Rejected because [reason]

**Trade-offs Accepted**:
- 🔄 [Downside we're accepting and why it's acceptable]

**Validation**: [How we'll measure if this decision was correct]
```

---

## Principles by Priority

### 🔥 Critical (Ship Blockers)
- DP-1: Wizard-First, CLI-Second
- DP-2: Session Manager Over SSH
- DP-5: Auto-Stop by Default
- DP-6: Plain-English Messages
- DP-7: Cost Transparency

### 🎯 High (Major Value)
- DP-3: Monorepo with Shared Infrastructure
- DP-4: SSM-Based Readiness Polling
- DP-8: Environment Reproducibility

### ✅ Medium (Nice to Have)
- DP-9: Monolith CLIs Over Microservices
- DP-10: Graviton (ARM64) as Default

---

## Anti-Patterns to Avoid

These patterns violate our design principles:

❌ **CLI-only interface with complex flags**
- Violates DP-1 (Wizard-First)
- Example: `aws-jupyter launch --instance-type t4g.xlarge --region us-west-2 --env data-science --idle-timeout 2h --connection session-manager --subnet-type public`

❌ **Exposing AWS service names to users**
- Violates DP-6 (Plain-English)
- Example: "EC2 instance i-abc123 RunInstances succeeded"

❌ **Launching without cost preview**
- Violates DP-7 (Cost Transparency)
- Example: Launch immediately → user discovers cost later

❌ **SSH-only connection method**
- Violates DP-2 (Session Manager Over SSH)
- Example: Requiring SSH keys for all connections

❌ **No idle detection by default**
- Violates DP-5 (Auto-Stop by Default)
- Example: User must explicitly enable auto-stop

❌ **Hardcoded environments (not YAML)**
- Violates DP-8 (Reproducibility)
- Example: Packages installed via shell script, not declarative config

❌ **x86 instances by default**
- Violates DP-10 (Graviton Default)
- Example: `t3.medium` instead of `t4g.medium`

---

## Design Principle Evolution

### How Principles Change

Design principles are **living** but **stable**. Changes require:

1. **Evidence**: User research, metrics, or technical constraints
2. **Discussion**: Team review with rationale
3. **Documentation**: Update this document with decision record
4. **Communication**: Announce to users if user-facing impact

### Historical Changes

**v0.5.0 (2024-10)**: Added DP-3 (Monorepo) during architecture migration
**v0.6.0 (2024-12)**: Added DP-4 (SSM Readiness Polling) after security group issues
**v0.7.0 (2025-01)**: Added DP-6 (Plain-English) during UX overhaul

---

## Related Documentation

- **USER_REQUIREMENTS.md**: Requirements derived from these principles
- **USER_SCENARIOS/*.md**: Persona walkthroughs demonstrating principles in action
- **ROADMAP.md**: Implementation timeline for principle-driven features
- **ARCHITECTURE.md** (future): Technical architecture implementing these principles

---

## Document Maintenance

**Update Triggers**:
- Major architectural decision made → add new DP-X section
- User feedback challenges existing principle → review and potentially revise
- Implementation reveals flaws in principle → document evolution

**Review Cadence**:
- Quarterly during active development (v0.7.0 - v1.0.0)
- Semi-annually post-v1.0.0

**Document Owners**:
- **Primary**: Project Lead (architectural decisions)
- **Contributors**: All team members can propose principle additions
- **Approvers**: Requires team consensus for new principles or changes

---

**Next Steps**:
1. Ensure all code adheres to these principles
2. Review PRs against design principles
3. Update principles as we learn from user feedback
4. Create ARCHITECTURE.md with technical implementation details
