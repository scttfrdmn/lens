# Tool Support Strategy

This document outlines the strategy for expanding Lens to support a comprehensive suite of research tools beyond the initial Jupyter/RStudio/VSCode offerings.

## Overview

Lens aims to support the full spectrum of research software - from open-source command-line tools to commercial GUI applications - making any research tool accessible via a simple CLI command.

---

## Tool Categories

### By Access Method

#### 1. **Web-Based Tools** (Current Architecture)
- Access via browser
- Port forwarding via SSH or Session Manager
- Examples: Jupyter, RStudio, VSCode, OpenRefine

**Current Support**: ✅ Fully implemented

#### 2. **Desktop GUI Tools** (Requires DCV)
- Full desktop environment required
- GPU acceleration for some tools
- Examples: MATLAB, ArcGIS, QGIS, ParaView, ImageJ

**Current Support**: ⏳ Planned (NICE DCV in v0.10.0)

#### 3. **Hybrid Tools** (Web + Desktop)
- Can work in both modes
- Examples: ParaView (has web viewer), Octave (has web UI)

**Current Support**: ⏳ Partial

---

### By License Type

#### 1. **Open Source** (Easiest)
- No licensing concerns
- Can build custom AMIs
- Free to use
- Examples: QGIS, ParaView, Octave, OpenRefine, ImageJ/Fiji

**Implementation**: Standard app template

#### 2. **Cloud-Authenticated (Modern Pattern)**
- User logs in with institutional/vendor credentials
- License validated automatically via cloud
- No license configuration needed
- Examples: MATLAB (MathWorks account), ArcGIS (ArcGIS Online), Mathematica (Wolfram ID)

**Implementation**: Just install software + DCV desktop (SIMPLEST!)

#### 3. **AWS Marketplace (Pay-as-you-go)**
- Software cost included in EC2 hourly rate
- Automatic billing through AWS
- No license management needed
- Examples: Various commercial tools

**Implementation**: AMI marketplace integration required

#### 4. **Legacy BYOL (License Files/Servers)**
- Older software with traditional licensing
- User provides license key/file or server address
- Becoming less common
- Examples: Older versions of commercial software, some specialized tools

**Implementation**: License configuration system (only if needed)

#### 5. **Subscription/Cloud Native**
- Vendor-hosted licensing
- May require API keys or account setup
- Examples: Some cloud-based research platforms

**Implementation**: Case-by-case integration

---

## ⚡ Simplified Licensing Reality

**Modern commercial tools** (MATLAB, ArcGIS, Mathematica, etc.) increasingly use **cloud authentication**:
1. Install software on AMI
2. User launches application via DCV
3. User logs in with credentials (university email, MathWorks account, etc.)
4. Application validates license automatically via cloud

**This means**: Most commercial tools are as simple to support as open-source tools - just need the software installed and DCV working. No complex license management needed!

---

## Tool Authentication Patterns

Understanding which tools use modern cloud authentication vs legacy licensing helps prioritize implementation.

### ✅ Cloud Authentication (Simple - Just Install + DCV)

**Commercial Tools - Users Log In:**
- **MATLAB** - MathWorks account or institutional credentials (OAuth/SAML)
- **Mathematica** - Wolfram ID account login
- **ArcGIS Pro** - ArcGIS Online credentials (most common for academic users)
- **Geneious** - Geneious account with subscription
- **Origin** - OriginLab account login (newer versions)
- **GraphPad Prism** - GraphPad account login

**How it works:**
1. Launch tool via DCV desktop
2. Tool prompts for login when opened
3. User enters credentials
4. Tool validates with vendor's cloud service
5. Ready to use!

**Implementation effort**: LOW - Just install software on AMI

### 🔓 Open Source (Simplest - No Authentication)

**No licensing concerns:**
- QGIS, ParaView, ImageJ/Fiji, Octave, OpenRefine
- Blender, MeshLab, CloudCompare
- CellProfiler, Cytoscape, UGENE
- SNAP, Maxima, OpenFOAM

**Implementation effort**: LOWEST - Install and go

### 🔐 Legacy License Files/Servers (Complex - Deferred to v0.16+)

**Require Configuration:**
- **Stata** - License file or network license server
- **SPSS** - License file or license server
- **SAS** - License file (though has cloud options)
- **Ansys** - License server typically
- **COMSOL** - License file or server
- **ENVI/ERDAS** - License files

**How it works:**
1. User provides license file or server address
2. Lens uploads/configures license
3. Tool validates against file/server
4. Ready to use

**Implementation effort**: HIGHER - Needs license management infrastructure

### 💰 AWS Marketplace (Medium Complexity)

**License included in hourly cost:**
- Various tools available with pay-as-you-go pricing
- MATLAB, Mathematica, and others have marketplace options
- No license management needed, but higher hourly costs

**Implementation effort**: MEDIUM - Need marketplace AMI integration

---

## Tool Inventory

### Requested Tools

| Tool | Type | License | Domain | GUI Required | Priority |
|------|------|---------|--------|--------------|----------|
| **MATLAB** | Commercial | Marketplace/BYOL | Engineering, Data Science | Yes | 🔥 Highest |
| **ArcGIS** | Commercial | Marketplace/BYOL | GIS, Geography | Yes | 🔥 High |
| **QGIS** | Open Source | Free | GIS, Geography | Yes | 🔥 High |
| **ParaView** | Open Source | Free | Visualization, Engineering | Yes/Web | 🔥 High |
| **OpenRefine** | Open Source | Free | Data Cleaning | Web | Medium |
| **Octave** | Open Source | Free | Math, Engineering | Yes/Web | Medium |

### Additional High-Value Tools

#### **Bioinformatics & Life Sciences**
| Tool | Type | License | GUI | Priority |
|------|------|---------|-----|----------|
| **ImageJ/Fiji** | Open Source | Free | Yes | 🔥 High |
| **PyMOL** | Open/Commercial | Freemium | Yes | High |
| **Geneious** | Commercial | Subscription | Yes | High |
| **CellProfiler** | Open Source | Free | Yes | Medium |
| **Cytoscape** | Open Source | Free | Yes | Medium |
| **UGENE** | Open Source | Free | Yes | Low |

#### **Statistics & Economics**
| Tool | Type | License | GUI | Priority |
|------|------|---------|-----|----------|
| **Stata** | Commercial | BYOL | Yes | 🔥 High |
| **SPSS** | Commercial | BYOL/Subscription | Yes | High |
| **SAS** | Commercial | Marketplace | Yes/Web | Medium |
| **EViews** | Commercial | BYOL | Yes | Low |

#### **Mathematics & Symbolic Computation**
| Tool | Type | License | GUI | Priority |
|------|------|---------|-----|----------|
| **Mathematica** | Commercial | Marketplace/BYOL | Yes | 🔥 High |
| **Maple** | Commercial | BYOL | Yes | Medium |
| **Maxima** | Open Source | Free | Yes | Low |

#### **Engineering & Simulation**
| Tool | Type | License | GUI | Priority |
|------|------|---------|-----|----------|
| **Ansys** | Commercial | Marketplace/BYOL | Yes | Medium |
| **COMSOL** | Commercial | BYOL | Yes | Medium |
| **OpenFOAM** | Open Source | Free | CLI | Low |

#### **3D Visualization & Modeling**
| Tool | Type | License | GUI | Priority |
|------|------|---------|-----|----------|
| **Blender** | Open Source | Free | Yes | Medium |
| **MeshLab** | Open Source | Free | Yes | Low |
| **CloudCompare** | Open Source | Free | Yes | Low |

#### **Remote Sensing & Geospatial**
| Tool | Type | License | GUI | Priority |
|------|------|---------|-----|----------|
| **ENVI** | Commercial | BYOL | Yes | Medium |
| **ERDAS IMAGINE** | Commercial | BYOL | Yes | Low |
| **SNAP** (ESA) | Open Source | Free | Yes | Medium |

#### **Data Analysis & Graphing**
| Tool | Type | License | GUI | Priority |
|------|------|---------|-----|----------|
| **OriginPro** | Commercial | BYOL | Yes | Low |
| **GraphPad Prism** | Commercial | Subscription | Yes | Medium |
| **Tableau Desktop** | Commercial | Subscription | Yes | Low |

---

## Architectural Approaches

### Option 1: Individual App Model (Current)
**Structure**: Separate binary for each tool (`lens-matlab`, `lens-arcgis`, etc.)

**Pros**:
- Clear, predictable UX
- Tool-specific optimizations
- Easy to maintain per-tool

**Cons**:
- Many binaries to maintain
- Slow to add new tools
- Large overall footprint

**Best For**: High-priority, heavily-used tools (MATLAB, ArcGIS, QGIS)

### Option 2: Universal Launcher with Catalog
**Structure**: Single `lens-tool` command with tool catalog

```bash
lens-tool launch matlab
lens-tool launch arcgis
lens-tool list --category gis
lens-tool search visualization
```

**Pros**:
- Easy to add new tools
- Consistent UX across all tools
- Smaller footprint

**Cons**:
- Less tool-specific customization
- More complex architecture
- One-size-fits-all approach

**Best For**: Long-tail of less-common tools

### Option 3: Hybrid Model (Recommended)
**Structure**: Individual apps for major tools + catalog for others

```bash
# Major tools get dedicated commands
lens-matlab launch
lens-qgis launch

# Others use universal launcher
lens-tool launch paraview
lens-tool launch octave
```

**Pros**:
- Best of both worlds
- Flexibility for different tool types
- Can optimize where it matters

**Cons**:
- Two systems to maintain
- Need clear criteria for which approach

**Criteria for Individual App**:
- Used by >1000 researchers
- Requires complex setup/licensing
- Benefits from tool-specific features
- Has multiple deployment options (marketplace, BYOL, etc.)

---

## Implementation Components

### 1. AMI Catalog System

**Purpose**: Central database of available AMIs for each tool

**Structure**:
```yaml
tools:
  matlab:
    name: "MATLAB"
    vendor: "MathWorks"
    category: ["mathematics", "engineering", "data-science"]
    access_method: "gui"
    amis:
      - type: "marketplace"
        ami_search: "matlab-r2024b"
        license: "included"
        cost_markup: "high"
      - type: "byol"
        ami_search: "matlab-byol-r2024b"
        license: "user-provided"
        cost_markup: "none"
      - type: "community"
        ami_id: "ami-xxxxx"
        license: "user-provided"
    license_config:
      type: "file"  # or "server", "key", "account"
      prompt: "Path to license file or license server address"
```

**Features**:
- Searchable by tool name, category, vendor
- Supports multiple AMI options per tool
- Includes pricing/cost information
- License requirement documentation

### 2. License Manager (Minimal - Most Tools Don't Need This!)

**Purpose**: Handle licensing for tools that need it

**Reality Check**: Modern commercial tools (MATLAB, ArcGIS, Mathematica) use cloud authentication - users just log in with credentials. License management is mostly unnecessary!

**Only needed for**:
- Legacy tools with license files
- Some specialized on-premise tools
- Tools without cloud authentication

**Capabilities** (if needed):
- Store optional license file references
- Document login procedures per tool
- Provide pre-launch reminders ("Have your MathWorks credentials ready")

**Example Flow** (Modern MATLAB):
```bash
$ lens-matlab launch

✓ Launching MATLAB R2024b
✓ Desktop ready at https://54.123.45.67:8443

💡 Reminder: You'll need to log in with your MathWorks or institutional
   credentials when MATLAB starts.

Opening desktop...
```

**Example Flow** (Legacy tool with license file):
```bash
$ lens-legacy-tool launch --license-file ~/tool.lic

✓ Uploading license file...
✓ Launching tool...
```

### 3. NICE DCV Integration

**Purpose**: Enable GUI tool support

**Architecture**:
- Base DCV desktop environment
- Tool-specific DCV images with pre-installed software
- Automatic DCV client connection

**User Experience**:
```bash
$ lens-matlab launch
...
✓ MATLAB desktop ready!

Connection methods:
1. Browser: https://12.34.56.78:8443
2. DCV Client: dcv://12.34.56.78
3. SSH tunnel: ssh -L 8443:localhost:8443 ...

Opening browser to MATLAB desktop...
```

### 4. Tool Discovery System

**Commands**:
```bash
# List all available tools
lens tools list

# Filter by category
lens tools list --category gis
lens tools list --category bioinformatics

# Search by name or keyword
lens tools search matlab
lens tools search "image analysis"

# Get detailed info
lens tools info matlab
```

**Output Example**:
```
$ lens tools info matlab

MATLAB R2024b
By MathWorks

Category: Mathematics, Engineering, Data Science
Access: Desktop GUI (NICE DCV)
License: Cloud Authentication (log in with MathWorks/institutional credentials)

How it works:
1. Launch MATLAB desktop
2. Log in with your credentials (university email or MathWorks account)
3. License validated automatically via cloud
4. Start working!

No license configuration needed - just have your credentials ready.

Recommended Instance Types:
- Light work: t3.large ($0.0832/hr)
- Standard: c6i.xlarge ($0.17/hr)
- Heavy computation: c6i.4xlarge ($0.68/hr)
- GPU work: g4dn.xlarge ($0.526/hr)

To launch:
  lens-matlab launch
  lens-matlab launch --license-file ~/matlab.lic
```

---

## Phased Rollout Plan

### Phase 1: v0.10.0 - GUI Foundation (2-3 months)

**Goal**: Enable GUI tool support with DCV

**Deliverables**:
- `lens-dcv-desktop` - Generic Ubuntu desktop
- NICE DCV integration fully working
- User can manually install any GUI tool
- Pre-installed: Firefox, VS Code, terminal

**Environments**:
- `general-desktop` - Basic Ubuntu desktop
- `gpu-desktop` - With CUDA for GPU tools

**Success Criteria**:
- User can launch DCV desktop
- Can install QGIS manually and use it
- Can install ImageJ and analyze images
- GPU tools can use GPU acceleration

**Note**: v0.11.0-v0.13.0 focus on package managers, collaboration, and cost management (see ROADMAP.md)

### Phase 2: v0.14.0 - Open Source GUI Tools

**Goal**: Add pre-configured open-source GUI tools

**New Applications**:
1. **lens-qgis** (GIS analysis)
   - Pre-installed QGIS with common plugins
   - Sample datasets included
   - Environments: `basic-gis`, `advanced-gis`

2. **lens-paraview** (Visualization)
   - ParaView with OSPRay rendering
   - Environments: `visualization`, `gpu-visualization`

3. **lens-imagej** (Image analysis)
   - Fiji distribution with common plugins
   - Environments: `microscopy`, `cell-analysis`

4. **lens-octave** (MATLAB alternative)
   - Octave with web UI option
   - Environments: `mathematics`, `signal-processing`

5. **lens-openrefine** (Data cleaning)
   - Web-based interface
   - Sample datasets and transformations

**Infrastructure**:
- AMI catalog system (alpha)
- Tool discovery commands
- Shared DCV base images

**Why This Is Simple**: All open source tools - just install and go!

### Phase 3: v0.15.0 - Cloud-Authenticated Commercial Tools ⚡

**Goal**: Support modern commercial tools with cloud authentication

**Why This Is Easy**: These tools handle licensing via user login - just install software + DCV!

**New Applications**:
1. **lens-matlab** (Cloud auth - MathWorks account)
   - User logs in with MathWorks or institutional credentials
   - No license configuration needed
   - Environments: `engineering`, `data-science`, `control-systems`

2. **lens-mathematica** (Cloud auth - Wolfram ID)
   - User logs in with Wolfram account
   - Automatic license validation
   - Environments: `symbolic-math`, `data-science`

3. **lens-arcgis** (Cloud auth - ArcGIS Online)
   - User logs in with ArcGIS Online credentials
   - Environments: `urban-planning`, `environmental`, `remote-sensing`

4. **lens-geneious** (Cloud auth - Geneious account)
   - Subscription-based cloud authentication
   - Environments: `bioinformatics`, `genomics`

**Infrastructure**:
- AMI marketplace search/selection (for marketplace options)
- Pricing calculator with software costs
- Simple pre-launch credential reminders (no complex license manager needed!)

**User Experience**:
```bash
$ lens-matlab launch

✓ Launching MATLAB R2024b
✓ Desktop ready!

💡 Reminder: Log in with your MathWorks or institutional credentials when MATLAB starts.

Opening desktop...
```

### Phase 4: v0.16.0 - Legacy License & Specialized Tools ⏸️

**Goal**: Support tools requiring traditional license files/servers (lower priority)

**New Applications**:
1. **lens-stata** (BYOL - license file)
   - Upload license file at launch
   - Environments: `econometrics`, `panel-data`, `survey-analysis`

2. **lens-spss** (BYOL - license file/server)
   - Configure license file or server
   - Environments: `social-science`, `survey-research`

3. **Specialized tools** (Various licensing)
   - PyMOL, Ansys, COMSOL, etc.
   - Custom configuration per tool

**Infrastructure** (only if needed):
- License file upload/storage
- License server configuration
- License validation helpers

**Why Later**: Requires more complex license management. Many vendors migrating to cloud authentication anyway.

---

## Tool Priority Matrix

### Tier 1: Open Source GUI Tools (v0.14)
**Criteria**: Open source + high academic demand

1. **QGIS** - Most popular open GIS tool
2. **ParaView** - Standard for large dataset visualization
3. **ImageJ/Fiji** - Standard for biological image analysis
4. **Octave** - MATLAB alternative for teaching
5. **OpenRefine** - Data cleaning is universal need

**Why These First**: Open source means no licensing complexity. Just install and go!

### Tier 2: Cloud-Authenticated Commercial (v0.15) 🎯 FOCUS
**Criteria**: High demand + modern cloud authentication

1. **MATLAB** - Most requested (MathWorks account login)
2. **Mathematica** - Popular in physics/math (Wolfram ID)
3. **ArcGIS** - Geography/GIS (ArcGIS Online credentials)
4. **Geneious** - Bioinformatics (subscription-based)

**Why These Second**: Simple to support - just install software + DCV. User logs in with credentials. No license configuration needed!

### Tier 3: Legacy License Tools (v0.16+) ⏸️ DEFERRED
**Criteria**: Older licensing models (license files/servers)

1. **Stata** - Economics (license files)
2. **SPSS** - Social sciences (license files/servers)
3. **Specialized tools** - Ansys, COMSOL, PyMOL, etc.

**Why Later**: Require license configuration system. Less urgent as many vendors are migrating to cloud authentication.

---

## Design Decisions

### Individual App vs Catalog

**Get Individual Apps** (lens-TOOL):
- MATLAB
- ArcGIS
- QGIS
- ParaView
- ImageJ/Fiji
- Stata
- Mathematica

**Use Catalog** (lens-tool launch TOOL):
- Octave
- OpenRefine
- SPSS
- PyMOL
- Geneious
- Ansys
- COMSOL
- All others

**Rationale**: Top 7 likely cover 80% of use cases

### AMI Management Strategy

**For Open Source Tools**:
- Build and maintain our own AMIs
- Update quarterly
- Store AMI IDs in tool catalog

**For Commercial Tools**:
- Query AWS Marketplace dynamically
- Let users choose from available versions
- Cache marketplace query results

**For BYOL Tools**:
- Provide base AMI with software pre-installed
- User configures license at launch
- Or point to community AMIs

---

## User Experience Examples

### Example 1: MATLAB (Cloud Authentication - Simplest!)
```bash
$ lens-matlab launch

✓ Launching MATLAB R2024b
✓ Instance i-xxxxx starting (c6i.xlarge, $0.17/hr)
✓ MATLAB desktop ready!

🌐 Browser: https://54.123.45.67:8443
   Opening MATLAB desktop...

---
MATLAB window opens via DCV:

MATLAB Login Screen:
Email: user@university.edu
Password: ********

✓ Signed in
✓ License validated automatically via MathWorks cloud
MATLAB R2024b ready!
```

**Note**: User authenticates directly in MATLAB - no pre-configuration needed!

### Example 2: QGIS (Open Source)
```bash
$ lens-qgis launch

Environment:
1. Basic GIS - QGIS with essential plugins
2. Advanced GIS - QGIS + GRASS + SAGA + PostGIS
3. Remote Sensing - QGIS + Orfeo Toolbox + SNAP

Select: 2

✓ Launching QGIS Advanced GIS environment
✓ Instance ready with QGIS 3.34, GRASS 8.3, SAGA 9.3
✓ Desktop ready at https://54.123.45.67:8443
```

### Example 3: ArcGIS (Cloud Authentication)
```bash
$ lens-arcgis launch

✓ Launching ArcGIS Pro 3.2
✓ Instance i-xxxxx starting (c6i.xlarge, $0.17/hr)
✓ ArcGIS desktop ready!

🌐 Browser: https://54.123.45.67:8443
   Opening ArcGIS desktop...

---
ArcGIS Pro window opens via DCV:

ArcGIS Pro Sign In:
ArcGIS Online username: researcher@university.edu
Password: ********

✓ Signed in
✓ License validated via ArcGIS Online
ArcGIS Pro ready!
```

**Note**: Most users authenticate via ArcGIS Online. Legacy license server support available in v0.16+ if needed.

### Example 4: Tool Discovery
```bash
$ lens tools search gis

Found 3 GIS tools:

1. QGIS 3.34 (lens-qgis)
   Open source desktop GIS
   Cost: Instance only (~$0.17/hr)

2. ArcGIS Pro 3.2 (lens-arcgis)
   Professional GIS software (cloud authentication)
   Cost: Instance (~$0.17/hr) + ArcGIS Online subscription

3. GRASS GIS 8.3 (lens-tool launch grass-gis)
   Open source GIS and remote sensing
   Cost: Instance only (~$0.08/hr)

For more info: lens tools info <tool-name>
```

---

## Technical Considerations

### AMI Catalog Format

**Location**: `~/.lens/tool-catalog.yaml` (with online updates)

**Structure**:
```yaml
version: "1.0"
last_updated: "2025-10-25"

tools:
  matlab:
    display_name: "MATLAB"
    vendor: "MathWorks"
    categories: ["mathematics", "engineering", "data-science"]
    access_method: "gui"
    gpu_support: true

    deployment_options:
      - id: "cloud-auth"
        name: "Cloud Authentication (Recommended)"
        license: "user-login"
        ami_id: "ami-xxxxx"  # MATLAB with MathWorks login support
        cost_markup: 0.0
        authentication: "User logs in with MathWorks/institutional credentials"

      - id: "marketplace"
        name: "AWS Marketplace"
        license: "included"
        ami_search_pattern: "MATLAB-R{version}-*"
        cost_markup: 0.50  # per hour

      - id: "legacy-license"
        name: "Legacy License File (v0.16+)"
        license: "user-provided"
        ami_id: "ami-xxxxx"
        cost_markup: 0.0
        license_types: ["file", "server"]
        note: "Most users should use cloud-auth instead"

    recommended_instances:
      light: ["t3.large", "t3.xlarge"]
      standard: ["c6i.xlarge", "c6i.2xlarge"]
      heavy: ["c6i.4xlarge", "c6i.8xlarge"]
      gpu: ["g4dn.xlarge", "g4dn.2xlarge"]

    environments:
      - name: "engineering"
        description: "For engineering and simulation work"
        toolboxes: ["Control System", "Signal Processing", "Simulink"]

      - name: "data-science"
        description: "For data analysis and machine learning"
        toolboxes: ["Statistics", "Machine Learning", "Deep Learning"]
```

### License Storage (Optional - Legacy Tools Only)

**Note**: Most modern tools (MATLAB, ArcGIS, Mathematica, Geneious) use cloud authentication and don't need license storage. This is only for legacy tools in v0.16+.

**Security**: Encrypt license files at rest

**Location**: `~/.lens/licenses/` (chmod 700)

**Format**:
```
~/.lens/licenses/
├── stata.lic (encrypted)
├── spss-server.conf (encrypted)
└── ansys.lic (encrypted)
```

**Config Reference** (for legacy tools):
```yaml
# in ~/.lens/config.yaml
licenses:
  stata:
    type: "file"
    path: "~/.lens/licenses/stata.lic"
  spss:
    type: "server"
    server: "license.university.edu:27000"
```

---

## Success Metrics

### v0.14 (Open Source GUI Tools)
- ✅ 5+ open-source GUI tools supported
- ✅ Users can launch and use QGIS without AWS knowledge
- ✅ ImageJ workflows work end-to-end
- ✅ ParaView can visualize large datasets (10GB+)
- ✅ Octave provides free MATLAB alternative

### v0.15 (Cloud-Authenticated Commercial Tools)
- ✅ MATLAB launches successfully with MathWorks login
- ✅ Users authenticate directly in application (no pre-config)
- ✅ ArcGIS Pro working with ArcGIS Online credentials
- ✅ Mathematica and Geneious cloud auth working
- ✅ Documentation clearly explains "just log in" workflow

### v0.16 (Legacy License Tools)
- ✅ Stata license file upload working
- ✅ SPSS license server configuration working
- ✅ License validation helpers available
- ✅ Clear documentation for IT admins

### v1.0 (Production)
- ✅ 20+ research tools supported
- ✅ Tool catalog searchable and comprehensive
- ✅ Community can contribute new tools
- ✅ Used across 10+ research domains

---

## Open Questions

1. **AMI Updates**: How do we keep AMIs current with software updates?
   - Auto-check for newer marketplace AMIs?
   - Rebuild our AMIs quarterly?
   - Let users specify versions?

2. **Cloud Authentication Flow**: How do we optimize the login experience?
   - Pre-populate credential helpers?
   - Document institutional login flows (SAML/OAuth)?
   - Provide troubleshooting for common auth issues?

3. **Cost Transparency**: How do we show costs clearly?
   - Separate software vs infrastructure costs?
   - Show monthly projections?
   - Compare marketplace vs cloud-auth BYOL?

4. **GPU Instances**: When do we recommend GPU?
   - For ParaView: optional but better
   - For deep learning: required
   - For general compute: usually not worth it

5. **Community AMIs**: How do we handle user-contributed AMIs?
   - Verification process?
   - Security scanning?
   - Support policy?

---

## Next Steps

1. **Immediate** (v0.10.0):
   - Complete NICE DCV desktop implementation
   - Test with manual QGIS installation
   - Design AMI catalog format

2. **Short-term** (v0.14.0):
   - Build AMIs for top 5 open-source tools
   - Implement tool catalog system
   - Add `lens-qgis`, `lens-paraview`, `lens-imagej` commands

3. **Medium-term** (v0.15.0) 🎯:
   - Build MATLAB AMI with cloud auth support
   - Add credential reminder system
   - Add `lens-matlab`, `lens-arcgis`, `lens-mathematica` commands
   - Document institutional login flows

4. **Long-term** (v0.16.0+):
   - Legacy license file/server support
   - Community AMI contributions
   - Tool ecosystem platform
