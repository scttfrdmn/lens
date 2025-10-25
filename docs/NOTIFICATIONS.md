# Notification Hooks

Lens supports flexible notification hooks that let you receive alerts for instance lifecycle events via Slack, email, desktop notifications, or any custom script.

## Overview

Notification hooks are shell commands that execute when specific events occur (launch, stop, connect, etc.). Each hook receives event details as environment variables, making it easy to integrate with external services.

## Configuration

Add hooks to your `~/.lens/config.yaml` file:

```yaml
hooks:
  on_launch_started: 'echo "Launching instance ${AWS_IDE_INSTANCE_ID}"'
  on_launch_completed: 'echo "Instance ${AWS_IDE_INSTANCE_ID} is ready!"'
  on_launch_failed: 'echo "Launch failed: ${AWS_IDE_ERROR}"'
  on_stop_started: 'echo "Stopping instance ${AWS_IDE_INSTANCE_ID}"'
  on_stop_completed: 'echo "Instance ${AWS_IDE_INSTANCE_ID} stopped"'
  on_stop_failed: 'echo "Stop failed: ${AWS_IDE_ERROR}"'
  on_connect_started: 'echo "Connecting to ${AWS_IDE_INSTANCE_ID}"'
  on_connect_failed: 'echo "Connection failed: ${AWS_IDE_ERROR}"'
```

## Available Events

| Event | When It Fires | Environment Variables |
|-------|---------------|----------------------|
| `on_launch_started` | When launching begins | Instance type, environment, region |
| `on_launch_completed` | When instance is ready | Instance ID, IP address, connection URL |
| `on_launch_failed` | When launch fails | Error message |
| `on_stop_started` | When stopping begins | Instance ID |
| `on_stop_completed` | When instance is stopped | Instance ID, total runtime |
| `on_stop_failed` | When stop fails | Instance ID, error message |
| `on_connect_started` | When connecting to instance | Instance ID, connection method |
| `on_connect_failed` | When connection fails | Instance ID, error message |

## Environment Variables

All hooks receive these environment variables:

- `AWS_IDE_EVENT`: Event type (e.g., "launch_completed")
- `AWS_IDE_INSTANCE_ID`: EC2 instance ID
- `AWS_IDE_INSTANCE_TYPE`: Instance type (e.g., "t4g.medium")
- `AWS_IDE_ENVIRONMENT`: Environment name (e.g., "data-science")
- `AWS_IDE_REGION`: AWS region
- `AWS_IDE_TIMESTAMP`: Event timestamp (RFC3339 format)
- `AWS_IDE_APP`: App name ("jupyter", "rstudio", or "vscode")
- `AWS_IDE_ERROR`: Error message (only for failed events)

## Examples

### Slack Notifications

Post messages to a Slack channel using a webhook:

```yaml
hooks:
  on_launch_completed: |
    curl -X POST -H 'Content-type: application/json' \
      --data "{\"text\":\"✅ Instance ${AWS_IDE_INSTANCE_ID} is ready! (${AWS_IDE_APP}, ${AWS_IDE_INSTANCE_TYPE})\"}" \
      https://hooks.slack.com/services/YOUR/WEBHOOK/URL

  on_launch_failed: |
    curl -X POST -H 'Content-type: application/json' \
      --data "{\"text\":\"❌ Launch failed: ${AWS_IDE_ERROR}\"}" \
      https://hooks.slack.com/services/YOUR/WEBHOOK/URL
```

**How to get a Slack webhook URL:**
1. Go to https://api.slack.com/messaging/webhooks
2. Create a new incoming webhook for your workspace
3. Copy the webhook URL

### Email Notifications

Send emails using the system's mail command:

```yaml
hooks:
  on_launch_completed: |
    echo "Your ${AWS_IDE_APP} environment is ready at instance ${AWS_IDE_INSTANCE_ID}" | \
      mail -s "Lens: Environment Ready" your-email@example.com

  on_stop_completed: |
    echo "Your ${AWS_IDE_APP} instance ${AWS_IDE_INSTANCE_ID} has been stopped" | \
      mail -s "Lens: Instance Stopped" your-email@example.com
```

**Note:** Requires `mail` or `sendmail` to be configured on your system. On macOS, you can use `mail` directly. On Linux, you may need to install and configure `mailutils` or `postfix`.

### Desktop Notifications

#### macOS

```yaml
hooks:
  on_launch_completed: |
    osascript -e "display notification \"${AWS_IDE_APP} instance ready!\" with title \"Lens\""

  on_launch_failed: |
    osascript -e "display notification \"Launch failed: ${AWS_IDE_ERROR}\" with title \"Lens\" sound name \"Basso\""
```

#### Linux (using notify-send)

```yaml
hooks:
  on_launch_completed: |
    notify-send "Lens" "${AWS_IDE_APP} instance ${AWS_IDE_INSTANCE_ID} is ready!"

  on_launch_failed: |
    notify-send -u critical "Lens" "Launch failed: ${AWS_IDE_ERROR}"
```

**Note:** Requires `libnotify` package on Linux.

### Logging to File

Keep a log of all events:

```yaml
hooks:
  on_launch_started: 'echo "[$(date)] Launch started: ${AWS_IDE_INSTANCE_TYPE}" >> ~/.lens/events.log'
  on_launch_completed: 'echo "[$(date)] Launch completed: ${AWS_IDE_INSTANCE_ID}" >> ~/.lens/events.log'
  on_stop_completed: 'echo "[$(date)] Stop completed: ${AWS_IDE_INSTANCE_ID}" >> ~/.lens/events.log'
```

### Multiple Actions

Combine multiple actions in one hook:

```yaml
hooks:
  on_launch_completed: |
    # Log to file
    echo "[$(date)] ${AWS_IDE_APP} ready: ${AWS_IDE_INSTANCE_ID}" >> ~/.lens/events.log

    # Send desktop notification
    osascript -e "display notification \"Your environment is ready!\" with title \"Lens\""

    # Post to Slack
    curl -X POST -H 'Content-type: application/json' \
      --data "{\"text\":\"✅ ${AWS_IDE_APP} ready: ${AWS_IDE_INSTANCE_ID}\"}" \
      https://hooks.slack.com/services/YOUR/WEBHOOK/URL
```

### Cost Alerts

Track costs and send alerts when they exceed thresholds:

```yaml
hooks:
  on_launch_completed: |
    # Calculate current month cost
    COST=$(lens-jupyter costs --format json | jq '.total_cost')
    if (( $(echo "$COST > 50" | bc -l) )); then
      echo "⚠️ Monthly cost ($${COST}) exceeded $50 threshold!" | \
        mail -s "Lens: Cost Alert" lab-pi@example.com
    fi
```

### Lab PI Dashboard

Post all lab activity to a dedicated Slack channel:

```yaml
hooks:
  on_launch_started: |
    curl -X POST -H 'Content-type: application/json' \
      --data "{\"text\":\"🚀 ${AWS_IDE_APP} launching: ${AWS_IDE_INSTANCE_TYPE} (${AWS_IDE_ENVIRONMENT})\"}" \
      https://hooks.slack.com/services/LAB/DASHBOARD/WEBHOOK

  on_launch_completed: |
    curl -X POST -H 'Content-type: application/json' \
      --data "{\"text\":\"✅ ${AWS_IDE_APP} ready: ${AWS_IDE_INSTANCE_ID}\"}" \
      https://hooks.slack.com/services/LAB/DASHBOARD/WEBHOOK

  on_stop_completed: |
    curl -X POST -H 'Content-type: application/json' \
      --data "{\"text\":\"🛑 ${AWS_IDE_APP} stopped: ${AWS_IDE_INSTANCE_ID}\"}" \
      https://hooks.slack.com/services/LAB/DASHBOARD/WEBHOOK
```

## Advanced Configuration

### Conditional Hooks

Use shell logic to conditionally execute hooks:

```yaml
hooks:
  on_launch_completed: |
    # Only notify for production instances
    if [[ "${AWS_IDE_ENVIRONMENT}" == "production" ]]; then
      echo "Production instance launched" | mail -s "Alert" admin@example.com
    fi
```

### Custom Scripts

Call external scripts for complex logic:

```yaml
hooks:
  on_launch_completed: '/path/to/your/notify-team.sh "${AWS_IDE_INSTANCE_ID}" "${AWS_IDE_APP}"'
```

Example `notify-team.sh`:
```bash
#!/bin/bash
INSTANCE_ID=$1
APP=$2

# Your custom logic here
curl -X POST https://your-api.com/notifications \
  -d "instance_id=${INSTANCE_ID}" \
  -d "app=${APP}"
```

### Timeout

Hooks have a 30-second timeout to prevent blocking the main operation. If your hook takes longer, run it in the background:

```yaml
hooks:
  on_launch_completed: 'nohup /path/to/slow-script.sh "${AWS_IDE_INSTANCE_ID}" > /dev/null 2>&1 &'
```

## Security Considerations

1. **Webhook URLs**: Store sensitive webhook URLs in environment variables instead of directly in config:
   ```yaml
   hooks:
     on_launch_completed: 'curl -X POST "${SLACK_WEBHOOK_URL}" -d "..."'
   ```

2. **Credentials**: Never put passwords or API keys directly in hooks. Use AWS Secrets Manager or environment variables.

3. **Command Injection**: Hooks execute as shell commands with your user permissions. Be careful with user input.

4. **File Permissions**: The `~/.lens/config.yaml` file should have restrictive permissions (600) to prevent unauthorized access.

## Troubleshooting

### Hook Not Executing

1. Check that your `~/.lens/config.yaml` is valid YAML
2. Ensure the hook command is quoted properly (use `|` for multi-line commands)
3. Test the command manually in your shell first
4. Check for syntax errors in your shell command

### Hook Fails Silently

Hooks are designed to not block main operations. To see hook errors:

1. Test the command manually:
   ```bash
   export AWS_IDE_INSTANCE_ID="i-test123"
   export AWS_IDE_APP="jupyter"
   # Run your hook command here
   ```

2. Add explicit logging to your hook:
   ```yaml
   hooks:
     on_launch_completed: |
       {
         curl -X POST https://... 2>&1
       } >> ~/.lens/hook-errors.log
   ```

### Webhook Returns Error

- **Slack**: Verify webhook URL is correct and not expired
- **Email**: Ensure mail command is installed and configured
- **Desktop**: Check notification service is running

## Disabling Hooks

To temporarily disable all hooks, rename or remove the hooks section from `~/.lens/config.yaml`:

```yaml
# hooks:  # Commented out to disable
#   on_launch_completed: '...'
```

Or remove specific hooks you don't want:

```yaml
hooks:
  on_launch_completed: '...'  # Keep this one
  # on_launch_failed: '...'    # Disabled this one
```

## Best Practices

1. **Start Simple**: Begin with basic notifications and add complexity as needed
2. **Test First**: Always test hooks manually before adding to config
3. **Use Timeouts**: For long-running hooks, use background execution (`&`)
4. **Log Important Events**: Keep a log file of critical events for auditing
5. **Team Coordination**: For lab environments, coordinate notification channels with your team
6. **Cost Monitoring**: Set up alerts when costs exceed your lab budget

## Related Commands

- `lens-jupyter config show` - View current configuration
- `lens-jupyter config edit` - Edit configuration file
- `lens-jupyter costs` - View cost information for use in hooks

## Future Enhancements

We're considering adding built-in notification providers in future versions:
- Direct Slack integration (no webhook URL needed)
- AWS SES email integration
- Microsoft Teams support
- Discord webhooks

Vote for features you'd like to see: https://github.com/scttfrdmn/lens/issues
