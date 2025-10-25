# Package Manager Migration Guide

## Repository Migration: lens-jupyter → lens

The lens-jupyter project has been transformed into the **Lens monorepo** (`scttfrdmn/lens`). This document provides guidance for updating existing package manager submissions.

### 🎯 What Changed

- **Repository URL:** `github.com/scttfrdmn/lens-jupyter` → `github.com/scttfrdmn/lens`
- **Release URLs:** All release artifacts now come from lens repository
- **Binary location:** No change - `lens-jupyter` and `lens-jupyter.exe` remain the same
- **Functionality:** 100% backward compatible - no breaking changes

### 📦 Existing Package Manager PRs

#### 1. Scoop (ScoopInstaller/Main)

**PR Link:** https://github.com/ScoopInstaller/Main/pull/7267

**Status:** Needs update for monorepo

**Required Changes:**
```json
{
    "version": "0.5.0",
    "homepage": "https://github.com/scttfrdmn/lens",
    "architecture": {
        "64bit": {
            "url": "https://github.com/scttfrdmn/lens/releases/download/v0.5.0/lens-jupyter_Windows_x86_64.zip",
            "hash": "sha256:<from checksums.txt after v0.5.0 release>"
        },
        "arm64": {
            "url": "https://github.com/scttfrdmn/lens/releases/download/v0.5.0/lens-jupyter_Windows_arm64.zip",
            "hash": "sha256:<from checksums.txt after v0.5.0 release>"
        }
    },
    "checkver": {
        "github": "https://github.com/scttfrdmn/lens"
    },
    "autoupdate": {
        "architecture": {
            "64bit": {
                "url": "https://github.com/scttfrdmn/lens/releases/download/v$version/lens-jupyter_Windows_x86_64.zip"
            },
            "arm64": {
                "url": "https://github.com/scttfrdmn/lens/releases/download/v$version/lens-jupyter_Windows_arm64.zip"
            }
        },
        "hash": {
            "url": "$baseurl/checksums.txt"
        }
    }
}
```

**Action Items:**
1. Wait for v0.5.0 release to complete and artifacts to be published
2. Get SHA256 hashes from: https://github.com/scttfrdmn/lens/releases/download/v0.5.0/checksums.txt
3. Update PR #7267 with new manifest or
4. Close old PR and submit fresh one with updated manifest

**Comment for PR Update:**
```markdown
## Update: Repository Migration

The lens-jupyter project has been migrated to a monorepo structure at
https://github.com/scttfrdmn/lens.

**Changes in this update:**
- Updated repository URLs to point to lens
- Updated to v0.5.0 (monorepo release)
- No functional changes to the binary
- Autoupdate will continue to work with new repository

The lens-jupyter binary remains functionally identical and 100% backward compatible.
```

---

#### 2. Conda-forge (conda-forge/staged-recipes)

**PR Link:** https://github.com/conda-forge/staged-recipes/pull/31241

**Status:** Needs update for monorepo

**Required Changes:**

The conda-forge recipe needs to be updated to point to the new repository. Here's what needs to change in `recipes/lens-jupyter/meta.yaml`:

```yaml
{% set name = "lens-jupyter" %}
{% set version = "0.5.0" %}

package:
  name: {{ name|lower }}
  version: {{ version }}

source:
  url: https://github.com/scttfrdmn/lens/archive/refs/tags/v{{ version }}.tar.gz
  sha256: <hash of source tarball>

build:
  number: 0
  script:
    - cd apps/jupyter  # NEW: Navigate to jupyter app directory
    - go build -v -o $PREFIX/bin/{{ name }} ./cmd/lens-jupyter  # Updated path

requirements:
  build:
    - {{ compiler('go-cgo') }}
    - go >=1.22
  run:
    - aws-cli

test:
  commands:
    - lens-jupyter --version

about:
  home: https://github.com/scttfrdmn/lens
  license: Apache-2.0
  license_family: Apache
  license_file: LICENSE
  summary: CLI tool for launching Jupyter Lab instances on AWS EC2
  description: |
    lens-jupyter is part of the Lens toolkit, providing a simple CLI
    for launching and managing Jupyter Lab instances on AWS EC2 Graviton
    processors with automatic idle detection and cost optimization.
  doc_url: https://github.com/scttfrdmn/lens/blob/main/apps/jupyter/README.md
  dev_url: https://github.com/scttfrdmn/lens

extra:
  recipe-maintainers:
    - scttfrdmn
```

**Key Changes:**
1. **Source URL:** Points to lens repository
2. **Build script:** Adds `cd apps/jupyter` to navigate to app directory
3. **Build command:** Updated path `./cmd/lens-jupyter`
4. **URLs:** All links updated to lens repository
5. **Description:** Notes it's part of Lens toolkit

**Action Items:**
1. Wait for v0.5.0 release to complete
2. Get SHA256 hash of source tarball:
   ```bash
   curl -sL https://github.com/scttfrdmn/lens/archive/refs/tags/v0.5.0.tar.gz | sha256sum
   ```
3. Update PR #31241 with new recipe or
4. Close old PR and submit fresh one with updated recipe

**Comment for PR Update:**
```markdown
## Update: Repository Migration

The lens-jupyter project has been migrated to a monorepo structure at
https://github.com/scttfrdmn/lens.

**Changes in this update:**
- Updated repository URLs to point to lens
- Updated build script to navigate to `apps/jupyter/` directory
- Updated to v0.5.0 (monorepo release)
- Updated all documentation links
- No functional changes to the binary

The lens-jupyter binary remains functionally identical and 100% backward compatible.
The monorepo structure allows sharing infrastructure with other Lens tools like lens-rstudio.
```

---

### 🔄 Automated Updates

Both package managers have autoupdate mechanisms:

**Scoop:** The `autoupdate` section will automatically generate PRs for future releases using the lens repository URLs.

**Conda-forge:** Once the feedstock is created, the regro-cf-autotick-bot will automatically create PRs for new releases. The feedstock's `recipe/meta.yaml` will need the updated URLs.

---

### ⚡ Quick Checklist

- [ ] Wait for v0.5.0 GitHub release to complete
- [ ] Download checksums.txt from v0.5.0 release
- [ ] Update Scoop PR #7267 with new manifest
- [ ] Calculate source tarball SHA256 for conda
- [ ] Update Conda PR #31241 with new recipe
- [ ] Verify both PRs reference lens repository
- [ ] Test installations after PRs are merged

---

### 📞 Questions?

If package maintainers have questions about the migration:
- **GitHub Issues:** https://github.com/scttfrdmn/lens/issues
- **Discussions:** https://github.com/scttfrdmn/lens/discussions
- **Migration docs:** This file in the repository

---

### 🎉 Benefits of Migration

The monorepo structure provides:
- Shared infrastructure for multiple IDE tools (Jupyter, RStudio, etc.)
- Consistent release management across all tools
- Unified documentation and contribution process
- Better code reuse and maintenance

Users of lens-jupyter will see no changes - the tool works exactly as before!
