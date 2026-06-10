# Deployment

## Docker Compose

```yaml
services:
  ccx:
    image: crpi-i19l8zl0ugidq97v.cn-hangzhou.personal.cr.aliyuncs.com/bene/ccx:latest
    ports:
      - '3000:3000'
    volumes:
      - ./.config:/app/.config
    environment:
      - ENV=production
      - PROXY_ACCESS_KEY=your-proxy-key
      # Admin API key (optional, falls back to PROXY_ACCESS_KEY if not set)
      # - ADMIN_ACCESS_KEY=your-admin-secret-key
    restart: unless-stopped
```

## System Service

### Linux (systemd)

```ini
[Unit]
Description=CCX AI API Gateway
After=network.target

[Service]
Type=simple
ExecStart=/opt/ccx/ccx
WorkingDirectory=/opt/ccx
Environment=PROXY_ACCESS_KEY=your-proxy-key
# Admin API key (optional, falls back to PROXY_ACCESS_KEY if not set)
#Environment=ADMIN_ACCESS_KEY=your-admin-secret-key
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

### macOS (launchd)

See `docs/service/com.ccx.gateway.plist` for reference.

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | 3000 | Server port |
| `ENV` | production | Runtime environment |
| `PROXY_ACCESS_KEY` | - | Proxy access key (required) |
| `PROXY_ACCESS_KEYS` | - | Additional proxy access keys (optional, comma- or newline-separated; does not grant admin access) |
| `ADMIN_ACCESS_KEY` | - | Admin console key (optional) |
| `QUIET_POLLING_LOGS` | true | Suppress polling logs |
| `MAX_REQUEST_BODY_SIZE_MB` | 50 | Max request body size |

See `ENVIRONMENT.md` in the project root for the full list.
