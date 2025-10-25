# Lens User Requirements

> **Purpose**: This document defines the authoritative requirements for Lens based on academic researcher needs. Every requirement maps to user personas, quantified pain points, and measurable success criteria.

**Last Updated**: 2025-10-20
**Status**: Living Document
**Authority**: This document is the source of truth for feature prioritization

---

## Document Structure

This document organizes requirements into these categories:

1. **Ease of Use** - Non-technical researchers must succeed without training
2. **Cost Management** - Academic budgets are limited; cost control is critical
3. **Research Tools** - Support diverse research workflows and domains
4. **Reproducibility** - Enable reproducible research and collaboration
5. **Security & Compliance** - Meet institutional security requirements
6. **Performance** - Researchers need responsive, reliable tools

Each requirement follows this format:
- **Requirement ID**: Unique identifier
- **User Story**: "As a [persona], I need [feature] because [reason]"
- **Pain Point**: Quantified current problem
- **Success Metric**: Measurable outcome
- **Priority**: Critical / High / Medium / Low
- **Phase**: Target roadmap phase (v0.7.0 - v1.0.0)
- **Related Personas**: Which personas benefit
- **Related Issues**: GitHub issues implementing this requirement

---

## 1. Ease of Use Requirements

### REQ-1.1: Beginner-Friendly Onboarding

**User Story**: As a Solo Researcher with limited technical skills, I need to launch my first instance in under 5 minutes without reading documentation, because I have analysis deadlines and can't spend days learning new tools.

**Pain Point**:
- **Before**: Manual EC2 setup requires understanding VPCs, security groups, SSH keys, IAM roles = 8+ hours for first launch
- **Technical knowledge barrier**: 70% of potential users don't know what "SSH" or "security group" means

**Success Metric**:
- ✅ Non-technical researcher can launch first instance in < 5 minutes
- ✅ Zero AWS/technical jargon in core user flows
- ✅ Interactive wizard asks questions in plain English
- ✅ Default selections follow best practices (no configuration required)

**Priority**: CRITICAL
**Phase**: v0.7.0 - User Experience & Accessibility
**Related Personas**: Solo Researcher, Graduate Student, Course Instructor
**Related Issues**: #1, #2, #3
**Related Documentation**: `USER_SCENARIOS/01_SOLO_RESEARCHER_WALKTHROUGH.md` - Pain Point #1

---

### REQ-1.2: Plain-English Error Messages

**User Story**: As a Graduate Student, I need error messages I can understand and act on, because cryptic AWS error codes block my research progress and I don't have time to become an AWS expert.

**Pain Point**:
- **Before**: AWS error "UnauthorizedOperation: You are not authorized to perform this operation" = 30 minutes searching AWS docs
- **Context loss**: Errors don't explain what the user was trying to do or how to fix it
- **Technical jargon**: Terms like "IAM policy", "VPC", "subnet" are meaningless to researchers

**Success Metric**:
- ✅ Zero AWS error codes visible to users
- ✅ Every error message has 3 parts: (1) What went wrong, (2) Why, (3) How to fix
- ✅ Next steps are actionable commands, not conceptual explanations
- ✅ 90% of users can resolve common errors without support

**Priority**: CRITICAL
**Phase**: v0.7.0 - User Experience & Accessibility (COMPLETE)
**Related Personas**: All personas
**Related Issues**: (Implemented in v0.7.0)
**Related Documentation**: `USER_SCENARIOS/03_GRAD_STUDENT_WALKTHROUGH.md` - Pain Point #2

---

### REQ-1.3: Wizard as Default Interaction

**User Story**: As a Solo Researcher, I need the wizard to launch automatically when I run `lens-jupyter`, because I don't know what flags to use and need guidance every time.

**Pain Point**:
- **Before**: Running `lens-jupyter` with no args shows help text = user doesn't know what to do next
- **Hidden wizard**: Users don't discover the `wizard` command exists
- **Flag overload**: 15+ CLI flags overwhelm beginners

**Success Metric**:
- ✅ `lens-jupyter` (no args) launches interactive wizard
- ✅ `lens-jupyter launch` uses saved preferences or sensible defaults
- ✅ Wizard usage > 80% of first-time users
- ✅ Advanced users can skip with `--no-wizard` or `launch` command

**Priority**: HIGH
**Phase**: v0.7.0 - User Experience & Accessibility
**Related Personas**: Solo Researcher, Graduate Student, Course Instructor
**Related Issues**: #1
**Related Documentation**: `USER_SCENARIOS/01_SOLO_RESEARCHER_WALKTHROUGH.md` - Pain Point #3

---

### REQ-1.4: Remember User Preferences

**User Story**: As a Data Scientist, I need Lens to remember my preferred instance type, region, and environment, because I launch the same configuration 3-4 times per week and don't want to re-enter everything.

**Pain Point**:
- **Before**: Every launch requires specifying `--instance-type t4g.xlarge --region us-west-2 --env ml-pytorch`
- **Inconsistency**: Forgetting to specify preferred settings leads to using wrong instance type = unexpected costs
- **Cognitive load**: Remembering exact flag syntax every time

**Success Metric**:
- ✅ Wizard asks "Use same settings as last time?" on subsequent launches
- ✅ Preferences saved in `~/.lens/preferences.yaml`
- ✅ Preferences per app (Jupyter vs RStudio may have different defaults)
- ✅ 70% reduction in keystrokes for repeat users

**Priority**: HIGH
**Phase**: v0.7.0 - User Experience & Accessibility
**Related Personas**: All personas
**Related Issues**: #3
**Related Documentation**: `USER_SCENARIOS/04_DATA_SCIENTIST_WALKTHROUGH.md` - Pain Point #3

---

### REQ-1.5: Quickstart for Instant Launch

**User Story**: As a Course Instructor preparing for class, I need to launch a standard environment in one command with zero prompts, because I'm launching 5 minutes before class starts and don't have time for a wizard.

**Pain Point**:
- **Before**: Even with wizard, need to answer 4-5 questions = 60 seconds minimum
- **Time pressure**: Instructors need instant launch for demos
- **Known configuration**: Experienced users know exactly what they want

**Success Metric**:
- ✅ `lens-jupyter quickstart` launches in < 10 seconds (non-interactive)
- ✅ Uses sensible defaults: t4g.medium, us-east-1, data-science env
- ✅ Can override: `lens-jupyter quickstart --env ml-pytorch --size large`
- ✅ Instructor can demo in class without wizard interruption

**Priority**: HIGH
**Phase**: v0.7.0 - User Experience & Accessibility
**Related Personas**: Course Instructor, Data Scientist, Lab Manager
**Related Issues**: #2
**Related Documentation**: `USER_SCENARIOS/05_INSTRUCTOR_WALKTHROUGH.md` - Pain Point #1

---

## 2. Cost Management Requirements

### REQ-2.1: Cost Preview Before Launch

**User Story**: As a Graduate Student with a $500/semester budget, I need to see estimated hourly and monthly costs before launching, because I can't afford surprise bills and my advisor will cut my access if I overspend.

**Pain Point**:
- **Before**: Launch instance → check bill 3 days later → realize t3.xlarge costs $3.20/day = $96/month = budget gone in 5 months
- **Budget anxiety**: 85% of academic researchers cite cost uncertainty as #1 barrier to cloud adoption
- **No recovery**: Once money is spent, it's gone; can't undo

**Success Metric**:
- ✅ Wizard shows cost estimate before confirming launch
- ✅ Display: "Hourly: $0.10/hr | Daily (24/7): $2.40 | Monthly (24/7): $72.00"
- ✅ Warning if selected instance would exceed 50% of monthly budget
- ✅ Confirmation required for instances > $0.50/hour
- ✅ 90% of users report feeling "confident about costs" in surveys

**Priority**: CRITICAL
**Phase**: v0.7.0 - User Experience & Accessibility
**Related Personas**: Graduate Student, Solo Researcher, Course Instructor
**Related Issues**: #4 (email notifications include costs)
**Related Documentation**: `USER_SCENARIOS/03_GRAD_STUDENT_WALKTHROUGH.md` - Pain Point #1

---

### REQ-2.2: Automatic Idle Shutdown

**User Story**: As a Lab PI managing 5 students, I need instances to auto-stop when idle, because students forget to shut down and I've wasted $1,200 on idle instances running overnight.

**Pain Point**:
- **Before**: Student launches t3.xlarge for 2-hour analysis → forgets to terminate → runs for 72 hours = $230 wasted
- **Forgetfulness**: 60% of cloud waste is from idle instances running unnoticed
- **Budget impact**: $1,200/year wasted across 5 students = 12% of total lab budget

**Success Metric**:
- ✅ Auto-stop enabled by default with 2-hour idle timeout
- ✅ Multi-signal idle detection: CPU < 10%, no active kernels, no SSH sessions
- ✅ Warning email 10 minutes before auto-stop
- ✅ 80% cost reduction vs 24/7 operation for typical usage patterns
- ✅ 95% of idle waste eliminated

**Priority**: CRITICAL
**Phase**: v0.6.0 - Cost Optimization (COMPLETE)
**Related Personas**: Lab PI, Graduate Student, Lab Manager
**Related Issues**: (Implemented in v0.6.0)
**Related Documentation**: `USER_SCENARIOS/02_LAB_PI_WALKTHROUGH.md` - Pain Point #2

---

### REQ-2.3: Budget Tracking and Alerts

**User Story**: As a Lab Manager overseeing 10 labs with $50,000/year budget, I need real-time budget tracking with alerts, because surprise AWS bills have caused grant budget overruns and I need early warnings to prevent overspending.

**Pain Point**:
- **Before**: Check AWS Cost Explorer once per week → discover $4,000 overage from 2 weeks ago → scramble to reallocate budget
- **No attribution**: Can't tell which lab or researcher caused the overage
- **Late detection**: AWS bill arrives 7-10 days after month ends = too late to correct

**Success Metric**:
- ✅ Set budget threshold in config: `cost_alert_threshold: 5000` ($5,000/month)
- ✅ Email alert at 50%, 75%, 90%, 100% of budget
- ✅ `lens-jupyter costs` shows current month spend vs budget
- ✅ Per-lab cost breakdown with tagging
- ✅ 100% of budget overages detected within 24 hours

**Priority**: HIGH
**Phase**: v0.11.0 - Cost Management for Labs
**Related Personas**: Lab PI, Lab Manager
**Related Issues**: #19, #20, #21
**Related Documentation**: `USER_SCENARIOS/06_LAB_MANAGER_WALKTHROUGH.md` - Pain Point #1

---

### REQ-2.4: Cost Reporting for Grants

**User Story**: As a Lab PI preparing grant renewals, I need to generate cost reports showing AWS spending by project, researcher, and date range, because grant administrators require detailed budget justifications and I currently spend 20 hours manually compiling this data from AWS bills.

**Pain Point**:
- **Before**: Download 12 months of AWS bills → manually filter by project tags → copy into Excel → calculate totals = 20 hours work
- **Grant requirements**: Must show "Personnel: $X, Cloud Compute: $Y, Storage: $Z"
- **Multiple grants**: 3 active grants × 4 quarterly reports = 12 reports/year

**Success Metric**:
- ✅ `lens-jupyter costs --report --from 2025-01-01 --to 2025-12-31 --project NIH-R01-2025 --format pdf`
- ✅ Report includes: total cost, cost by researcher, cost by instance type, monthly trend
- ✅ Export formats: PDF (for grant submissions), CSV (for Excel analysis)
- ✅ Report generation in < 60 seconds
- ✅ 95% time reduction (20 hours → 1 hour)

**Priority**: HIGH
**Phase**: v0.11.0 - Cost Management for Labs
**Related Personas**: Lab PI, Lab Manager
**Related Issues**: #21
**Related Documentation**: `USER_SCENARIOS/02_LAB_PI_WALKTHROUGH.md` - Pain Point #4

---

## 3. Research Tools Requirements

### REQ-3.1: Jupyter Lab Support

**User Story**: As a Data Scientist, I need a full-featured Jupyter Lab environment with common data science packages, because that's my primary analysis tool and I need it to "just work" without manual package installation.

**Pain Point**:
- **Before**: Launch blank EC2 → install Python → install Jupyter → install pandas, matplotlib, scikit-learn, etc. = 3 hours setup
- **Version conflicts**: pip installs create dependency hell
- **Reproducibility**: Colleagues can't recreate my environment

**Success Metric**:
- ✅ 6+ built-in environments (data-science, ml-pytorch, deep-learning, etc.)
- ✅ Each environment lists included packages
- ✅ Launch to working Jupyter Lab in < 3 minutes
- ✅ `data-science` environment includes 50+ common packages (pandas, numpy, scipy, matplotlib, scikit-learn, etc.)

**Priority**: CRITICAL
**Phase**: v0.1.0 - v0.4.0 (COMPLETE)
**Related Personas**: Solo Researcher, Data Scientist, Graduate Student
**Related Issues**: (Implemented)
**Related Documentation**: `USER_SCENARIOS/04_DATA_SCIENTIST_WALKTHROUGH.md` - Pain Point #1

---

### REQ-3.2: RStudio Server Support

**User Story**: As a Social Science Researcher, I need RStudio Server with tidyverse pre-installed, because R is the standard tool in my field and I need to analyze survey data with familiar tools.

**Pain Point**:
- **Before**: R users forced to use Jupyter (unfamiliar interface) or manually set up RStudio Server (complex)
- **Domain mismatch**: Social scientists, economists, biologists primarily use R, not Python
- **Tidyverse essential**: 90% of R workflows depend on tidyverse packages

**Success Metric**:
- ✅ `lens-rstudio launch` provides full RStudio Server
- ✅ Built-in environments: minimal, tidyverse, bioconductor, shiny
- ✅ Feature parity with lens-jupyter (all commands work identically)
- ✅ Tidyverse environment has 100+ R packages pre-installed

**Priority**: CRITICAL
**Phase**: v0.5.0 - Monorepo Stabilization (COMPLETE)
**Related Personas**: Solo Researcher, Graduate Student, Lab PI (R-focused labs)
**Related Issues**: (Implemented)
**Related Documentation**: `USER_SCENARIOS/01_SOLO_RESEARCHER_WALKTHROUGH.md` (if R-focused)

---

### REQ-3.3: VSCode Server Support

**User Story**: As a Graduate Student learning to code, I need VSCode's familiar interface in the cloud, because I've used VSCode Desktop for classes and want the same experience without installing anything on my laptop.

**Pain Point**:
- **Before**: Jupyter notebooks great for data analysis, but not for writing packages, debugging, or multi-file projects
- **Learning curve**: Students already know VSCode from coursework
- **Versatility**: Need general-purpose IDE for Python, R, Julia, shell scripts, etc.

**Success Metric**:
- ✅ `lens-vscode launch` provides code-server (VSCode in browser)
- ✅ Built-in environments: web-dev, python-dev, go-dev, fullstack
- ✅ Extensions auto-installed (Python, R, Jupyter, GitLens, etc.)
- ✅ 100% feature parity with lens-jupyter and lens-rstudio

**Priority**: CRITICAL
**Phase**: v0.6.0 - Testing & Quality (COMPLETE)
**Related Personas**: Graduate Student, Data Scientist, Solo Researcher
**Related Issues**: (Implemented)
**Related Documentation**: `USER_SCENARIOS/03_GRAD_STUDENT_WALKTHROUGH.md` - Pain Point #4

---

### REQ-3.4: Additional Research Tool Support

**User Story**: As a Research Lab PI, I need to support diverse research tools (Streamlit, Zeppelin, Quarto, NICE DCV), because my lab has 5 researchers using different workflows and I need unified management.

**Pain Point**:
- **Before**: Each tool requires different setup process = PI can't support all researchers
- **Tool diversity**: Bioinformatician needs GUI tools (ImageJ), data scientist needs Jupyter, web developer needs Streamlit
- **Fragmentation**: 5 different management approaches = no cost tracking, no standardization

**Success Metric**:
- ✅ Add 6+ additional tools: Amazon Q, Streamlit, Zeppelin, Theia, Quarto, NICE DCV Desktop
- ✅ Unified CLI: `aws-streamlit launch` works identically to `lens-jupyter launch`
- ✅ Shared infrastructure: cost tracking, auto-stop, SSH all work across tools
- ✅ Lab PI can manage all tools from single dashboard

**Priority**: HIGH
**Phase**: v0.8.0 - Additional Research Tools
**Related Personas**: Lab PI, Lab Manager, Data Scientist
**Related Issues**: #5, #6, #7, #8, #9
**Related Documentation**: `USER_SCENARIOS/02_LAB_PI_WALKTHROUGH.md` - Pain Point #3

---

## 4. Reproducibility Requirements

### REQ-4.1: Environment Export and Import

**User Story**: As a Solo Researcher preparing to publish a paper, I need to export my exact environment configuration, because reviewers require reproducible analysis and I need to share my setup in supplementary materials.

**Pain Point**:
- **Before**: Paper accepted → reviewer asks "Can you share your analysis environment?" → no easy way to export exact package versions
- **Reproducibility crisis**: 70% of published analyses can't be reproduced due to missing environment details
- **Manual export**: `pip freeze > requirements.txt` doesn't capture system packages, R packages, or configurations

**Success Metric**:
- ✅ `lens-jupyter env export` creates complete environment snapshot
- ✅ Export includes: Python/R packages + versions, system packages, Jupyter extensions, environment variables
- ✅ `lens-jupyter env import environment.yml` recreates identical environment
- ✅ Published environments can be reproduced 5 years later
- ✅ 95% package version match on import

**Priority**: HIGH
**Phase**: v0.9.0 - Package Managers & Reproducible Environments
**Related Personas**: Solo Researcher, Data Scientist, Graduate Student
**Related Issues**: #12
**Related Documentation**: `USER_SCENARIOS/01_SOLO_RESEARCHER_WALKTHROUGH.md` - Pain Point #4

---

### REQ-4.2: Conda and BioConda Integration

**User Story**: As a Bioinformatics Researcher, I need BioConda packages for genomics analysis, because 500+ bioinformatics tools are only available via BioConda and manual compilation takes days.

**Pain Point**:
- **Before**: Install GATK, BWA, SAMtools manually → dependency conflicts → version mismatches → analysis fails
- **Compilation time**: Building genomics tools from source = 8-12 hours
- **Domain standard**: BioConda is the de facto package manager in bioinformatics

**Success Metric**:
- ✅ Full conda environment support with `environment.yml` files
- ✅ BioConda channel enabled by default in bioinformatics environments
- ✅ Pre-configured environments: genomics-pipeline, rna-seq, single-cell
- ✅ `lens-jupyter launch --env genomics-pipeline` includes 50+ BioConda tools
- ✅ Package caching reduces subsequent launch time by 70%

**Priority**: HIGH
**Phase**: v0.9.0 - Package Managers & Reproducible Environments
**Related Personas**: Solo Researcher (bioinformatics focus), Data Scientist
**Related Issues**: #10, #11
**Related Documentation**: New walkthrough needed for bioinformatics persona

---

### REQ-4.3: Domain-Specific Environment Templates

**User Story**: As a Social Science PhD Student, I need a pre-configured environment for survey analysis (pandas, statsmodels, matplotlib), because I don't know which packages I need and want to start analyzing data immediately.

**Pain Point**:
- **Before**: Ask advisor "What packages do I need?" → install 20 packages one by one → discover missing dependencies → reinstall
- **Domain knowledge**: New PhD students don't know field-standard tools
- **Time waste**: 4-6 hours setup before doing any actual research

**Success Metric**:
- ✅ 10+ domain-specific templates by v1.0.0:
  - `genomics-pipeline` (biology)
  - `rna-seq` (biology)
  - `survey-analysis` (social science)
  - `text-analysis` (social science)
  - `econometrics` (economics)
  - `time-series` (economics/finance)
  - `climate-analysis` (climate science)
  - `geospatial` (geography/climate)
  - `physics-sim` (physics/engineering)
- ✅ Each template documented with typical use cases
- ✅ Community can contribute templates
- ✅ 80% of users find a suitable template without customization

**Priority**: HIGH
**Phase**: v0.9.0 - Package Managers & Reproducible Environments
**Related Personas**: Graduate Student, Solo Researcher
**Related Issues**: #13
**Related Documentation**: `USER_SCENARIOS/03_GRAD_STUDENT_WALKTHROUGH.md` - Pain Point #5

---

## 5. Collaboration Requirements

### REQ-5.1: Instance Sharing with Lab Members

**User Story**: As a Graduate Student, I need to share my running Jupyter instance with my advisor for a 30-minute code review, because screen-sharing doesn't work well and my advisor needs to interact with my notebook.

**Pain Point**:
- **Before**: Advisor wants to review analysis → export notebook → email → advisor imports → different environment → code doesn't run
- **Zoom limitations**: Screen-sharing for code review is slow and frustrating
- **Security**: Can't give advisor my AWS credentials

**Success Metric**:
- ✅ `lens-jupyter share i-abc123 --email advisor@university.edu --duration 2h --read-only`
- ✅ Advisor receives email with one-click access link
- ✅ Read-only mode prevents accidental changes
- ✅ Access automatically expires after duration
- ✅ Access logs show who accessed when

**Priority**: MEDIUM
**Phase**: v0.10.0 - Collaboration Features
**Related Personas**: Graduate Student, Lab PI, Data Scientist
**Related Issues**: #15
**Related Documentation**: `USER_SCENARIOS/03_GRAD_STUDENT_WALKTHROUGH.md` - Pain Point #6

---

### REQ-5.2: Lab-Wide Configuration Templates

**User Story**: As a Lab PI, I need to set default configurations for my entire lab (region, instance types, environments, budgets), because I'm managing 5 students and want consistent, cost-controlled setups without micromanaging.

**Pain Point**:
- **Before**: Each student uses different region, instance types → inconsistent costs, hard to track
- **Budget control**: Student launches p3.8xlarge GPU instance by accident → $12/hour
- **Onboarding**: Each new student requires 1-hour configuration session with PI

**Success Metric**:
- ✅ PI creates lab config: `lens-jupyter config init --lab chen-physics-lab`
- ✅ Config sets: allowed instance types, max instance size, default region, approved environments, budget per student
- ✅ Students inherit lab config: `lens-jupyter launch --lab chen-physics-lab`
- ✅ Students can't launch unapproved instance types
- ✅ 90% reduction in PI configuration time

**Priority**: MEDIUM
**Phase**: v0.10.0 - Collaboration Features
**Related Personas**: Lab PI, Lab Manager
**Related Issues**: #16
**Related Documentation**: `USER_SCENARIOS/02_LAB_PI_WALKTHROUGH.md` - Pain Point #5

---

### REQ-5.3: S3 Data Sync and Backup

**User Story**: As a Data Scientist, I need automatic S3 sync of my project folders, because I've lost work twice when instances were accidentally terminated and manual backups are unreliable.

**Pain Point**:
- **Before**: Work on analysis for 8 hours → instance terminated by mistake → lose all unsaved work
- **Manual backup**: Remember to `scp` files locally = unreliable, slow
- **Data loss**: 40% of researchers report losing work due to instance termination

**Success Metric**:
- ✅ Configure S3 sync: `lens-jupyter config set s3_sync_bucket my-research-data`
- ✅ Auto-sync every 15 minutes: `/home/ubuntu/projects/ → s3://my-research-data/projects/`
- ✅ On-demand sync: `lens-jupyter sync now`
- ✅ Automatic restore on launch: detect previous synced data → ask "Restore previous work?"
- ✅ Version history: rollback to sync from 1 hour ago, 1 day ago, etc.
- ✅ 100% data loss prevention

**Priority**: HIGH
**Phase**: v0.10.0 - Collaboration Features
**Related Personas**: All personas
**Related Issues**: #17, #18
**Related Documentation**: `USER_SCENARIOS/04_DATA_SCIENTIST_WALKTHROUGH.md` - Pain Point #4

---

## 6. Security & Compliance Requirements

### REQ-6.1: Session Manager (No SSH Keys)

**User Story**: As a Lab Manager at a university, I need to comply with institutional security policies that prohibit SSH keys, because our IT security office requires all remote access to use centrally managed authentication.

**Pain Point**:
- **Before**: SSH key management is security risk → keys stolen, shared, or leaked → unauthorized access
- **Compliance**: Universities require MFA, audit logs, centralized access control
- **Key sprawl**: 5 students × 3 instance types = 15 SSH keys to manage

**Success Metric**:
- ✅ Session Manager enabled by default: `lens-jupyter launch --connection session-manager`
- ✅ No SSH keys created or stored
- ✅ Uses IAM authentication (existing university IAM integration)
- ✅ All connections logged in CloudTrail for audits
- ✅ Works with university MFA requirements
- ✅ 100% of security audit requirements met

**Priority**: HIGH
**Phase**: v0.5.0 - Monorepo Stabilization (COMPLETE)
**Related Personas**: Lab Manager, Lab PI
**Related Issues**: (Implemented)
**Related Documentation**: `USER_SCENARIOS/06_LAB_MANAGER_WALKTHROUGH.md` - Pain Point #2

---

### REQ-6.2: Private Subnet Support

**User Story**: As a Lab Manager handling sensitive health data, I need instances in private subnets with no public IP addresses, because HIPAA compliance requires research computing infrastructure to be isolated from the public internet.

**Pain Point**:
- **Before**: Public IP addresses on instances = potential attack surface
- **Compliance requirements**: HIPAA, FERPA, IRB data protection policies require network isolation
- **Manual NAT Gateway**: Setting up NAT Gateways correctly is complex

**Success Metric**:
- ✅ Launch in private subnet: `lens-jupyter launch --subnet-type private`
- ✅ Automatic NAT Gateway creation if needed
- ✅ No public IP assigned to instance
- ✅ Access via Session Manager (works without public IP)
- ✅ Egress for package installation via NAT Gateway
- ✅ 100% HIPAA network isolation requirements met

**Priority**: MEDIUM
**Phase**: v0.5.0 - Monorepo Stabilization (COMPLETE)
**Related Personas**: Lab Manager, Lab PI (sensitive data)
**Related Issues**: (Implemented)
**Related Documentation**: `USER_SCENARIOS/06_LAB_MANAGER_WALKTHROUGH.md` - Pain Point #3

---

### REQ-6.3: Audit Logging for Compliance

**User Story**: As a Lab Manager preparing for a security audit, I need detailed logs of all instance launches, terminations, and access, because our IRB requires demonstrating who accessed research data and when.

**Pain Point**:
- **Before**: Auditor asks "Who had access to sensitive dataset on March 15?" → no records → 20 hours reconstructing timeline from memory
- **Compliance gaps**: IRBs require audit trails for human subjects research
- **Accountability**: Can't prove compliance = research project shutdown

**Success Metric**:
- ✅ All AWS API calls logged to CloudTrail (automatic)
- ✅ `lens-jupyter audit --from 2025-03-01 --to 2025-03-31` generates report
- ✅ Report shows: who launched what, when, from where (IP), what actions taken
- ✅ Export formats: PDF (for auditors), CSV (for analysis)
- ✅ 95% time reduction in audit prep (20 hours → 1 hour)

**Priority**: MEDIUM
**Phase**: v0.7.0 - Security & Compliance (future)
**Related Personas**: Lab Manager, Lab PI
**Related Issues**: #29
**Related Documentation**: `USER_SCENARIOS/06_LAB_MANAGER_WALKTHROUGH.md` - Pain Point #4

---

## 7. Performance Requirements

### REQ-7.1: Fast Launch Times

**User Story**: As a Course Instructor starting class in 5 minutes, I need my demo environment to launch in under 3 minutes, because I can't have students waiting 10 minutes while I set up.

**Pain Point**:
- **Before**: Launch → wait 5-10 minutes → students lose focus → awkward silence
- **Unreliability**: Sometimes takes 2 minutes, sometimes takes 8 minutes → can't plan
- **First impression**: Slow launch makes tool seem unreliable

**Success Metric**:
- ✅ Average launch time < 3 minutes (from command to Jupyter accessible)
- ✅ 95% of launches complete in < 4 minutes
- ✅ Progress indicators show ETA
- ✅ Streaming cloud-init logs provide feedback
- ✅ SSM readiness polling (2-3 minutes typical)

**Priority**: HIGH
**Phase**: v0.6.0 - Testing & Quality (COMPLETE)
**Related Personas**: Course Instructor, All personas
**Related Issues**: (Implemented - SSM polling, progress streaming)
**Related Documentation**: `USER_SCENARIOS/05_INSTRUCTOR_WALKTHROUGH.md` - Pain Point #2

---

### REQ-7.2: Responsive CLI Commands

**User Story**: As a Data Scientist checking instance status, I need `lens-jupyter status` to respond in under 2 seconds, because I check status frequently throughout the day and slow commands break my flow.

**Pain Point**:
- **Before**: `lens-jupyter status` takes 5-8 seconds → disrupts workflow
- **Impatience**: Users assume command failed and hit Ctrl+C
- **Frequency**: Status checked 10-20 times per session

**Success Metric**:
- ✅ All CLI commands respond in < 2 seconds
- ✅ `status`, `list`, `info` use cached data where possible
- ✅ Long operations (launch, terminate) show immediate acknowledgment + background progress
- ✅ 100% of commands under 2 seconds (excluding launch/terminate operations)

**Priority**: MEDIUM
**Phase**: v1.0.0 - Production Ready
**Related Personas**: All personas
**Related Issues**: #30
**Related Documentation**: (All walkthroughs - implicit expectation)

---

### REQ-7.3: GPU Support for ML Workloads

**User Story**: As a Graduate Student training deep learning models, I need GPU instances (g4dn, p3, p4) with CUDA pre-configured, because training on CPU takes 20 hours vs 2 hours on GPU and I have deadlines.

**Pain Point**:
- **Before**: CPU training is 10x slower → miss paper deadlines → research progress stalled
- **Setup complexity**: Installing CUDA drivers manually = 4-6 hours troubleshooting
- **Cost optimization**: GPU instances expensive → need auto-stop even more critical

**Success Metric**:
- ✅ GPU instance types supported: g4dn (affordable), p3 (performance), p4 (latest)
- ✅ CUDA pre-installed and configured in GPU environments
- ✅ `lens-jupyter launch --instance-type g4dn.xlarge --env deep-learning-gpu`
- ✅ PyTorch/TensorFlow automatically detect GPU
- ✅ Cost warnings for expensive GPU instances
- ✅ Auto-stop essential (GPU idle = $1-12/hour waste)

**Priority**: MEDIUM
**Phase**: v0.8.0 - Additional Research Tools (backlog)
**Related Personas**: Graduate Student, Data Scientist
**Related Issues**: #28
**Related Documentation**: `USER_SCENARIOS/03_GRAD_STUDENT_WALKTHROUGH.md` - Pain Point #7

---

## Summary: Priority Matrix

### CRITICAL (Ship Blockers)
| ID | Requirement | Phase | Status |
|----|-------------|-------|--------|
| REQ-1.1 | Beginner-Friendly Onboarding | v0.7.0 | In Progress |
| REQ-1.2 | Plain-English Error Messages | v0.7.0 | ✅ Complete |
| REQ-2.1 | Cost Preview Before Launch | v0.7.0 | Planned |
| REQ-2.2 | Automatic Idle Shutdown | v0.6.0 | ✅ Complete |
| REQ-3.1 | Jupyter Lab Support | v0.1.0 | ✅ Complete |
| REQ-3.2 | RStudio Server Support | v0.5.0 | ✅ Complete |
| REQ-3.3 | VSCode Server Support | v0.6.0 | ✅ Complete |

### HIGH (Major Value)
| ID | Requirement | Phase | Status |
|----|-------------|-------|--------|
| REQ-1.3 | Wizard as Default | v0.7.0 | Planned |
| REQ-1.4 | Remember Preferences | v0.7.0 | Planned |
| REQ-1.5 | Quickstart Command | v0.7.0 | Planned |
| REQ-2.3 | Budget Tracking | v0.11.0 | Planned |
| REQ-2.4 | Cost Reporting | v0.11.0 | Planned |
| REQ-3.4 | Additional Research Tools | v0.8.0 | Planned |
| REQ-4.1 | Environment Export/Import | v0.9.0 | Planned |
| REQ-4.2 | Conda/BioConda | v0.9.0 | Planned |
| REQ-4.3 | Domain Templates | v0.9.0 | Planned |
| REQ-5.3 | S3 Sync and Backup | v0.10.0 | Planned |
| REQ-6.1 | Session Manager | v0.5.0 | ✅ Complete |
| REQ-7.1 | Fast Launch Times | v0.6.0 | ✅ Complete |

### MEDIUM (Nice to Have)
| ID | Requirement | Phase | Status |
|----|-------------|-------|--------|
| REQ-5.1 | Instance Sharing | v0.10.0 | Planned |
| REQ-5.2 | Lab Config Templates | v0.10.0 | Planned |
| REQ-6.2 | Private Subnets | v0.5.0 | ✅ Complete |
| REQ-6.3 | Audit Logging | Future | Backlog |
| REQ-7.2 | Responsive CLI | v1.0.0 | Planned |
| REQ-7.3 | GPU Support | v0.8.0 | Backlog |

---

## Document Maintenance

**Update Triggers**:
- New persona walkthrough created → extract requirements
- User feedback indicates missing requirement → add and prioritize
- Feature implemented → mark status as "Complete"
- Roadmap phase changes → update phase assignments

**Review Cadence**:
- Weekly during active development (v0.7.0 - v0.9.0)
- Monthly during maintenance (post-v1.0.0)

**Document Owners**:
- **Primary**: Project Lead
- **Contributors**: All team members can propose requirements
- **Approvers**: Project Lead + at least one persona representative (actual researcher)

---

**Next Steps**:
1. Create `docs/USER_SCENARIOS/` with 6 persona walkthroughs referencing these requirements
2. Create GitHub issues for each requirement (30+ issues)
3. Link issues back to this document
4. Update ROADMAP.md to reference requirement IDs
