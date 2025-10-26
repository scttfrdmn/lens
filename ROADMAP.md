# Lens Roadmap

This document outlines the future development plans for the Lens project. The project follows semantic versioning and is organized into phases.

## ⚠️ Version Number Clarification

**Important**: Roadmap phase numbers (v0.7.0, v0.8.0, etc.) represent feature planning phases and may not match actual release version numbers due to development priorities and release timing.

**Current Status**:
- v0.8.0 release = Rebranding (AWS IDE → Lens)
- v0.9.0 release = Completed v0.7.0 roadmap goals (User Experience)
- v0.10.0 planned = Will address v0.8.0 roadmap goals (Additional Tools)

We maintain flexibility to ship features when ready, not strictly by roadmap order.

---

## 📊 Current Status: Production-Ready Platform

**Project Evolution:**
- ✅ **v0.1.0-v0.4.0**: lens-jupyter (single app) - Feature complete
- ✅ **v0.5.0**: Monorepo established with RStudio feature parity
- ✅ **v0.6.0**: Comprehensive testing infrastructure
- ✅ **v0.6.1**: Unified config system and cost tracking
- ✅ **v0.6.2**: Full feature parity - config and costs commands for all tools
- ✅ **v0.6.3**: Documentation polish and updates
- 🚀 **Current Focus**: User experience for non-technical academic researchers

**Target Audience:** Academic researchers who need cloud-based analysis tools but may not be technically savvy

**Monorepo Structure:**
- `pkg/` - Shared AWS infrastructure library (well-tested)
- `apps/jupyter/` - Jupyter Lab launcher (production ready)
- `apps/rstudio/` - RStudio Server launcher (production ready)
- `apps/vscode/` - VSCode Server launcher (production ready)

**Coverage Status:**
- pkg: Comprehensive unit, integration, and E2E tests
- All three apps: Production ready with full feature parity
- Testing: Unit → Integration → Smoke → E2E test pyramid

---

## 🎯 v0.5.0 - Monorepo Stabilization (Current - Q1 2025)

**Status:** In Progress
**Target Date:** February 2025
**Priority:** CRITICAL

### Goals
Complete the monorepo migration and establish RStudio as a first-class citizen alongside Jupyter.

### Tasks

**Completed:**
- [x] Transform codebase into Go workspace
- [x] Extract shared infrastructure into pkg/
- [x] Create apps/jupyter with full feature set
- [x] Create apps/rstudio with basic implementation
- [x] Fix all test failures
- [x] Update CI/CD for monorepo structure
- [x] Remove legacy code and duplication
- [x] Update documentation for monorepo
- [x] **SSM-based readiness polling** for secure service health checks
  - [x] Created pkg/aws/ssm.go with SSMClient implementation
  - [x] Added PollServiceReadinessViaSSM() to pkg/readiness/poller.go
  - [x] Migrated VSCode to SSM polling (port 8080)
  - [x] Migrated Jupyter to SSM polling (port 8888)
  - [x] Migrated RStudio to SSM polling (port 8787)
  - [x] All apps tested end-to-end with SSM readiness checks
  - [x] Works regardless of security group configuration
- [x] **Progress streaming improvements**
  - [x] Real-time cloud-init progress via SSH
  - [x] Concurrent progress streaming and SSM polling
  - [x] Enhanced user experience during launch

**Completed:**
- [x] **RStudio Feature Parity**
  - [x] Ported all 10 commands from Jupyter to RStudio
  - [x] Created 4 R-specific environments (minimal, tidyverse, bioconductor, shiny)
  - [x] Added 933 lines of unit tests (27 test functions)
  - [x] Fixed port configuration bug (8888 → 8787)
  - [x] Updated README with current capabilities
  - [x] Verified all commands work end-to-end

**Completed:**
- [x] **Shared Infrastructure Testing** (Value-Focused Approach)
  - [x] Unit tests for pure functions (formatDuration, cleanupStateFile, etc.)
  - [x] Config management fully tested (environment, state, keys)
  - [x] Integration tests for AWS operations (EC2, IAM, AMI, SSM)
  - [x] E2E tests for complete workflows (all 3 IDE types)
  - [x] Testing pyramid properly balanced: unit → integration → E2E

**Completed:**
- [x] Documentation
  - [x] Updated CHANGELOG with all v0.5.0 changes
  - [x] Migration guide not needed (no users yet, solo dev project)
  - [x] RStudio README comprehensively updated
  - [x] ROADMAP reflects current state and value-focused testing

### Success Criteria
- ✅ Both apps build and test successfully
- ✅ CI/CD pipeline working for monorepo
- ✅ RStudio has feature parity with Jupyter
- ✅ Shared library appropriately tested (value over coverage metrics)
- ✅ Complete documentation for both apps

---

## 🧪 v0.6.0 - Testing & Quality (Q3-Q4 2025)

**Status:** ✅ Complete
**Completion Date:** October 2025
**Priority:** HIGH

### Goals
Establish comprehensive testing infrastructure with unit, integration, and E2E tests.

### Tasks

**Completed:**
- [x] **Test Infrastructure**
  - [x] Created comprehensive Makefile with test targets (unit, integration, smoke, e2e, regression)
  - [x] Set up LocalStack Docker Compose configuration for AWS mocking
  - [x] Added build tags for test separation (-tags=integration, -tags=smoke, -tags=e2e)
  - [x] Configured appropriate timeouts for each test type

- [x] **Unit Tests for SSM and Readiness**
  - [x] pkg/aws/ssm_test.go (354 lines, 14 test functions)
  - [x] pkg/readiness/poller_test.go (461 lines, 12 test functions)
  - [x] Tests for CommandResult struct, ServiceConfig, polling configuration

- [x] **Integration Tests with LocalStack**
  - [x] pkg/aws/ec2_integration_test.go (602 lines, 12 test functions)
  - [x] Tests for EC2Client and SSMClient initialization
  - [x] Verified LocalStack compatibility
  - [x] Integration tests run with `make test-integration`
  - [x] EC2 LaunchInstance tests (valid params, auto-subnet, invalid params)
  - [x] Instance operations tests (start/stop/terminate/get info)
  - [x] AMI operations tests (create/list/delete)
  - [x] IAM role tests (GetOrCreateSessionManagerRole with idempotency)
  - [x] Error handling tests (non-existent resources, invalid params)
  - [x] AMI selector integration tests (multiple OS/arch combinations)

- [x] **Smoke Tests (Quick Real AWS Validation)**
  - [x] pkg/aws/ec2_smoke_test.go (370 lines, 7 test functions)
  - [x] Basic EC2 client creation and connectivity
  - [x] IAM role verification (Session Manager)
  - [x] Subnet discovery (public subnets)
  - [x] Instance type availability checks
  - [x] AMI discovery (Ubuntu 24.04, Amazon Linux)
  - [x] Availability zone compatibility
  - [x] Quick launch prerequisite validation
  - [x] All tests passing with AWS_PROFILE=aws
  - [x] Run with `make test-smoke`

- [x] **End-to-End Tests**
  - [x] pkg/e2etest/helpers.go (500+ lines) - Shared E2E test infrastructure
  - [x] Complete launch → connect → terminate flow (Jupyter) - 4 test functions
  - [x] Complete launch → connect → terminate flow (RStudio) - 4 test functions
  - [x] Complete launch → connect → terminate flow (VSCode) - 5 test functions
  - [x] Test Session Manager connection methods (UseSessionMgr flag)
  - [x] Test environment generation and customization (MultipleEnvironments tests)
  - [x] Instance lifecycle testing (stop/start operations)
  - [x] Multi-architecture testing (ARM64 Graviton + x86)
  - [x] Custom port configuration testing
  - [x] SSM-based readiness polling (no security group dependency)
  - [x] Automatic cleanup with defer statements
  - [x] Run with `make test-e2e`

- [x] **Documentation**
  - [x] Updated TESTING.md with LocalStack setup instructions
  - [x] Documented testing strategy pyramid (unit → integration → smoke → E2E)
  - [x] Added LocalStack vs Moto comparison
  - [x] Documented SSM readiness polling end-to-end test results

**Current Coverage Status:**
- pkg/aws: 2.5% (struct tests only; most functions require AWS/LocalStack)
- pkg/cli: 0.0% (Cobra commands; hard to unit test without extensive mocking)
- pkg/config: 84.7% (✅ already well-tested)
- pkg/readiness: 0.0% (struct tests only; polling functions require network)
- apps/jupyter/cli: 27.8%
- apps/rstudio/cli: 0.0%

**Testing Strategy:**
- **Unit Tests**: Structs, data validation, pure functions (no AWS/network)
- **Integration Tests**: AWS API calls via LocalStack (no real AWS costs)
- **Smoke Tests**: Quick real AWS validation (<10 min, minimal cost)
- **E2E Tests**: Full workflows with real AWS (longer, higher cost)

### Success Criteria
- ✅ Comprehensive Makefile with all test types
- ✅ LocalStack integration working
- ✅ Unit tests for testable modules
- ✅ Smoke tests for quick validation
- ✅ Integration tests covering major AWS operations
- ✅ E2E tests for all three apps (Jupyter, RStudio, VSCode)
- ✅ CI/CD running unit + integration tests automatically

---

## ✨ v0.7.0 Roadmap Phase - User Experience & Accessibility

**Status:** ✅ Complete (Shipped in v0.9.0 Release)
**Completion Date:** October 2025
**Priority:** 🔥 CRITICAL - Non-technical researchers need this

**Note**: This roadmap phase was completed and released as v0.9.0 (after v0.8.0 rebranding release).

### Goals
Make Lens accessible to non-technical academic researchers with guided setup and plain-language interactions.

### Tasks

**Interactive Launch Wizard** 🎯 HIGHEST PRIORITY ✅ COMPLETE
- [x] Add `wizard` command that guides users through setup
  - [x] "What type of analysis do you want to do?"
    - Data science (Python + R)
    - Statistical analysis (R-focused)
    - Machine learning / Deep learning
    - Web development (VS Code)
  - [x] Plain language options (no technical jargon)
  - [x] Automatic "best practice" selections with recommendations
  - [x] Show cost estimate before launching (hourly + monthly)
  - [x] Implemented for all 3 apps (Jupyter, RStudio, VSCode)
  - [x] Uses survey library for interactive prompts
  - [x] Auto-stop configuration with idle timeout options
  - [x] Optional instance naming
- [x] Add `quickstart` command for each app
  - [x] `lens-jupyter quickstart` → instant launch with sensible defaults
  - [x] `lens-rstudio quickstart` → instant R environment
  - [x] `lens-vscode quickstart` → instant coding environment

**Better Error Messages** 🎯 HIGH PRIORITY ✅ COMPLETE
- [x] Created pkg/errors package for beginner-friendly error handling
  - [x] FriendlyError struct with Title, Explanation, NextSteps
  - [x] Emoji indicators for clarity (❌ error, 💡 next steps, 🔍 technical details)
  - [x] Plain-English error messages (no AWS jargon)
- [x] Common error patterns implemented:
  - [x] AWS credentials → "Can't connect to AWS" with setup instructions
  - [x] Instance states → Clear explanations ("turned off", "starting up")
  - [x] Permission errors → "You don't have permission to do that"
  - [x] Region mismatches → "Wrong AWS region" with switching instructions
  - [x] Quotas → "Too many instances running" with cleanup guidance
  - [x] Network issues → "Connection timed out" with troubleshooting
  - [x] Missing environments → Suggestions to use wizard or list envs
- [x] TranslateAWSError() automatically converts technical errors
- [x] Next steps with specific commands users can run
- [ ] `doctor` command (deferred - errors package provides diagnostic info)

**Enhanced Visual Output** 🎯 HIGH PRIORITY ✅ COMPLETE
- [x] Color-coded output (success=green, warning=yellow, error=red)
- [x] Progress infrastructure with plain text
  - [x] Created pkg/output package with progress bars and ETA
  - [x] Estimated time messages ("this should take 2-3 minutes")
- [x] Friendly completion messages
  - [x] "✓ Your Jupyter Lab is ready!"
  - [x] Structured connection instructions
  - [x] Clear next steps
- [x] Remove technical AWS jargon from all output
  - [x] "Starting environment" not "EC2 RunInstances API call"
  - [x] "Waiting for environment to boot up" not "Polling SSM"
  - [x] Transformed all 27 functions across 3 apps (VSCode, Jupyter, RStudio)

**Status Notifications** ⚠️ DEFERRED (Low Priority)
- [x] Notification hooks infrastructure complete
  - [x] Generic hook system for all lifecycle events
  - [x] Environment variables passed to hooks
  - [x] Documentation with Slack/email/desktop examples (docs/NOTIFICATIONS.md)
  - [x] Users can configure custom notification scripts
- [ ] Built-in email notifications (deferred to v0.8+)
  - [ ] "Your Jupyter Lab is ready at: http://..."
  - [ ] "Your instance will stop in 30 minutes due to idle timeout"
- [ ] Built-in Slack webhook support (deferred to v0.8+)
- [ ] Desktop notifications (investigate feasibility)

**Default Behavior Improvements** ✅ COMPLETE
- [x] Make wizard the default when no arguments provided
  - [x] `lens-jupyter` → launches wizard
  - [x] `lens-jupyter launch` → uses last settings or defaults
- [x] Remember user preferences
  - [x] Last used environment, instance type, region
  - [x] Offer to reuse settings: "Use same settings as last time? (Y/n)"
  - [x] Preferences stored in `~/.lens/{app}-preferences.json`
  - [x] Quick relaunch skips wizard questions when reusing settings

### Success Criteria
- ✅ Non-technical researcher can launch Jupyter in <2 minutes without reading docs
- ✅ All error messages are in plain English with clear next steps
- ✅ Progress is visible and understandable
- ✅ No AWS/technical jargon in user-facing output

---

## 🔬 v0.10.0 - GUI Foundation & Tool Expansion (was v0.8.0 roadmap phase)

**Status:** Next - Planned for v0.10.0 Release
**Target:** 2-3 months
**Priority:** 🔥 HIGH - Researchers need variety of tools + GUI support

**📖 See Also**: `docs/TOOL_SUPPORT_STRATEGY.md` for comprehensive tool expansion plan

### Goals
1. Enable GUI tool support with NICE DCV
2. Add web-based research tools (Streamlit, Zeppelin, Q Developer)
3. Establish foundation for future GUI tools (MATLAB, ArcGIS, QGIS, etc.)

### Applications to Add

**Amazon Q Developer** 🎯 HIGHEST PRIORITY
- [ ] Create `apps/q-developer/` (standalone Amazon Q IDE)
- [ ] AI-assisted coding environment for researchers
- [ ] Code suggestions, explanations, and documentation
- [ ] Perfect for researchers learning to code
- [ ] AWS native - seamless integration
- [ ] Built-in environments:
  - [ ] `research-coding` - Python + R with Q assistance
  - [ ] `data-analysis` - Data science with AI help
  - [ ] `learning` - Beginner-friendly with Q explanations
- [ ] **Already added to VSCode** - Q Developer extension included in all VSCode environments

**Streamlit** 🎯 HIGHEST PRIORITY
- [ ] Create `apps/streamlit/`
- [ ] Turn Python scripts into interactive web apps
- [ ] Perfect for sharing analysis with non-programmers
- [ ] Excellent for academic collaboration
- [ ] Built-in environments:
  - [ ] `data-viz` - Data visualization apps
  - [ ] `ml-demo` - Machine learning demos
  - [ ] `dashboard` - Analysis dashboards

**Apache Zeppelin** 🎯 HIGH PRIORITY
- [ ] Create `apps/zeppelin/`
- [ ] Multi-language notebooks (Python, R, Scala, SQL)
- [ ] Built-in visualization and charting
- [ ] Popular in big data research
- [ ] Built-in environments:
  - [ ] `data-engineering` - ETL and data processing
  - [ ] `sql-analytics` - Database analysis
  - [ ] `spark` - Big data processing

**Theia IDE** 🎯 MEDIUM PRIORITY
- [ ] Create `apps/theia/`
- [ ] Full IDE similar to VSCode but more extensible
- [ ] Good for researchers who code occasionally
- [ ] Built-in environments:
  - [ ] `python` - Python development
  - [ ] `r-dev` - R package development
  - [ ] `multi-lang` - Multiple languages

**Quarto** 🎯 MEDIUM PRIORITY
- [ ] Create `apps/quarto/`
- [ ] Academic publishing platform
- [ ] Create papers, presentations, websites from code
- [ ] Works with Jupyter and R
- [ ] Important for reproducible research
- [ ] Built-in environments:
  - [ ] `academic-paper` - LaTeX + code
  - [ ] `presentation` - Reveal.js slides
  - [ ] `website` - Research group websites

**Observable Framework** 🎯 LOW PRIORITY
- [ ] Create `apps/observable/`
- [ ] Interactive JavaScript notebooks
- [ ] Excellent for data visualization
- [ ] Used in data journalism and research
- [ ] Built-in environments:
  - [ ] `data-viz` - D3.js visualizations
  - [ ] `interactive` - Interactive analysis

**NICE DCV Desktop** 🎯 HIGH PRIORITY
- [ ] Create `apps/dcv-desktop/` (full Linux desktop via DCV)
- [ ] AWS native remote desktop protocol (NICE DCV)
- [ ] Low-latency, high-quality desktop streaming
- [ ] Essential for GUI applications (MATLAB, Igor Pro, ImageJ, etc.)
- [ ] GPU support for visualization and computation
- [ ] Built-in environments:
  - [ ] `matlab-desktop` - MATLAB with full GUI
  - [ ] `data-viz-desktop` - ParaView, Visit, Tableau Desktop
  - [ ] `image-analysis` - ImageJ, Fiji, QuPath, CellProfiler
  - [ ] `bioinformatics-gui` - Geneious, CLC Workbench, UGENE
  - [ ] `general-desktop` - Ubuntu desktop with research tools
  - [ ] `gpu-workstation` - CUDA, visualization, ML (GPU instances)
- [ ] Benefits:
  - [ ] No need for X11 forwarding or VNC
  - [ ] Works from browser or native client
  - [ ] Clipboard and file transfer support
  - [ ] Multi-monitor support
  - [ ] Better than traditional remote desktop protocols

### Implementation Order (v0.10.0)
1. **NICE DCV Desktop** (critical foundation for GUI apps)
2. **Streamlit** (most requested web tool)
3. **Amazon Q Developer** (AWS native, already partially integrated)
4. **Zeppelin** (fills notebook alternative niche)

### Success Criteria
- ✅ NICE DCV desktop working with browser and native client access
- ✅ At least 2-3 new web-based tools added
- ✅ Users can manually install GUI tools on DCV desktop
- ✅ GPU acceleration working for compatible tools
- ✅ Documentation for each tool with academic use cases

---

## 📦 v0.11.0 - Package Managers & Reproducible Environments (was v0.9.0 roadmap phase)

**Status:** Planned for v0.11.0 Release
**Target:** 3-4 months
**Priority:** MEDIUM-HIGH - Reproducibility is critical for research

### Goals
Make it easy for researchers to create reproducible, domain-specific environments.

### Tasks

**Complete Conda Integration** 🎯 HIGH PRIORITY
- [ ] Full conda environment support
  - [ ] Support environment.yml files
  - [ ] Automatic conda environment activation
  - [ ] Package caching to speed up launches
- [ ] BioConda integration for bioinformatics
  - [ ] Pre-configured bioinformatics environments
  - [ ] Common genomics tools pre-installed
- [ ] Conda forge channel support
- [ ] Environment export/import
  - [ ] `lens-jupyter env export` → saves environment.yml
  - [ ] `lens-jupyter env import environment.yml` → recreates environment

**System Package Management** 🎯 MEDIUM PRIORITY
- [ ] Declarative package installation in environments
  - [ ] apt packages (Ubuntu/Debian)
  - [ ] yum packages (Amazon Linux)
- [ ] Custom package lists per environment
  - [ ] `system_packages:` section in environment YAML
  - [ ] Version pinning support
- [ ] Package caching for faster rebuilds

**Domain-Specific Environment Templates** 🎯 HIGH PRIORITY
- [ ] **Biology/Genomics**
  - [ ] Genomics pipeline (GATK, BWA, SAMtools)
  - [ ] RNA-seq analysis (DESeq2, edgeR)
  - [ ] Single-cell analysis (Seurat, Scanpy)
- [ ] **Social Sciences**
  - [ ] Survey analysis (pandas, statsmodels)
  - [ ] Text analysis (NLTK, spaCy)
  - [ ] Network analysis (NetworkX, igraph)
- [ ] **Economics/Finance**
  - [ ] Econometrics (statsmodels, linearmodels)
  - [ ] Time series (prophet, ARIMA)
  - [ ] Financial modeling (QuantLib)
- [ ] **Climate Science**
  - [ ] Climate data analysis (xarray, iris)
  - [ ] Geospatial tools (GDAL, rasterio)
  - [ ] Visualization (cartopy, proplot)
- [ ] **Physics/Engineering**
  - [ ] Scientific computing (NumPy, SciPy)
  - [ ] Symbolic math (SymPy, Maxima)
  - [ ] Simulation tools (OpenFOAM, FEniCS)

**Environment Sharing** 🎯 MEDIUM PRIORITY
- [ ] Export environments to shareable format
  - [ ] Include all packages and versions
  - [ ] Include system packages
  - [ ] Include custom configurations
- [ ] Import environments from collaborators
- [ ] Community environment repository
  - [ ] Share environments with research community
  - [ ] Upvote/review system
  - [ ] Search by research domain

**Easy Package Installation Commands**
- [ ] `lens-jupyter packages install pandas matplotlib seaborn`
- [ ] `lens-rstudio packages install tidyverse ggplot2`
- [ ] Automatic dependency resolution
- [ ] Conflict detection and resolution

### Success Criteria
- ✅ Conda fully integrated with all major channels
- ✅ At least 10 domain-specific templates available
- ✅ Environment export/import works reliably
- ✅ Researchers can recreate exact environments from papers
- ✅ Package installation is simple and fast

---

## 🤝 v0.12.0 - Collaboration Features (was v0.10.0 roadmap phase)

**Status:** Planned for v0.12.0 Release
**Target:** 4-5 months
**Priority:** MEDIUM - Academic research is collaborative

### Goals
Enable research teams to collaborate effectively using shared cloud resources.

### Tasks

**Instance Sharing** 🎯 HIGH PRIORITY
- [ ] Share running instances with lab members
  - [ ] Generate time-limited access tokens
  - [ ] Read-only vs full access modes
  - [ ] Revoke access tokens
- [ ] Email invitations with one-click access
- [ ] Access logs for shared instances

**Team Workspaces** 🎯 MEDIUM PRIORITY
- [ ] Lab-wide configuration templates
  - [ ] PI sets defaults for entire lab
  - [ ] Lab members inherit settings
  - [ ] Override mechanism for special cases
- [ ] Shared environments
  - [ ] Lab-specific environment library
  - [ ] Version control for shared environments
- [ ] Resource quotas per team member
  - [ ] Set limits on instance types
  - [ ] Set limits on running hours
  - [ ] Budget allocation per researcher

**Data Sync & Backup** 🎯 HIGH PRIORITY
- [ ] S3 integration for datasets
  - [ ] Automatic sync of project folders
  - [ ] Configurable sync schedules
  - [ ] Bandwidth optimization
- [ ] Automatic notebook backups
  - [ ] Daily backups to S3
  - [ ] Version history
  - [ ] Easy restore from backup
- [ ] Project folder management
  - [ ] Shared project directories
  - [ ] Permissions management

**JupyterHub Support** 🎯 LOW PRIORITY (Future)
- [ ] Multi-user Jupyter on single instance
  - [ ] Create apps/jupyterhub
  - [ ] Authentication integration
  - [ ] User management
  - [ ] Resource quotas per user
- [ ] Ideal for teaching/workshops
- [ ] Class-wide deployments

### Success Criteria
- ✅ Lab members can easily share instances
- ✅ Team configuration templates work reliably
- ✅ Data sync prevents work loss
- ✅ Collaboration features are secure
- ✅ Works across all IDE types

---

## 💰 v0.13.0 - Cost Management for Labs (was v0.11.0 roadmap phase)

**Status:** Planned for v0.13.0 Release
**Target:** 5-6 months
**Priority:** MEDIUM - Academic budgets are tight

### Goals
Help research labs manage cloud spending with budget tracking and optimization.

### Tasks

**Budget Alerts** 🎯 HIGH PRIORITY
- [ ] Set monthly budget per lab/project
  - [ ] Configure budget threshold in config
  - [ ] Per-researcher budgets
  - [ ] Per-project budgets
- [ ] Email alerts when approaching limit
  - [ ] Warning at 50%, 75%, 90% of budget
  - [ ] Daily digest for lab managers
- [ ] Auto-stop when budget exceeded
  - [ ] Configurable enforcement
  - [ ] Grace period before hard stop
  - [ ] Emergency override mechanism

**Cost Reporting** 🎯 HIGH PRIORITY
- [ ] Generate reports for grant reporting
  - [ ] PDF and CSV export
  - [ ] Customizable date ranges
  - [ ] Include breakdown by researcher
  - [ ] Include breakdown by project
- [ ] Per-project cost breakdown
  - [ ] Tag instances with project codes
  - [ ] Aggregate costs by project
  - [ ] Export for grant renewals
- [ ] Per-student/researcher tracking
  - [ ] Individual usage reports
  - [ ] Cost allocation by user
  - [ ] Usage trends over time

**Optimization Recommendations** 🎯 MEDIUM PRIORITY
- [ ] Usage pattern analysis
  - [ ] "Your instance has been idle 80% of the time"
  - [ ] Suggest smaller instance types
  - [ ] Identify over-provisioned resources
- [ ] Spot instance suggestions
  - [ ] Analyze workload suitability for Spot
  - [ ] Estimate savings with Spot instances
  - [ ] One-click conversion to Spot
- [ ] Reserved instance analysis
  - [ ] Identify steady-state workloads
  - [ ] Calculate RI savings potential
  - [ ] Recommend RI purchase strategy

**Cost Forecasting**
- [ ] Predict monthly costs based on current usage
- [ ] Alert when trending over budget
- [ ] Seasonal adjustment (academic calendar)
- [ ] Multi-year cost projections

### Success Criteria
- ✅ Labs can set and enforce budgets
- ✅ Cost reports suitable for grant reporting
- ✅ Researchers see their individual usage
- ✅ Cost optimization recommendations save 20%+
- ✅ No surprise bills

---

## 🖥️ v0.14.0 - Open Source GUI Tools

**Status:** Planned for v0.14.0 Release
**Target:** 6-7 months
**Priority:** HIGH - Widely used open-source research tools

**📖 See Also**: `docs/TOOL_SUPPORT_STRATEGY.md` - Tool Priority Tier 1

### Goals
Add pre-configured open-source GUI applications that are widely used in academic research.

### Applications

**lens-qgis** 🎯 HIGHEST PRIORITY
- [ ] QGIS desktop GIS application
- [ ] Pre-installed common plugins
- [ ] Environments:
  - [ ] `basic-gis` - Essential GIS tools
  - [ ] `advanced-gis` - QGIS + GRASS + SAGA + PostGIS
  - [ ] `remote-sensing` - QGIS + Orfeo Toolbox + SNAP
- [ ] Sample datasets included

**lens-paraview** 🎯 HIGH PRIORITY
- [ ] ParaView for scientific visualization
- [ ] OSPRay rendering support
- [ ] GPU acceleration support
- [ ] Environments:
  - [ ] `visualization` - Standard ParaView
  - [ ] `gpu-visualization` - With GPU rendering
  - [ ] `large-data` - Optimized for datasets >10GB

**lens-imagej** 🎯 HIGH PRIORITY
- [ ] ImageJ/Fiji for image analysis
- [ ] Pre-installed common plugins
- [ ] Environments:
  - [ ] `microscopy` - Fluorescence microscopy tools
  - [ ] `cell-analysis` - Cell segmentation and tracking
  - [ ] `3d-imaging` - 3D reconstruction tools

**lens-octave** 🎯 MEDIUM PRIORITY
- [ ] GNU Octave (MATLAB alternative)
- [ ] Web UI and desktop GUI options
- [ ] Environments:
  - [ ] `mathematics` - Basic math and plotting
  - [ ] `signal-processing` - Signal analysis tools
  - [ ] `control-systems` - Control theory

**lens-openrefine** 🎯 MEDIUM PRIORITY
- [ ] OpenRefine for data cleaning
- [ ] Web-based interface
- [ ] Environments:
  - [ ] `data-cleaning` - Basic cleaning and transformation
  - [ ] `reconciliation` - Entity matching and linking

### Success Criteria
- ✅ 5 open-source GUI tools fully supported
- ✅ Each tool has 2-3 domain-specific environments
- ✅ Tools work seamlessly via NICE DCV
- ✅ GPU acceleration working where applicable
- ✅ Documentation with research use cases

---

## 💼 v0.15.0 - Cloud-Authenticated Commercial Tools

**Status:** Planned for v0.15.0 Release
**Target:** 7-9 months
**Priority:** HIGH - High-demand commercial tools with simple licensing

**📖 See Also**: `docs/TOOL_SUPPORT_STRATEGY.md` - Tool Priority Tier 2

### Goals
Support high-demand commercial research tools that use modern cloud authentication (user logs in with credentials - simple!).

### Why This Is Easy
Modern commercial tools handle licensing via cloud authentication:
1. Install software on AMI
2. Launch via DCV desktop
3. User logs in with institutional/vendor credentials
4. Software validates license automatically

**No license configuration system needed!** Just DCV + installed software.

### Applications

**lens-matlab** 🎯 HIGHEST PRIORITY
- [ ] MATLAB desktop via NICE DCV
- [ ] Cloud authentication (user logs in with MathWorks/institutional credentials)
- [ ] No license configuration needed - MATLAB handles authentication
- [ ] Environments:
  - [ ] `engineering` - Simulink, Control System, Signal Processing
  - [ ] `data-science` - Statistics, Machine Learning, Deep Learning
  - [ ] `computational-biology` - Bioinformatics, Image Processing
  - [ ] `finance` - Financial Toolbox, Econometrics
- [ ] GPU support for deep learning workloads

**lens-mathematica** 🎯 HIGH PRIORITY
- [ ] Wolfram Mathematica desktop via DCV
- [ ] Cloud authentication (user logs in with Wolfram ID)
- [ ] Wolfram Cloud integration
- [ ] Environments:
  - [ ] `symbolic-math` - Symbolic computation
  - [ ] `data-science` - Statistical analysis
  - [ ] `physics` - Mathematical physics tools

**lens-arcgis** 🎯 HIGH PRIORITY
- [ ] ArcGIS Pro desktop via DCV
- [ ] Cloud authentication (ArcGIS Online or enterprise portal credentials)
- [ ] Environments:
  - [ ] `urban-planning` - City and regional planning
  - [ ] `environmental` - Environmental analysis
  - [ ] `remote-sensing` - Satellite imagery analysis

**lens-geneious** 🎯 MEDIUM PRIORITY
- [ ] Geneious Prime for bioinformatics
- [ ] Cloud authentication (Geneious account)
- [ ] Environments:
  - [ ] `genomics` - DNA/RNA sequence analysis
  - [ ] `molecular-biology` - General molecular biology

### Success Criteria
- ✅ 3-4 major cloud-authenticated commercial tools working
- ✅ User can log in with credentials and start working immediately
- ✅ No complex license configuration needed
- ✅ Documentation shows login process for each tool
- ✅ AMIs built with latest software versions

---

## 🗺️ v0.16.0 - Legacy License & Specialized Tools

**Status:** Planned for v0.16.0 Release
**Target:** 9-12 months
**Priority:** LOW-MEDIUM - Legacy licensing and niche tools

**📖 See Also**: `docs/TOOL_SUPPORT_STRATEGY.md` - Tool Priority Tiers 3 & 4

### Goals
Support tools that still use traditional licensing (license files/servers) and other specialized domain tools.

**Note**: This phase requires building license configuration infrastructure - deferred until after cloud-authenticated tools are done.

### Applications by Domain

**Statistics & Social Sciences** (Legacy Licensing)
- [ ] **lens-stata** - Stata (BYOL license files)
  - License file configuration
  - Network license server support (if needed)
- [ ] **lens-spss** - SPSS Statistics (BYOL or subscription)
  - License configuration
- [ ] SAS (via lens-tool catalog)

**Bioinformatics**
- [ ] **lens-pymol** - PyMOL (open source + commercial)
- [ ] CellProfiler (via lens-tool catalog)

**Engineering & Simulation**
- [ ] Ansys (via lens-tool catalog, marketplace)
- [ ] COMSOL Multiphysics (via lens-tool catalog, BYOL)
- [ ] OpenFOAM (via lens-tool catalog, open source)

**3D & Visualization**
- [ ] Blender (via lens-tool catalog, open source)
- [ ] MeshLab (via lens-tool catalog, open source)

**Remote Sensing**
- [ ] ENVI/IDL (via lens-tool catalog, BYOL)
- [ ] SNAP (ESA) (via lens-tool catalog, open source)

### Tool Catalog System
- [ ] Universal `lens-tool` launcher for long-tail tools
- [ ] Community AMI contributions
- [ ] Tool discovery: `lens tools search bioinformatics`
- [ ] Tool info: `lens tools info geneious`

### Infrastructure Needed
- [ ] License file upload and storage system
- [ ] License server configuration
- [ ] Pre-launch license validation (optional)

### Success Criteria
- ✅ Legacy license tools (Stata, SPSS) working
- ✅ License file configuration documented for IT admins
- ✅ Tool catalog system operational for long-tail tools
- ✅ Community can contribute tool definitions
- ✅ Clear separation: cloud-auth tools (easy) vs legacy (complex)

---

## 🚀 v1.0.0 - Production Ready for Academia

**Status:** Planned
**Target:** 6-9 months
**Priority:** HIGH - Ready for wide academic adoption

### Goals
Achieve production-grade stability and completeness for academic research use.

### Requirements for v1.0.0
- [ ] **Research Tools:** 6+ IDE types (Jupyter, RStudio, VSCode, Streamlit, Zeppelin, Theia)
- [ ] **User Experience:** Non-technical researchers can use without training
- [ ] **Documentation:** Video tutorials and step-by-step guides for each tool
- [ ] **Stability:** No critical bugs, <2s command response time
- [ ] **Reproducibility:** Full environment export/import for research reproducibility
- [ ] **Cost Management:** Budget tracking and optimization for labs
- [ ] **Collaboration:** Instance sharing and team features working
- [ ] **Domain Templates:** 10+ research domain-specific environments
- [ ] **Test Coverage:** 60%+ overall coverage
- [ ] **User Base:** 100+ active researchers across multiple institutions
- [ ] **Feedback:** Incorporate feedback from academic beta testers

### Final Polish for Academic Users
- [ ] **Video Tutorials**
  - [ ] Getting started for complete beginners
  - [ ] Each IDE type with research examples
  - [ ] Common workflows (data analysis, paper writing, etc.)
  - [ ] Troubleshooting common issues
- [ ] **Documentation**
  - [ ] Research-focused examples (not developer examples)
  - [ ] Domain-specific guides (biology, social science, etc.)
  - [ ] Screenshots and step-by-step instructions
  - [ ] FAQ for non-technical users
- [ ] **Performance**
  - [ ] Optimize launch times (<5 minutes typical)
  - [ ] Reduce memory usage
  - [ ] Binary size <50MB
- [ ] **Community**
  - [ ] Environment template repository
  - [ ] User showcase (research enabled by Lens)
  - [ ] Active discussion forum for researchers

### Success Criteria for Academic Adoption
- ✅ Non-technical researcher can launch and use any tool in <5 minutes
- ✅ Widely adopted across 5+ universities/research institutions
- ✅ Used in published research (reproducibility)
- ✅ Positive feedback from diverse research domains
- ✅ Active community sharing environments and tips
- ✅ Cost savings documented (vs commercial alternatives)
- ✅ Cited in research papers' methods sections

---

## 📋 Backlog & Ideas

Features under consideration but not yet scheduled:

### Additional IDE Support
- [ ] PyCharm Server
- [ ] IntelliJ IDEA Server
- [ ] RStudio Connect
- [ ] Observable Framework
- [ ] Quarto publishing platform

### Advanced Features
- [ ] GPU instance support
- [ ] Auto-scaling based on load
- [ ] Kubernetes deployment option
- [ ] Container-based environments
- [ ] Spot instance support
- [ ] Reserved instance recommendations

### Integrations
- [ ] Git integration for version control
- [ ] Database connection management
- [ ] S3 data sync
- [ ] Secrets management (AWS Secrets Manager)
- [ ] Parameter Store integration
- [ ] GitHub Codespaces-like experience

### Developer Tools
- [ ] Plugin system for extensibility
- [ ] REST API server
- [ ] Web UI dashboard (unified across IDEs)
- [ ] Mobile app (iOS/Android)
- [ ] VS Code extension
- [ ] Terraform provider

---

## 🤝 Contributing

We welcome contributions! Here's how to get involved:

1. **Check the roadmap** - Find features you're interested in
2. **Open an issue** - Discuss the feature before implementing
3. **Submit a PR** - Follow our [contributing guidelines](CONTRIBUTING.md)
4. **Join discussions** - Participate in feature planning

### Priority Areas for Contributions
- User experience improvements (v0.7.0) - Make it easier for non-technical researchers
- Additional research tools (v0.8.0) - Streamlit, Zeppelin, Theia, Quarto
- Domain-specific environment templates (v0.9.0) - Biology, social science, etc.
- Documentation with research examples
- Video tutorials for academic users
- Bug fixes and usability improvements

---

## 📅 Release Cycle

- **Major versions** (x.0.0): Significant new features, possible breaking changes
- **Minor versions** (0.x.0): New features, backward compatible
- **Patch versions** (0.0.x): Bug fixes, documentation updates

**Typical Timeline:**
- Minor releases: Every 6-8 weeks
- Patch releases: As needed
- Major releases: When significant milestones achieved

---

## 📊 Metrics & Goals

### Project Health Metrics
- Test Coverage: Target 60% by v1.0.0 (currently 18.7%)
- Response Time: <2s for all commands
- Binary Size: <50MB for all platforms (currently ~44MB)
- Memory Usage: <100MB for typical operations
- Startup Time: <100ms

### Community Metrics (Academic Focus)
- Active Researchers: 100+ across 5+ institutions by v1.0.0
- Research Domains: 10+ represented (biology, social science, etc.)
- GitHub Stars: 500+ by v1.0.0
- Contributors: 10+ by v1.0.0 (including academic contributors)
- Issues Resolved: 90%+ within 30 days
- IDE Types Supported: 6+ by v1.0.0 (Jupyter, RStudio, VSCode, Streamlit, Zeppelin, Theia)
- Environment Templates: 20+ domain-specific templates

### Monorepo Health
- Shared library stability: No breaking changes
- Cross-IDE compatibility: All features work identically
- Documentation coverage: 100% of features documented
- CI/CD reliability: 99%+ green builds

---

## 🔄 Versioning Strategy

We follow [Semantic Versioning 2.0.0](https://semver.org/):

- **MAJOR** version when making incompatible API changes
- **MINOR** version when adding functionality in a backward compatible manner
- **PATCH** version when making backward compatible bug fixes

**Special Considerations for Monorepo:**
- Version numbers apply to the entire project (all apps share version)
- Breaking changes in pkg/ trigger major version bump
- New IDE support triggers minor version bump
- Bug fixes in any app trigger patch version bump

---

## 📝 Last Updated

- **Date:** October 2025
- **Current Release:** v0.9.0 (completed v0.7.0 roadmap phase - User Experience & Accessibility)
- **Next Release:** v0.10.0 (will address v0.8.0 roadmap phase - Additional Research Tools)
- **Next Focus:** Streamlit, Q Developer, NICE DCV, Zeppelin
- **Next Review:** December 2025
- **Project Status:** Production-ready platform with 3 IDE types, excellent UX for academic researchers

**Version Alignment Note**: Roadmap phases are planning documents. Actual releases may ship features from different phases based on priorities and readiness.

---

## 💬 Feedback

Have suggestions for the roadmap?
- Open an issue: https://github.com/scttfrdmn/lens/issues
- Start a discussion: https://github.com/scttfrdmn/lens/discussions

We prioritize features based on:
1. **Academic researcher needs** - Ease of use for non-technical users
2. **Research domain coverage** - Support for diverse research fields
3. **Reproducibility** - Enable reproducible research workflows
4. **Cost efficiency** - Help labs manage limited budgets
5. **Multi-IDE compatibility** - Consistent experience across tools
6. **Implementation complexity** vs impact
7. **Community contributions** - Especially from academic users
