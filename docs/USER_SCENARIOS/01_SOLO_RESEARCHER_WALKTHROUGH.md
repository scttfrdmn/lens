# Solo Researcher Walkthrough: Dr. Maria Rodriguez

> **Persona**: Computational biologist analyzing genomic data
> **Technical Level**: 2/5 (Can use R/Python, basic command line, intimidated by AWS)
> **Budget**: $1,200/year from discretionary grant funds
> **Primary Pain Point**: Laptop insufficient for large datasets, AWS too complex and scary

---

## Profile

**Name**: Dr. Maria Rodriguez
**Position**: Postdoctoral Research Associate, Molecular Biology
**Institution**: Mid-size state university
**Research Focus**: Comparative genomics of antibiotic-resistant bacteria
**Age**: 32
**Location**: Phoenix, Arizona (US West Coast, uses us-west-2)

### Technical Background
- **Comfortable with**: R, basic Python, Excel, analyzing CSV files
- **Uncomfortable with**: Command line beyond basic commands, AWS console, infrastructure
- **Has used**: Personal laptop (8GB RAM), lab desktop occasionally
- **Never used**: Cloud computing (AWS/GCP/Azure), Docker, server administration

### Research Context
- **Daily workflow**: Download bacterial genome sequences from NCBI → compare against reference genomes → identify resistance genes
- **Dataset sizes**: 50-500 GB per analysis project
- **Computation needs**: Intermittent (2-3x per week, 2-4 hours per session)
- **Collaboration**: Solo work, occasional data sharing with advisor
- **Publications**: 2 papers/year, reviewers increasingly demand reproducible analysis

### Budget Reality
- **Annual budget**: $1,200 discretionary funds from postdoc training grant
- **Other costs**: Conferences ($500), lab supplies ($300), subscriptions ($200)
- **Available for compute**: ~$200/quarter = $50/month realistic budget
- **Anxiety level**: HIGH - one mistake could blow entire budget
- **Advisor oversight**: Advisor trusts Maria but won't approve overspending

---

## Current Situation (Before Lens)

### Existing Setup
**Primary workstation**: 2019 MacBook Pro
- 8 GB RAM (insufficient for large genome alignments)
- 512 GB SSD (runs out of space frequently)
- Analysis that should take 2 hours takes 8-12 hours
- Laptop unusable during long computations

**Problems**:
1. **Insufficient compute**: Genome alignment of 50GB dataset → laptop freezes for 8 hours
2. **Data management**: Constantly deleting files to free up space
3. **Opportunity cost**: Can't use laptop for writing during analysis runs
4. **Reproducibility**: Colleagues can't recreate analysis (different R package versions)

### Previous AWS Attempts
Maria tried to use AWS EC2 directly (September 2024):

**Day 1** (Friday, 3 hours):
- Googled "AWS genomics analysis"
- Created AWS account, got lost in AWS Console
- Tried to launch EC2 instance:
  - What's a VPC? What's a subnet?
  - What's a security group?
  - Which AMI should I use?
  - What instance type do I need?
- Eventually launched t2.micro by accident (too small)
- Couldn't SSH (security group misconfigured)
- **Gave up after 3 hours**

**Day 2** (Monday):
- Watched 2 YouTube tutorials on EC2
- Successfully launched t3.xlarge
- Forgot to configure storage → only 8GB root volume (need 200GB)
- Terminated and relaunched with 200GB
- **Finally working after 2 more hours**

**Day 3** (Tuesday - Friday):
- Installed R, Bioconductor packages (2 hours)
- Ran analysis successfully (4 hours)
- Downloaded results
- **Forgot to terminate instance**

**Day 10** (Next Monday):
- Received AWS bill: **$320 for 7 days** (instance ran continuously)
- Panic! That's 64% of quarterly budget gone
- Terminated instance immediately
- **Decided AWS was "too expensive and complicated"**

### Pain Points Summary

| Pain Point | Current Impact | Quantified Cost |
|-----------|----------------|----------------|
| Laptop insufficient | Analysis takes 8 hours vs 2 hours | 6 hours lost per analysis |
| No cloud access | Can't run large datasets | 3 projects delayed |
| AWS too complex | 5 hours learning, still failed | 5 hours wasted |
| Forgot to stop instance | $320 surprise bill | 64% of quarterly budget |
| Can't reproduce analysis | Paper revisions difficult | 8 hours per revision |

**Total impact**:
- **Time lost**: 6 hours per week × 40 weeks = 240 hours/year (6 work weeks)
- **Money wasted**: $320 mistake (64% of budget) → risk aversion prevents cloud usage
- **Research delayed**: 3 projects waiting for compute access

---

## Lens Workflow

### Initial Setup (First Time Use)

**Monday Morning, 9:00 AM** - Maria decides to try Lens after colleague recommendation

#### Step 1: Installation (5 minutes)

Maria follows the installation guide:

```bash
# macOS installation
brew tap scttfrdmn/tap
brew install lens-jupyter
```

**Output**:
```
==> Tapping scttfrdmn/tap
==> Downloading https://github.com/scttfrdmn/lens/releases/download/v0.7.2/lens-jupyter_0.7.2_darwin_amd64.tar.gz
==> Installing lens-jupyter from scttfrdmn/tap
🍺  /opt/homebrew/Cellar/lens-jupyter/0.7.2: 1 file, 14.2MB
```

**Maria's reaction**: ✅ "That was easy! Just like installing normal software."

#### Step 2: AWS Configuration (10 minutes)

Maria has AWS account from previous failed attempt:

```bash
aws configure
```

**Prompts**:
```
AWS Access Key ID [None]: AKIAIOSFODNN7EXAMPLE
AWS Secret Access Key [None]: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
Default region name [None]: us-west-2
Default output format [None]: json
```

**Maria's reaction**: ✅ "Okay, I remember this part from last time. At least I have credentials saved."

#### Step 3: First Launch with Wizard (3 minutes)

```bash
lens-jupyter
```

**Lens launches interactive wizard:**

```
🚀 Welcome to AWS Jupyter!

Let's get you set up. I'll ask a few quick questions.

❓ What type of analysis do you want to do?

  1) Data science (Python + R)
  2) Statistical analysis (R-focused)
  3) Machine learning / Deep learning
  4) Bioinformatics (genomics, RNA-seq)  ← Maria selects this
  5) Other

Your choice: 4

Great! I'll set up a bioinformatics environment with:
  - Python 3.11 + R 4.3
  - BioConda packages (BWA, SAMtools, GATK)
  - Jupyter Lab with genomics visualization tools

❓ How much compute power do you need?

  1) Light (for small datasets < 10GB)
  2) Medium (for typical datasets 10-100GB)  ← Maria selects this
  3) Heavy (for large datasets > 100GB or many samples)
  4) GPU (for deep learning models)

Your choice: 2

Perfect! I recommend: t4g.large
  - 2 CPUs, 8 GB RAM
  - Good for typical genomic analysis workflows

❓ How long will your analysis run?

  1) Quick analysis (< 1 hour)
  2) Typical workflow (2-4 hours)  ← Maria selects this
  3) Long computation (4-8 hours)
  4) Very long (8+ hours)

Your choice: 2

Got it! I'll set up auto-stop after 2 hours of being idle.
This saves money when you forget to shut down.

💰 Cost Estimate:
   Instance: t4g.large
   Hourly: $0.0672/hr
   Your session (4 hours): ~$0.27
   If you ran 24/7: $48.38/month
   With auto-stop (2h/day): ~$4.03/month ✨

This fits comfortably in your $50/month budget!

❓ Ready to launch?

  y) Yes, launch my environment!
  n) No, let me change something

Your choice: y
```

**Maria's reaction**: ✅ "Wow, this is actually explaining things! And it shows me the cost BEFORE launching. That's exactly what I needed last time."

#### Step 4: Launch Progress (2 minutes)

```
🚀 Launching your bioinformatics environment...

✓ Creating secure network setup
✓ Launching instance in us-west-2
✓ Installing bioinformatics tools (this takes 2-3 minutes)

⏳ Installing packages:
   ▸ cloud-init: Setting up Python 3.11...
   ▸ cloud-init: Installing conda and BioConda...
   ▸ cloud-init: Installing BWA, SAMtools, GATK, BLAST...
   ▸ cloud-init: Configuring Jupyter Lab...

✓ Your environment is ready!

🌐 Jupyter Lab URL: http://54.186.132.45:8888
🔑 Token: 2a91c4f6e4b7d8e9f1234567890abcde

📋 To connect:
   1. Click the URL above (or copy-paste into browser)
   2. The token is already included in the URL

💡 Tips:
   - Your environment will auto-stop after 2 hours idle (saves money!)
   - To stop manually: lens-jupyter stop
   - To check costs: lens-jupyter costs
   - Need help? lens-jupyter --help

📊 Instance Details:
   ID: i-0abc123def456789
   Type: t4g.large
   Region: us-west-2
   Environment: bioinformatics
```

**Maria's reaction**: ✅ "That was SO much easier than last time! It took 15 minutes total vs 5 hours, and I actually understand what's happening."

**Total setup time**: **18 minutes** (vs 5 hours with manual AWS)
**Success rate**: **100%** (vs failed on first attempt)

---

## Daily Workflow: Typical Analysis Session

### Scenario: Comparing 20 bacterial genomes

**Tuesday, 10:00 AM** - Maria needs to analyze new genome sequences

#### Step 1: Check Running Instances (5 seconds)

```bash
lens-jupyter status
```

**Output**:
```
📊 Your Jupyter Environments:

i-0abc123def456789 (bioinformatics) - STOPPED
  Last used: Yesterday at 2:15 PM
  Runtime: 3.2 hours
  Cost yesterday: $0.22
```

**Maria's reaction**: ✅ "Good, it auto-stopped like it said it would! And I can see it only cost $0.22 yesterday."

#### Step 2: Restart Instance (30 seconds)

```bash
lens-jupyter start
```

**Output**:
```
🚀 Starting your environment...

✓ Instance starting (takes ~30 seconds)
✓ Jupyter Lab is ready!

🌐 URL: http://54.186.132.45:8888
🔑 Same token as before: 2a91c4f6e4b7d8e9f1234567890abcde

💡 Your previous notebooks and data are exactly as you left them.
```

#### Step 3: Analysis Work (3 hours)

Maria opens Jupyter Lab in browser:

1. **Upload genome sequences** (15 minutes)
   - Drag-drop FASTA files from NCBI
   - 50 GB total

2. **Run alignment pipeline** (2 hours)
   ```python
   # In Jupyter notebook
   !bwa index reference_genome.fasta
   !bwa mem -t 2 reference_genome.fasta sample1.fastq > sample1.sam
   # ... repeat for 20 samples
   ```

3. **Visualize results** (30 minutes)
   ```python
   import pandas as pd
   import matplotlib.pyplot as plt

   results = pd.read_csv('alignment_results.csv')
   results.plot.scatter(x='coverage', y='identity')
   plt.savefig('resistance_genes.png')
   ```

4. **Download results** (15 minutes)
   - Download CSV files
   - Download plots
   - Download updated notebooks

#### Step 4: Cleanup (10 seconds)

```bash
lens-jupyter stop
```

**Output**:
```
⏸️  Stopping your environment to save money...

✓ Jupyter Lab stopped
✓ Instance stopped

💰 Today's Session:
   Runtime: 3.5 hours
   Cost: $0.24
   Monthly projection: $7.20 (well under your $50 budget!)

💾 Your data is saved. Restart anytime with: lens-jupyter start
```

**Maria's reaction**: ✅ "I remembered to stop it this time! And I can see exactly how much it cost - $0.24 is totally reasonable."

### Daily Workflow Summary

| Step | Time | vs Laptop | vs Manual AWS |
|------|------|-----------|---------------|
| Start instance | 30 sec | N/A | 30 sec |
| Upload data | 15 min | 15 min | 15 min |
| Run analysis | 2 hours | **8 hours** | 2 hours |
| Visualize | 30 min | 30 min | 30 min |
| Download results | 15 min | N/A | 15 min |
| Stop instance | 10 sec | N/A | 10 sec (if remembered) |
| **TOTAL** | **3.2 hours** | **9.75 hours** | **3.2 hours (if successful)** |

**Time saved**: **6.5 hours per analysis** (67% faster)
**Cost**: **$0.24 per session**
**Peace of mind**: Auto-stop prevents runaway costs

---

## Cost Management

### Monthly Cost Tracking

**End of Month 1** - Maria checks her spending:

```bash
lens-jupyter costs
```

**Output**:
```
💰 Cost Summary - September 2025

📊 Instances:
   i-0abc123def456789 (bioinformatics)
     Type: t4g.large ($0.0672/hr)
     Total runtime: 42.5 hours (13% of month)
     Total cost: $2.86
     Utilization: Active 42.5h, Stopped 677.5h

💡 Cost Breakdown:
   Compute (EC2): $2.86
   Storage (20GB EBS): $2.00
   Data Transfer: $0.14
   ─────────────────────
   Total: $5.00

📈 Comparison:
   If ran 24/7: $48.38
   Your actual: $5.00
   Savings: $43.38 (90%)

🎯 Budget Status:
   Monthly budget: $50.00
   Spent: $5.00 (10%)
   Remaining: $45.00

✅ You're doing great! Well under budget.
```

**Maria's reaction**: ✅ "Only $5 for the entire month! That's amazing. I ran 12 analysis sessions and stayed way under budget."

### Cost Comparison: Month 1

| Scenario | Cost | Notes |
|----------|------|-------|
| **Laptop only** | $0 | But 6.5 hours wasted per analysis × 12 = 78 hours lost |
| **Manual AWS (with mistake)** | $320 | Maria's actual experience (forgot to stop) |
| **Manual AWS (perfect use)** | $28 | If Maria never forgot to stop (unrealistic) |
| **Lens with auto-stop** | **$5** | Actual cost with automated cost control |

**ROI Analysis**:
- **Time saved**: 78 hours per month @ $25/hr academic time value = **$1,950 value**
- **Money spent**: $5
- **ROI**: 390:1

---

## Pain Points & Solutions

### Pain #1: Laptop Too Slow for Genomic Analysis

**Before Lens**:
- Analysis runtime: 8 hours on laptop
- Laptop unusable during computation
- Frequently runs out of RAM → crashes
- Limited to datasets < 20GB

**With Lens**:
- Analysis runtime: 2 hours on t4g.large
- Can continue working on laptop during analysis
- 8GB RAM sufficient for most workflows
- Datasets up to 100GB feasible

**Success Metric**:
- ✅ 75% time reduction (8 hours → 2 hours)
- ✅ 4x larger datasets possible
- ✅ Laptop remains usable for email/writing

**Related GitHub Issues**:
- #28 - GPU support for even larger datasets
- #13 - Domain-specific bioinformatics templates

**Related Requirements**: REQ-3.1 (Jupyter Lab Support), REQ-7.1 (Fast Launch Times)

---

### Pain #2: AWS Too Complex and Intimidating

**Before Lens**:
- 5 hours to learn AWS basics (VPC, security groups, SSH)
- Still failed on first attempt
- Anxiety about making expensive mistakes
- Gave up and avoided cloud computing

**With Lens**:
- 18 minutes from installation to working environment
- Wizard asks plain-English questions
- No need to understand AWS concepts
- Success on first try

**Success Metric**:
- ✅ 95% time reduction in setup (5 hours → 15 minutes)
- ✅ 100% success rate (vs 0% with manual AWS)
- ✅ Zero AWS jargon in user interface
- ✅ Colleague adoption: Maria recommended to 3 other postdocs, all successful

**Related GitHub Issues**:
- #1 - Make wizard the default
- #2 - Add quickstart for repeat users
- #3 - Remember user preferences

**Related Requirements**: REQ-1.1 (Beginner-Friendly Onboarding), REQ-1.2 (Plain-English Errors)

---

### Pain #3: Cost Anxiety and Budget Blowouts

**Before Lens**:
- $320 surprise bill from forgotten instance (64% of budget)
- Anxiety prevented using cloud computing
- No visibility into costs until bill arrived
- Couldn't trust self to remember to stop instances

**With Lens**:
- Cost preview before launching
- Auto-stop prevents runaway costs
- Real-time cost tracking
- Monthly cost: $5 (10% of budget) vs $320 mistake

**Success Metric**:
- ✅ 98% cost reduction ($320 → $5 per month)
- ✅ 100% budget protection (auto-stop prevents overruns)
- ✅ Cost visibility: Can check spending anytime
- ✅ Confidence: No longer afraid to use cloud

**Related GitHub Issues**:
- #19 - Budget alerts at 50%/75%/90%
- #20 - Email notifications before auto-stop
- #21 - Cost reports for grant reporting

**Related Requirements**: REQ-2.1 (Cost Preview), REQ-2.2 (Auto-Stop), REQ-2.3 (Budget Tracking)

---

### Pain #4: Analysis Not Reproducible

**Before Lens**:
- Colleague: "Can you share your analysis?"
- Maria: "Sure!" → emails R script
- Colleague: "I'm getting different results"
- Problem: Different R package versions, different BioConda versions
- Solution: 8 hours emailing back-and-forth about environment details

**With Lens**:
- Maria exports environment:
  ```bash
  lens-jupyter env export > maria-genomics-analysis.yaml
  ```
- Emails YAML file to colleague
- Colleague launches identical environment:
  ```bash
  lens-jupyter launch --env maria-genomics-analysis.yaml
  ```
- Analysis reproduces perfectly

**Success Metric**:
- ✅ 95% time reduction (8 hours → 30 minutes)
- ✅ 100% reproducibility (identical package versions)
- ✅ Shareable with colleagues
- ✅ Suitable for paper supplements

**Related GitHub Issues**:
- #12 - Environment export/import
- #13 - BioConda integration
- #14 - Community environment templates

**Related Requirements**: REQ-4.1 (Environment Export), REQ-4.2 (BioConda Integration)

---

### Pain #5: No Collaboration with Advisor

**Before Lens**:
- Advisor: "Can I review your analysis?"
- Maria: Exports notebook → emails → advisor imports → different results
- Or: Screen-share via Zoom (slow, frustrating)

**With Lens** (future v0.10.0):
- Maria shares instance with advisor:
  ```bash
  lens-jupyter share i-0abc123def456789 --email advisor@university.edu --duration 2h --read-only
  ```
- Advisor receives email with one-click access link
- Advisor reviews analysis in real-time
- Read-only mode prevents accidental changes
- Access expires automatically after 2 hours

**Success Metric**:
- ⏳ 90% time reduction in advisor reviews (pending implementation)
- ⏳ Real-time collaboration vs async email exchange
- ⏳ No "works on my machine" problems

**Related GitHub Issues**:
- #15 - Instance sharing with lab members
- #17 - S3 sync for data backup

**Related Requirements**: REQ-5.1 (Instance Sharing), REQ-5.3 (S3 Sync)

---

## Success Metrics: 3 Months Later

### Time Savings

| Activity | Before Lens | With Lens | Improvement |
|----------|----------------|--------------|-------------|
| Initial AWS setup | 5 hours (failed) | 15 minutes | 95% faster |
| Per-analysis runtime | 8 hours | 2 hours | 75% faster |
| Analyses per month | 8 (limited by time) | 12 | 50% more |
| Environment sharing | 8 hours troubleshooting | 30 minutes | 94% faster |
| **Total time saved** | - | **78 hours/month** | **= 2 weeks/month** |

### Cost Savings

| Month | Scenario | Cost | Notes |
|-------|----------|------|-------|
| Sept | Manual AWS (mistake) | $320 | Forgot to stop instance |
| Oct | Lens | $5 | 12 analysis sessions with auto-stop |
| Nov | Lens | $6 | 14 analysis sessions |
| Dec | Lens | $5 | 11 analysis sessions |
| **Q4 Total** | **Lens** | **$16** | **vs $960 if repeated mistakes** |

**Savings**: $944 per quarter (98% reduction)

### Research Output

| Metric | Before Lens | After Lens (3 months) | Improvement |
|--------|----------------|--------------------------|-------------|
| Datasets analyzed | 8 per quarter | 35 per quarter | 4.4x more |
| Papers drafted | 0.5 per quarter | 1.5 per quarter | 3x more |
| Collaborations | 1 (difficult) | 4 (easy) | 4x more |
| Confidence level | 3/10 (afraid of cloud) | 9/10 (comfortable) | +200% |

### Qualitative Feedback

**Maria's quote (3 months later)**:
> "Lens transformed my research. I was terrified of cloud computing after that $320 bill, but now I use it 3-4 times per week without thinking about it. The wizard makes it feel like launching a normal app, not configuring a server. And auto-stop means I sleep soundly knowing I won't wake up to a $500 bill. I've analyzed 4x more datasets this quarter than last year's entire year. I recommended it to 3 colleagues and they all love it too."

**Advisor's feedback**:
> "Maria's productivity has increased dramatically. She's analyzing datasets that were previously impossible on her laptop, and the reproducible environments mean I can verify her analysis myself. The cost is minimal compared to the research value. This is exactly what academic researchers need."

---

## Technical Details

### Typical Instance Configuration

**Instance Type**: t4g.large (ARM64 Graviton3)
- 2 vCPUs
- 8 GB RAM
- $0.0672/hour = $1.61/day (24/7) = $48.38/month (24/7)
- With auto-stop (2h/day average): **$4/month**

**Storage**: 20 GB EBS (general purpose SSD)
- $0.10/GB/month = $2.00/month
- Persists when instance stopped
- Enough for notebooks + intermediate files
- Large datasets stored temporarily during analysis

**Environment**: bioinformatics
- Python 3.11 + R 4.3
- BioConda packages: BWA, SAMtools, GATK, BLAST, bcftools
- Python packages: pandas, biopython, numpy, scipy, matplotlib
- R packages: Biostrings, GenomicRanges, ggplot2
- Jupyter Lab with bioinformatics extensions

**Region**: us-west-2 (Oregon)
- Closest to Arizona (low latency)
- Lower costs than us-east-1
- Good availability of instance types

### Usage Pattern

**Frequency**: 3-4x per week
**Duration per session**: 2-4 hours
**Monthly runtime**: 40-50 hours (vs 720 hours if left running 24/7)
**Utilization**: 6-7% (vs 100% if no auto-stop)

**Cost calculation**:
- Compute: 45 hours × $0.0672/hr = $3.02
- Storage: 20 GB × $0.10/GB = $2.00
- Data transfer: ~$0.10
- **Total: ~$5/month**

---

## Lessons Learned

### What Works Well

1. **Wizard interface**: Maria never reads documentation, wizard guides her through every decision
2. **Cost preview**: Seeing "$0.27 for 4 hours" before launching eliminates anxiety
3. **Auto-stop**: Maria never worries about forgetting to stop instances
4. **Plain English**: "Environment" instead of "EC2 instance" makes it less intimidating
5. **Fast restart**: 30-second restart makes it feel like resuming work, not launching a server

### What Could Be Better (Feature Requests)

1. **Email notifications** (#20): "Your analysis is done, Jupyter is ready at: http://..."
   - Would be helpful when launching at start of meeting (environment ready when meeting ends)

2. **Data persistence** (#17): S3 sync for automatic backup
   - Current: Must manually download results before stopping
   - Future: Automatic sync to S3, restore on next launch

3. **Larger datasets** (#28): GPU instances for very large genomes
   - Current: t4g.large handles most bacteria (2-5 MB genomes)
   - Future: Need GPU for human genome (3,000 MB) or metagenomic samples

4. **Conda integration** (#11): Import existing environment.yml
   - Current: Use built-in bioinformatics environment
   - Future: Import custom environments from Conda

5. **Quickstart command** (#2): Skip wizard for repeat launches
   - Current: Wizard is helpful but adds 2 minutes
   - Future: `lens-jupyter quickstart --last` to reuse previous config

---

## Conclusion

Lens transformed Maria from a **cloud-avoider** to a **cloud-enthusiast**:

### Before Lens
- ❌ Terrified of AWS after $320 mistake
- ❌ Limited to laptop (8-hour analyses, frequent crashes)
- ❌ Avoided cloud computing entirely
- ❌ Research progress blocked by compute limitations
- ❌ Budget anxiety prevented exploration

### After Lens
- ✅ Uses cloud 3-4x per week confidently
- ✅ Analyses run 4x faster (2 hours vs 8 hours)
- ✅ 4x research output (35 datasets vs 8 per quarter)
- ✅ $5/month cost (98% cheaper than manual AWS)
- ✅ Zero budget anxiety (auto-stop provides safety net)
- ✅ Recommended to 3 colleagues (all adopted successfully)

**Key Success Factors**:
1. **Wizard interface** - Maria never had to read AWS documentation
2. **Cost transparency** - Preview before launch eliminated anxiety
3. **Auto-stop** - Safety net prevents runaway costs
4. **Plain English** - No intimidating AWS jargon
5. **Fast setup** - 15 minutes vs 5 hours (95% faster)

**ROI**:
- **Time**: 78 hours saved per month (2 work weeks)
- **Cost**: $5/month vs $320 mistakes
- **Research**: 4x more datasets analyzed
- **Confidence**: 3/10 → 9/10

Maria represents the **core target user**: academic researcher with limited technical skills and tight budget who needs cloud compute to be **simple, safe, and affordable**. Lens succeeds by meeting all three requirements.
