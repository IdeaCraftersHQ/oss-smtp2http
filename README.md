SMTP2HTTP (email-to-web)
========================
smtp2http is a simple smtp server that resends the incoming email to the configured web endpoint (webhook) as a basic http post request.

## Features

- Convert SMTP emails to HTTP webhook calls
- **Postmark webhook format support** - Compatible with Postmark's webhook JSON structure
- **Environment variable configuration** - Configure via `.env` file or environment variables
- SPF validation and domain filtering
- Attachment and embedded file support (base64 encoded)
- Docker support for easy deployment

Dev 
===
- `go mod vendor`
- `go build`

Dev with Docker
==============
Locally :
- `go mod vendor`
- `docker build -f Dockerfile.dev -t smtp2http-dev .`
- `docker run -p 25:25 smtp2http-dev --timeout.read=50 --timeout.write=50 --webhook=http://some.hook/api`

Or build it as it comes from the repo :
- `docker build -t smtp2http .`
- `docker run -p 25:25 smtp2http --timeout.read=50 --timeout.write=50 --webhook=http://some.hook/api`

The `timeout` options are of course optional but make it easier to test in local with `telnet localhost 25`

## Testing

Test the SMTP server using telnet:
```bash
telnet localhost 25
HELO zeus
MAIL FROM:<email@from.com>
RCPT TO:<youremail@example.com>
DATA
your mail content
.
```

Docker (production)
=====
**Docker images arn't available online for now**
**See "Dev with Docker" above**
- `docker run -p 25:25 smtp2http --webhook=http://some.hook/api`

## Configuration

### Command Line Usage
```bash
# Basic usage
smtp2http --listen=:25 --webhook=http://localhost:8080/api/smtp-hook

# Postmark format with authentication
smtp2http --postmark --postmark-token=your-server-token

# View all options
smtp2http --help
```

### Environment Variables
Copy `.env.example` to `.env` and configure:

```bash
# SMTP Server Settings
SMTP_LISTEN_ADDR=:25
WEBHOOK_URL=https://your-webhook.com/api

# Postmark Integration
USE_POSTMARK_FORMAT=true
POSTMARK_SERVER_TOKEN=your-server-token

# Optional settings
ALLOWED_DOMAIN=yourdomain.com
SMTP_USER=username
SMTP_PASS=password
```

### Postmark Integration

When using `--postmark` flag:
- Email data is converted to Postmark's webhook JSON format
- Webhook URL defaults to `https://api.postmarkapp.com/email`
- Include `--postmark-token` to send the `X-Postmark-Server-Token` header
- All Postmark-specific fields are properly mapped (FromFull, ToFull, etc.)

```bash
# Simple Postmark setup (uses Postmark API endpoint)
smtp2http --postmark --postmark-token=your-server-token

# Custom webhook with Postmark format
smtp2http --postmark --postmark-token=your-token --webhook=https://your-app.com/webhook
```

Contribution
============
Original repo from @alash3al
Thanks to @aranajuan


