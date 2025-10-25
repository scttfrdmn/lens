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

#### 2. **AWS Marketplace (Pay-as-you-go)**
- Software cost included in EC2 hourly rate
- Automatic billing through AWS
- No license management needed
- Examples: MATLAB (marketplace), ArcGIS Enterprise, SAS, Mathematica

**Implementation**: AMI marketplace integration required

#### 3. **BYOL (Bring Your Own License)**
- User provides license key/file
- May need license server configuration
- Lower instance costs (no software markup)
- Examples: MATLAB (BYOL), ArcGIS Desktop, Stata, SPSS

**Implementation**: License configuration system required

#### 4. **Subscription/Cloud Native**
- Vendor-hosted licensing
- May require API keys or account setup
- Examples: Some cloud-based research platforms

**Implementation**: Case-by-case integration

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

### 2. License Configuration Manager

**Purpose**: Handle different licensing scenarios

**Capabilities**:
- Prompt for license keys/files
- Upload license files to instance
- Configure license servers
- Validate licenses before launch
- Store encrypted license info in `~/.lens/licenses/`

**Example Flow**:
```bash
$ lens-matlab launch

Selecting MATLAB deployment option:
1. AWS Marketplace (includes license, $X.XX/hr)
2. Bring Your Own License (requires license, $Y.YY/hr)

Choice: 2

License Configuration:
1. License file
2. License server
3. Network license

Choice: 1

License file path: ~/matlab.lic
✓ License file validated
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

Deployment Options:
1. Campus License Server (Recommended for Academic Users)
   - Most common for universities and research institutions
   - Requires campus license server address (e.g., license.university.edu:27000)
   - Cost: Base instance only (~$0.17/hr for c6i.xlarge)
   - No software markup
   - AMI: ami-yyyyy (auto-selected)

2. AWS Marketplace
   - License included in instance cost
   - Cost: +$0.50/hour over base instance (~$0.67/hr total)
   - Good for users without campus access
   - AMI: ami-xxxxx (auto-selected)

3. License File (BYOL)
   - For personal/professional MATLAB licenses
   - Requires .lic license file
   - Cost: Base instance only
   - AMI: ami-yyyyy (auto-selected)

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

### Phase 2: v0.11.0 - Open Source GUI Tools (2-3 months)

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

**Infrastructure**:
- AMI catalog system (alpha)
- Tool discovery commands
- Shared DCV base images

### Phase 3: v0.12.0 - AWS Marketplace Tools (3-4 months)

**Goal**: Support commercial tools via AWS Marketplace

**New Applications**:
1. **lens-matlab** (marketplace + BYOL)
   - AWS Marketplace AMI integration
   - BYOL license file support
   - Environments: `engineering`, `data-science`, `control-systems`

2. **lens-mathematica** (marketplace + BYOL)
   - Wolfram Cloud integration
   - Environments: `symbolic-math`, `data-science`

**Infrastructure**:
- AMI marketplace search/selection
- Pricing calculator with software costs
- License configuration manager (basic)

### Phase 4: v0.13.0 - BYOL Commercial Tools (3-4 months)

**Goal**: Support BYOL licensing for major commercial tools

**New Applications**:
1. **lens-arcgis** (marketplace + BYOL)
   - ArcGIS Pro or ArcGIS Desktop
   - License server configuration
   - Environments: `urban-planning`, `environmental`, `remote-sensing`

2. **lens-stata** (BYOL)
   - License file or server
   - Environments: `econometrics`, `panel-data`, `survey-analysis`

3. **lens-spss** (BYOL)
   - License configuration
   - Environments: `social-science`, `survey-research`

**Infrastructure**:
- Advanced license manager (server config)
- License validation
- Multi-seat license support

### Phase 5: v0.14.0 - Specialized Domain Tools (4-5 months)

**Goal**: Support highly specialized research tools

**New Applications**:
1. **lens-geneious** (Bioinformatics)
2. **lens-pymol** (Molecular visualization)
3. **lens-ansys** (Engineering simulation)
4. **lens-comsol** (Multiphysics)

**Infrastructure**:
- Tool catalog (stable)
- Community AMI contributions
- Custom AMI builder for labs

---

## Tool Priority Matrix

### Tier 1: Must-Have (v0.10-0.11)
**Criteria**: Open source + high academic demand

1. **QGIS** - Most popular open GIS tool
2. **ParaView** - Standard for large dataset visualization
3. **ImageJ/Fiji** - Standard for biological image analysis
4. **Octave** - MATLAB alternative for teaching
5. **OpenRefine** - Data cleaning is universal need

### Tier 2: High-Value Commercial (v0.12)
**Criteria**: High demand + AWS Marketplace available

1. **MATLAB** - Most requested commercial tool
2. **Mathematica** - Popular in physics/math

### Tier 3: BYOL Commercial (v0.13)
**Criteria**: Many institutions have licenses

1. **ArcGIS** - Standard in geography departments
2. **Stata** - Standard in economics
3. **SPSS** - Standard in social sciences

### Tier 4: Specialized (v0.14+)
**Criteria**: Specific domains, smaller user bases

1. Engineering: Ansys, COMSOL
2. Bioinformatics: Geneious, PyMOL
3. Remote Sensing: ENVI, ERDAS
4. Statistics: GraphPad Prism, OriginPro

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

### Example 1: MATLAB (Campus License - Most Common)
```bash
$ lens-matlab launch

MATLAB Deployment Options:
1. Campus License Server (recommended for academic users)
2. AWS Marketplace (license included) - $0.50/hr software + instance
3. License File (BYOL)

Select: 1

Campus License Server Configuration:
Enter license server address: license.university.edu:27000

✓ Testing license server connection...
✓ License server reachable (10 licenses available)
✓ Launching MATLAB R2024b
✓ Instance i-xxxxx starting (c6i.xlarge, $0.17/hr)
✓ MATLAB desktop ready!

🌐 Browser: https://54.123.45.67:8443
   Auto-opening in 5 seconds...
```

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

### Example 3: ArcGIS (BYOL)
```bash
$ lens-arcgis launch

ArcGIS Pro requires a license. Choose licensing method:
1. License file
2. Named user (ArcGIS Online)
3. Concurrent use license server

Select: 3

License server address: license.university.edu:27000

✓ Testing license server connection...
✓ License server reachable
✓ Launching ArcGIS Pro 3.2
✓ Desktop ready at https://54.123.45.67:8443
```

### Example 4: Tool Discovery
```bash
$ lens tools search gis

Found 3 GIS tools:

1. QGIS 3.34 (lens-qgis)
   Open source desktop GIS
   Cost: Instance only (~$0.17/hr)

2. ArcGIS Pro 3.2 (lens-arcgis)
   Professional GIS software
   Cost: Instance + license (BYOL) or Marketplace

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
      - id: "marketplace"
        name: "AWS Marketplace"
        license: "included"
        ami_search_pattern: "MATLAB-R{version}-*"
        cost_markup: 0.50  # per hour

      - id: "byol"
        name: "Bring Your Own License"
        license: "user"
        ami_id: "ami-xxxxx"  # or search pattern
        cost_markup: 0.0
        license_types: ["file", "server"]

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

### License Storage

**Security**: Encrypt license files at rest

**Location**: `~/.lens/licenses/` (chmod 700)

**Format**:
```
~/.lens/licenses/
├── matlab.lic (encrypted)
├── arcgis-server.conf (encrypted)
└── stata.lic (encrypted)
```

**Config Reference**:
```yaml
# in ~/.lens/config.yaml
licenses:
  matlab:
    type: "file"
    path: "~/.lens/licenses/matlab.lic"
  arcgis:
    type: "server"
    server: "license.university.edu:27000"
```

---

## Success Metrics

### v0.11 (Open Source Tools)
- ✅ 5+ open-source GUI tools supported
- ✅ Users can launch and use QGIS without AWS knowledge
- ✅ ImageJ workflows work end-to-end
- ✅ ParaView can visualize large datasets (10GB+)

### v0.12 (Marketplace Tools)
- ✅ MATLAB marketplace integration working
- ✅ Users can compare marketplace vs BYOL costs
- ✅ Automatic marketplace AMI discovery
- ✅ Pricing calculator includes software costs

### v0.13 (BYOL Tools)
- ✅ ArcGIS license file/server configuration working
- ✅ Stata and SPSS license setup working
- ✅ License validation before launch
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

2. **License Validation**: How much validation do we do?
   - Validate license files before launch?
   - Or launch and let tool validate?
   - What about network license servers?

3. **Cost Transparency**: How do we show costs clearly?
   - Separate software vs infrastructure costs?
   - Show monthly projections?
   - Compare marketplace vs BYOL?

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

2. **Short-term** (v0.11.0):
   - Build AMIs for top 5 open-source tools
   - Implement tool catalog system
   - Add `lens-qgis` command

3. **Medium-term** (v0.12.0):
   - AWS Marketplace API integration
   - License configuration manager
   - Add `lens-matlab` command

4. **Long-term** (v0.13.0+):
   - BYOL license support
   - Community AMI contributions
   - Tool ecosystem platform
