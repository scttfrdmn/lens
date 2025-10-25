# Scoop Package for lens-jupyter

This directory contains the Scoop manifest for lens-jupyter.

**Note:** lens-jupyter is now part of the [Lens monorepo](https://github.com/scttfrdmn/lens). The binary and functionality remain unchanged.

## For Users

Install lens-jupyter via Scoop:

```powershell
scoop install https://raw.githubusercontent.com/scttfrdmn/lens/main/scoop/lens-jupyter.json
```

Or after it's added to the main bucket:

```powershell
scoop install lens-jupyter
```

## For Maintainers

### Testing the Manifest Locally

Before submitting to Scoop, test the manifest:

```powershell
# Install Scoop if you haven't already
# Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
# irm get.scoop.sh | iex

# Test the manifest
scoop install .\scoop\lens-jupyter.json

# Verify it works
lens-jupyter --version

# Uninstall to clean up
scoop uninstall lens-jupyter
```

### Updating the Manifest

When releasing a new version:

1. The `autoupdate` section will automatically handle new releases
2. Or manually update:
   - Change `version` field
   - Update URLs to point to new release
   - Update `hash` values from checksums.txt

### Submitting to Scoop Main Bucket

1. Fork https://github.com/ScoopInstaller/Main
2. Copy `lens-jupyter.json` to `bucket/` directory
3. Test locally: `scoop install ./bucket/lens-jupyter.json`
4. Submit PR to ScoopInstaller/Main

## Manifest Features

- **Automatic Updates**: The `autoupdate` section allows Scoop to automatically generate PRs for new releases
- **Multi-Architecture**: Supports both x86_64 and ARM64 Windows
- **Hash Verification**: Uses SHA256 checksums from release
- **Version Checking**: Automatically checks GitHub for new releases
