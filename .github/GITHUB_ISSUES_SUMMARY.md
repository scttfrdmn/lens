# GitHub Issues Summary

This document provides a comprehensive mapping of all GitHub issues to the project's personas, roadmap phases, and requirements. It serves as the central traceability document connecting ROADMAP.md → Issues → USER_SCENARIOS/ → USER_REQUIREMENTS.md.

**Last Updated:** 2025-10-25
**Total Issues Planned:** 30 issues across 6 phases + 1 critical bug

---

## 🐛 CRITICAL BUGS (Fix Immediately)

### Issue #7: State changes not recorded during start/stop/terminate operations
**Priority:** 🔥 CRITICAL - Blocks core functionality
**Status:** Created 2025-10-25
**Labels:** `bug`, `priority: critical`, `area: config`, `persona: lab-pi`, `persona: graduate-student`, `persona: instructor`, `technical-debt`
**Blocks:** Issues #25-28 (v0.11.0 Cost Management features)

**Impact:**
Cost tracking is fundamentally broken. The `RecordStateChange()` method exists in `pkg/config/state.go` but is never called during lifecycle operations (start, stop, terminate, launch). This causes:
- ❌ Inaccurate cost calculations (treats all time as running time)
- ❌ No utilization tracking (can't calculate running vs. stopped hours)
- ❌ Broken effective cost calculations
- ❌ Unable to show savings from stop/start vs. 24/7 operation

**Personas Impacted:**
- 🔥 **Lab PI** (CRITICAL) - Can't track $15K budget accurately
- 🔥 **Graduate Student** (HIGH) - Can't see actual costs, causes budget anxiety
- 🔥 **Course Instructor** (HIGH) - Can't manage class budget
- **Research Computing Manager** (MEDIUM) - Institutional cost visibility broken

**Related Requirements:**
- **REQ-12.1** (Budget Tracking) - BLOCKED by this bug
- **REQ-12.4** (Cost Reporting) - Inaccurate without state changes
- **REQ-6.1** (Auto-Stop Idle Instances) - Cost savings can't be measured

**Related Pain Points:**
- Lab PI Pain #1 ($4K cost overrun, blind to spending) - USER_SCENARIOS/03
- Graduate Student Pain #3 (budget anxiety) - USER_SCENARIOS/02
- Instructor Pain #7 (no warning before exceeding budget) - USER_SCENARIOS/04

**Fix Scope:**
Add `instance.RecordStateChange()` + `cfg.Save()` in:
- `internal/cli/launch.go` → record "running" after successful launch
- `internal/cli/start.go` → record "running" after instance starts
- `internal/cli/stop.go` → record "stopped" after instance stops
- `internal/cli/terminate.go` → record "terminated" after termination

**Testing Requirements:**
- Unit tests for state change recording
- Integration tests verifying state persists across operations
- Cost calculation tests using recorded state changes

**Timeline:** 2-3 hours for fix + tests (MUST complete before v0.8.0 work)

---

## 📊 Issue Overview by Phase

| Phase | Issue Count | Priority | Target |
|-------|------------|----------|--------|
| v0.7.0 - User Experience & Accessibility | 6 | 🔥 CRITICAL | 1-2 months |
| v0.8.0 - Additional Research Tools | 7 | 🔥 HIGH | 2-3 months |
| v0.9.0 - Package Managers & Reproducibility | 6 | MEDIUM-HIGH | 2-3 months |
| v0.10.0 - Collaboration Features | 5 | MEDIUM | 2-3 months |
| v0.11.0 - Cost Management for Labs | 4 | MEDIUM | 1-2 months |
| v1.0.0 - Production Ready | 2 | HIGH | 6-9 months |

---

## 🎯 v0.7.0 - User Experience & Accessibility (6 Issues)

### Issue #1: Add quickstart command for instant launch
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.7-ux`, `area: cli`, `persona: solo-researcher`, `persona: graduate-student`, `persona: instructor`
**Personas:** Solo Researcher, Graduate Student, Course Instructor
**Requirements:** REQ-1.1 (Beginner-Friendly Onboarding), REQ-7.1 (Fast Launch Times)
**Pain Points:**
- Solo Researcher Pain #1 (5hr → 15min setup) - USER_SCENARIOS/01
- Graduate Student Pain #2 (GPU access friction) - USER_SCENARIOS/02
- Instructor Pain #1 (8hr Week 1 installation) - USER_SCENARIOS/04

**Description:**
```
As a Solo Researcher, I need to launch an environment instantly with sensible defaults without answering any questions, because I just want to start analyzing my data immediately.

**Current Pain Point:**
Even with the wizard, users must answer 6-8 questions before launch. For users who just want to get started quickly, this adds friction.

**Success Metric:**
✅ User can launch in <30 seconds with zero questions
✅ 90% of quickstart launches succeed on first try
✅ Defaults work for 80% of common use cases

**Proposed Solution:**
Add `quickstart` command to all three apps:
- `lens-jupyter quickstart` → t4g.medium, data-science-python env, auto-stop 2hrs
- `lens-rstudio quickstart` → t4g.medium, tidyverse env, auto-stop 2hrs
- `lens-vscode quickstart` → t4g.medium, general-dev env, auto-stop 2hrs
```

---

### Issue #2: Make wizard the default when no subcommand is provided
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.7-ux`, `area: cli`, `area: wizard`, `persona: solo-researcher`, `persona: graduate-student`, `persona: instructor`
**Personas:** Solo Researcher, Graduate Student, Course Instructor
**Requirements:** REQ-1.1 (Beginner-Friendly Onboarding), REQ-1.2 (Plain-English Interface)
**Pain Points:**
- Solo Researcher Pain #1 (command discovery) - USER_SCENARIOS/01
- Instructor Pain #1 (student onboarding) - USER_SCENARIOS/04

**Description:**
```
As a Graduate Student with limited CLI experience, I need the wizard to launch automatically when I run `lens-jupyter`, because I don't know what subcommands are available.

**Current Pain Point:**
Running `lens-jupyter` shows help text with command list. Non-technical users don't understand subcommands and get confused.

**Success Metric:**
✅ `lens-jupyter` (no args) launches wizard
✅ 95% of first-time users successfully launch
✅ Power users can skip with `--no-wizard`

**Proposed Solution:**
- `lens-jupyter` → launches wizard
- `lens-jupyter launch` → uses last settings or defaults
- `lens-jupyter --no-wizard launch` → CLI mode
- Update all three apps (jupyter, rstudio, vscode)
```

---

### Issue #3: Remember user preferences between launches
**Priority:** MEDIUM
**Labels:** `enhancement`, `phase: 0.7-ux`, `area: config`, `persona: solo-researcher`, `persona: graduate-student`
**Personas:** Solo Researcher, Graduate Student
**Requirements:** REQ-1.3 (Quick Re-launch), REQ-4.1 (Configuration Persistence)
**Pain Points:**
- Solo Researcher Pain #3 (repetitive configuration) - USER_SCENARIOS/01
- Graduate Student Pain #4 (context switching overhead) - USER_SCENARIOS/02

**Description:**
```
As a Solo Researcher who launches similar environments weekly, I need the tool to remember my previous settings, because I use the same instance type and environment every time.

**Current Pain Point:**
Every launch requires re-selecting instance type, environment, region, and other settings even when using the same configuration repeatedly.

**Success Metric:**
✅ Preference memory reduces wizard time from 90s to 15s
✅ "Use last settings?" appears as first question
✅ Users can override individual settings if needed

**Proposed Solution:**
- Store last launch config in `~/.aws-ide/preferences.yaml`
- Wizard asks: "Use same settings as last time? (Y/n)"
- If yes → launch immediately with previous config
- If no → proceed through wizard questions
- Track: instance type, environment, region, auto-stop duration
```

---

### Issue #4: Add email notifications for instance lifecycle events
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.7-ux`, `area: output`, `persona: graduate-student`, `persona: lab-pi`
**Personas:** Graduate Student, Lab PI
**Requirements:** REQ-5.4 (Status Notifications), REQ-6.1 (Auto-Stop Idle Instances)
**Pain Points:**
- Graduate Student Pain #3 (forgetting to stop instances) - USER_SCENARIOS/02
- Lab PI Pain #2 (no visibility into team usage) - USER_SCENARIOS/03

**Description:**
```
As a Graduate Student running long analyses, I need email notifications when my environment is ready and when it's about to auto-stop, because I'm often working on other tasks while waiting.

**Current Pain Point:**
Users must actively check if instance is ready. Long-running analyses may auto-stop without warning, losing progress.

**Success Metric:**
✅ 95% email delivery rate within 1 minute of event
✅ 80% reduction in "forgot to stop" cost overruns
✅ Optional (opt-in via config flag)

**Proposed Solution:**
Add config option: `notifications.email: user@example.com`
Send emails for:
- "Your Jupyter Lab is ready at: http://..."
- "Your instance will stop in 30 minutes due to idle timeout"
- "Your instance has stopped (idle timeout reached)"
- Use AWS SES for email delivery
```

---

### Issue #5: Add Slack webhook support for lab notifications
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.7-ux`, `area: output`, `persona: lab-pi`, `persona: instructor`
**Personas:** Lab PI, Course Instructor
**Requirements:** REQ-5.4 (Status Notifications), REQ-8.3 (Team Communication)
**Pain Points:**
- Lab PI Pain #2 (no visibility into team usage) - USER_SCENARIOS/03
- Instructor Pain #4 (monitoring 30 student environments) - USER_SCENARIOS/04

**Description:**
```
As a Lab PI managing 7 researchers, I need Slack notifications when lab members launch/stop instances, because I need to monitor lab resource usage without checking manually.

**Current Pain Point:**
PIs have no visibility into when team members are using resources, making budget management reactive instead of proactive.

**Success Metric:**
✅ Slack messages delivered within 30 seconds
✅ Lab-wide dashboard shows all active instances
✅ Weekly summary of usage and costs

**Proposed Solution:**
Add config option: `notifications.slack_webhook: https://hooks.slack.com/...`
Send messages for:
- "Alice launched Jupyter (t4g.xlarge, $0.134/hr estimated)"
- "Bob's RStudio auto-stopped after 2hrs idle"
- "Weekly lab usage: $47.23 of $500 budget (9%)"
- Include user, app type, instance type, estimated cost
```

---

### Issue #6: Investigate desktop notifications for instance readiness
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.7-ux`, `area: output`, `persona: solo-researcher`, `persona: graduate-student`, `question`
**Personas:** Solo Researcher, Graduate Student
**Requirements:** REQ-5.4 (Status Notifications), REQ-7.1 (Fast Launch Times)
**Pain Points:**
- Solo Researcher Pain #1 (multitasking during launch) - USER_SCENARIOS/01
- Graduate Student Pain #4 (context switching) - USER_SCENARIOS/02

**Description:**
```
As a Solo Researcher waiting for my environment to launch, I need desktop notifications so I can work on other tasks without constantly checking the terminal.

**Current Pain Point:**
Launch takes 2-3 minutes. Users must keep terminal visible to see when ready, preventing multitasking.

**Success Metric:**
✅ Notifications work on macOS, Linux, Windows
✅ 90% of users find notifications helpful (user survey)
✅ Optional (opt-in via config flag)

**Research Needed:**
- Cross-platform notification libraries (Go)
- macOS: terminal-notifier or native
- Linux: libnotify / notify-send
- Windows: Windows 10 toast notifications
- Permissions/security considerations
- Fallback if notifications unavailable
```

---

## 🔬 v0.8.0 - Additional Research Tools (7 Issues)

### Issue #7: Add Amazon Q Developer standalone app (lens-q-developer)
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.8-tools`, `area: lens-q-developer`, `persona: solo-researcher`, `persona: graduate-student`, `persona: instructor`
**Personas:** Solo Researcher, Graduate Student, Course Instructor
**Requirements:** REQ-2.4 (AI-Assisted Development), REQ-3.6 (New Research Tools)
**Pain Points:**
- Solo Researcher Pain #5 (learning curve for coding) - USER_SCENARIOS/01
- Graduate Student Pain #5 (debugging without help) - USER_SCENARIOS/02
- Instructor Pain #3 (teaching students to code) - USER_SCENARIOS/04

**Description:**
```
As a Solo Researcher learning to code for data analysis, I need AI-assisted coding with explanations, because I often get stuck and don't have a coding mentor.

**Current Pain Point:**
Q Developer extension is available in VSCode, but not as standalone. Researchers who prefer Jupyter/RStudio don't have access to AI assistance.

**Success Metric:**
✅ Standalone Q Developer app launched in <3 minutes
✅ AI code suggestions work in Python, R, SQL
✅ Three built-in environments (research-coding, data-analysis, learning)

**Proposed Solution:**
Create `apps/q-developer/` following monorepo pattern:
- Port 9000 (Amazon Q Developer IDE)
- Environments:
  - `research-coding` - Python + R with Q assistance
  - `data-analysis` - Data science with AI help
  - `learning` - Beginner-friendly with Q explanations
- Leverage AWS native integration
- CLI: `lens-q-developer launch --env research-coding`
```

---

### Issue #8: Add Streamlit app for interactive data apps (lens-streamlit)
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.8-tools`, `area: streamlit`, `persona: solo-researcher`, `persona: lab-pi`, `persona: instructor`, `use-case: collaboration`, `use-case: teaching`
**Personas:** Solo Researcher, Lab PI, Course Instructor
**Requirements:** REQ-2.3 (Interactive Visualizations), REQ-3.6 (New Research Tools), REQ-8.2 (Shareable Results)
**Pain Points:**
- Solo Researcher Pain #6 (sharing results with non-coders) - USER_SCENARIOS/01
- Lab PI Pain #4 (team collaboration friction) - USER_SCENARIOS/03
- Instructor Pain #2 (student result presentation) - USER_SCENARIOS/04

**Description:**
```
As a Solo Researcher presenting to non-technical collaborators, I need to turn my Python analysis into an interactive web app, because static notebooks don't engage stakeholders.

**Current Pain Point:**
Sharing analysis requires collaborators to understand code. Notebooks (`.ipynb`) don't run interactively for non-programmers.

**Success Metric:**
✅ Python script → web app in <5 minutes
✅ Shareable URL for lab members/collaborators
✅ No Docker/deployment knowledge required

**Proposed Solution:**
Create `apps/streamlit/` following monorepo pattern:
- Port 8501 (Streamlit default)
- Environments:
  - `data-viz` - Data visualization apps (Plotly, Altair)
  - `ml-demo` - Machine learning demos (scikit-learn, TensorFlow)
  - `dashboard` - Analysis dashboards (pandas, streamlit-aggrid)
- Auto-detect Streamlit scripts in project
- CLI: `lens-streamlit launch --script my_app.py`
- Built-in tunnel for sharing (AWS SSM + public URL)
```

---

### Issue #9: Add Apache Zeppelin for multi-language notebooks (lens-zeppelin)
**Priority:** MEDIUM
**Labels:** `enhancement`, `phase: 0.8-tools`, `area: zeppelin`, `persona: lab-pi`, `persona: it-admin`, `use-case: bioinformatics`, `use-case: data-engineering`
**Personas:** Lab PI, Research Computing Manager
**Requirements:** REQ-3.6 (New Research Tools), REQ-3.7 (Multi-Language Support)
**Pain Points:**
- Lab PI Pain #5 (diverse tool requirements across team) - USER_SCENARIOS/03
- Research Computing Manager Pain #3 (supporting 100+ researchers) - USER_SCENARIOS/05

**Description:**
```
As a Lab PI with researchers using Python, R, Scala, and SQL, I need a notebook that supports all languages, because different analyses require different tools.

**Current Pain Point:**
Jupyter focuses on Python, RStudio on R. Labs need multiple separate tools for multi-language workflows (Python → SQL → R pipeline).

**Success Metric:**
✅ Python, R, Scala, SQL in same notebook
✅ Built-in visualizations without custom libraries
✅ Three domain-specific environments

**Proposed Solution:**
Create `apps/zeppelin/` following monorepo pattern:
- Port 8080 (Zeppelin default)
- Environments:
  - `data-engineering` - ETL and data processing (Spark, SQL)
  - `sql-analytics` - Database analysis (PostgreSQL, MySQL connectors)
  - `spark` - Big data processing (PySpark, SparkR, Scala)
- Pre-configure interpreters (Python, R, SQL, Scala)
- CLI: `lens-zeppelin launch --env spark`
```

---

### Issue #10: Add NICE DCV Desktop for GUI research applications (lens-dcv-desktop)
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.8-tools`, `area: dcv-desktop`, `aws: gpu`, `persona: lab-pi`, `persona: graduate-student`, `use-case: bioinformatics`, `use-case: machine-learning`, `use-case: visualization`
**Personas:** Lab PI, Graduate Student, Solo Researcher
**Requirements:** REQ-3.6 (New Research Tools), REQ-3.8 (GUI Application Support), REQ-6.3 (GPU Instance Support)
**Pain Points:**
- Graduate Student Pain #2 (GUI apps like MATLAB, ImageJ) - USER_SCENARIOS/02
- Lab PI Pain #5 (diverse tools - MATLAB, Igor Pro) - USER_SCENARIOS/03
- Solo Researcher Pain #4 (bioinformatics GUI tools) - USER_SCENARIOS/01

**Description:**
```
As a Graduate Student using MATLAB, ImageJ, and other GUI tools, I need a full Linux desktop environment on AWS, because command-line alternatives don't exist for my discipline.

**Current Pain Point:**
Many research tools require GUI (MATLAB, Igor Pro, ImageJ, Geneious). X11 forwarding is slow and unreliable. VNC is complicated to set up.

**Success Metric:**
✅ Full desktop in browser with <100ms latency
✅ GPU support for visualization and computation
✅ Copy/paste and file transfer working
✅ 6 domain-specific desktop environments

**Proposed Solution:**
Create `apps/dcv-desktop/` following monorepo pattern:
- AWS NICE DCV (native AWS remote desktop)
- Port 8443 (DCV web client)
- Environments:
  - `matlab-desktop` - MATLAB with full GUI
  - `data-viz-desktop` - ParaView, Visit, Tableau
  - `image-analysis` - ImageJ, Fiji, QuPath, CellProfiler
  - `bioinformatics-gui` - Geneious, CLC Workbench, UGENE
  - `general-desktop` - Ubuntu desktop with research tools
  - `gpu-workstation` - CUDA, visualization, ML (GPU instances)
- GPU instance type selection (p3, g4dn, g5)
- CLI: `lens-dcv-desktop launch --env matlab-desktop --gpu`
```

---

### Issue #11: Add Theia IDE for extensible cloud development (lens-theia)
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.8-tools`, `area: theia`, `persona: solo-researcher`, `persona: graduate-student`
**Personas:** Solo Researcher, Graduate Student
**Requirements:** REQ-3.6 (New Research Tools), REQ-3.4 (VSCode Alternative)
**Pain Points:**
- Solo Researcher Pain #7 (occasional coding needs) - USER_SCENARIOS/01
- Graduate Student Pain #6 (IDE learning curve) - USER_SCENARIOS/02

**Description:**
```
As a Solo Researcher who codes occasionally, I need a full-featured IDE that's easier to extend than VSCode, because I want custom workflows without complex configuration.

**Current Pain Point:**
VSCode requires extension marketplace access and complex settings. Theia offers better extensibility for custom research workflows.

**Success Metric:**
✅ Launch Theia in <3 minutes
✅ Pre-configured extensions for Python, R
✅ Three research-focused environments

**Proposed Solution:**
Create `apps/theia/` following monorepo pattern:
- Port 3000 (Theia default)
- Environments:
  - `python` - Python development (PyLance equivalent)
  - `r-dev` - R package development (devtools, roxygen2)
  - `multi-lang` - Multiple languages (Python, R, Julia)
- Pre-install popular extensions
- CLI: `lens-theia launch --env python`
```

---

### Issue #12: Add Quarto for academic publishing (lens-quarto)
**Priority:** MEDIUM
**Labels:** `enhancement`, `phase: 0.8-tools`, `area: quarto`, `persona: graduate-student`, `persona: lab-pi`, `persona: instructor`, `use-case: reproducibility`, `use-case: teaching`
**Personas:** Graduate Student, Lab PI, Course Instructor
**Requirements:** REQ-3.6 (New Research Tools), REQ-9.1 (Reproducible Research), REQ-9.2 (Academic Publishing)
**Pain Points:**
- Graduate Student Pain #7 (paper writing with code) - USER_SCENARIOS/02
- Lab PI Pain #6 (reproducible papers) - USER_SCENARIOS/03
- Instructor Pain #5 (course materials generation) - USER_SCENARIOS/04

**Description:**
```
As a Graduate Student writing my dissertation, I need to create papers with embedded code and figures, because manually updating figures when data changes is error-prone.

**Current Pain Point:**
Academic writing workflow: Write in Word → Generate figures in R/Python → Manually insert → Update when reviewers request changes → Repeat. Not reproducible.

**Success Metric:**
✅ Code + text → PDF/HTML/Word in one command
✅ Works with Jupyter notebooks and R Markdown
✅ Three publication-focused environments

**Proposed Solution:**
Create `apps/quarto/` following monorepo pattern:
- Port 4000 (Quarto preview server)
- Environments:
  - `academic-paper` - LaTeX + code (PDF output)
  - `presentation` - Reveal.js slides with live code
  - `website` - Research group websites (Hugo/Jekyll integration)
- Pre-configure LaTeX, bibliography management (BibTeX)
- CLI: `lens-quarto launch --env academic-paper`
- Support for .qmd, .Rmd, .ipynb source files
```

---

### Issue #13: Add Observable Framework for data visualization (lens-observable)
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.8-tools`, `area: observable`, `persona: solo-researcher`, `use-case: visualization`, `use-case: collaboration`
**Personas:** Solo Researcher, Lab PI
**Requirements:** REQ-3.6 (New Research Tools), REQ-2.3 (Interactive Visualizations)
**Pain Points:**
- Solo Researcher Pain #6 (sharing interactive visualizations) - USER_SCENARIOS/01
- Lab PI Pain #4 (collaboration and communication) - USER_SCENARIOS/03

**Description:**
```
As a Solo Researcher creating data visualizations for publications, I need interactive JavaScript notebooks, because static figures don't convey complex patterns effectively.

**Current Pain Point:**
Python/R visualizations are static images in papers. Interactive D3.js visualizations require web development expertise.

**Success Metric:**
✅ Data → interactive visualization in <30 minutes
✅ Shareable URLs for collaborators/reviewers
✅ Embed in websites and papers (iframe)

**Proposed Solution:**
Create `apps/observable/` following monorepo pattern:
- Port 3000 (Observable Framework)
- Environments:
  - `data-viz` - D3.js visualizations with examples
  - `interactive` - Interactive analysis dashboards
- Pre-configure D3.js, Vega-Lite, Plot libraries
- CLI: `lens-observable launch --env data-viz`
- Export to static HTML for publication
```

---

## 📦 v0.9.0 - Package Managers & Reproducibility (6 Issues)

### Issue #14: Add full conda environment support with environment.yml
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.9-reproducibility`, `area: environments`, `persona: graduate-student`, `persona: lab-pi`, `use-case: reproducibility`, `use-case: bioinformatics`
**Personas:** Graduate Student, Lab PI, Solo Researcher
**Requirements:** REQ-9.1 (Reproducible Research), REQ-9.3 (Package Management), REQ-4.3 (Environment Export/Import)
**Pain Points:**
- Graduate Student Pain #8 (dependency conflicts) - USER_SCENARIOS/02
- Lab PI Pain #6 (reproducibility for papers) - USER_SCENARIOS/03
- Solo Researcher Pain #8 (complex software installation) - USER_SCENARIOS/01

**Description:**
```
As a Graduate Student publishing research, I need to export my exact package environment, because reviewers and readers must be able to reproduce my results.

**Current Pain Point:**
Pre-built environments don't include all needed packages. Manual conda install is slow and forgotten. Papers lack reproducible environment specs.

**Success Metric:**
✅ `environment.yml` → working environment in <10 minutes
✅ `lens-jupyter env export` captures exact environment
✅ Package caching speeds up repeated launches by 80%

**Proposed Solution:**
- Support `environment.yml` files in launch wizard
- Auto-detect environment.yml in project directory
- New commands:
  - `lens-jupyter env export > environment.yml`
  - `lens-jupyter launch --env-file environment.yml`
- Package caching in S3 for faster rebuilds
- Support conda-forge and bioconda channels
```

---

### Issue #15: Add BioConda integration for bioinformatics workflows
**Priority:** MEDIUM
**Labels:** `enhancement`, `phase: 0.9-reproducibility`, `area: environments`, `persona: solo-researcher`, `persona: graduate-student`, `use-case: bioinformatics`
**Personas:** Solo Researcher, Graduate Student
**Requirements:** REQ-9.3 (Package Management), REQ-10.1 (Domain-Specific Environments)
**Pain Points:**
- Solo Researcher Pain #4 (bioinformatics tools setup) - USER_SCENARIOS/01
- Graduate Student Pain #9 (genomics pipeline dependencies) - USER_SCENARIOS/02

**Description:**
```
As a Solo Researcher in bioinformatics, I need BioConda pre-configured, because genomics tools are notoriously difficult to install manually.

**Current Pain Point:**
Bioinformatics tools (GATK, BWA, SAMtools) have complex dependencies. Conda solves this but requires bioconda channel configuration.

**Success Metric:**
✅ BioConda channel enabled by default in bio environments
✅ 50+ common genomics tools available via conda
✅ Pre-configured environments for genomics, RNA-seq, single-cell

**Proposed Solution:**
- Add bioconda channel to all bioinformatics environments
- Create specialized environments:
  - `genomics-pipeline` - GATK, BWA, SAMtools, BCFtools
  - `rna-seq-analysis` - DESeq2, edgeR, Salmon, Kallisto
  - `single-cell` - Seurat, Scanpy, Monocle
- Document bioconda usage in bioinformatics guide
- Test environment builds in CI/CD
```

---

### Issue #16: Add domain-specific environment templates (10+ research domains)
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.9-reproducibility`, `area: environments`, `persona: solo-researcher`, `persona: graduate-student`, `persona: lab-pi`, `persona: instructor`, `use-case: data-analysis`, `use-case: machine-learning`, `use-case: bioinformatics`, `use-case: statistics`
**Personas:** Solo Researcher, Graduate Student, Lab PI, Course Instructor
**Requirements:** REQ-10.1 (Domain-Specific Environments), REQ-10.2 (Research Use Case Coverage)
**Pain Points:**
- Solo Researcher Pain #8 (tool discovery and setup) - USER_SCENARIOS/01
- Graduate Student Pain #10 (domain-specific dependencies) - USER_SCENARIOS/02
- Lab PI Pain #5 (diverse researcher needs) - USER_SCENARIOS/03
- Instructor Pain #6 (discipline-specific teaching) - USER_SCENARIOS/04

**Description:**
```
As a Solo Researcher in climate science, I need a pre-configured environment with xarray, iris, and geospatial tools, because I don't know what packages my field uses.

**Current Pain Point:**
Generic "data science" environments miss domain-specific tools. Researchers waste hours discovering and installing specialized packages.

**Success Metric:**
✅ 10+ research domains covered with 2-3 environments each
✅ 90% of researchers find a suitable starting template
✅ Community can contribute new templates

**Proposed Solution:**
Create domain-specific environment templates:

**Biology/Genomics:**
- `genomics-pipeline` - GATK, BWA, SAMtools
- `rna-seq` - DESeq2, edgeR, Salmon
- `single-cell` - Seurat, Scanpy

**Social Sciences:**
- `survey-analysis` - pandas, statsmodels, matplotlib
- `text-analysis` - NLTK, spaCy, gensim
- `network-analysis` - NetworkX, igraph, graph-tool

**Economics/Finance:**
- `econometrics` - statsmodels, linearmodels, pandas
- `time-series` - prophet, ARIMA, statsmodels
- `financial-modeling` - QuantLib, pandas, numpy

**Climate Science:**
- `climate-data` - xarray, iris, netCDF4
- `geospatial` - GDAL, rasterio, geopandas
- `climate-viz` - cartopy, proplot, matplotlib

**Physics/Engineering:**
- `scientific-computing` - NumPy, SciPy, SymPy
- `simulation` - OpenFOAM, FEniCS, COMSOL
- `symbolic-math` - SymPy, Maxima, Sage

Store templates in `pkg/config/environments/` with documentation
```

---

### Issue #17: Add environment export/import for reproducibility
**Priority:** MEDIUM
**Labels:** `enhancement`, `phase: 0.9-reproducibility`, `area: environments`, `area: cli`, `persona: graduate-student`, `persona: lab-pi`, `use-case: reproducibility`
**Personas:** Graduate Student, Lab PI
**Requirements:** REQ-9.1 (Reproducible Research), REQ-4.3 (Environment Export/Import)
**Pain Points:**
- Graduate Student Pain #8 (paper reproducibility) - USER_SCENARIOS/02
- Lab PI Pain #6 (ensuring research reproducibility) - USER_SCENARIOS/03

**Description:**
```
As a Graduate Student submitting my paper, I need to export my exact environment specification, because journals now require reproducible computational methods.

**Current Pain Point:**
No easy way to capture complete environment (packages, versions, system deps). "Works on my machine" is not acceptable for published research.

**Success Metric:**
✅ One command exports complete environment
✅ Colleagues can recreate exact environment from export
✅ Export includes conda packages, system packages, custom configs

**Proposed Solution:**
New commands for all apps:
- `lens-jupyter env export > my-analysis-env.yml`
  - Includes conda packages with exact versions
  - Includes system packages (apt/yum)
  - Includes custom configuration
- `lens-jupyter launch --import-env my-analysis-env.yml`
  - Recreates exact environment
  - Validates all packages available
  - Warns if versions unavailable
- Format: Extended environment.yml with metadata
```

---

### Issue #18: Add system package management (apt/yum) to environments
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.9-reproducibility`, `area: environments`, `technical-debt`, `persona: graduate-student`, `use-case: reproducibility`
**Personas:** Graduate Student, Lab PI
**Requirements:** REQ-9.3 (Package Management), REQ-4.2 (Custom Environments)
**Pain Points:**
- Graduate Student Pain #10 (system dependencies) - USER_SCENARIOS/02

**Description:**
```
As a Graduate Student, I need to declare system packages in my environment file, because my Python packages depend on system libraries (e.g., GDAL, HDF5).

**Current Pain Point:**
Conda packages have system dependencies. Environments only specify Python/R packages. Users must manually install system deps or environment fails.

**Success Metric:**
✅ Declare system packages in environment YAML
✅ Automatic installation during environment setup
✅ Version pinning for system packages

**Proposed Solution:**
Extend environment.yml format:
```yaml
name: my-research-env
channels:
  - conda-forge
dependencies:
  - python=3.11
  - pandas=2.0.0
system_packages:
  apt:  # for Ubuntu/Debian
    - libgdal-dev=3.6.2
    - libhdf5-dev=1.12.2
  yum:  # for Amazon Linux
    - gdal-devel-3.6.2
    - hdf5-devel-1.12.2
```
- Install system packages before conda environment
- Cache system packages in AMI for common tools
- Support version pinning for reproducibility
```

---

### Issue #19: Add community environment repository with sharing
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.9-reproducibility`, `area: environments`, `use-case: collaboration`, `use-case: reproducibility`, `persona: lab-pi`, `persona: instructor`, `persona: it-admin`
**Personas:** Lab PI, Course Instructor, Research Computing Manager
**Requirements:** REQ-8.4 (Community Sharing), REQ-10.1 (Domain-Specific Environments)
**Pain Points:**
- Lab PI Pain #4 (sharing configurations across team) - USER_SCENARIOS/03
- Instructor Pain #3 (distributing consistent environments) - USER_SCENARIOS/04
- Research Computing Manager Pain #4 (supporting diverse disciplines) - USER_SCENARIOS/05

**Description:**
```
As a Lab PI, I need to share my lab's custom environments with the research community, because reproducibility benefits everyone.

**Current Pain Point:**
Each lab reinvents the wheel creating domain-specific environments. No way to discover what others have created.

**Success Metric:**
✅ Search environments by research domain
✅ Upvote/review system for quality
✅ 50+ community-contributed environments by v1.0

**Proposed Solution:**
GitHub repository for community environments:
- `aws-ide/environments` repo
- Directory structure: `biology/genomics/GATK-pipeline.yml`
- Metadata: description, author, domain, tools included
- CLI integration:
  - `lens-jupyter env search genomics` → lists community environments
  - `lens-jupyter env install community/biology/genomics/GATK-pipeline`
- Upvote/star system on GitHub
- Testing: Community environments tested in CI/CD
- Documentation: Each environment includes README with use cases
```

---

## 🤝 v0.10.0 - Collaboration Features (5 Issues)

### Issue #20: Add instance sharing with time-limited access tokens
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.10-collaboration`, `area: aws`, `security`, `persona: lab-pi`, `persona: instructor`, `use-case: collaboration`, `use-case: teaching`
**Personas:** Lab PI, Course Instructor, Graduate Student
**Requirements:** REQ-8.1 (Instance Sharing), REQ-11.2 (Access Control), REQ-11.3 (Secure Collaboration)
**Pain Points:**
- Lab PI Pain #4 (no collaboration features) - USER_SCENARIOS/03
- Instructor Pain #4 (can't help stuck students) - USER_SCENARIOS/04
- Graduate Student Pain #11 (can't share with advisor) - USER_SCENARIOS/02

**Description:**
```
As a Lab PI, I need to access a student's running instance to help debug their analysis, because screen sharing doesn't give me the full environment.

**Current Pain Point:**
No way to share access to running instance. Researchers must share AWS credentials (insecure) or recreate environment locally (time-consuming).

**Success Metric:**
✅ Generate shareable link in <10 seconds
✅ Read-only and full-access modes
✅ Tokens expire after configurable time (1hr, 24hr, 1 week)
✅ Revoke access tokens instantly

**Proposed Solution:**
New commands for all apps:
- `lens-jupyter share --instance i-abc123 --access read-only --expires 24h`
  - Generates URL: `https://share.aws-ide.io/abc123-token`
  - Creates temporary IAM policy with SSM Session Manager access
  - Recipient clicks URL → authenticated via Session Manager → read-only terminal
- `lens-jupyter share --access full --expires 1h`
  - Full read-write access to notebooks
  - Jupyter token included in URL
- `lens-jupyter share revoke --token abc123`
  - Immediately revokes access
- Access logs: who accessed, when, what actions
- Email invitation option: `--email colleague@university.edu`
```

---

### Issue #21: Add team workspaces with lab-wide configuration templates
**Priority:** MEDIUM
**Labels:** `enhancement`, `phase: 0.10-collaboration`, `area: config`, `persona: lab-pi`, `persona: it-admin`, `use-case: collaboration`
**Personas:** Lab PI, Research Computing Manager
**Requirements:** REQ-8.5 (Team Workspaces), REQ-11.1 (Centralized Configuration), REQ-12.2 (Budget Allocation)
**Pain Points:**
- Lab PI Pain #2 (no centralized control) - USER_SCENARIOS/03
- Research Computing Manager Pain #5 (1,250 researchers, inconsistent setups) - USER_SCENARIOS/05

**Description:**
```
As a Lab PI managing 7 researchers, I need to set default configurations for my entire lab, because I want consistent environments and centralized budget control.

**Current Pain Point:**
Each lab member configures independently. No way to enforce defaults (region, instance types, budgets). PI has no visibility or control.

**Success Metric:**
✅ Lab-wide defaults apply to all members
✅ Members can override for special cases
✅ 90% reduction in configuration time for new lab members

**Proposed Solution:**
Hierarchical configuration system:
- Lab configuration: `~/.aws-ide/lab-config.yml`
```yaml
lab:
  name: "Chen Lab"
  budget: $15000
  defaults:
    region: us-east-1
    instance_types: [t4g.medium, t4g.large, t4g.xlarge]  # whitelist
    auto_stop: 2h
    environments: [bioconductor, tidyverse, genomics-pipeline]
  members:
    - alice@uni.edu:
        budget: $2000
        can_override: false  # must use defaults
    - bob@uni.edu:
        budget: $3000
        can_override: true  # can override for special cases
```
- PI sets lab config, members inherit
- New command: `lens-jupyter config init-lab`
- Members see: "Using Chen Lab defaults (override with --no-lab-config)"
```

---

### Issue #22: Add S3 integration for automatic data sync and backup
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.10-collaboration`, `area: aws`, `aws: s3`, `persona: solo-researcher`, `persona: graduate-student`, `persona: lab-pi`, `use-case: collaboration`, `use-case: reproducibility`
**Personas:** Solo Researcher, Graduate Student, Lab PI
**Requirements:** REQ-8.6 (Data Persistence), REQ-9.4 (Backup and Recovery)
**Pain Points:**
- Solo Researcher Pain #2 (lost work when instance terminated) - USER_SCENARIOS/01
- Graduate Student Pain #3 (forgot to save notebooks before stopping) - USER_SCENARIOS/02
- Lab PI Pain #7 (no backups, student lost 2 weeks of work) - USER_SCENARIOS/03

**Description:**
```
As a Solo Researcher, I need automatic backup of my notebooks to S3, because I've lost work by forgetting to save before terminating instances.

**Current Pain Point:**
Work stored on instance EBS volume. If instance terminated, work lost. Users must manually copy files to local machine or S3.

**Success Metric:**
✅ Zero data loss incidents
✅ Automatic sync every 15 minutes
✅ Version history for notebooks (restore previous versions)
✅ Configurable sync folders and schedule

**Proposed Solution:**
S3 sync integration:
- Config option:
```yaml
backup:
  enabled: true
  s3_bucket: s3://my-research-bucket/aws-ide-backups/
  sync_folders:
    - ~/notebooks
    - ~/data
  schedule: "*/15 * * * *"  # every 15 minutes
  versioning: true
```
- Automatic sync runs in background
- On instance stop/terminate: final sync before shutdown
- Restore on launch: `lens-jupyter launch --restore-from s3://...`
- CLI commands:
  - `lens-jupyter backup now` - force immediate backup
  - `lens-jupyter backup list` - show backup history
  - `lens-jupyter backup restore --date 2025-10-15` - restore from date
- Bandwidth optimization: rsync-style incremental sync
- Encryption: S3 server-side encryption by default
```

---

### Issue #23: Add JupyterHub support for multi-user environments
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.10-collaboration`, `area: jupyterhub`, `persona: instructor`, `persona: it-admin`, `use-case: teaching`
**Personas:** Course Instructor, Research Computing Manager
**Requirements:** REQ-8.7 (Multi-User Environments), REQ-13.1 (Teaching Support)
**Pain Points:**
- Instructor Pain #2 (managing 30 individual instances is expensive) - USER_SCENARIOS/04
- Research Computing Manager Pain #6 (institutional deployment for 100+ users) - USER_SCENARIOS/05

**Description:**
```
As a Course Instructor teaching 30 students, I need multi-user Jupyter on a single instance, because launching 30 separate instances is expensive and hard to manage.

**Current Pain Point:**
One instance per student = 30 instances = 30× cost + 30× management overhead. JupyterHub allows one instance with 30 users.

**Success Metric:**
✅ 30 students on single r6g.4xlarge (16 cores, 128GB RAM)
✅ 80% cost reduction vs individual instances
✅ Per-user quotas (CPU, memory, storage)

**Proposed Solution:**
Create `apps/jupyterhub/` following monorepo pattern:
- Port 8000 (JupyterHub)
- Authentication:
  - GitHub OAuth (most common for courses)
  - AWS Cognito integration
  - LDAP for institutions
- Per-user resource quotas:
  - CPU: 0.5-2 cores per user
  - Memory: 2-8GB per user
  - Storage: 10-50GB per user
- Environments:
  - `course-teaching` - Class-wide deployment
  - `workshop` - Short-term workshop (<1 week)
  - `lab-shared` - Small lab (5-10 users)
- Admin dashboard for instructor:
  - Monitor user activity
  - Restart individual user servers
  - View resource usage
- CLI: `lens-jupyterhub launch --users 30 --instance-type r6g.4xlarge`
- User management: CSV upload of student emails
```

---

### Issue #24: Add project folder permissions and shared directories
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.10-collaboration`, `area: config`, `persona: lab-pi`, `use-case: collaboration`
**Personas:** Lab PI, Research Computing Manager
**Requirements:** REQ-8.5 (Team Workspaces), REQ-8.6 (Data Persistence)
**Pain Points:**
- Lab PI Pain #4 (no shared project folders) - USER_SCENARIOS/03

**Description:**
```
As a Lab PI, I need shared project folders accessible to my entire lab, because we're analyzing the same datasets collaboratively.

**Current Pain Point:**
Each researcher has isolated environment. No shared storage for common datasets. Duplicate data storage = wasted S3 costs.

**Success Metric:**
✅ Shared folders accessible to all lab members
✅ 50% reduction in duplicate data storage
✅ Permissions: read-only vs read-write per user

**Proposed Solution:**
Shared S3-backed folders:
```yaml
lab:
  name: "Chen Lab"
  shared_folders:
    - path: /shared/datasets
      s3_bucket: s3://chen-lab-shared/datasets
      permissions:
        - user: "*"  # all lab members
          access: read-only
    - path: /shared/projects/project-alpha
      s3_bucket: s3://chen-lab-shared/project-alpha
      permissions:
        - user: alice@uni.edu
          access: read-write
        - user: bob@uni.edu
          access: read-write
        - user: "*"
          access: read-only
```
- Mount shared folders at launch time
- S3FS or similar for S3-backed filesystem
- Lazy loading for large datasets
- Caching for frequently accessed data
```

---

## 💰 v0.11.0 - Cost Management for Labs (4 Issues)

### Issue #25: Add budget alerts with email notifications
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.11-cost-mgmt`, `area: costs`, `persona: lab-pi`, `persona: instructor`, `use-case: cost-optimization`
**Personas:** Lab PI, Course Instructor, Research Computing Manager
**Requirements:** REQ-12.1 (Budget Tracking), REQ-12.3 (Cost Alerts)
**Pain Points:**
- Lab PI Pain #1 ($4,000 cost overrun, blind to spending) - USER_SCENARIOS/03
- Instructor Pain #7 (no warning before exceeding $2K budget) - USER_SCENARIOS/04
- Research Computing Manager Pain #2 ($47K overspend across institution) - USER_SCENARIOS/05

**Description:**
```
As a Lab PI with a $15,000 annual budget, I need email alerts when my lab approaches the budget limit, because I had a $4,000 overrun last year.

**Current Pain Point:**
No visibility into spending until AWS bill arrives. By then, damage is done. Need proactive alerts to take action before exceeding budget.

**Success Metric:**
✅ Alerts at 50%, 75%, 90% of budget
✅ Daily digest for lab managers
✅ Auto-stop option when budget exceeded
✅ 100% of cost overruns prevented

**Proposed Solution:**
Budget configuration:
```yaml
budget:
  monthly: $1250  # $15K / 12 months
  per_researcher:
    alice@uni.edu: $300
    bob@uni.edu: $400
  alerts:
    email: pi@university.edu
    thresholds: [50, 75, 90, 100]
  enforcement:
    auto_stop_at_100: true
    grace_period: 24h  # warning before hard stop
    emergency_override: true  # PI can override
```
- Email alerts:
  - "⚠️ Chen Lab budget: 50% used ($625 of $1,250 this month)"
  - "🔥 URGENT: 90% budget used, auto-stop in 24 hours"
- Daily digest: "Lab spending: $47/day (on track for $1,410 this month)"
- Per-researcher tracking: "Alice: $127 of $300 (42%)"
- Auto-stop behavior:
  - At 100%: send warning email, 24hr grace period
  - After grace period: stop all running instances
  - PI override: `lens-jupyter config budget override --reason "grant deadline"`
- CLI: `lens-jupyter costs budget-status`
```

---

### Issue #26: Add cost reporting for grant reporting and allocation
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 0.11-cost-mgmt`, `area: costs`, `persona: lab-pi`, `persona: it-admin`, `use-case: cost-optimization`
**Personas:** Lab PI, Research Computing Manager
**Requirements:** REQ-12.4 (Cost Reporting), REQ-12.5 (Grant Reporting)
**Pain Points:**
- Lab PI Pain #1 (need reports for grant renewals) - USER_SCENARIOS/03
- Research Computing Manager Pain #2 (institutional reporting requirements) - USER_SCENARIOS/05

**Description:**
```
As a Lab PI submitting grant renewal, I need detailed cost reports by project and researcher, because the grant requires documentation of computational resources used.

**Current Pain Point:**
AWS Cost Explorer is too complex for academic users. Need reports formatted for grant applications with per-project and per-researcher breakdowns.

**Success Metric:**
✅ PDF and CSV export for grant applications
✅ Customizable date ranges (grant period)
✅ Breakdown by researcher, project, instance type
✅ Suitable for NSF/NIH grant reporting

**Proposed Solution:**
Cost reporting commands:
```bash
# Generate grant report
lens-jupyter costs report \
  --start-date 2024-09-01 \
  --end-date 2025-08-31 \
  --format pdf \
  --output grant-renewal-2025.pdf

# Per-project breakdown
lens-jupyter costs report \
  --project project-alpha \
  --breakdown researcher \
  --format csv

# Lab summary for PI
lens-jupyter costs summary --monthly
```

Report includes:
- Executive summary: Total costs, trend, efficiency metrics
- Per-researcher breakdown: Alice ($487), Bob ($693), etc.
- Per-project breakdown: Project Alpha ($1,247), Project Beta ($934)
- Instance type breakdown: t4g.medium (45%), t4g.xlarge (30%)
- Timeline visualization: spending over grant period
- Optimization recommendations: "Switch to Spot for 60% savings"
- Grant-ready formatting: matches NSF/NIH requirements

Tagging system for projects:
```bash
lens-jupyter launch --project project-alpha --grant NSF-12345
```
- Tags propagate to EC2 instances
- AWS Cost Allocation Tags for reporting
- Query by project code in reports
```

---

### Issue #27: Add usage pattern analysis and optimization recommendations
**Priority:** MEDIUM
**Labels:** `enhancement`, `phase: 0.11-cost-mgmt`, `area: costs`, `persona: lab-pi`, `persona: it-admin`, `use-case: cost-optimization`
**Personas:** Lab PI, Research Computing Manager
**Requirements:** REQ-12.6 (Cost Optimization), REQ-6.2 (Spot Instance Support)
**Pain Points:**
- Lab PI Pain #8 (don't know how to optimize costs) - USER_SCENARIOS/03
- Research Computing Manager Pain #7 (no institutional optimization guidance) - USER_SCENARIOS/05

**Description:**
```
As a Lab PI, I need recommendations on how to reduce costs, because I don't understand AWS pricing well enough to optimize on my own.

**Current Pain Point:**
Researchers don't know about Spot instances, right-sizing, or Reserved Instances. Miss 50-90% potential savings.

**Success Metric:**
✅ Actionable recommendations: "Switch to Spot for 70% savings"
✅ One-click optimization: accept recommendation → automatically apply
✅ 20%+ cost reduction for average lab

**Proposed Solution:**
Usage analysis command:
```bash
lens-jupyter costs analyze
```

Output:
```
💡 Cost Optimization Recommendations for Chen Lab

1. ⚠️ HIGH IMPACT: Spot Instances (Est. savings: $340/month, 68%)
   - Your workloads are 95% interruptible
   - Current: t4g.xlarge On-Demand ($0.134/hr)
   - Recommendation: t4g.xlarge Spot ($0.040/hr)
   - Action: `lens-jupyter config set default.spot_instances true`

2. 💡 MEDIUM IMPACT: Right-sizing (Est. savings: $80/month, 16%)
   - Alice's instance idle 80% of the time
   - Current: t4g.xlarge (4 vCPU, 16GB)
   - Recommendation: t4g.large (2 vCPU, 8GB)
   - Action: `lens-jupyter resize --instance i-abc123 --type t4g.large`

3. 💵 LOW IMPACT: Reserved Instances (Est. savings: $45/month, 9%)
   - Bob runs 24/7 for 3+ months
   - Current: On-Demand pricing
   - Recommendation: 3-month RI (30% discount)
   - Action: Contact AWS for RI purchase

4. ⏱️ AUTO-STOP: Instance left running (Cost: $67 wasted this month)
   - 3 instances running >12hrs with no activity
   - Recommendation: Enable auto-stop after 2hr idle
   - Action: `lens-jupyter config set auto_stop.enabled true`

Total Potential Savings: $465/month (93% of current spend)
```

Analysis includes:
- Spot instance suitability analysis (workload patterns)
- Idle time detection (right-sizing recommendations)
- Reserved Instance ROI calculation
- Auto-stop opportunities
- One-click accept: `lens-jupyter costs optimize --accept-all`
```

---

### Issue #28: Add cost forecasting with academic calendar awareness
**Priority:** LOW
**Labels:** `enhancement`, `phase: 0.11-cost-mgmt`, `area: costs`, `persona: lab-pi`, `persona: instructor`, `use-case: cost-optimization`
**Personas:** Lab PI, Course Instructor
**Requirements:** REQ-12.7 (Cost Forecasting)
**Pain Points:**
- Lab PI Pain #1 (unpredictable monthly costs) - USER_SCENARIOS/03
- Instructor Pain #7 (semester budget planning) - USER_SCENARIOS/04

**Description:**
```
As a Lab PI, I need to predict my monthly costs based on current usage, because I need to plan my annual budget and avoid overruns.

**Current Pain Point:**
Academic research has seasonal patterns (summer intensive work, winter break lull). Need forecasting that understands academic calendar.

**Success Metric:**
✅ Predict monthly costs within ±10% accuracy
✅ Alert when trending over budget
✅ Academic calendar adjustment (summer surge, winter lull)

**Proposed Solution:**
Cost forecasting command:
```bash
lens-jupyter costs forecast --months 6
```

Output:
```
📊 Cost Forecast for Chen Lab (6 months)

Current Usage: $487/month
Trend: ↗️ Increasing 15% month-over-month

Month        Forecast   Budget   Status   Notes
------------------------------------------------------
Nov 2025     $560      $1,250    ✅ OK    Conference deadline surge
Dec 2025     $180      $1,250    ✅ OK    Winter break lull
Jan 2026     $520      $1,250    ✅ OK    Spring semester start
Feb 2026     $580      $1,250    ✅ OK
Mar 2026     $610      $1,250    ✅ OK    Grant renewal analysis
Apr 2026     $590      $1,250    ✅ OK

6-month total: $3,040 of $7,500 budget (41%)

⚠️ Alert: Trending over annual budget by 8% if current growth continues
💡 Recommendation: Enable auto-stop to reduce forecast by $120/month
```

Academic calendar awareness:
- Configure academic calendar:
```yaml
forecasting:
  academic_calendar:
    fall_semester: [2025-09-01, 2025-12-15]
    winter_break: [2025-12-15, 2026-01-15]
    spring_semester: [2026-01-15, 2026-05-15]
    summer_research: [2026-05-15, 2026-08-31]
  seasonal_patterns:
    conference_deadlines: [2025-11-01, 2026-03-01]  # surge
    winter_break: [2025-12-15, 2026-01-05]  # lull
```
- Forecast adjusts for known patterns
- Alert: "Conference deadline approaching, expect 20% cost increase"
```

---

## 🚀 v1.0.0 - Production Ready (2 Issues)

### Issue #29: Create video tutorials for all research tools and workflows
**Priority:** HIGH
**Labels:** `documentation`, `phase: 1.0-production`, `area: docs`, `persona: solo-researcher`, `persona: graduate-student`, `persona: instructor`, `use-case: teaching`
**Personas:** Solo Researcher, Graduate Student, Course Instructor
**Requirements:** REQ-1.4 (Video Tutorials), REQ-13.2 (Educational Resources)
**Pain Points:**
- Solo Researcher Pain #1 (learning curve for cloud) - USER_SCENARIOS/01
- Graduate Student Pain #1 (no time to read documentation) - USER_SCENARIOS/02
- Instructor Pain #8 (students need self-serve resources) - USER_SCENARIOS/04

**Description:**
```
As a Solo Researcher new to cloud computing, I need video tutorials showing complete workflows, because I learn better from watching than reading documentation.

**Current Pain Point:**
Text documentation exists but non-technical researchers prefer video. No visual walkthroughs of complete research workflows.

**Success Metric:**
✅ 15+ video tutorials covering all major use cases
✅ Each video <10 minutes (short, focused)
✅ 90% of users can complete task after watching video
✅ Hosted on YouTube with transcripts

**Proposed Solution:**
Create video tutorial series:

**Getting Started (3 videos):**
1. "Your First Jupyter Lab in 5 Minutes" (Solo Researcher)
2. "GPU Instance for Deep Learning" (Graduate Student)
3. "Setting Up a Class of 30 Students" (Instructor)

**Research Workflows (8 videos):**
4. "Genomics Analysis Pipeline (GATK, BWA, SAMtools)"
5. "Machine Learning with PyTorch on GPU"
6. "R Statistical Analysis and Visualization"
7. "Creating Interactive Data Apps with Streamlit"
8. "Reproducible Research with Quarto"
9. "Collaborative Analysis with Shared Instances"
10. "Cost Management for Research Labs"
11. "Bioinformatics Analysis with BioConda"

**Troubleshooting (4 videos):**
12. "Common Errors and How to Fix Them"
13. "Optimizing Costs with Spot Instances"
14. "Backing Up Your Work to S3"
15. "Setting Up Team Workspaces"

Production specs:
- Screen recordings with voiceover
- Real research examples (not toy datasets)
- Step-by-step with pauses
- Closed captions for accessibility
- Downloadable scripts/code from video
- Host on YouTube, embed in docs
- Create video landing page: docs/videos/
```

---

### Issue #30: Conduct academic beta testing program and incorporate feedback
**Priority:** HIGH
**Labels:** `enhancement`, `phase: 1.0-production`, `persona: solo-researcher`, `persona: graduate-student`, `persona: lab-pi`, `persona: instructor`, `persona: it-admin`, `use-case: data-analysis`, `use-case: machine-learning`, `use-case: bioinformatics`, `use-case: teaching`
**Personas:** All personas
**Requirements:** REQ-14.1 (Academic Adoption), REQ-14.2 (User Feedback)
**Pain Points:** All pain points across all personas - validation of solutions

**Description:**
```
As the AWS IDE project maintainer, I need to recruit 100+ academic researchers to beta test the tool, because real-world usage will reveal issues and missing features.

**Current Pain Point:**
Tool is developed in relative isolation. Need real feedback from diverse research domains and institutions.

**Success Metric:**
✅ 100+ active beta testers across 5+ institutions
✅ 10+ research domains represented
✅ 50+ bug reports and feature requests collected
✅ Incorporate top 20 requests into v1.0
✅ User satisfaction: 4.5+/5.0 average rating

**Proposed Solution:**
Academic Beta Testing Program:

**Phase 1: Recruitment (Month 1)**
- Outreach to universities:
  - Research computing centers
  - Data science departments
  - Bioinformatics programs
  - Social science methods labs
- Target institutions:
  - Large R1 universities (5+)
  - Liberal arts colleges (3+)
  - Community colleges (2+)
- Recruit diverse personas:
  - 30 Solo Researchers
  - 30 Graduate Students
  - 20 Lab PIs
  - 15 Course Instructors
  - 5 Research Computing Managers

**Phase 2: Onboarding (Month 2)**
- Onboarding workshops (virtual):
  - "Getting Started with AWS IDE" (for all users)
  - "Teaching with AWS IDE" (for instructors)
  - "Managing Lab Costs" (for PIs)
- Provide beta testing resources:
  - AWS credits ($500/tester for computing costs)
  - Dedicated Slack channel for support
  - Weekly office hours (Q&A sessions)

**Phase 3: Testing (Months 3-6)**
- Structured feedback collection:
  - Weekly usage surveys (5 min)
  - Monthly in-depth interviews (30 min)
  - Bug report template on GitHub
  - Feature request template on GitHub
- Usage tracking (anonymized):
  - Which features are most used?
  - Where do users get stuck?
  - What errors occur frequently?
- Community building:
  - Beta tester showcase: published papers using AWS IDE
  - Environment sharing: best practices and templates
  - Peer support: experienced users help newcomers

**Phase 4: Analysis & Iteration (Month 7)**
- Analyze feedback:
  - Top 20 bug fixes (prioritize for v1.0)
  - Top 20 feature requests (prioritize for v1.0)
  - Usability issues (fix before v1.0)
- User satisfaction survey:
  - Would you recommend AWS IDE? (NPS score)
  - Rate overall experience (1-5 stars)
  - What's missing for your research?
- Document success stories:
  - "How AWS IDE enabled my Nature paper"
  - "We saved $12K this year vs commercial tools"
  - "My students published their first papers"

**Deliverables:**
- Beta testing report (PDF, 20-30 pages)
- Prioritized feature backlog from feedback
- Academic testimonials and case studies
- Community of 100+ researchers for launch
```

---

## 📋 Traceability Matrix

### Issues → Personas Mapping

| Persona | Issue Count | Key Issues |
|---------|-------------|------------|
| Solo Researcher | 21 | #1, #2, #3, #7, #8, #14, #16, #22 |
| Graduate Student | 20 | #1, #2, #3, #4, #7, #10, #12, #14, #16, #17, #20, #22 |
| Lab PI | 18 | #5, #8, #10, #12, #14, #16, #17, #19, #20, #21, #22, #25, #26, #27 |
| Course Instructor | 12 | #1, #2, #5, #8, #12, #16, #19, #20, #23, #25, #28, #29 |
| Research Computing Manager | 8 | #9, #19, #21, #24, #25, #26, #27, #30 |

### Issues → Roadmap Phases Mapping

| Phase | Issues |
|-------|--------|
| v0.7.0 - User Experience & Accessibility | #1, #2, #3, #4, #5, #6 |
| v0.8.0 - Additional Research Tools | #7, #8, #9, #10, #11, #12, #13 |
| v0.9.0 - Package Managers & Reproducibility | #14, #15, #16, #17, #18, #19 |
| v0.10.0 - Collaboration Features | #20, #21, #22, #23, #24 |
| v0.11.0 - Cost Management for Labs | #25, #26, #27, #28 |
| v1.0.0 - Production Ready | #29, #30 |

### Issues → USER_REQUIREMENTS.md Mapping

| Requirement | Related Issues |
|-------------|----------------|
| REQ-1.1 (Beginner-Friendly Onboarding) | #1, #2 |
| REQ-1.2 (Plain-English Interface) | #2 |
| REQ-1.3 (Quick Re-launch) | #3 |
| REQ-1.4 (Video Tutorials) | #29 |
| REQ-2.3 (Interactive Visualizations) | #8, #13 |
| REQ-2.4 (AI-Assisted Development) | #7 |
| REQ-3.6 (New Research Tools) | #7, #8, #9, #10, #11, #12, #13 |
| REQ-3.7 (Multi-Language Support) | #9 |
| REQ-3.8 (GUI Application Support) | #10 |
| REQ-4.2 (Custom Environments) | #18 |
| REQ-4.3 (Environment Export/Import) | #14, #17 |
| REQ-5.4 (Status Notifications) | #4, #5, #6 |
| REQ-6.1 (Auto-Stop Idle Instances) | #4 |
| REQ-6.2 (Spot Instance Support) | #27 |
| REQ-6.3 (GPU Instance Support) | #10 |
| REQ-7.1 (Fast Launch Times) | #1, #6 |
| REQ-8.1 (Instance Sharing) | #20 |
| REQ-8.2 (Shareable Results) | #8 |
| REQ-8.3 (Team Communication) | #5 |
| REQ-8.4 (Community Sharing) | #19 |
| REQ-8.5 (Team Workspaces) | #21, #24 |
| REQ-8.6 (Data Persistence) | #22, #24 |
| REQ-8.7 (Multi-User Environments) | #23 |
| REQ-9.1 (Reproducible Research) | #12, #14, #17 |
| REQ-9.3 (Package Management) | #14, #15, #18 |
| REQ-9.4 (Backup and Recovery) | #22 |
| REQ-10.1 (Domain-Specific Environments) | #15, #16, #19 |
| REQ-10.2 (Research Use Case Coverage) | #16 |
| REQ-11.1 (Centralized Configuration) | #21 |
| REQ-11.2 (Access Control) | #20 |
| REQ-11.3 (Secure Collaboration) | #20 |
| REQ-12.1 (Budget Tracking) | #25 |
| REQ-12.2 (Budget Allocation) | #21 |
| REQ-12.3 (Cost Alerts) | #25 |
| REQ-12.4 (Cost Reporting) | #26 |
| REQ-12.5 (Grant Reporting) | #26 |
| REQ-12.6 (Cost Optimization) | #27 |
| REQ-12.7 (Cost Forecasting) | #28 |
| REQ-13.1 (Teaching Support) | #23 |
| REQ-13.2 (Educational Resources) | #29 |
| REQ-14.1 (Academic Adoption) | #30 |
| REQ-14.2 (User Feedback) | #30 |

### Issues → USER_SCENARIOS Walkthrough Mapping

| Walkthrough | Pain Points Addressed | Related Issues |
|-------------|----------------------|----------------|
| 01_SOLO_RESEARCHER | Pain #1 (setup time), #2 (lost work), #3 (repetitive config), #4 (bioinformatics), #5 (learning), #6 (sharing), #7 (occasional coding), #8 (software install) | #1, #2, #3, #7, #8, #13, #14, #15, #16, #22 |
| 02_GRADUATE_STUDENT | Pain #1 (no docs time), #2 (GPU/GUI), #3 (forgot to stop), #4 (context switch), #5 (debugging), #6 (IDE learning), #7 (paper writing), #8 (dependencies), #9 (genomics), #10 (domain tools), #11 (sharing with advisor) | #1, #2, #3, #4, #7, #10, #11, #12, #14, #15, #16, #17, #18, #20, #22, #29 |
| 03_LAB_PI | Pain #1 (cost overrun), #2 (no visibility), #4 (collaboration), #5 (diverse tools), #6 (reproducibility), #7 (no backups), #8 (optimization) | #5, #8, #9, #10, #12, #14, #16, #17, #19, #20, #21, #22, #24, #25, #26, #27, #28 |
| 04_COURSE_INSTRUCTOR | Pain #1 (8hr install), #2 (managing instances), #3 (teaching code), #4 (monitoring), #5 (course materials), #6 (discipline-specific), #7 (budget planning), #8 (self-serve) | #1, #2, #5, #8, #12, #16, #19, #20, #23, #25, #28, #29 |
| 05_RESEARCH_COMPUTING_MANAGER | Pain #2 (overspend), #3 (supporting diverse), #4 (disciplines), #5 (inconsistent setups), #6 (institutional), #7 (optimization) | #9, #19, #21, #24, #25, #26, #27, #30 |

---

## 📊 Priority Distribution

| Priority | Count | Issues |
|----------|-------|--------|
| 🔥 HIGH | 14 | #1, #2, #7, #8, #10, #14, #16, #20, #22, #25, #26, #29, #30 |
| MEDIUM | 8 | #3, #9, #12, #15, #17, #21, #27 |
| LOW | 8 | #4, #5, #6, #11, #13, #18, #19, #23, #24, #28 |

---

## 🔄 Dependency Graph

Issues with dependencies (must complete X before Y):

- **#14 (conda support)** → #15 (BioConda), #16 (domain templates), #17 (env export/import), #18 (system packages)
- **#2 (wizard default)** → #3 (remember preferences)
- **#7 (Q Developer)**, #8 (Streamlit), #9 (Zeppelin), #10 (DCV Desktop), #11 (Theia), #12 (Quarto), #13 (Observable) → #29 (video tutorials) - need tools built before creating tutorials
- **#1-28** → #30 (beta testing) - beta testing requires most features implemented
- **#20 (instance sharing)** → #23 (JupyterHub) - sharing is foundation for multi-user
- **#21 (team workspaces)** → #24 (shared folders) - workspaces enable shared storage

---

## 📝 Next Steps

1. **Create GitHub issues** using this summary as reference (Week 3)
2. **Configure GitHub Project board** with custom fields (Week 4)
3. **Prioritize v0.7.0 issues** for immediate work (#1, #2, #3)
4. **Begin v0.8.0 planning** once v0.7.0 complete
5. **Regular review** of this traceability document as issues evolve

---

**Document Purpose:** This summary serves as the authoritative mapping between ROADMAP.md, GitHub Issues, USER_SCENARIOS/, and USER_REQUIREMENTS.md. It enables bidirectional traceability and ensures all development work aligns with user needs.
