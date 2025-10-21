# Course Instructor Walkthrough: Prof. David Lee

> **Persona**: Data Science instructor teaching "Introduction to Data Science" (CS 301)
> **Technical Level**: 4/5 (Strong technical background, limited time during semester)
> **Budget**: $2,000/semester for 30 students from department teaching budget
> **Primary Pain Point**: "Works on my machine" problems with 30 students on different OS, package conflicts, unreliable free-tier tools

---

## Profile

**Name**: Prof. David Lee
**Position**: Assistant Professor, Computer Science Department
**Institution**: Public state university (teaching-focused, 3-2 teaching load)
**Courses**: Data Science (CS 301), Machine Learning (CS 402), Senior Capstone (CS 495)
**Age**: 38
**Location**: Seattle, Washington (uses us-west-2)

### Teaching Context

**Course**: CS 301 - Introduction to Data Science (Spring 2025)
- **Enrollment**: 30 students (mix of CS, Statistics, Business majors)
- **Level**: Undergraduate junior/senior level
- **Prerequisites**: Programming (Python), Statistics basics
- **Format**: 2× 75-minute lectures/week + 1× 50-minute lab session

**Student demographics**:
- 40% CS majors (comfortable with command line, Git)
- 30% Statistics majors (strong R skills, weak Python/command line)
- 20% Business majors (minimal programming, Excel-heavy background)
- 10% Other (Biology, Economics, etc. - varying technical skills)

**Operating system diversity**:
- 50% MacOS (M1/M2 Macs common = ARM64 architecture issues)
- 40% Windows (package installation nightmares)
- 10% Linux (usually work fine, but obscure distros cause problems)

### Teaching Philosophy & Constraints

**Philosophy**: "Students should spend time learning data science, not fighting with package installations"

**Time constraints**:
- 6 hours/week teaching (lectures + lab)
- 4 hours/week prep (lesson plans, assignments)
- 3 hours/week office hours (mostly troubleshooting student setups)
- 2 hours/week grading
- **Total**: 15 hours/week on this one course
- **Available for infrastructure**: 0 hours (wants to eliminate this)

**Budget reality**:
- Department allocation: $2,000/semester for cloud computing
- **Per student**: $2,000 ÷ 30 students = $66.67/student
- **Per week**: $2,000 ÷ 15 weeks = $133.33/week
- **Must stay under budget**: No flexibility (department won't approve overages)

---

## Current Situation (Before AWS IDE)

### Previous Teaching Approach (Fall 2024)

#### Week 1: Installation Hell

**Lecture 1 (Tuesday)**: Course introduction, Python basics review
**Homework (due Thursday)**: "Install Anaconda, Jupyter, pandas, matplotlib"

**Thursday lab session**: Instead of teaching, spent entire 50 minutes troubleshooting:

**Installation problems observed** (27 of 30 students had issues):

1. **MacOS M1/M2 students (15 students)**:
   - Rosetta translation issues with Intel-compiled packages
   - TensorFlow doesn't support ARM64 natively → weird workarounds
   - 5 students: 2+ hours troubleshooting, still not working

2. **Windows students (12 students)**:
   - Anaconda PATH not set correctly → "jupyter: command not found"
   - Windows Defender blocking conda.exe → permission errors
   - One student's antivirus deleted Python.exe thinking it was malware

3. **Linux students (3 students)**:
   - Mostly okay, but one student on Arch Linux had dependency conflicts
   - Another on Ubuntu 18.04 (ancient) couldn't install latest Jupyter

**Prof. Lee's experience**:
- Week 1 lecture: 10 minutes content, 65 minutes troubleshooting
- Week 1 lab: 0 minutes content, 50 minutes troubleshooting
- Office hours: 20+ students came, queue out the door
- **Total time lost**: 8 hours troubleshooting installations
- **Students frustrated**: "Why is this so hard? I thought we'd be doing data science..."

#### Week 2-4: Continued Problems

**Ongoing issues**:
- "My code works on my laptop but not on my friend's laptop"
- "Professor, I updated Anaconda and now nothing works"
- "I accidentally deleted a package and broke everything"
- Every assignment: 5-10 students can't submit because "environment broken"

**Real incident (Week 3)**:
- Assignment: Analyze dataset with pandas + matplotlib
- Student Sarah: "My matplotlib plots are blank - nothing shows up"
- Prof. Lee debugging: 45 minutes later, discovered Sarah had matplotlib 2.1 (ancient) vs class using 3.8
- Solution: Reinstall Anaconda → 30 minutes → other packages broke
- **Time wasted**: 75 minutes for one student

**Office hours transformation**:
- Intended: Help students understand data science concepts
- Reality: 80% of time spent on environment troubleshooting
- Prof. Lee: "I'm an IT support technician, not a professor"

#### Week 5: Google Colab Experiment

**Prof. Lee's decision**: "Enough. We're switching to Google Colab."

**Tuesday lecture**:
- "Everyone, we're moving to Google Colab. Go to colab.google.com"
- 30 students open browsers
- "Create a new notebook and run: `import pandas`"
- **Success!** All 30 students working

**Prof. Lee's relief**: "Finally, I can teach data science!"

#### Week 6-12: Colab Limitations Emerge

**Problems discovered**:

1. **Frequent disconnections** (every 90 minutes on free tier):
   - Student working on assignment for 2 hours
   - Connection drops → lost all variables/data in memory
   - Student restarts kernel, re-runs all cells (10 minutes)
   - Frustration builds

2. **Cannot install custom packages**:
   - Assignment: Use specialized NLP library (spaCy language models)
   - Colab: Can install, but takes 5 minutes every session (large download)
   - 30 students × 5 minutes = 2.5 hours of class time waiting

3. **No persistent storage**:
   - Students upload datasets every session
   - Large dataset (500MB) → 10 minutes to upload
   - Or: use Google Drive (confusing for students, permission issues)

4. **Limited resources**:
   - Final project: Train simple neural network
   - Colab free tier: Throttled GPU access, only 12 hours/day
   - Students: "I'm trying to train my model but Colab says I've used up my GPU quota"
   - Prof. Lee: "Well, you'll have to wait until tomorrow..."

5. **Inconsistent environment**:
   - Colab auto-updates packages weekly
   - Week 10: Colab updated pandas 1.5 → 2.0 (breaking API changes)
   - Student assignments from Week 5 stopped working
   - Prof. Lee spent 2 hours updating all course materials

**Prof. Lee's frustration**:
> "Colab solved the installation problem but created new problems. Disconnections during assignments, GPU quota limits, inconsistent environments. And I can't give students a challenging final project because Colab is too limited. There has to be a better way."

### Pain Points Summary (Fall 2024)

| Pain Point | Impact | Time Lost | Student Impact |
|-----------|--------|-----------|----------------|
| Installation problems (Week 1) | Can't start teaching | 8 hours | 27/30 students blocked |
| Ongoing environment issues | Every assignment delayed | 4 hours/week × 15 weeks = 60 hours | 5-10 students/week stuck |
| Colab disconnections | Frustration, lost work | 2 hours/week TA support | Daily complaints |
| Colab limitations | Can't assign challenging projects | Watered-down curriculum | Reduced learning outcomes |
| Inconsistent environments | "Works on my machine" continues | 3 hours/week office hours | Collaboration difficult |
| No reproducibility | Can't verify student work | 10 hours grading (running code manually) | Academic integrity concerns |

**Total impact**:
- **Prof. Lee's time**: 8 hours (Week 1) + 60 hours (ongoing) + 45 hours (Colab issues) = **113 hours/semester**
- **Teaching quality**: Reduced (can't assign realistic projects)
- **Student satisfaction**: Mixed (better than installation hell, but still frustrated)

---

## AWS IDE Workflow (Spring 2025)

### Pre-Semester Preparation

#### December 2024: Prof. Lee Discovers AWS IDE

**Conference talk**: Hears Prof. Chen present on using AWS IDE for teaching
**Intrigued**: "Auto-stop would solve budget concerns. Identical environments would solve 'works on my machine'."

**Evaluation** (Christmas break):

```bash
# Prof. Lee tests on personal MacBook
brew install aws-jupyter

# Launches with wizard
aws-jupyter

# Estimates costs
# 30 students × 3 hours/week × 15 weeks × $0.05/hour = $675
# Under $2,000 budget ✓

# Configures AWS account
aws configure
```

**Proof of concept**:
- Creates test environment with course packages
- Launches, runs sample notebook
- Verifies auto-stop works
- **Decision**: "Let's do it."

#### January 2025: Course Environment Design

**Prof. Lee creates course environment**:

```bash
# Create course-specific environment
cat > cs301-data-science-spring2025.yaml <<EOF
name: cs301-data-science-spring2025
description: CS 301 Introduction to Data Science - Spring 2025
instructor: Prof. David Lee (david.lee@university.edu)

packages:
  system:
    - git
    - vim

  python:
    version: "3.11"
    packages:
      # Core data science
      - numpy==1.26.4
      - pandas==2.1.4
      - matplotlib==3.8.2
      - seaborn==0.13.0
      - scikit-learn==1.4.0

      # Jupyter
      - jupyter==1.0.0
      - ipywidgets==8.1.1

      # Additional libraries (used in later assignments)
      - requests==2.31.0
      - beautifulsoup4==4.12.2
      - sqlalchemy==2.0.25
      - plotly==5.18.0

      # Final project (ML)
      - tensorflow==2.15.0
      - keras==2.15.0

jupyter_extensions:
  - jupyterlab-git
  - jupyterlab-plotly

notes: |
  This environment is tested and frozen for Spring 2025 semester.
  All students will use identical package versions to ensure
  reproducibility and eliminate "works on my machine" problems.
EOF
```

**Key decisions**:
1. **Freeze package versions**: No mid-semester updates (learned from Colab pain)
2. **ARM64-compatible**: Works on M1/M2 Macs (t4g instances)
3. **Include all semester packages**: Pre-install everything students will need
4. **Cost-optimize**: t4g.medium ($0.0336/hr) sufficient for course work

#### Test Run (2 weeks before semester)

**Prof. Lee invites 3 TAs to test**:

```bash
# TAs each launch environment
aws-jupyter launch --env cs301-data-science-spring2025.yaml --instance-type t4g.medium

# Output (2 minutes later):
✓ Your environment is ready!
🌐 Jupyter Lab: http://44.234.12.78:8888

# TAs test all course notebooks (Week 1-5 content)
# Result: Everything works, zero issues
```

**TA feedback**:
- "Way easier than setting up Anaconda"
- "I'm on M1 Mac - usually have package problems, but this just worked"
- "The auto-stop is clever - I forgot to stop and it only cost $0.07"

**Prof. Lee's confidence**: ✅ Ready for class

---

### Semester Begins: Week 1

#### Tuesday Lecture (Jan 14, 2025)

**Prof. Lee's introduction**:
> "Welcome to CS 301! This semester, we're using AWS IDE - cloud-based Jupyter Lab. Why cloud?
> 1. **No installation headaches**: No Anaconda, no 'works on my machine' problems
> 2. **Identical environments**: Everyone has the exact same packages
> 3. **More powerful**: Cloud instances faster than most laptops
> 4. **Real-world skills**: Industry uses cloud, you should learn it
>
> Setup takes 15 minutes. We'll do it together in Thursday's lab."

**Student reactions**:
- CS majors: "Cool, I've heard of AWS"
- Stats majors: "Is this hard?"
- Business majors: "What's cloud? What's AWS?"
- Prof. Lee: "Don't worry, there's a wizard that guides you. You'll be fine."

#### Thursday Lab (Jan 16, 2025): Group Setup Session

**Prof. Lee's plan**: 50-minute lab session, get all 30 students working

**9:00 AM - Lab starts**

**Step 1** (5 minutes): "Everyone go to github.com/university-cs/cs301-spring2025"
- Repository has installation instructions for Mac, Windows, Linux
- Students install AWS IDE (Homebrew, Scoop, apt)

**Step 2** (10 minutes): "Now we set up AWS credentials"
- Prof. Lee projects screen
- Shows: `aws configure`
- **Problem**: Students don't have AWS accounts yet
- **Solution** (prepared): Prof. Lee created 30 AWS IAM users in department account
  - Each student gets email with credentials
  - Students configure with their individual credentials

**Step 3** (15 minutes): "Launch your first environment"

```bash
# Prof. Lee projects command
aws-jupyter launch --env cs301-data-science-spring2025.yaml
```

**Real-time results**:
- 3 minutes: 28 students have Jupyter Lab open
- 2 students stuck:
  - Student 1: Forgot to run `aws configure` → TA helps, fixed in 2 minutes
  - Student 2: AWS credentials typo → re-ran `aws configure`, fixed in 1 minute
- **9:18 AM**: All 30 students have Jupyter Lab running

**Prof. Lee's internal reaction**: "18 minutes to get 30 students working. Last semester, Week 1 took 8 hours. This is incredible."

**Step 4** (10 minutes): "Let's verify everything works"

Prof. Lee distributes test notebook (week01-test.ipynb):

```python
# week01-test.ipynb
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

print("Success! Your environment works.")
print(f"Python version: {import sys; sys.version}")
print(f"Pandas version: {pd.__version__}")
print(f"NumPy version: {np.__version__}")

# Create simple plot
x = np.linspace(0, 10, 100)
y = np.sin(x)
plt.plot(x, y)
plt.title("If you see this plot, you're all set!")
plt.show()
```

**Results**:
- All 30 students: "Success! Your environment works."
- All 30 students see the sine wave plot
- **Zero issues**

**Prof. Lee's closing** (5 minutes):
> "Great! You're all set up. A few important things:
> - Your instance will auto-stop after 2 hours idle. This is good - saves money.
> - When you're done working, run: `aws-jupyter stop`
> - To restart next time: `aws-jupyter start`
> - Check your costs anytime: `aws-jupyter costs`
>
> Next Tuesday, we dive into pandas DataFrames. See you then!"

**9:50 AM - Lab ends**

**Prof. Lee's reflection**:
> "We spent 18 minutes on setup and had 32 minutes for actual content. Last semester, I spent entire 50 minutes troubleshooting installations and still had 3 students not working. This is a game-changer."

---

### Week 2-8: Teaching Without Infrastructure Friction

#### Typical Week

**Tuesday lecture**:
- 0 minutes: Environment troubleshooting
- 75 minutes: Actual data science content
- **No office hours queue** for "my setup is broken"

**Thursday lab**:
- Students arrive, restart instances: `aws-jupyter start` (30 seconds)
- 50 minutes: Hands-on exercises
- End of lab: "Remember to stop your instances!"

**Prof. Lee's weekly routine**:
- Monday: Prepare lecture materials
- Tuesday: Teach lecture
- Wednesday: Office hours (concepts only, no troubleshooting)
- Thursday: Teach lab
- Friday: Grade assignments (by running student code in own AWS IDE instance)

**Time spent on infrastructure**: ~0 hours/week (vs 4 hours/week previously)

#### Week 4 Incident: A Student Forgets to Stop

**Monday morning**: Student Email
> "Prof. Lee, I forgot to stop my instance Friday. It ran all weekend. Am I in trouble?"

**Prof. Lee checks**:
- Instance auto-stopped Saturday at 11:15 AM (2 hours after last activity)
- Runtime: ~17 hours total (Friday 3 PM - Saturday 11 AM)
- Cost: 17 hours × $0.0336/hr = **$0.57**

**Prof. Lee's response**:
> "No worries! Auto-stop kicked in, so it only ran 17 hours instead of the full weekend (72 hours). Your cost was $0.57. This is exactly why we use auto-stop - mistakes happen, and this limits the damage. Just try to remember to stop manually when you're done to save even more. Thanks for letting me know!"

**Comparison to nightmare scenario**:
- If no auto-stop: 72 hours × $0.0336/hr = $2.42
- If student used expensive instance (t3.xlarge): 72 hours × $0.1664/hr = $11.98
- **With AWS IDE auto-stop**: $0.57 (95% cost avoidance)

**Prof. Lee's takeaway**: "Auto-stop just saved a student from a $12 mistake. This is protecting my budget."

---

### Week 10: Assignment - Web Scraping with Beautiful Soup

**Assignment**: Scrape a website, analyze data, visualize trends

**Student workflow**:

```bash
# Monday: Student starts assignment
aws-jupyter start  # 30 seconds

# Open Jupyter, work for 2 hours
# Write scraping script, test on website
# Download data, clean with pandas

# Student goes to dinner (forgets to stop)
# Auto-stop triggers at 8:30 PM (2 hours idle)
```

**Wednesday**: Student continues

```bash
# Restart instance
aws-jupyter start  # 30 seconds

# Continue working (2 hours)
# Finish analysis, create visualizations
# Download final notebook

# Stop manually
aws-jupyter stop
```

**Key points**:
- **Zero "works on my machine"**: All 30 students have identical environments
- **No package installation friction**: Beautiful Soup pre-installed
- **Reproducibility**: Prof. Lee can run any student's notebook in his own AWS IDE instance

**Grading experience** (Friday):

Prof. Lee launches his instance with course environment:

```bash
aws-jupyter launch --env cs301-data-science-spring2025.yaml
```

- Uploads first student's notebook
- Runs all cells → works perfectly
- Grades, repeats for all 30 students
- **Zero "I can't run this code" issues**

**Grading time**: 3 hours (vs 5 hours previously, with 10 hours troubleshooting last semester)

---

### Week 15: Final Project - ML Model Training

**Assignment**: Train a neural network for image classification

**Previous semester (Colab)**:
- Students hit GPU quota limits
- Training took 6 hours on CPU → many students couldn't complete
- Prof. Lee watered down project to fit Colab constraints

**This semester (AWS IDE)**:

Students use GPU instances:

```bash
aws-jupyter launch --instance-type g4dn.xlarge --env cs301-data-science-spring2025.yaml
```

**Configuration**:
- GPU: NVIDIA T4 (Colab-equivalent hardware)
- Cost: $0.526/hour
- Student budget impact: 6-hour training = $3.16 (within $66 budget)

**Results**:
- All 30 students successfully trained models
- No quota limits, no throttling
- Training time: 1.5 hours on GPU (vs 6 hours on CPU)
- **Success rate**: 100% (vs 60% last semester with Colab limitations)

**Student feedback**:
- "Way faster than I expected"
- "Didn't have any connectivity issues like Colab"
- "The auto-stop saved me - I went to get coffee and forgot about it, but it stopped after 2 hours"

---

## Cost Management: Semester Budget Tracking

### Monthly Spending (30 Students)

**January (4 weeks)**:

| Week | Usage Pattern | Hours/Student | Cost/Student | Total Cost |
|------|--------------|---------------|--------------|------------|
| 1 | Setup + test | 1.5 | $0.05 | $1.50 (30 students) |
| 2 | Assignment 1 | 3 | $0.10 | $3.00 |
| 3 | Assignment 2 | 3 | $0.10 | $3.00 |
| 4 | Midterm prep | 4 | $0.13 | $3.90 |
| **Jan Total** | - | **11.5** | **$0.39** | **$11.70** |

**February-April (similar pattern)**: ~$45/month

**May (finals, 2 weeks)**:

| Week | Usage Pattern | Hours/Student | Cost/Student | Total Cost |
|------|--------------|---------------|--------------|------------|
| 14 | Final project start | 6 | $0.20 | $6.00 |
| 15 | Final project (GPU) | 6 (GPU) | $3.16 | $94.80 |
| **May Total** | - | **12** | **$3.36** | **$100.80** |

### Semester Total (15 Weeks)

| Month | Total Cost | Budget % | Notes |
|-------|-----------|----------|-------|
| Jan | $11.70 | 0.6% | Setup + early assignments |
| Feb | $42.50 | 2.1% | Regular assignments |
| Mar | $48.30 | 2.4% | Midterm + assignments |
| Apr | $45.80 | 2.3% | Regular assignments |
| May | $100.80 | 5.0% | Final project (GPU usage spike) |
| **TOTAL** | **$249.10** | **12.5%** | Under budget by 87.5%! |

**Budget Performance**:
- Allocated: $2,000
- Spent: $249.10
- Under budget: $1,750.90 (87.5% savings!)
- **Per student**: $8.30 (vs $66.67 budgeted)

**Prof. Lee's reaction**:
> "We used 12.5% of the budget. I was conservative in my estimate, but this is incredible. The department gave me $2,000 and I'm returning $1,750. They're going to think I didn't use cloud computing at all. Meanwhile, students had a better experience than last year's Colab + Anaconda chaos."

---

## Pain Points & Solutions

### Pain #1: Installation Hell (Week 1 Chaos)

**Before AWS IDE**:
- Week 1: 8 hours troubleshooting 27/30 students' installations
- MacOS M1/M2: ARM64 package conflicts
- Windows: PATH issues, antivirus interference
- Ongoing: 60 hours/semester helping students fix broken environments

**With AWS IDE**:
- Week 1: 18 minutes to get 30/30 students working
- Zero OS-specific issues (cloud instances are Linux-based)
- Ongoing: 0 hours environment troubleshooting

**Success Metric**:
- ✅ 96% time reduction in setup (8 hours → 18 minutes)
- ✅ 100% student success rate (30/30 vs 27/30)
- ✅ Zero ongoing environment issues (60 hours → 0 hours saved)

**Related GitHub Issues**: #1 (Wizard default), #2 (Quickstart command)
**Related Requirements**: REQ-1.1 (Beginner-Friendly Onboarding)

---

### Pain #2: "Works on My Machine" Syndrome

**Before AWS IDE**:
- Every assignment: 5-10 students couldn't run example code
- Different Anaconda versions, package conflicts
- Grading nightmare: Prof. Lee couldn't run student code 30% of the time
- Collaboration impossible (students had different environments)

**With AWS IDE**:
- All 30 students have byte-for-byte identical environments
- Example code runs on everyone's machine
- Grading seamless: Prof. Lee launches same environment, runs all notebooks successfully
- Students can collaborate (share notebooks, they just work)

**Real example** (Week 7 group project):
- 3 students working together on data pipeline
- Student A writes pandas code → Student B adds visualization → Student C adds analysis
- All working in identical environments
- **Zero integration problems**

**Success Metric**:
- ✅ 100% code compatibility (vs 70% before)
- ✅ Zero grading issues (vs 30% code wouldn't run)
- ✅ Collaboration seamless (group projects actually work)

**Related GitHub Issues**: #12 (Environment export/import)
**Related Requirements**: REQ-4.1 (Environment Reproducibility)

---

### Pain #3: Colab Limitations and Disconnections

**Before AWS IDE (using Colab)**:
- Disconnections every 90 minutes (lost work, frustration)
- GPU quota limits (final projects limited)
- Inconsistent environment (auto-updates broke assignments mid-semester)
- Cannot install custom packages persistently

**With AWS IDE**:
- Sessions persist until manually stopped (students work 4+ hours uninterrupted)
- Dedicated GPU (no quotas, no throttling)
- Frozen environment (package versions locked for semester)
- Full control (students can install anything)

**Real example** (Final project):
- Student training neural network (1.5 hours)
- Colab would have disconnected 2× (lost training progress)
- AWS IDE: Uninterrupted training, model completes
- Cost: $0.79 (1.5 hours × $0.526/hr)

**Success Metric**:
- ✅ Zero disconnections (vs 40% of students hit Colab disconnection issues)
- ✅ 100% final project completion (vs 60% with Colab)
- ✅ More challenging assignments possible (not limited by free-tier constraints)

**Related GitHub Issues**: #4 (Email notifications)
**Related Requirements**: REQ-7.1 (Fast Launch), REQ-7.3 (GPU Support)

---

### Pain #4: Budget Anxiety and Cost Overruns

**Before AWS IDE**:
- No AWS usage (previous semester) due to fear of student cost overruns
- Colab "free" but severe limitations
- Anaconda "free" but 8 hours Week 1 setup = $800 value of Prof. Lee's time

**With AWS IDE**:
- Auto-stop prevents cost overruns
- Spent $249 of $2,000 budget (87.5% under)
- GPU usage possible within budget
- **Peace of mind**: Prof. Lee confident costs won't spike

**Student mistake mitigation**:
- Student forgets to stop → auto-stop after 2 hours → max $0.07 damage
- Worst case scenario (student launches p3.8xlarge and forgets):
  - Without auto-stop: 72 hours × $12.24/hr = $881
  - With auto-stop: 2 hours × $12.24/hr = $24.48
  - **97% damage prevention**

**Success Metric**:
- ✅ 87.5% under budget ($249 vs $2,000)
- ✅ Zero cost overruns (auto-stop works 100% of time)
- ✅ GPU usage affordable ($3.16 per student for final project)
- ✅ Budget predictability (no surprises)

**Related GitHub Issues**: #19 (Budget alerts), #20 (Email before auto-stop)
**Related Requirements**: REQ-2.1 (Cost Preview), REQ-2.2 (Auto-Stop)

---

### Pain #5: Grading Burden and Academic Integrity

**Before AWS IDE**:
- 30% of student code wouldn't run on Prof. Lee's machine (different environments)
- Spent 2-3 hours per assignment troubleshooting why code won't run
- Academic integrity concerns: Did student actually write this, or copy from web? (Can't verify if code doesn't run)
- Grading time: 10 hours/assignment (5 hours running code + 5 hours actual grading)

**With AWS IDE**:
- Prof. Lee launches identical environment
- All student code runs perfectly
- Can execute every student's notebook start-to-finish
- Academic integrity: Can verify code actually produces claimed results

**Grading workflow**:

```bash
# Prof. Lee Friday afternoon
aws-jupyter launch --env cs301-data-science-spring2025.yaml

# Download all 30 student notebooks from LMS
# Run each notebook:
# - Verify outputs match submission
# - Check code quality
# - Grade

# Stop instance
aws-jupyter stop
```

**Grading time**: 3 hours (vs 10 hours previously)

**Success Metric**:
- ✅ 100% student code runs (vs 70% before)
- ✅ 70% grading time reduction (10 hours → 3 hours)
- ✅ Academic integrity verifiable
- ✅ Can provide better feedback (actually ran code, see results)

**Related GitHub Issues**: #12 (Environment export for reproducible assignments)
**Related Requirements**: REQ-4.1 (Reproducibility)

---

## Success Metrics: End of Semester

### Teaching Quality Impact

| Metric | Fall 2024 (Colab/Anaconda) | Spring 2025 (AWS IDE) | Improvement |
|--------|----------------------------|----------------------|-------------|
| Week 1 setup time | 8 hours | 18 minutes | 96% faster |
| Students working after Week 1 | 27/30 (90%) | 30/30 (100%) | +10% |
| Office hours for troubleshooting | 60 hours/semester | 0 hours | 100% elimination |
| Grading time per assignment | 10 hours | 3 hours | 70% reduction |
| Final project completion rate | 18/30 (60%) | 30/30 (100%) | +40% |
| "Works on my machine" issues | Weekly (5-10 students) | Zero | 100% elimination |

### Student Satisfaction

**End-of-semester course evaluation**:

**Question**: "The course computing environment (AWS IDE) was:"

| Response | Count | % |
|----------|-------|---|
| Excellent - worked flawlessly | 23 | 77% |
| Good - minor issues | 6 | 20% |
| Fair - some problems | 1 | 3% |
| Poor - major problems | 0 | 0% |

**Student comments**:
- "Way better than Anaconda hell last semester (I took CS 201). Setup took 15 minutes, everything just worked."
- "I'm on a Windows laptop and usually have package problems. AWS IDE saved me so much frustration."
- "Auto-stop is genius - I forgot to stop twice and it only cost like 10 cents total."
- "Being able to use GPUs for final project was amazing. Would have been impossible on my laptop."
- "Prof. Lee actually taught data science instead of spending time troubleshooting installations."

**Negative feedback** (1 student):
- "I wish we could keep our environments after semester ends. Mine auto-deleted after 30 days."
  - Prof. Lee's note: This is AWS IDE default cleanup, could be configured differently

### Budget Performance

**Final Costs**:
- Budget: $2,000
- Actual: $249.10
- Under budget: $1,750.90 (87.5%)
- Per student: $8.30 (vs $66.67 budgeted)

**Department Chair reaction**:
> "David, you submitted $249 in cloud computing expenses for a 30-student course. Your budget was $2,000. Can you explain?"
>
> **Prof. Lee**: "AWS IDE has auto-stop built-in, so students only pay for time actually used. Plus ARM-based instances are cheaper. We accomplished more than last semester (GPU-enabled final projects) at 12% of budget."
>
> **Chair**: "This is impressive. I'm approving your request to use AWS IDE for CS 402 (Machine Learning) next fall. That course needs GPUs heavily - let's budget $3,000 and see where you land."

### Prof. Lee's Time Savings

| Activity | Time Before (hours/semester) | Time After (hours/semester) | Savings |
|----------|----------------------------|----------------------------|---------|
| Week 1 setup | 8 | 0.3 | 7.7 |
| Ongoing troubleshooting | 60 | 0 | 60 |
| Office hours (environment issues) | 45 | 0 | 45 |
| Grading (running student code) | 70 | 25 | 45 |
| **TOTAL** | **183 hours** | **25.3 hours** | **157.7 hours saved** |

**ROI Analysis**:
- Time saved: 157.7 hours
- Value @ $50/hour (teaching rate): **$7,885**
- AWS IDE cost: $249
- **Net value**: $7,885 - $249 = **$7,636 (3,066% ROI)**

**Prof. Lee's quality of life**:
- More time for research (2 papers published this semester vs 1 last semester)
- More time for advising (took on 2 new undergraduate research students)
- Less stress (no more IT support role)
- Better teaching evaluations (could focus on content, not troubleshooting)

---

## Lessons Learned & Best Practices

### What Worked Exceptionally Well

1. **Group setup session (Week 1 lab)**: Getting everyone set up together eliminated isolated troubleshooting
2. **Frozen environment**: Locking package versions for entire semester prevented mid-semester breakage
3. **Pre-installed all packages**: Students never needed to install anything (no "pip install" friction)
4. **Course-specific environment**: `cs301-data-science-spring2025.yaml` in GitHub repo = reproducible semester
5. **Auto-stop**: Eliminated budget anxiety, protected from student mistakes

### Teaching Strategies Developed

**Week 1**: "Setup Sprint"
- Dedicate entire lab session to setup
- Have TAs ready for 1-on-1 help
- Don't move forward until 100% students working

**Ongoing**: "Stop Reminder Routine"
- End every lab session with: "Remember: `aws-jupyter stop`"
- Emphasize auto-stop is safety net, not primary shutdown method
- Weekly cost check-in: "Run `aws-jupyter costs` and make sure you're under $15/month"

**Grading**: "Run Every Notebook"
- Launch course environment
- Execute every student notebook start-to-finish
- Can verify results match student claims (academic integrity)

### Future Enhancements Wanted

1. **Class management dashboard** (feature request):
   - Prof. Lee wants to see all 30 students' instances at once
   - Current: Must ask students individually
   - Desired: `aws-jupyter class-status --students cs301-roster.csv`

2. **Bulk environment distribution** (#16):
   - Current: Students manually launch with `--env` flag
   - Desired: Prof. Lee provisions 30 accounts with environment pre-configured
   - Students just run `aws-jupyter start` on Day 1

3. **Cost limits per student** (#19):
   - Current: Honor system ($15/month guideline)
   - Desired: Hard limit with email alert
   - Prof. Lee sets $20/student semester limit, auto-alert at $15

4. **Semester cleanup policy**:
   - Current: Instances auto-delete after 30 days inactive
   - Desired: Keep student environments available for 1 year (alumni access for portfolio/resume)

5. **Assignment templates** (feature request):
   - Desired: Pre-configured notebooks with starter code
   - Distribute via `aws-jupyter assignment import week03-assignment.ipynb`

---

## Technical Details

### Course Environment Specifications

**Instance type**: t4g.medium (ARM64 Graviton)
- 2 vCPUs
- 4 GB RAM
- Cost: $0.0336/hour = $0.81/day (24/7) = $24.19/month (24/7)
- With auto-stop (3h/week average): **$0.40/month per student**

**Storage**: 20 GB EBS per student
- Cost: $2.00/month per student
- Enough for course notebooks + datasets

**GPU instance** (final project only): g4dn.xlarge
- NVIDIA T4 GPU (16GB VRAM)
- Cost: $0.526/hour
- Usage: 6 hours per student for final project = $3.16

**Total cost per student**:
- Regular usage: $0.40/month × 4 months = $1.60
- Storage: $2.00/month × 4 months = $8.00
- GPU (final): $3.16
- **Total**: $12.76 (vs $66.67 budgeted)

### Usage Patterns

**Typical student weekly usage**:
- 1× lecture follow-along (1.5 hours)
- 1× lab session (1 hour)
- 1× homework (2-3 hours)
- **Total**: 4.5-5.5 hours/week
- **Monthly hours**: 18-22 hours
- **Monthly cost**: 20 hours × $0.0336/hr = **$0.67/month**

**Final project spike** (Week 15):
- GPU instance: 6 hours
- Cost: $3.16
- Total semester cost: $2.68 (regular) + $3.16 (GPU) + $8.00 (storage) = **$13.84 per student**

---

## Conclusion

AWS IDE transformed Prof. Lee's teaching experience from "IT support technician" to "data science educator":

### Key Outcomes
- ✅ **96% setup time reduction** (8 hours → 18 minutes)
- ✅ **100% student success rate** (30/30 working vs 27/30)
- ✅ **157 hours saved** (entire semester, teaching + grading)
- ✅ **87.5% under budget** ($249 vs $2,000)
- ✅ **100% reproducibility** (no "works on my machine")
- ✅ **40% improvement in final project completion** (100% vs 60%)

### ROI Summary
- **Time value**: $7,885 (157 hours × $50/hour)
- **Budget savings**: $1,751 (returned to department)
- **AWS IDE cost**: $249
- **Net value**: $9,636 benefit
- **ROI**: 3,869%

### Success Factors
1. **Wizard interface**: Non-technical students successful
2. **Frozen environment**: No mid-semester package update disasters
3. **Auto-stop**: Budget protection, student mistake mitigation
4. **Reproducibility**: Grading seamless, academic integrity verifiable
5. **GPU access**: Challenging final projects possible

**Prof. Lee's quote** (end of semester):
> "I've taught this course 4 times. This was the best semester yet. I actually taught data science instead of troubleshooting package installations. Students completed more challenging projects (GPU-enabled neural networks) at lower cost ($249 vs $2,000 budget). And I got 157 hours back - that's 4 weeks of my life I can spend on research, advising, or family. AWS IDE should be the standard for teaching data science."

**Department adoption**:
- Fall 2025: 3 more CS professors adopting AWS IDE for their courses
- Department now budgeting centrally ($10,000/year for all courses)
- University IT considering campus-wide AWS IDE deployment

Prof. Lee represents the **Course Instructor persona**: teaching 30+ students with diverse technical backgrounds and OS environments, tight budget, no time for infrastructure support, needs reproducibility for grading and academic integrity. AWS IDE addresses all requirements while saving massive time and money.

**Related GitHub Issues**: #1, #2, #4, #12, #13, #16, #19, #20, #28
