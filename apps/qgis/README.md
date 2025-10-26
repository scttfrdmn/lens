# lens-qgis

Launch QGIS (Geographic Information System) on AWS with browser-based remote desktop access.

## Overview

`lens-qgis` is an **application-centric** tool that makes it dead simple to launch and use QGIS for GIS analysis. You don't manage a generic desktop - you launch QGIS.

```bash
lens-qgis launch        # Launch QGIS (not "a desktop where you install QGIS")
lens-qgis connect       # Connect to QGIS desktop
```

## Why lens-qgis?

- **No local installation**: QGIS runs in the cloud, access via browser
- **Scalable**: Use t3.xlarge for basic work, g4dn.xlarge for large rasters
- **Pre-configured**: QGIS + common plugins already installed
- **Cost-effective**: Auto-stop on idle, pay only for what you use
- **Secure**: AWS Session Manager, no exposed ports

## QGIS Environments

### basic-gis (Default)
```bash
lens-qgis launch --env basic-gis
```
- **Instance**: t3.xlarge (4 vCPU, 16GB RAM)
- **Cost**: ~$0.17/hour
- **Includes**: QGIS + essential plugins
- **Use for**: General GIS analysis, map creation

### advanced-gis
```bash
lens-qgis launch --env advanced-gis
```
- **Instance**: t3.xlarge
- **Cost**: ~$0.17/hour
- **Includes**: QGIS + GRASS GIS + SAGA GIS + PostGIS
- **Use for**: Advanced spatial analysis, terrain modeling

### remote-sensing
```bash
lens-qgis launch --env remote-sensing
```
- **Instance**: g4dn.xlarge (4 vCPU, 16GB RAM, NVIDIA T4 GPU)
- **Cost**: ~$0.53/hour
- **Includes**: QGIS + Orfeo Toolbox + SNAP + GPU acceleration
- **Use for**: Satellite imagery, large raster processing

## Quick Start

### Interactive Wizard (Recommended)
```bash
lens-qgis
```

### Quick Launch with Defaults
```bash
lens-qgis quickstart
```

### Launch Specific Environment
```bash
lens-qgis launch --env remote-sensing --instance-type g4dn.xlarge
```

## Usage

### Launch QGIS
```bash
lens-qgis launch \
  --env basic-gis \
  --instance-type t3.xlarge \
  --profile myprofile
```

### Connect to QGIS
```bash
lens-qgis connect <instance-id>
```

This will:
1. Start SSM port forwarding for DCV (port 8443)
2. Open browser to `https://localhost:8443`
3. Display login credentials
4. QGIS desktop appears in browser - ready to use!

### List QGIS Instances
```bash
lens-qgis list
```

### Stop QGIS
```bash
lens-qgis stop <instance-id>
```

### Terminate QGIS
```bash
lens-qgis terminate <instance-id>
```

## Configuration

### Auto-stop on Idle
QGIS automatically stops after inactivity to save costs:
```bash
lens-qgis launch --idle-timeout 30m
```

### GPU Instances
For large raster datasets and satellite imagery:
```bash
lens-qgis launch --env remote-sensing --instance-type g4dn.xlarge
```

### Storage
Adjust EBS volume size for large GIS datasets:
```bash
lens-qgis launch --ebs-size 100
```

## Architecture

### Application-Centric Design

Lens follows an **application-centric** model where each tool (`lens-qgis`, `lens-matlab`, `lens-paraview`) is a dedicated command for that specific application.

**This is NOT a generic desktop platform.** If you want a generic Ubuntu/Rocky desktop, use:
- Standard cloud desktop solutions
- AWS WorkSpaces
- Standard Linux distro AMIs

**Lens's value** is making specific research applications trivially easy to launch and use.

### Technical Stack

- **Remote Desktop**: NICE DCV (AWS Desktop Cloud Visualization)
- **Port**: 8443 (HTTPS/WebSocket)
- **Connection**: AWS Session Manager port forwarding (secure, no exposed ports)
- **Desktop Environment**: XFCE (lightweight, fast)
- **AMI**: Ubuntu 22.04 + QGIS + DCV pre-installed
- **GPU Support**: Optional, for remote-sensing environment

## Cost Optimization

### Tips
1. **Use basic-gis** for most work (~$0.17/hour vs ~$0.53/hour GPU)
2. **Enable auto-stop** to prevent idle costs
3. **Stop when not using** - stopped instances only cost for storage
4. **Right-size storage** - only pay for what you need

### Cost Examples
| Environment | Instance | Hours/Day | Days/Month | Monthly Cost |
|-------------|----------|-----------|------------|--------------|
| basic-gis | t3.xlarge | 8 | 20 | ~$27 |
| remote-sensing | g4dn.xlarge | 4 | 10 | ~$21 |
| advanced-gis | t3.xlarge | 8 | 20 | ~$27 |

## Development Status

**Version**: 0.10.0 (Development)

This is the initial application-centric implementation of lens-qgis as part of the v0.10.0 GUI Foundation & Tool Expansion release.

## Contributing

See the main Lens repository for contribution guidelines.

## License

Copyright © 2025 Lens Project
