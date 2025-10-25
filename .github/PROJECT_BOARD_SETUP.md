# GitHub Project Board Setup

## Overview
This document describes the setup and configuration of the GitHub Project board for Lens (Lens) development.

**Project**: Lens Development (Project #6)
**URL**: https://github.com/users/scttfrdmn/projects/6

## Custom Fields

The following custom fields have been configured programmatically:

### 1. Persona (Single Select)
Identifies the primary user persona for each issue.

**Options:**
- **Solo Researcher** (Blue) - Individual researcher
- **Graduate Student** (Green) - PhD/Masters student
- **Lab PI** (Yellow) - Principal Investigator
- **Course Instructor** (Orange) - Teaching faculty
- **Research Computing Manager** (Red) - IT/Admin

### 2. Phase (Single Select)
Maps issues to roadmap phases.

**Options:**
- **v0.7.0 - User Experience** (Blue) - Quickstart & ease-of-use
- **v0.8.0 - Research Tools** (Green) - Additional research applications
- **v0.9.0 - Reproducibility** (Yellow) - Package managers & environments
- **v0.10.0 - Collaboration** (Orange) - Team features & sharing
- **v0.11.0 - Cost Management** (Red) - Budget tracking & optimization
- **v1.0.0 - Production** (Purple) - Beta testing & documentation
- **Backlog** (Gray) - Future work

### 3. Estimate (Number)
Story points or hour estimate for implementation.

### 4. ROI (Single Select)
Return on investment assessment.

**Options:**
- **High** (Green) - High return on investment
- **Medium** (Yellow) - Medium return on investment
- **Low** (Red) - Low return on investment

## Project Views

The following views should be created manually through the GitHub UI:

### View 1: Kanban (Default)
**Layout:** Board
**Group by:** Status
**Sort by:** Priority (High → Low)

**Columns:**
- Todo
- In Progress
- Done

**Purpose:** Default workflow view for tracking issue status.

### View 2: By Phase
**Layout:** Board
**Group by:** Phase
**Sort by:** Priority (High → Low)

**Purpose:** Roadmap view showing issues organized by development phase.

### View 3: By Persona
**Layout:** Board
**Group by:** Persona
**Sort by:** Phase, then Priority

**Purpose:** User-centric view showing which personas each issue addresses.

### View 4: Current Sprint (v0.7.0)
**Layout:** Table
**Filter:**
- Phase = "v0.7.0 - User Experience"
- Status ≠ "Done"

**Sort by:** Priority (High → Low), then Estimate

**Columns to show:**
- Title
- Status
- Priority (from labels)
- Persona
- Estimate
- Assignees

**Purpose:** Focus view for active sprint work.

### View 5: Backlog
**Layout:** Table
**Filter:**
- Status = "Todo"
- Phase ≠ "v0.7.0 - User Experience" OR Phase = "Backlog"

**Sort by:** Phase, then ROI (High → Low), then Priority

**Columns to show:**
- Title
- Phase
- Priority (from labels)
- Persona
- ROI
- Estimate

**Purpose:** Future work planning and prioritization.

## How to Create Views Manually

### Step-by-Step Instructions:

1. **Navigate to the Project Board:**
   https://github.com/users/scttfrdmn/projects/6

2. **Create a New View:**
   - Click the "+" icon next to existing views
   - Choose a layout (Board or Table)
   - Name the view

3. **Configure Grouping (for Board views):**
   - Click "Group by" dropdown
   - Select the desired field (Status, Phase, or Persona)

4. **Configure Filters:**
   - Click the filter icon
   - Add filter conditions (e.g., Phase = "v0.7.0")
   - Multiple filters can be combined

5. **Configure Sorting:**
   - Click "Sort" dropdown
   - Choose primary and secondary sort fields
   - Toggle ascending/descending order

6. **Configure Columns (for Table views):**
   - Click "Fields" button
   - Check/uncheck fields to show/hide
   - Drag to reorder columns

## Using the Project Board

### Setting Field Values

Once views are created, you'll need to populate custom field values for each issue:

#### Automated from Labels:
- **Priority**: Extracted from `priority: high/medium/low` labels
- **Phase**: Extracted from `phase: X.X.X` labels
- **Persona**: Extracted from `persona: xxx` labels

#### Manual Entry Needed:
- **Estimate**: Add story points or hour estimates
- **ROI**: Assess return on investment (High/Medium/Low)

### Workflow

1. **Planning:**
   - Use "Backlog" view to review upcoming work
   - Set Estimate and ROI for prioritization
   - Move high-priority items to current sprint

2. **Sprint Execution:**
   - Use "Current Sprint" view for daily standup
   - Update Status as work progresses (Todo → In Progress → Done)
   - Assign team members

3. **Tracking:**
   - Use "By Phase" view to monitor roadmap progress
   - Use "By Persona" view to ensure balanced user coverage
   - Use "Kanban" view for general workflow

### Traceability

Each issue is linked back to:
- **User Personas**: Documented in `.github/personas/`
- **Requirements**: Documented in `USER_REQUIREMENTS.md`
- **Roadmap**: Documented in `ROADMAP.md`
- **Issue Summary**: Documented in `.github/GITHUB_ISSUES_SUMMARY.md`

This creates a 5-layer traceability system:
```
Personas → Scenarios → Requirements → Issues → Pull Requests
```

## Maintenance

### Adding New Issues

When creating new issues:

1. Use appropriate issue template (`.github/ISSUE_TEMPLATE/`)
2. Add all relevant labels (priority, phase, persona, use-case, area)
3. Add to project board automatically via GitHub Actions
4. Set custom fields (Persona, Phase, Estimate, ROI)

### Updating the Board

- **Weekly**: Review "Current Sprint" view and update status
- **Monthly**: Review "By Phase" view and adjust roadmap
- **Quarterly**: Review "Backlog" view and reprioritize based on ROI

## Related Documentation

- [Project Alignment Strategy](../PROJECT_ALIGNMENT_STRATEGY.md)
- [User Requirements](../USER_REQUIREMENTS.md)
- [Roadmap](../ROADMAP.md)
- [GitHub Issues Summary](.github/GITHUB_ISSUES_SUMMARY.md)
- [Design Principles](../DESIGN_PRINCIPLES.md)

## Automation

The following GitHub Actions workflows assist with project management:

- **`.github/workflows/labels.yml`**: Syncs labels from `.github/labels.yml`
- **`.github/workflows/project-automation.yml`**: Auto-adds issues to project (if configured)

## Notes

- Custom fields were created programmatically via GitHub GraphQL API
- Views must be created manually through the UI (API limitations)
- The project board integrates with all issue templates and labels
- This setup follows the organizational pattern from the lfr-tools project
