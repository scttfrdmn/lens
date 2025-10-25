# Graduate Student Walkthrough: Alex Kim

> **Persona**: 3rd-year PhD student in machine learning
> **Technical Level**: 3/5 (Comfortable with Python, Git, basic cloud concepts)
> **Budget**: $500/semester from advisor's grant (strictly enforced)
> **Primary Pain Point**: Need GPU for deep learning but advisor worried about cost overruns

---

## Profile

**Name**: Alex Kim
**Position**: PhD Student (Year 3 of 5), Computer Science
**Institution**: Large research university
**Research Focus**: Computer vision for medical imaging (cancer detection in X-rays)
**Age**: 26
**Location**: Boston, Massachusetts (uses us-east-1)

### Technical Background
- **Comfortable with**: Python, PyTorch, Jupyter notebooks, Git, command line
- **Uncomfortable with**: AWS infrastructure, production deployments, cost optimization
- **Has used**: Google Colab (free tier), lab GPU server (shared, always busy)
- **Never used**: AWS EC2, managing cloud costs, production ML infrastructure

### Academic Context
- **Advisor**: Prof. Sarah Chen (cautious about cloud spending after previous student overspend)
- **Lab**: 5 PhD students, 2 postdocs, shared resources
- **Deadline pressure**: Dissertation defense in 18 months, needs 3 more publications
- **Competition**: Other students need lab GPU server, scheduling conflicts common

### Budget Reality
- **Semester budget**: $500 (from advisor's NIH grant)
- **Advisor's rules**:
  - Must get approval for instances > $1/hour
  - Weekly cost reports required
  - Auto-termination of idle instances (after previous student left p3.8xlarge running for weekend = $800)
- **Pressure**: Can't ask for more money; must make $500 last 4 months
- **Consequence of overspending**: Lost cloud access, back to shared lab GPU

---

## Current Situation (Before Lens)

### Existing Setup

**Primary option**: Lab GPU server
- NVIDIA RTX 3090 (24GB GPU RAM)
- Shared among 7 researchers
- Scheduling via Slack ("Who's using the GPU tonight?")
- No scheduling system = conflicts common

**Typical scenario**:
- Alex schedules GPU for Tuesday night
- Starts training run at 6 PM
- Another student "just needs to run quick test" at 8 PM → kills Alex's job
- Argument in lab Slack channel
- Advisor intervenes: "We need a better solution"

**Backup option**: Google Colab
- Free tier: Frequent disconnections after 90 minutes
- Pro tier ($10/month): Still limited to 12-hour sessions
- No persistent storage
- Can't install custom CUDA libraries

### Pain Points

1. **GPU access bottleneck**: 7 researchers sharing 1 GPU = frequent conflicts
2. **No long training runs**: Lab GPU monopolized if training takes > 4 hours
3. **Colab limitations**: Disconnects during critical experiments
4. **Advisor anxiety**: Previous student's $800 mistake made advisor nervous about cloud
5. **Budget pressure**: $500 must last entire semester (4 months = $125/month)

### Previous Cloud Attempt (October 2024)

**Day 1**: Alex got advisor approval to try AWS
- Launched p3.2xlarge (V100 GPU, $3.06/hour)
- Trained model for 4 hours = $12.24
- Forgot to terminate over weekend
- Monday: $220 bill (72 hours)
- **Advisor cut cloud access**: "You can use the lab GPU"

**Result**: Back to scheduling conflicts, no cloud access for 3 months

---

## Lens Workflow

### Getting Advisor Buy-In

**Meeting with Advisor** (Monday, 10 AM)

**Alex**: "I found a tool called Lens that has auto-stop built-in. It prevents the runaway cost problem."

**Advisor**: "How do I know you won't forget again?"

**Alex**: "It automatically stops the instance after 2 hours idle. Even if I forget, max cost is maybe $10 instead of $220."

**Advisor**: "Show me a cost estimate for one month."

**Alex launches Lens wizard**:

```bash
lens-jupyter
```

**Wizard shows**:
```
💰 Cost Estimate:
   Instance: g4dn.xlarge (GPU)
   Hourly: $0.526/hr
   Your typical training run (6 hours): $3.16
   If you ran 3 training sessions/week:
     - Weekly: $9.48
     - Monthly: $37.92

   With auto-stop after 2h idle: ~$38/month
   Your semester budget: $500
   This gives you: 13 months of compute at this rate ✅
```

**Advisor**: "Okay, that's reasonable. But I want weekly cost reports."

**Alex**: "Lens has a `costs` command. I can send you weekly reports."

**Advisor**: ✅ Approved

---

## Daily Workflow

### Scenario: Training Cancer Detection Model

**Tuesday, 2 PM** - Alex needs to train a new model architecture

#### Step 1: Launch GPU Instance

```bash
lens-jupyter quickstart --gpu --env deep-learning
```

**Output** (30 seconds):
```
🚀 Launching GPU environment (g4dn.xlarge)...

✓ Your environment is ready!
🌐 Jupyter Lab: http://3.87.22.145:8888
🎮 GPU: NVIDIA T4 (16GB) detected and ready

💰 Cost reminder:
   $0.526/hour = $3.16 for 6-hour training
   Auto-stop after 2h idle (prevents runaway costs)
```

#### Step 2: Upload Training Code

Alex uses Jupyter Lab:
- Upload training script (`train_cancer_detector.py`)
- Upload dataset (3GB medical images from S3)
- Configure hyperparameters in notebook

#### Step 3: Start Training

```python
# In Jupyter notebook
import torch
import torchvision
from train_cancer_detector import CancerNet, train_model

# Verify GPU available
assert torch.cuda.is_available(), "GPU not detected!"
print(f"Using GPU: {torch.cuda.get_device_name(0)}")  # NVIDIA T4

# Load data
train_loader = get_data_loader('medical_images/', batch_size=32)

# Train model
model = CancerNet().cuda()
train_model(model, train_loader, epochs=50, lr=0.001)
# Training starts... Alex can walk away
```

**Training progress visible in Jupyter**:
```
Epoch 1/50: 100%|██████████| Loss: 0.523, Acc: 78.2%
Epoch 2/50: 100%|██████████| Loss: 0.412, Acc: 82.1%
...
```

#### Step 4: Do Other Work While Training (5 hours)

Alex leaves instance running and:
- Attends class (2 hours)
- Writes dissertation chapter (2 hours)
- Lab meeting (1 hour)

**Meanwhile**: Instance keeps training uninterrupted

#### Step 5: Check Results (Evening)

Alex returns at 7 PM, opens Jupyter:

```python
# Training complete!
# Epoch 50/50: Loss: 0.089, Acc: 96.7%
# Model saved to: models/cancer_net_v3.pth
```

Downloads trained model and visualizations.

#### Step 6: Stop Instance

```bash
lens-jupyter stop
```

**Output**:
```
⏸️  Stopping GPU environment...

✓ Instance stopped

💰 Session Summary:
   Runtime: 5.2 hours
   Cost: $2.74
   Your semester spending: $12.48 / $500 (2.5%)

📊 Still well under budget!
```

---

## Cost Management: Weekly Report to Advisor

**Friday afternoon** - Alex sends weekly report

```bash
lens-jupyter costs --this-week --format summary
```

**Output**:
```
💰 Weekly Cost Report: Mar 18-24, 2025

📊 Instances Used:
   1. i-0f3a2c8 (deep-learning-gpu) - g4dn.xlarge
      Sessions: 3
      Total runtime: 16.8 hours
      Cost: $8.84

💡 Usage Breakdown:
   Mon: No usage
   Tue: 5.2h → $2.74
   Wed: 6.1h → $3.21
   Thu: 5.5h → $2.89
   Fri: No usage

📈 Month-to-Date (March):
   Week 1: $9.12
   Week 2: $8.84
   Week 3: $11.20
   Week 4: [in progress]
   ──────────────
   Total: $29.16

🎯 Budget Status:
   Semester budget: $500
   Spent (6 weeks): $58.45
   Remaining: $441.55
   Burn rate: $9.74/week
   Projected semester total: $155.84 (69% under budget!)

✅ On track!
```

**Alex emails to advisor**: "Week 3 report attached. $8.84 this week, $58 total for semester (12% of budget). Three successful training runs."

**Advisor's response**: "Great, keep it up."

---

## Pain Points & Solutions

### Pain #1: GPU Access Conflicts

**Before Lens**:
- 7 researchers sharing 1 lab GPU
- Scheduling via Slack (informal, conflicts common)
- Training interrupted by other students
- Can't run overnight (someone else scheduled)

**With Lens**:
- Dedicated GPU when needed (g4dn.xlarge)
- No scheduling conflicts
- Can run overnight
- Stop when done, cost only $0.53/hour

**Success Metric**:
- ✅ Zero scheduling conflicts (dedicated resource)
- ✅ 100% training completion (no interruptions)
- ✅ Overnight runs possible
- ✅ 3x more experiments per week (limited by budget, not access)

**Related GitHub Issues**: #28 (GPU support)
**Related Requirements**: REQ-7.3 (GPU Support)

---

### Pain #2: Advisor Budget Anxiety

**Before Lens**:
- Previous student's $800 mistake → advisor nervous
- Alex had cloud access revoked
- Required pre-approval for every launch (slow)
- Advisor micromanaged spending

**With Lens**:
- Auto-stop prevents runaway costs
- Weekly reports build trust
- Spending predictable ($9-10/week)
- Advisor gave Alex autonomy back

**Success Metric**:
- ✅ Zero cost overruns (auto-stop works 100% of time)
- ✅ Advisor trust restored (weekly reports show control)
- ✅ 90% under budget (using $155 of $500)
- ✅ Autonomy: Alex doesn't need approval for each launch

**Related GitHub Issues**: #19 (Budget alerts), #21 (Cost reports)
**Related Requirements**: REQ-2.2 (Auto-Stop), REQ-2.4 (Cost Reporting)

---

### Pain #3: Google Colab Limitations

**Before Lens**:
- Colab free tier: Disconnects after 90 minutes
- Colab Pro: Max 12-hour sessions
- Lost training progress during disconnection
- Can't install custom CUDA libraries

**With Lens**:
- Sessions persist until Alex explicitly stops
- Full control over CUDA environment
- Can install any library
- Persistent storage (EBS)

**Success Metric**:
- ✅ Zero training interruptions (vs 40% on Colab)
- ✅ Custom environment with specific CUDA/PyTorch versions
- ✅ Training runs 24+ hours if needed
- ✅ Full root access for system packages

**Related GitHub Issues**: #12 (Environment export/import)
**Related Requirements**: REQ-4.1 (Reproducibility)

---

### Pain #4: Long Training Times Block Laptop

**Before Lens**:
- Training on laptop GPU (GTX 1650): 24 hours
- Laptop unusable during training
- Can't take laptop to class/meetings
- Must choose: train model OR use laptop

**With Lens**:
- Training on g4dn.xlarge: 5 hours
- Laptop free for other work
- Can close laptop, training continues
- Total freedom

**Success Metric**:
- ✅ 80% faster training (24h → 5h)
- ✅ Laptop always available
- ✅ Can work from anywhere (training in cloud)
- ✅ 3x more experiments (faster iteration)

**Related GitHub Issues**: #4 (Email notifications when training done)
**Related Requirements**: REQ-7.3 (GPU Performance)

---

### Pain #5: Reproducibility for Paper Submissions

**Before Lens**:
- Reviewer: "Can you share training environment?"
- Alex: "I used Colab... not sure exact package versions"
- Reviewer: "Results don't reproduce"
- Paper rejected

**With Lens**:
- Alex exports environment:
  ```bash
  lens-jupyter env export > cancer-detection-env.yaml
  ```
- Includes in paper supplementary materials
- Reviewers can launch identical environment
- Results reproduce perfectly

**Success Metric**:
- ✅ 100% reproducibility (identical packages)
- ✅ Reviewer satisfaction (no reproduction issues)
- ✅ Faster paper acceptance
- ✅ Can recreate environment 5 years later

**Related GitHub Issues**: #12 (Env export), #14 (Community templates)
**Related Requirements**: REQ-4.1 (Environment Export)

---

## Success Metrics: One Semester Later

### Research Productivity

| Metric | Before Lens | After Lens (4 months) | Improvement |
|--------|----------------|--------------------------|-------------|
| Training runs/week | 2-3 (limited by GPU access) | 6-8 (limited by time, not access) | 3x more |
| Papers submitted | 1 | 3 | 3x more |
| Experiments tried | 12 per month | 40 per month | 3.3x more |
| Training interruptions | 40% (Colab disconnects + conflicts) | 0% | 100% reliable |

### Cost Management

| Month | Sessions | Total Hours | Cost | Budget % |
|-------|----------|-------------|------|----------|
| Jan | 12 | 68.5h | $36.03 | 7.2% |
| Feb | 14 | 78.2h | $41.13 | 8.2% |
| Mar | 16 | 89.1h | $46.87 | 9.4% |
| Apr | 14 | 73.9h | $38.87 | 7.8% |
| **Total** | **56** | **309.7h** | **$162.90** | **32.6%** |

**Budget performance**:
- Allocated: $500
- Spent: $163
- Remaining: $337 (67% under budget!)
- Advisor's reaction: "This is exactly what I wanted - predictable, controlled spending."

### Advisor Satisfaction

**Advisor's feedback**:
> "Alex's use of Lens is a model for the lab. Weekly cost reports give me visibility, auto-stop prevents mistakes, and spending is 67% under budget. Three other students in my lab are now using it. The $163 cost is negligible compared to Alex's 3x productivity increase. We're making cloud computing work within academic budgets."

### Qualitative Impact

**Alex's quote**:
> "Lens gave me GPU access without the guilt. Before, I felt bad asking for the lab GPU because I knew others needed it too. And Colab was frustrating with constant disconnections. Now I can experiment freely, train overnight, and my advisor trusts me because costs are under control. I've submitted 3 papers this semester - I was planning on 1. I'm actually ahead of my dissertation timeline now."

---

## Lessons Learned

### What Works Well

1. **Auto-stop eliminates trust issues**: Advisor comfortable giving autonomy
2. **Weekly cost reports**: Builds trust, shows responsible usage
3. **GPU instances**: g4dn.xlarge perfect for PhD research (not too expensive, not too slow)
4. **Cost estimates**: Advisor could evaluate ROI before approving
5. **Quickstart command**: Alex uses `lens-jupyter quickstart --gpu` for quick launches

### Feature Requests

1. **Email notification when training completes** (#4)
   - Current: Alex checks Jupyter manually
   - Desired: Email when training finishes (can walk away)

2. **S3 integration for datasets** (#17)
   - Current: Re-upload 3GB dataset each launch (15 minutes)
   - Desired: Auto-sync from S3 (1 minute)

3. **Spot instances** (#23)
   - Current: On-demand g4dn.xlarge = $0.526/hr
   - Desired: Spot = $0.158/hr (70% savings, can train 3x more)

4. **Budget alerts** (#19)
   - Current: Alex manually checks costs
   - Desired: Email at 50% of semester budget

5. **Team workspace** (#16)
   - Desired: Share environment with labmates
   - Use case: "I got this working, try my exact setup"

---

## Technical Details

### Typical Configuration

**Instance**: g4dn.xlarge
- GPU: NVIDIA T4 (16GB VRAM)
- CPU: 4 vCPUs
- RAM: 16 GB
- Cost: $0.526/hour
- Training: 5-6 hours typical

**Environment**: deep-learning-gpu
- CUDA 11.8
- PyTorch 2.0.1 with GPU support
- TensorFlow 2.13
- Jupyter Lab
- Medical imaging libraries: SimpleITK, nibabel, pydicom

**Storage**: 100GB EBS
- Models: 500MB-2GB each
- Datasets: 3-10GB temporary
- Cost: $10/month

**Monthly cost breakdown**:
- Compute: 80h × $0.526/hr = $42
- Storage: $10
- Data transfer: ~$1
- **Total**: ~$53/month average

### Usage Pattern

- **Sessions per week**: 4
- **Hours per session**: 5-6
- **Weekly runtime**: 20-24 hours
- **Monthly cost**: $40-45
- **Semester cost**: $160-180 (32-36% of $500 budget)

**Efficiency**: Under budget by 64-68%, allowing flexibility for larger experiments

---

## Conclusion

Lens solved Alex's GPU access problem while staying within advisor's strict budget constraints:

### Key Outcomes
- ✅ **3x research productivity** (3 papers vs 1 per semester)
- ✅ **Zero scheduling conflicts** (dedicated GPU access)
- ✅ **67% under budget** ($163 used of $500 allocated)
- ✅ **Advisor trust restored** (weekly reports + auto-stop)
- ✅ **100% training reliability** (no interruptions)

### Success Factors
1. **Auto-stop**: Eliminated advisor's fear of cost overruns
2. **Cost reporting**: Weekly reports built trust
3. **GPU access**: g4dn.xlarge perfect for deep learning research
4. **Budget control**: Stayed way under limits
5. **Autonomy**: No longer needed approval for each launch

**ROI**: $163 spent → 3 papers submitted → on track for timely PhD graduation

Alex represents the **graduate student persona**: technically capable but budget-constrained, needing advisor trust and GPU access to complete dissertation. Lens provides both.

**Related GitHub Issues**: #4, #12, #16, #17, #19, #21, #23, #28
