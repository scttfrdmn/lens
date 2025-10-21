# Research Computing Manager Walkthrough: Jennifer Martinez

> **Persona**: Director of Research Computing supporting 250+ faculty, 800+ graduate students
> **Technical Level**: 5/5 (Expert - 15 years systems administration, cloud architecture)
> **Budget**: $500,000/year institutional AWS budget across 50+ research groups
> **Primary Pain Point**: Lack of visibility, cost control, security compliance, and support burden across decentralized cloud usage

---

## Profile

**Name**: Jennifer Martinez
**Position**: Director, Office of Research Computing
**Institution**: Major R1 public research university (30,000 students, $400M research expenditures)
**Department**: Central IT, reports to CIO
**Age**: 45
**Location**: Austin, Texas (primary region: us-east-1)

### Organizational Context

**Office of Research Computing (ORC)**:
- **Mission**: Enable computational research across all disciplines
- **Staff**: 8 people
  - Director (Jennifer)
  - 2 HPC sys admins
  - 2 cloud architects
  - 2 user support specialists
  - 1 security/compliance officer
- **Budget**: $3M/year total ($500K cloud, $1.5M HPC cluster, $1M personnel)

**User base**:
- 250 faculty PIs (across 12 colleges)
- 800 graduate students
- 150 postdocs
- 50 undergraduate research assistants
- **Total**: ~1,250 researchers

**Disciplines served**:
- 30% Life Sciences (genomics, bioinformatics, medical imaging)
- 25% Physical Sciences (physics, chemistry, materials science)
- 20% Engineering (CFD, FEA, ML/AI)
- 15% Social Sciences (text analytics, survey analysis, econometrics)
- 10% Other (digital humanities, geosciences, climate)

### Institutional Priorities & Constraints

**Strategic goals** (from Provost):
1. **Enable cutting-edge research** (compete with top-10 universities)
2. **Cost-effective operations** (state funding declining, must stretch budget)
3. **Security & compliance** (HIPAA, FERPA, IRB requirements)
4. **Faculty satisfaction** (faculty surveys, retention)

**IT governance**:
- Security office mandates MFA, audit logging, vulnerability scanning
- Legal requires data sovereignty controls (HIPAA, export control)
- Procurement requires competitive pricing, vendor diversity
- Budget office demands quarterly financial reporting

**Political reality**:
- Faculty want autonomy ("Don't tell me how to do research!")
- IT wants control ("We need security and cost oversight!")
- Jennifer's role: Balance both (enable research + maintain guardrails)

---

## Current Situation (Before AWS IDE)

### Decentralized Cloud Chaos (2023-2024)

**How researchers currently use AWS**:

1. **40% use personal AWS accounts** (faculty credit card):
   - Pro: Complete autonomy, no bureaucracy
   - Con: No institutional visibility, no cost control, security unknown, support burden

2. **30% use departmental AWS accounts** (50+ separate accounts):
   - Each department has own AWS account
   - Pro: Department-level cost tracking
   - Con: 50 different configurations, inconsistent security, duplicated effort

3. **20% use central institutional AWS account** (managed by ORC):
   - Pro: Centralized billing, security controls
   - Con: Bureaucratic (4-week turnaround for new projects), not user-friendly

4. **10% avoid cloud entirely** (use ancient on-prem HPC cluster):
   - Reason: Cloud too complex, too expensive, don't know where to start

### Pain Points: The Big Three

#### Pain #1: Cost Visibility & Control

**Problem**: Jennifer has no idea what university is spending on AWS

**Known costs**:
- Central institutional AWS account: $180,000/year (visible)
- Departmental accounts (estimated): $250,000/year (partially visible via finance reports)
- Personal faculty accounts: **Unknown** (could be $50K, could be $500K)

**Cost incidents** (2024):

**January**: Biology department AWS bill = $48,000
- Normal: $15,000/month
- Spike: Postdoc launched 20× r5.24xlarge instances (768 cores total)
- Forgot to terminate for 3 weeks
- Cost: $33,000 overrun
- Impact: Department froze all AWS usage for 2 months (blocked 30 researchers)

**March**: Chemistry professor's personal account = $12,000 bill
- Professor shocked: "I thought I stopped everything!"
- Investigation: Left 5 instances running since December
- Professor paid out of pocket (embarrassed, never used cloud again)
- Chilling effect: 10 other faculty heard story, decided cloud "too risky"

**August**: Engineering lab = $25,000 overrun
- GPU instances left running over summer break
- Students graduated, no one terminated resources
- PI discovered in September (3 months later)
- Grant budget exhausted, research halted

**Annual waste estimate**: $150,000/year in forgotten/idle resources (30% of cloud spending)

#### Pain #2: Security & Compliance

**Problem**: Jennifer's security officer can't assess cloud security posture

**Security incidents** (2024):

**February**: Medical School HIPAA violation
- Researcher analyzing patient MRI scans on personal EC2 instance
- Instance in public subnet, SSH open to 0.0.0.0/0
- Security scan detected exposure
- Reported to HHS Office for Civil Rights
- Investigation: 6 months, $50,000 legal fees
- Penalty: $15,000 fine
- **Outcome**: Medical school banned all personal AWS accounts

**April**: Data breach (export-controlled research)
- Engineering lab working on DoD-funded project (export controlled)
- Stored data in S3 bucket without encryption, publicly readable
- Discovered by external security researcher (reported responsibly)
- DoD investigation, lab shut down for 3 months
- **Outcome**: University almost lost DoD funding eligibility

**June**: Compliance audit failure (NSF)
- NSF requires data management plans for all grants
- Auditor: "Where are your backups? What's your retention policy?"
- Researchers: "Uh... I think it's on S3? Maybe?"
- 50 NSF-funded projects couldn't demonstrate compliance
- **Outcome**: NSF required remediation plan, threatened funding holds

**Security officer's assessment**: "We have no idea what's running in the cloud. It's a compliance nightmare."

#### Pain #3: Support Burden & Expertise Gap

**Problem**: ORC support team overwhelmed by cloud questions

**Support ticket volume** (2024):
- HPC cluster: 800 tickets/year (manageable with 2 sys admins)
- AWS: 2,400 tickets/year (3x more, with 2 cloud architects)
- **Backlog**: 300 open tickets, 6-week average response time

**Common tickets**:
1. "How do I launch an EC2 instance?" (beginner) - 40% of tickets
2. "My instance won't start" (troubleshooting) - 25%
3. "How do I set up VPC/security groups?" (networking) - 20%
4. "My bill was $5,000, help!" (cost overrun) - 10%
5. "How do I make this HIPAA compliant?" (security) - 5%

**Support team burnout**:
- 60-hour weeks common (cannot keep up with demand)
- 1 cloud architect quit (accepted industry job for 50% raise)
- Remaining architect considering leaving
- Jennifer to CIO: "We need to hire 2 more people or change our approach"

**Faculty frustration**:
- 6-week ticket turnaround unacceptable (research moves fast)
- Complex AWS setup discourages cloud adoption
- Faculty resort to personal accounts (bypassing ORC) to avoid bureaucracy

---

## Discovery & Evaluation (Fall 2024)

### Jennifer Hears About AWS IDE

**November 2024**: Jennifer attends EDUCAUSE conference

**Session**: "Simplifying Academic Cloud Computing"
- Speaker: Prof. James Chen (we met him earlier!)
- Topic: How AWS IDE reduced lab support burden by 90%
- Key points:
  - Auto-stop prevents cost overruns
  - Plain-English interface (faculty don't need AWS expertise)
  - Environment reproducibility (compliance-friendly)
  - Centralized cost tracking

**Jennifer's reaction**: "This could solve all three problems - cost, security, support"

**Post-talk conversation** (coffee break):

**Jennifer**: "This sounds too good to be true. What's the catch?"

**Prof. Chen**: "There isn't one, really. My lab went from $4,000 cost overruns to zero. My students don't need my help anymore - the wizard walks them through setup. And I can see all spending in one place."

**Jennifer**: "What about security? HIPAA compliance?"

**Prof. Chen**: "Session Manager built-in - no SSH keys. Private subnets supported. CloudTrail logs everything. My security officer was happy."

**Jennifer**: "I need to test this."

### Pilot Program Design (December 2024)

**Jennifer's plan**:
1. **Phase 1 (January-February)**: Pilot with 3 friendly research groups (10-15 researchers)
2. **Phase 2 (March-April)**: Expand to 10 groups (50 researchers)
3. **Phase 3 (May-August)**: Campus-wide rollout (all researchers)

**Pilot selection criteria**:
- Mix of disciplines (life science, engineering, social science)
- Mix of technical levels (expert to beginner)
- Mix of compliance needs (HIPAA, export control, none)
- Faculty willing to provide feedback

**Selected pilot groups**:
1. **Biology lab** (Prof. Garcia) - 6 researchers, genomics, beginner-moderate technical level
2. **Engineering lab** (Prof. Patel) - 5 researchers, ML/AI, expert technical level
3. **Public health** (Prof. Johnson) - 4 researchers, survey analysis, HIPAA compliance required

**Pilot goals**:
- ✅ Reduce support tickets by 50%
- ✅ Achieve 100% cost visibility
- ✅ Pass security audit
- ✅ Maintain/improve researcher satisfaction

### Pilot Setup (January 2025)

**Week 1**: Infrastructure preparation

Jennifer's cloud architect (Tom) configures centralized AWS account:

```bash
# Create organization units
aws organizations create-organizational-unit \
  --parent-id r-ROOT \
  --name "Research-Pilot"

# Create 3 AWS accounts (one per pilot group)
aws organizations create-account \
  --email research-pilot-bio@university.edu \
  --account-name "Biology-Garcia-Pilot"

aws organizations create-account \
  --email research-pilot-eng@university.edu \
  --account-name "Engineering-Patel-Pilot"

aws organizations create-account \
  --email research-pilot-ph@university.edu \
  --account-name "PublicHealth-Johnson-Pilot"
```

**Week 2**: Pilot kickoff meeting (all 15 researchers + PIs)

**Jennifer's presentation** (30 minutes):
> "Welcome to the AWS IDE pilot. Why are we doing this?
> 1. **Cost control**: Auto-stop prevents runaway costs
> 2. **Ease of use**: Wizard guides you through setup - no AWS expertise needed
> 3. **Security**: Built-in compliance features (Session Manager, audit logs)
> 4. **Support**: We expect 50% fewer support tickets
>
> Your role: Use AWS IDE for 2 months, give us honest feedback."

**Researcher questions**:

**Q**: "Can I still use my personal AWS account if I want?"
**A**: "Yes, but we're asking you to try AWS IDE first. We think you'll prefer it."

**Q**: "What if I need something AWS IDE doesn't support?"
**A**: "Let us know - we can file feature requests or help with workarounds."

**Q**: "How much does this cost me?"
**A**: "Zero. University pays. We want feedback more than anything."

**Setup session** (1 hour): All 15 researchers install and launch first instance

**Results**:
- 15 minutes: All 15 researchers have AWS IDE installed
- 25 minutes: 13/15 launched first instance successfully
- 2 researchers had AWS credential issues (Tom helped, fixed in 10 minutes)
- **35 minutes**: All 15 researchers have running Jupyter Lab instances

**Post-setup feedback**:
- "Way easier than I expected" (beginner researcher)
- "The wizard is nice, but can I skip it?" (expert researcher) → Tom: "Yes, use `--no-wizard` flag"
- "I like seeing the cost estimate before launching" (PI watching budget)

---

## Pilot Results (January-February 2025)

### Quantitative Metrics

#### Support Ticket Reduction

| Category | Before Pilot (monthly avg) | During Pilot (monthly avg) | Reduction |
|----------|---------------------------|---------------------------|-----------|
| "How do I launch?" | 80 tickets | 8 tickets | 90% |
| Troubleshooting | 50 tickets | 12 tickets | 76% |
| Networking/VPC | 40 tickets | 4 tickets | 90% |
| Cost overruns | 20 tickets | 0 tickets | 100% |
| Security/compliance | 10 tickets | 2 tickets | 80% |
| **TOTAL** | **200 tickets/month** | **26 tickets/month** | **87%** |

**Tom's reaction** (cloud architect):
> "We went from drowning in tickets to actually having time to work on strategic projects. The 26 tickets we do get are legitimate edge cases, not 'how do I get started' questions."

#### Cost Visibility

**Before pilot** (January 2024):
- Known costs: $180,000 (central account)
- Unknown costs: Estimated $100,000-300,000 (personal/departmental accounts)
- **Visibility**: ~40-60%

**During pilot** (January-February 2025):
- All pilot usage in central account: 100% visible
- Pilot spending: $4,200 over 2 months (15 researchers)
- Projected annual (if all 1,250 researchers): $420,000/year
- **Visibility**: 100% for pilot group

**Cost breakdown** (pilot group, 2 months):

| Research Group | Researchers | Total Cost | Cost/Researcher | Notes |
|----------------|-------------|------------|----------------|-------|
| Biology (Garcia) | 6 | $1,800 | $300 | Genomics analysis, moderate usage |
| Engineering (Patel) | 5 | $2,100 | $420 | GPU usage for ML, high usage |
| Public Health (Johnson) | 4 | $300 | $75 | Statistical analysis, low usage |
| **TOTAL** | **15** | **$4,200** | **$280 avg** | **$1,680/year projected per researcher** |

**Jennifer's analysis**:
> "If we extrapolate to all 1,250 researchers: $280/month × 12 months × 1,250 = $4.2M/year potential spending. BUT:
> - Auto-stop should reduce actual spending by 70% (based on Prof. Chen's experience)
> - Realistic projection: $4.2M × 30% utilization = $1.26M/year
> - Current known + unknown spending: ~$430,000/year
> - We're enabling 3x more researchers at 3x the cost = same cost per active user
> - But with 100% visibility and control"

#### Security & Compliance

**Pilot security audit** (February 2025):

Security officer (Carlos) audits all 15 pilot researcher accounts:

**Findings**:

✅ **Pass**: All instances using Session Manager (no SSH key sprawl)
✅ **Pass**: All CloudTrail logging enabled (100% audit trail)
✅ **Pass**: No public IP addresses on Public Health HIPAA instances
✅ **Pass**: Auto-stop preventing resource abandonment (cost + security benefit)
❌ **Fail**: 2 researchers had S3 buckets without encryption enabled
❌ **Fail**: 1 researcher copied data to personal laptop without approval

**Pass rate**: 13/15 (87%) vs baseline 40% in personal AWS accounts

**Carlos's assessment**:
> "This is a massive improvement. Session Manager alone eliminates my #1 security concern (SSH key leaks). CloudTrail gives me audit logs for compliance. The 2 failures are researcher education issues, not AWS IDE problems. I recommend campus-wide rollout with security training."

### Qualitative Feedback

**Researcher survey** (end of February):

**Question**: "How satisfied are you with AWS IDE?"

| Response | Count | % |
|----------|-------|---|
| Very satisfied | 11 | 73% |
| Satisfied | 3 | 20% |
| Neutral | 1 | 7% |
| Dissatisfied | 0 | 0% |

**Researcher comments**:

**Positive**:
- "Setup took 15 minutes vs 3 hours with manual AWS" (Biology PhD student)
- "Auto-stop saved me from a $500 mistake - I forgot to stop a GPU instance" (Engineering postdoc)
- "I don't need to bother the help desk anymore" (Public Health faculty)
- "Cost preview before launching helps me budget" (Biology PI)

**Constructive criticism**:
- "I want to customize my environment more" (Engineering expert) → Response: Environment export/import feature coming
- "Can we have GPU quotas per lab?" (Engineering PI) → Response: Feature request filed (#19)

**PI feedback**:

**Prof. Garcia (Biology)**: "My students are more productive. They're doing research instead of fighting with AWS setup. And I can actually see what we're spending - that's huge for grant budget management."

**Prof. Patel (Engineering)**: "My group loves it. The auto-stop feature gives me peace of mind. We're training neural networks without worrying about forgot-to-terminate disasters."

**Prof. Johnson (Public Health)**: "Security audit passed - that's a first for our group. Previously we avoided cloud because HIPAA compliance was too complex. AWS IDE makes it straightforward."

---

## Campus-Wide Rollout Planning (March 2025)

### Pilot Success → Full Deployment

**Jennifer to CIO** (March 15, 2025):

**Memo**: "AWS IDE Pilot Results & Recommendation"

**Executive Summary**:
- **Support tickets**: 87% reduction (200 → 26 tickets/month)
- **Cost visibility**: 100% (vs 40-60% before)
- **Security**: 87% pass rate (vs 40% baseline)
- **Researcher satisfaction**: 93% satisfied/very satisfied
- **Recommendation**: Proceed with campus-wide rollout

**CIO response**: "Approved. Present to Provost for funding."

### Provost Presentation (March 20, 2025)

**Jennifer's pitch**:

**Slide 1**: Problem Statement
- Decentralized cloud usage = no visibility, security risks, cost overruns
- Support team overwhelmed (2,400 tickets/year, 6-week backlog)
- Faculty frustrated (complex AWS, long turnaround, fear of cost overruns)

**Slide 2**: AWS IDE Solution
- Centralized platform with distributed autonomy
- Auto-stop prevents cost overruns (100% elimination in pilot)
- Plain-English interface (87% support ticket reduction)
- Security built-in (Session Manager, CloudTrail)

**Slide 3**: Pilot Results (Numbers)
- 15 researchers, 2 months
- 87% support reduction, 100% cost visibility, 87% security pass rate
- $4,200 pilot spending ($280/researcher/month avg)

**Slide 4**: Scaled Projection
- 1,250 researchers × $280/month × 12 months = $4.2M potential
- With auto-stop (70% reduction): $1.26M/year realistic
- Current spending: $430,000/year (partial visibility)
- **Ask**: $800,000/year budget (centralized AWS account)

**Slide 5**: ROI
- Support cost avoidance: 87% reduction = 2 FTE savings = $180,000/year
- Security risk reduction: HIPAA violation avoidance = $15,000 fine avoided + $50,000 legal fees
- Faculty productivity: Estimated 500 hours/year saved across institution = $25,000 value
- **Total value**: $270,000/year savings + risk avoidance

**Slide 6**: Recommendation
- Approve $800,000/year AWS budget
- Mandate AWS IDE for all institutional cloud computing (phase out personal accounts over 12 months)
- Hire 1 additional cloud architect (backfill for support reduction)

**Provost questions**:

**Q**: "Why not let faculty use personal accounts if they want?"
**A**: "Security and compliance. We can't audit what we don't control. After 2 HIPAA/export control incidents last year, we need institutional oversight."

**Q**: "What if faculty resist?"
**A**: "Pilot feedback was 93% positive. We're making it easier, not harder. Faculty want simple, safe, budget-friendly cloud - AWS IDE delivers."

**Q**: "What about the $800K budget?"
**A**: "We're already spending ~$430K with limited visibility. $800K gives us headroom to support 3x more researchers with full control. And we're saving $270K/year in support costs."

**Provost**: "Approved. Let's do it."

---

## Campus-Wide Rollout (April-August 2025)

### Phase 1: Infrastructure (April 2025)

**Tom (cloud architect) + team configure**:

#### 1. AWS Organizations Structure

```
University AWS Root Account
├── Core OU (central IT services)
│   ├── Prod account (core services)
│   └── Dev account (testing)
├── Research OU
│   ├── Life Sciences OU
│   │   ├── Biology Dept
│   │   ├── Chemistry Dept
│   │   └── Medical School
│   ├── Physical Sciences OU
│   │   ├── Physics Dept
│   │   └── Materials Science
│   ├── Engineering OU
│   │   ├── Mechanical Eng
│   │   └── Electrical Eng
│   ├── Social Sciences OU
│   │   ├── Economics
│   │   ├── Psychology
│   │   └── Sociology
│   └── Other OU
│       ├── Geosciences
│       └── Digital Humanities
└── Teaching OU (course accounts)
    ├── CS Dept
    ├── Stats Dept
    └── Data Science Program
```

**50+ AWS accounts** (one per department/major research group)

#### 2. Service Control Policies (SCPs)

**Security guardrails** (applied to all Research OU accounts):

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Deny",
      "Action": [
        "ec2:RunInstances"
      ],
      "Resource": "arn:aws:ec2:*:*:instance/*",
      "Condition": {
        "StringNotEquals": {
          "ec2:InstanceMetadataServiceVersion": "2"  // Require IMDSv2
        }
      }
    },
    {
      "Effect": "Deny",
      "Action": [
        "s3:PutBucketPublicAccessBlock"
      ],
      "Resource": "*",
      "Condition": {
        "Bool": {
          "s3:BlockPublicAcls": "false"  // Block public S3 buckets
        }
      }
    },
    {
      "Effect": "Deny",
      "Action": [
        "ec2:*"
      ],
      "Resource": "*",
      "Condition": {
        "StringNotEquals": {
          "aws:RequestedRegion": [
            "us-east-1",
            "us-west-2"
          ]  // Limit to 2 regions (cost + compliance)
        }
      }
    }
  ]
}
```

**Cost controls**:
- Budget alerts at department level
- $10,000/month default department budget (adjustable)
- Email alerts to PI + Jennifer at 80%, 90%, 100%

#### 3. Centralized Logging & Monitoring

```bash
# CloudTrail (audit logs) sent to central S3 bucket
aws cloudtrail create-trail \
  --name university-audit-trail \
  --s3-bucket-name university-audit-logs \
  --is-organization-trail

# GuardDuty (threat detection)
aws guardduty enable-organization-admin-account \
  --admin-account-id <central-security-account>

# Cost & Usage Reports
aws cur put-report-definition \
  --report-definition '{ ... }'
```

#### 4. AWS IDE Environment Catalog

**Jennifer's team creates domain-specific environments**:

```bash
# Life Sciences
- genomics-analysis.yaml (BWA, SAMtools, GATK, BioConda)
- medical-imaging.yaml (ITK, SimpleITK, 3D Slicer)
- drug-discovery.yaml (RDKit, PyMOL, AutoDock)

# Physical Sciences
- molecular-dynamics.yaml (LAMMPS, GROMACS, VMD)
- quantum-chemistry.yaml (Gaussian, ORCA, NWChem)
- climate-modeling.yaml (NCL, CDO, xarray)

# Engineering
- cfd.yaml (OpenFOAM, ParaView, Fluent)
- ml-engineering.yaml (PyTorch, TensorFlow, CUDA)
- finite-element.yaml (FEniCS, deal.II, Gmsh)

# Social Sciences
- survey-analysis.yaml (R, tidyverse, ggplot2, stata)
- text-mining.yaml (NLTK, spaCy, transformers)
- econometrics.yaml (R, Stata, EViews)
```

**30+ domain-specific environments** curated by subject matter experts

### Phase 2: Training & Communication (May 2025)

#### 1. Documentation Website

Jennifer's team creates: **research-computing.university.edu/aws-ide**

**Pages**:
- Quick Start (15-minute setup guide)
- Video tutorials (one per discipline)
- FAQs
- Security & compliance guide
- Cost optimization tips
- Environment catalog
- Support contact info

#### 2. Training Sessions

**Weekly workshops** (May-August):
- "AWS IDE 101" (Tuesdays, 2-3 PM) - for beginners
- "Advanced AWS IDE" (Thursdays, 2-3 PM) - for experts
- Attendance: 300+ researchers over 4 months

#### 3. Faculty Champions Program

**Jennifer recruits 10 faculty "champions"** (including pilot PIs):
- Each champions AWS IDE in their college
- Provides peer support to colleagues
- Attends monthly "champion meetings" with Jennifer
- Incentive: $1,000 professional development fund

#### 4. Email Campaign

**Monthly emails** to all faculty/grad students:
- May: "AWS IDE is here - get started in 15 minutes"
- June: "Success stories from your colleagues"
- July: "AWS IDE saves researchers $150K in cost overruns"
- August: "Join 500+ researchers already using AWS IDE"

### Phase 3: Adoption Tracking (May-August 2025)

#### Adoption Metrics

| Month | New Users | Total Users | % of Target (1,250) | Monthly Cost |
|-------|-----------|-------------|---------------------|--------------|
| May | 85 | 100 | 8% | $28,000 |
| June | 120 | 220 | 18% | $61,600 |
| July | 180 | 400 | 32% | $112,000 |
| August | 150 | 550 | 44% | $154,000 |

**Adoption rate**: 44% in 4 months (550/1,250 researchers)

**Jennifer's assessment**: "Ahead of schedule. We projected 30% by end of summer, achieved 44%."

#### Discipline-Specific Adoption

| Discipline | Researchers | Adopted | % Adoption | Avg Monthly Cost/User |
|----------|-------------|---------|------------|----------------------|
| Life Sciences | 375 | 220 | 59% | $320 (high - genomics GPU usage) |
| Engineering | 312 | 160 | 51% | $380 (high - ML GPU usage) |
| Physical Sciences | 250 | 90 | 36% | $240 (moderate) |
| Social Sciences | 188 | 60 | 32% | $80 (low - light compute) |
| Other | 125 | 20 | 16% | $120 (low) |

**Patterns**:
- GPU-heavy disciplines (Life Sci, Engineering) adopt faster (better than HPC cluster)
- Social Sciences adopt slower (Stata/SPSS license issues, R users comfortable with RStudio Desktop)

---

## Success Metrics (6 Months Post-Rollout)

### Support Ticket Metrics

| Category | Before (annual) | After (projected annual) | Reduction |
|----------|----------------|-------------------------|-----------|
| "How do I launch?" | 960 | 125 | 87% |
| Troubleshooting | 600 | 180 | 70% |
| Networking/VPC | 480 | 60 | 88% |
| Cost overruns | 240 | 0 | 100% |
| Security/compliance | 120 | 35 | 71% |
| **TOTAL** | **2,400** | **400** | **83%** |

**Tom's team experience**:
> "We went from 2 overwhelmed architects to 2 productive architects. We now have time for strategic work - building out HPC-cloud hybrid workflows, optimizing costs, training researchers. Ticket backlog cleared for first time in 3 years."

### Cost Visibility & Control

**Before AWS IDE** (FY 2024):
- Central account spending: $180,000 (known)
- Departmental accounts: $250,000 (estimated from finance reports)
- Personal accounts: Unknown ($50,000-500,000 range)
- **Total estimated**: $430,000-930,000/year
- **Visibility**: ~40%

**After AWS IDE** (FY 2025, projected):
- All spending in central account: 100% visible
- Actual spending (first 6 months): $355,600
- Projected annual: $711,200
- Budget: $800,000
- **Utilization**: 89% of budget
- **Visibility**: 100%

**Cost avoidance** (first 6 months):
- Auto-stop prevented waste: Estimated $120,000 (based on 70% idle time reduction)
- No cost overrun incidents: $60,000 saved (vs 2 major incidents in 2024)
- **Total savings**: $180,000

**Jennifer's CFO report**:
> "FY 2025 cloud spending is tracking at $711K vs $800K budget (11% under). All spending is now visible and controlled. We estimate $180K in cost avoidance from auto-stop and eliminating overruns. Security incidents: zero (vs 3 in FY 2024). I recommend continuing this approach."

### Security & Compliance

**Security audit results** (August 2025):

| Metric | Before AWS IDE | After AWS IDE | Improvement |
|--------|---------------|---------------|-------------|
| Instances with Session Manager | 45% | 98% | +118% |
| Instances in private subnets (HIPAA) | 60% | 100% | +67% |
| CloudTrail logging enabled | 55% | 100% | +82% |
| SSH keys properly managed | 40% | 95% (5% using SSH) | +138% |
| S3 buckets encrypted | 65% | 92% | +42% |
| **Overall compliance score** | **53%** | **97%** | **+83%** |

**Carlos (security officer) report**:
> "This is the best security posture we've ever had for research computing. Session Manager eliminates SSH key sprawl. CloudTrail gives us complete audit trails. The 3% non-compliant instances are edge cases (researchers using legacy workflows). I'm confident presenting this to our next HIPAA audit."

**Compliance incidents**:
- FY 2024: 3 major incidents (HIPAA violation, export control breach, NSF audit failure)
- FY 2025 (6 months): 0 incidents
- **Incident reduction**: 100%

### Researcher Satisfaction

**Faculty survey** (end of August 2025):

**Question**: "How satisfied are you with institutional research computing support?"

| Response | FY 2024 | FY 2025 | Change |
|----------|---------|---------|--------|
| Very satisfied | 15% | 58% | +43% |
| Satisfied | 35% | 32% | -3% |
| Neutral | 25% | 8% | -17% |
| Dissatisfied | 18% | 2% | -16% |
| Very dissatisfied | 7% | 0% | -7% |

**Satisfaction rate**: 90% satisfied/very satisfied (vs 50% before)

**Faculty comments**:

**Positive themes**:
- "Setup is finally easy - 15 minutes vs 3 hours" (mentioned by 80% of respondents)
- "Auto-stop gives me peace of mind" (mentioned by 60%)
- "Support team responds quickly now" (mentioned by 55%)
- "I can see what I'm spending" (mentioned by 50%)
- "My students can set themselves up" (mentioned by 45%)

**Constructive feedback**:
- "I want to customize my environment more" (10 faculty) → Feature request filed
- "Can I use Spot instances?" (8 faculty) → Feature coming
- "Need more GPU quota" (5 faculty) → Addressed via quota increase

### Staff Morale & Retention

**ORC team survey** (anonymous):

**Before AWS IDE** (FY 2024):
- Workload: 4.8/10 (overwhelming)
- Job satisfaction: 5.2/10 (burnout)
- Considering leaving: 62.5% (5/8 staff)

**After AWS IDE** (FY 2025):
- Workload: 7.5/10 (manageable)
- Job satisfaction: 8.2/10 (engaged)
- Considering leaving: 12.5% (1/8 staff, unrelated reason - spouse relocation)

**Tom's reflection**:
> "This job is actually enjoyable now. I'm building cool stuff instead of answering 'how do I launch EC2' tickets all day. I was ready to quit 6 months ago. Now I'm excited to come to work."

---

## Strategic Impact (12-Month Reflection)

### Financial Impact (FY 2025)

| Category | Value | Notes |
|----------|-------|-------|
| **Costs** | | |
| AWS spending | $711,200 | All researchers combined |
| AWS IDE licenses | $0 | Open source |
| Additional staff | $0 | No new hires needed (support reduction) |
| **Subtotal Costs** | **$711,200** | |
| | | |
| **Savings** | | |
| Cost overrun avoidance | $180,000 | Auto-stop + visibility |
| Support FTE savings (2 FTE) | $180,000 | 83% ticket reduction |
| Security incident avoidance | $65,000 | No HIPAA/export control incidents |
| Faculty time savings | $125,000 | 500 hours saved × $250/hour |
| **Subtotal Savings** | **$550,000** | |
| | | |
| **Net Impact** | **-$161,200** | Positive ROI even with increased usage |

**CFO to Provost**:
> "The $800K AWS budget was approved with concern about cost. Actual spending was $711K, and we realized $550K in measurable savings and risk avoidance. Net cost: $161K for a platform supporting 550+ researchers (growing to 1,000+). This is one of the best IT investments we've made."

### Research Impact

**Quantified outcomes**:

**Papers published** (FY 2025 vs FY 2024):
- Papers with "computational analysis" in methods: 340 (vs 280 in FY 2024) = +21%
- Papers acknowledging AWS IDE: 85 (new)
- Papers with reproducible environments (supplementary materials): 62 (vs 15 in FY 2024) = +313%

**Grant success rate**:
- NSF grant proposals with data management plans: 100% (vs 65% in FY 2024)
- NIH grant proposals with computational methods: +18% (AWS resources cited)

**Research Computing Award**:
- University nominated for EDUCAUSE Award for Excellence in Research Computing (September 2025)
- Citation: "Innovative cloud platform enabling 550+ researchers with 100% cost visibility and security compliance"

### Institutional Reputation

**Recruit faculty candidates** (spring 2025 hiring):

**Before**: "Do you have GPU cluster access?"
**Now**: "Do you have cloud computing support?" → "Yes, full AWS access with institutional support via AWS IDE"

**Faculty retention**:
- 3 faculty cited research computing infrastructure as reason to stay (retention bonuses totaling $450K)
- ROI: $450K retention vs $1.5M replacement cost (recruiting new faculty)

**Student recruitment**:
- Graduate program brochures now highlight "state-of-the-art cloud computing infrastructure"
- Prospective PhD students in data-intensive fields (CS, bioinformatics, physics) cite computing as attraction

---

## Lessons Learned & Best Practices

### What Worked Exceptionally Well

1. **Pilot program**: Testing with 15 researchers before campus-wide rollout built confidence and refined approach
2. **Faculty champions**: Peer advocacy more effective than top-down mandate
3. **Domain-specific environments**: 30+ curated environments reduced setup friction
4. **Auto-stop as default**: Single feature addressed biggest concerns (cost control, security risk)
5. **Centralized visibility**: 100% spending transparency enabled data-driven decisions
6. **Security built-in**: Session Manager, CloudTrail eliminated major compliance gaps

### Implementation Challenges & Solutions

**Challenge 1**: Legacy workflows (researchers using old tools)
- **Solution**: Gradual migration, not forced; support personal accounts in parallel for 12 months
- **Result**: 44% adoption in 4 months (organic growth)

**Challenge 2**: Stata/SPSS license integration
- **Solution**: Worked with vendors on campus-wide license server accessible from cloud
- **Result**: Social sciences adoption increased from 10% to 32%

**Challenge 3**: HPC-cloud hybrid workflows
- **Solution**: Built connectors between on-prem cluster and AWS (data transfer, job submission)
- **Result**: Researchers use HPC for big jobs, AWS IDE for interactive analysis

**Challenge 4**: GPU quota limits
- **Solution**: Implemented approval workflow for GPU instances (5-minute turnaround)
- **Result**: No bottleneck, controlled costs

### Organizational Changes

**Before AWS IDE**:
- ORC team: 100% reactive (responding to tickets)
- Support model: "Help researchers figure out AWS"
- Satisfaction: 50% (faculty frustrated, staff burned out)

**After AWS IDE**:
- ORC team: 70% proactive (building infrastructure, curating environments, training)
- Support model: "Enable researcher self-sufficiency"
- Satisfaction: 90% faculty, 82% staff

**Cultural shift**:
- Faculty: From "cloud is too hard/expensive" → "cloud is my default"
- IT: From "we control everything" → "we enable autonomy within guardrails"
- Institution: From decentralized chaos → centralized visibility with distributed execution

---

## Future Roadmap (Jennifer's 3-Year Plan)

### Year 1 (FY 2026): Scale to 1,000+ Researchers

**Goals**:
- 80% adoption (1,000 of 1,250 researchers)
- $1.2M spending (within budget)
- <200 support tickets/month
- 0 security incidents

**Key initiatives**:
1. Decommission aging HPC cluster (15 years old) → cloud-first strategy
2. Expand environment catalog to 50+ domains
3. Build HPC-cloud hybrid scheduler
4. Implement chargeback model (departments pay for their usage)

### Year 2 (FY 2027): Advanced Features

**Goals**:
- Spot instance integration (70% cost savings for batch workloads)
- S3 data lake (centralized research data storage)
- GPU quota self-service (no approval workflow)
- Multi-cloud support (Azure for Microsoft-specific tools)

**Key initiatives**:
1. Cost optimization: Spot instances, reserved instances, savings plans
2. Data management: S3 lifecycle policies, Glacier archival
3. Collaboration: Instance sharing, team workspaces
4. Reproducibility: Environment version control, DOI assignment

### Year 3 (FY 2028): Ecosystem & Community

**Goals**:
- Environment marketplace (researchers share environments)
- Inter-institution collaboration (share with partner universities)
- Industry partnerships (AWS research credits program)
- Training pipeline (undergrads learn cloud skills)

**Key initiatives**:
1. Create "Research Computing Consortium" with 5 partner universities
2. Joint procurement (volume discounts)
3. Shared environment library (cross-institution reproducibility)
4. Cloud curriculum for undergraduate CS/data science programs

---

## Technical Details: Enterprise Architecture

### AWS Organizations Structure

```
Root Account (Billing + Organizations)
├── Security OU
│   ├── Audit Account (CloudTrail, Config, GuardDuty aggregation)
│   └── Log Archive Account (centralized logs)
├── Infrastructure OU
│   ├── Network Account (Transit Gateway, VPN, Direct Connect)
│   └── Shared Services Account (DNS, AD, LDAP)
├── Research OU (50+ accounts)
│   ├── Life Sciences OU
│   ├── Physical Sciences OU
│   ├── Engineering OU
│   ├── Social Sciences OU
│   └── Other OU
└── Teaching OU (30+ accounts)
    └── Course accounts per semester
```

**Total**: 85 AWS accounts (50 research + 30 teaching + 5 core)

### Centralized Services

**Identity & Access**:
- AWS SSO integrated with university LDAP/Active Directory
- Single sign-on: researchers use university NetID credentials
- Role-based access: faculty (admin), grad students (user), postdocs (power-user)
- MFA required for all access

**Networking**:
- Transit Gateway: Hub-and-spoke architecture
- Direct Connect: 10 Gbps link to campus network
- VPN: Backup connectivity
- Private connectivity to on-prem HPC cluster

**Cost Management**:
- AWS Cost & Usage Reports → Athena → QuickSight dashboards
- Budget alerts per department (email at 80%, 90%, 100%)
- Monthly chargeback reports (departments billed for their usage)
- Cost optimization recommendations (Compute Optimizer, Trusted Advisor)

**Security & Compliance**:
- CloudTrail: All API calls logged to central S3 bucket (7-year retention for compliance)
- GuardDuty: Threat detection across all accounts
- Security Hub: Aggregated security findings
- Config: Configuration compliance tracking
- Macie: S3 data classification (detects PII/PHI)

### Cost Tracking Architecture

**Tagging strategy**:
```
Key: Department, Value: Biology
Key: PI, Value: Garcia
Key: Project, Value: NIH-R01-2025
Key: User, Value: john.doe@university.edu
Key: CostCenter, Value: CC-12345
Key: Grant, Value: NSF-12345
```

**Automated tagging**:
- AWS IDE auto-tags instances with user info
- Lambda function enriches tags with department/PI from LDAP
- Cost Allocation Tags enabled in billing

**Dashboards** (QuickSight):
1. **Executive Dashboard** (Provost, CIO):
   - Total spending by college
   - Trend over time (monthly, quarterly)
   - Top 10 spenders (departments)

2. **Department Dashboard** (Dept chairs):
   - Spending by PI
   - Budget vs actual
   - Top cost drivers (instance types, services)

3. **PI Dashboard** (Faculty):
   - Spending by student/project
   - Daily usage patterns
   - Cost optimization recommendations

4. **IT Operations Dashboard** (Jennifer's team):
   - Instance counts, utilization
   - Security findings
   - Support ticket trends

### Disaster Recovery & Business Continuity

**Backup strategy**:
- All EBS volumes: Daily snapshots (7-day retention)
- S3 data: Cross-region replication (us-east-1 → us-west-2)
- CloudTrail logs: Replicated to Glacier (7-year retention)

**Disaster scenarios**:
1. **AZ failure**: Auto-restart instances in different AZ (5-minute RTO)
2. **Region failure**: Manual failover to us-west-2 (4-hour RTO, with DR plan)
3. **Ransomware**: Restore from snapshots (1-day RPO)

---

## Conclusion

AWS IDE transformed the university's research computing infrastructure from decentralized chaos to centralized visibility with distributed autonomy:

### Key Outcomes (12 Months)
- ✅ **83% support ticket reduction** (2,400 → 400 tickets/year)
- ✅ **100% cost visibility** (vs 40% before)
- ✅ **97% security compliance** (vs 53% before)
- ✅ **90% researcher satisfaction** (vs 50% before)
- ✅ **550+ researchers using cloud** (44% adoption in 6 months)
- ✅ **Zero security incidents** (vs 3 major incidents in previous year)
- ✅ **$550,000 realized savings** (support, cost avoidance, incident prevention)

### ROI Summary
- **Investment**: $711,200 (AWS spending)
- **Savings**: $550,000 (support, incidents, faculty time)
- **Net cost**: $161,200 for 550+ researchers
- **Cost per researcher**: $293/year (vs $430/year estimated before)
- **Intangible value**: Better security, reproducible research, faculty satisfaction

### Success Factors
1. **Pilot program**: Validated approach before large investment
2. **Auto-stop**: Single feature addressed multiple concerns (cost, security, waste)
3. **Centralized visibility**: 100% spending transparency enabled data-driven decisions
4. **Distributed execution**: Researchers self-sufficient, no IT bottleneck
5. **Security built-in**: Session Manager, CloudTrail, SCPs eliminated compliance gaps
6. **Faculty champions**: Peer advocacy drove organic adoption

### Strategic Impact

**Research outcomes**:
- +21% papers with computational methods
- +313% papers with reproducible environments
- 100% grant proposals with data management plans (vs 65%)

**Institutional benefits**:
- Faculty retention: 3 faculty cited computing infrastructure
- Student recruitment: Computing highlighted in recruiting materials
- Industry reputation: Nominated for EDUCAUSE Award

**Cultural transformation**:
- Faculty: From cloud-averse to cloud-first
- IT: From reactive support to proactive enablement
- Institution: From hidden cloud sprawl to managed platform

**Jennifer's reflection** (12-month retrospective):
> "One year ago, we had no visibility into cloud spending, 3 security incidents, and an overwhelmed support team. Today, we support 550+ researchers with 100% cost visibility, zero incidents, and a happy team. AWS IDE didn't just solve our technical problems - it changed how we think about research computing. We went from trying to control everything to enabling autonomy within smart guardrails. That's the future of academic IT."

Jennifer represents the **Research Computing Manager persona**: institutional-level oversight, managing hundreds of researchers, balancing security/compliance with researcher autonomy, limited support staff, need for cost visibility and control. AWS IDE addresses all requirements while transforming IT from bottleneck to enabler.

**Related GitHub Issues**: #12, #13, #14, #15, #16, #17, #19, #20, #21, #28, #29
