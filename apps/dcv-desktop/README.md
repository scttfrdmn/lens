# lens-dcv-desktop

Full Linux desktop environment with NICE DCV for GUI-based research applications.

## Overview

`lens-dcv-desktop` provides browser-based access to a full Linux desktop environment running on AWS EC2, powered by NICE DCV (Desktop Cloud Visualization). This enables researchers to run GUI applications that require visual interfaces, 3D rendering, or GPU acceleration.

## Use Cases

- **MATLAB**: Full IDE with Simulink support
- **QGIS**: Geographic information system analysis
- **ParaView**: Scientific visualization of large datasets
- **ImageJ/Fiji**: Biological image analysis
- **Geneious**: Bioinformatics workflows
- **General Research**: Any Linux GUI application

## Key Features

### NICE DCV Benefits
- 🌐 **Browser-based access** - No client installation required
- 🎮 **GPU acceleration** - OpenGL, DirectX, CUDA support
- 🚀 **High performance** - Up to 60 fps, low latency
- 🔒 **Secure** - AWS Session Manager (no exposed ports)
- 📋 **Full desktop features** - Copy/paste, file transfer, multi-monitor

### Desktop Environments

1. **general-desktop**
   - Ubuntu desktop with common research tools
   - Pre-installed: Firefox, VS Code, terminal utilities
   - Instance: t3.xlarge (4 vCPU, 16GB RAM)
   - Cost: ~$0.17/hour

2. **gpu-workstation**
   - CUDA-enabled desktop for GPU computing
   - Pre-installed: NVIDIA drivers, CUDA toolkit
   - Instance: g4dn.xlarge (4 vCPU, 16GB RAM, NVIDIA T4 GPU)
   - Cost: ~$0.53/hour

3. **matlab-desktop**
   - Desktop configured for MATLAB (user provides license)
   - Pre-installed: MATLAB dependencies
   - Instance: g4dn.xlarge (recommended for Simulink)
   - Cost: ~$0.53/hour + MATLAB license

4. **data-viz-desktop**
   - ParaView and visualization tools
   - Pre-installed: ParaView, Visit
   - Instance: g4dn.xlarge (GPU recommended)
   - Cost: ~$0.53/hour

5. **image-analysis**
   - ImageJ, Fiji, QuPath, CellProfiler
   - Pre-installed: Image analysis tools
   - Instance: t3.xlarge
   - Cost: ~$0.17/hour

6. **bioinformatics-gui**
   - Geneious, UGENE, bioinformatics tools
   - Pre-installed: Common bioinformatics GUI applications
   - Instance: t3.xlarge
   - Cost: ~$0.17/hour

## Quick Start

### Interactive Wizard (Recommended)
```bash
lens-dcv-desktop
```

### Quick Launch with Defaults
```bash
lens-dcv-desktop quickstart
```

### Launch Specific Environment
```bash
lens-dcv-desktop launch --env matlab-desktop --instance-type g4dn.xlarge
```

## Usage

### Launch Desktop
```bash
lens-dcv-desktop launch \
  --env general-desktop \
  --instance-type t3.xlarge \
  --profile myprofile
```

### Connect to Desktop
```bash
lens-dcv-desktop connect <instance-id>
```
This will:
1. Start SSM port forwarding for DCV (port 8443)
2. Open browser to `https://localhost:8443`
3. Display login credentials
4. Stream desktop to your browser

### List Desktops
```bash
lens-dcv-desktop list
```

### Stop Desktop
```bash
lens-dcv-desktop stop <instance-id>
```

### Terminate Desktop
```bash
lens-dcv-desktop terminate <instance-id>
```

## Desktop Environments

View available environments:
```bash
lens-dcv-desktop env
```

## Configuration

### Auto-stop on Idle
Desktops automatically stop after inactivity to save costs:
```bash
lens-dcv-desktop launch --idle-timeout 30m
```

### GPU Instances
For GPU-accelerated workloads:
```bash
lens-dcv-desktop launch --env gpu-workstation --instance-type g4dn.xlarge
```

Supported GPU instance families:
- **g4dn**: NVIDIA T4 (good price/performance)
- **g5**: NVIDIA A10G (better performance)
- **p3**: NVIDIA V100 (high-end compute)

### Storage
Adjust EBS volume size for large datasets:
```bash
lens-dcv-desktop launch --ebs-size 100
```

## Cost Optimization

### Tips
1. **Use t3 instances** for non-GPU workloads (~$0.17/hour vs ~$0.53/hour)
2. **Enable auto-stop** to prevent idle costs
3. **Stop when not using** - stopped instances only cost for storage
4. **Use Spot instances** for non-critical work (60-70% discount)

### Cost Comparison
| Environment | Instance Type | Cost/Hour | 8hrs/day × 20days |
|-------------|--------------|-----------|-------------------|
| General Desktop | t3.xlarge | $0.17 | $27/month |
| GPU Workstation | g4dn.xlarge | $0.53 | $85/month |
| MATLAB Desktop | g4dn.xlarge | $0.53 | $85/month |

**Note**: Desktop apps cost 4-12x more than web-based apps (Jupyter, RStudio) due to GUI overhead and GPU requirements, but still cost-effective vs. local workstations.

## Technical Details

### Architecture
- **Base**: Ubuntu 22.04 Desktop or AWS DCV AMI
- **Desktop Environment**: GNOME or XFCE
- **DCV Server**: NICE DCV 2023.0+
- **Port**: 8443 (HTTPS/WebSocket)
- **Connection**: AWS Session Manager port forwarding

### Instance Requirements
- **Minimum**: 4 vCPU, 16GB RAM (t3.xlarge)
- **GPU Apps**: g4dn.xlarge or higher
- **Large Datasets**: 32GB+ RAM (t3.2xlarge or g4dn.2xlarge)

### Network
- Uses AWS Session Manager (no public IP required)
- Port 8443 forwarded via SSM tunnel
- Secure end-to-end encryption

## Troubleshooting

### Desktop Won't Load
1. Check instance status: `lens-dcv-desktop status <instance-id>`
2. Verify DCV server is running
3. Check browser console for errors

### Poor Performance
1. Use GPU instance for 3D applications
2. Check network bandwidth (~5-10 Mbps required)
3. Reduce DCV quality settings if needed

### Can't Connect
1. Ensure AWS credentials are valid
2. Verify Session Manager plugin installed
3. Check security group allows SSM traffic

## Development Status

**Version**: 0.10.0 (Development)

This is the initial scaffold for lens-dcv-desktop. Full implementation coming soon!

## Contributing

See the main Lens repository for contribution guidelines.

## License

Copyright © 2025 Lens Project
