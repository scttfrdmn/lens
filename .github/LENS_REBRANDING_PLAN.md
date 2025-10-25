# Lens Rebranding Implementation Plan

**Issue**: #8
**Status**: In Progress
**Priority**: HIGH

## Confirmed Decisions

1. ✅ **Binary Names**: `lens-jupyter`, `lens-rstudio`, `lens-vscode`
2. ✅ **Repository Name**: `lens` (from `lens`)
3. ✅ **Config Directory**: `~/.lens/` (unified structure)
4. ✅ **Go Module Path**: `github.com/scttfrdmn/lens`
5. ✅ **Backward Compatibility**: Soft migration (auto-detect and migrate)

## Implementation Order

### Phase 1: Go Module Path Update
**Why First**: All imports depend on this; must be updated before anything else works.

1. Update `go.work` - workspace module path
2. Update `go.mod` files (root + 3 apps)
3. Update all import statements throughout codebase
4. Run `go mod tidy` to verify

### Phase 2: Binary Names and Directory Structure
**Why Second**: With imports fixed, we can rename executables and their directories.

1. Rename cmd directories:
   - `apps/jupyter/cmd/lens-jupyter/` → `apps/jupyter/cmd/lens-jupyter/`
   - `apps/rstudio/cmd/lens-rstudio/` → `apps/rstudio/cmd/lens-rstudio/`
   - `apps/vscode/cmd/lens-vscode/` → `apps/vscode/cmd/lens-vscode/`

2. Update Makefile build targets
3. Update all app-specific Makefiles
4. Test builds

### Phase 3: Configuration Paths
**Why Third**: With binaries building, update where they store config.

1. Update `pkg/config/state.go` - config directory paths
2. Update `pkg/config/keys.go` - key storage paths
3. Add migration logic to auto-migrate `~/.lens-jupyter` → `~/.lens/`
4. Update tests

### Phase 4: Code References
**Why Fourth**: Update all string literals and messages.

1. Update CLI help text and descriptions
2. Update error messages ("Lens" → "Lens")
3. Update log messages
4. Update comments and docstrings

### Phase 5: Documentation
**Why Fifth**: With code complete, update all documentation.

1. Update README.md (root)
2. Update app-specific READMEs
3. Update all docs/ markdown files
4. Update USER_SCENARIOS (all 5 personas)
5. Update ROADMAP and CHANGELOG

### Phase 6: CI/CD and Build Scripts
**Why Sixth**: Ensure automated builds work with new naming.

1. Update GitHub Actions workflows
2. Update build scripts
3. Update release process

### Phase 7: Repository Rename
**Why Last**: Once everything is merged and working, rename the repo.

1. Rename repository on GitHub
2. Update local remotes
3. Verify GitHub auto-redirects work

## Detailed Task List

### Phase 1: Go Module Path (30 min)

- [ ] Update `go.work`
- [ ] Update `go.mod` (root)
- [ ] Update `apps/jupyter/go.mod`
- [ ] Update `apps/rstudio/go.mod`
- [ ] Update `apps/vscode/go.mod`
- [ ] Find and replace all imports: `github.com/scttfrdmn/lens` → `github.com/scttfrdmn/lens`
- [ ] Run `go mod tidy` in all modules
- [ ] Verify no broken imports

### Phase 2: Binary Names (20 min)

- [ ] Rename `apps/jupyter/cmd/lens-jupyter/` → `apps/jupyter/cmd/lens-jupyter/`
- [ ] Rename `apps/rstudio/cmd/lens-rstudio/` → `apps/rstudio/cmd/lens-rstudio/`
- [ ] Rename `apps/vscode/cmd/lens-vscode/` → `apps/vscode/cmd/lens-vscode/`
- [ ] Update root `Makefile` - build targets and binary names
- [ ] Update `apps/jupyter/Makefile`
- [ ] Update `apps/rstudio/Makefile`
- [ ] Update `apps/vscode/Makefile`
- [ ] Test build all binaries: `make build`

### Phase 3: Configuration Paths (45 min)

- [ ] Update `pkg/config/state.go`:
  - Change config dir: `~/.lens-jupyter` → `~/.lens/`
  - Add unified structure: `~/.lens/state.json`, `~/.lens/keys/`, etc.
- [ ] Add migration function in `pkg/config/migrate.go`:
  ```go
  func MigrateFromLegacy() error {
    // Detect ~/.lens-jupyter, ~/.lens-rstudio, ~/.lens-vscode
    // Copy to ~/.lens/
    // Log migration
  }
  ```
- [ ] Call migration at app startup (in each main.go)
- [ ] Update `pkg/config/keys.go` - key storage paths
- [ ] Update all tests referencing config paths
- [ ] Test migration logic manually

### Phase 4: Code References (30 min)

- [ ] Search and replace "Lens" → "Lens" in:
  - Help text
  - Error messages
  - Log statements
  - Comments
- [ ] Update CLI descriptions in cobra commands
- [ ] Update version output strings
- [ ] Review and update any remaining references

### Phase 5: Documentation (60 min)

- [ ] Update `README.md`:
  - Title: "Lens" → "Lens"
  - Installation: `lens-jupyter` → `lens-jupyter`
  - All command examples
  - Project description
- [ ] Update `apps/jupyter/README.md`
- [ ] Update `apps/rstudio/README.md`
- [ ] Update `apps/vscode/README.md`
- [ ] Update `docs/USER_SCENARIOS/`:
  - All 5 persona walkthroughs
  - All command examples
- [ ] Update `USER_REQUIREMENTS.md`
- [ ] Update `DESIGN_PRINCIPLES.md`
- [ ] Update `ROADMAP.md`:
  - All command references
  - Tool names
- [ ] Update `CHANGELOG.md`:
  - Add v0.8.0 entry for rebranding
- [ ] Update `CONTRIBUTING.md`
- [ ] Update `.github/ISSUE_TEMPLATE/` (if any aws-* references)

### Phase 6: CI/CD (20 min)

- [ ] Update `.github/workflows/*.yml`:
  - Build job binary names
  - Artifact names
  - Any hardcoded paths
- [ ] Update `.github/labels.yml` (if needed)
- [ ] Test workflows (or verify on next push)

### Phase 7: Repository Rename (5 min)

- [ ] Rename repository: Settings → Repository name → `lens`
- [ ] Update local git remote:
  ```bash
  git remote set-url origin git@github.com:scttfrdmn/lens.git
  ```
- [ ] Verify old URL redirects work
- [ ] Update any external references (if any)

## Testing Checklist

After each phase, verify:

### Phase 1 Tests:
```bash
go build ./apps/jupyter/cmd/lens-jupyter  # Should fail (path doesn't exist yet)
go build ./...                           # Should succeed with no errors
```

### Phase 2 Tests:
```bash
make build                               # Should build lens-jupyter, lens-rstudio, lens-vscode
./bin/lens-jupyter version               # Should output version
./bin/lens-rstudio version
./bin/lens-vscode version
```

### Phase 3 Tests:
```bash
# Create fake old config
mkdir -p ~/.lens-jupyter
echo "test" > ~/.lens-jupyter/test.json

# Run migration
./bin/lens-jupyter version  # Should auto-migrate

# Verify migration
ls ~/.lens/                 # Should contain migrated files
```

### Phase 4 Tests:
```bash
./bin/lens-jupyter --help   # Should say "Lens" not "Lens"
./bin/lens-jupyter launch   # Error messages should say "Lens"
```

### Phase 5 Tests:
- Manually review all README files
- Check all command examples are updated
- Verify no aws-* references remain (except in CHANGELOG history)

### Phase 6 Tests:
- Trigger GitHub Actions workflow
- Verify builds succeed with new binary names

## Rollback Plan

If issues arise:

1. **During Development**: Use git to revert commits
2. **After Release**:
   - Keep backward compatibility in migration code
   - Document manual rollback steps
   - Provide `lens-jupyter` symlinks if needed

## Timeline

**Estimated Total Time**: 3-4 hours

- Phase 1: 30 min
- Phase 2: 20 min
- Phase 3: 45 min
- Phase 4: 30 min
- Phase 5: 60 min
- Phase 6: 20 min
- Phase 7: 5 min
- Testing: 30 min

**Suggested Schedule**: Complete in single session to avoid broken intermediate states.

## Notes

- All changes will be in a single commit (or small series of commits)
- Will create v0.8.0 release after completion
- GitHub provides automatic redirects for renamed repositories
- Migration code will be gentle (non-destructive, logs what it does)

## Related Issues

- Unblocks: All future development using new naming
- Updates: Issues #1-#6 (v0.7.0) will use lens-* in implementation
- Updates: Issue #8 documentation already uses lens-* naming

---

**Status**: Ready to implement
**Next Step**: Begin Phase 1 (Go Module Path Update)
