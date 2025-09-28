# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

smtp2http is a Go application that acts as an SMTP server which forwards incoming emails to HTTP webhooks as JSON payloads. It receives emails via SMTP and converts them to structured JSON data that gets posted to configured web endpoints.

## Architecture

The application consists of 4 main Go files:

- `main.go` - Entry point with SMTP server setup and main email processing handler
- `message.go` - Defines data structures for email messages, addresses, attachments, and embedded files
- `vars.go` - Command-line flag definitions and initialization
- `helper.go` - Utility functions for email address transformation

The flow is: SMTP email → Parse with go-smtpsrv → Convert to EmailMessage struct → POST as JSON to webhook

## Key Features

- SPF (Sender Policy Framework) validation
- Domain filtering (optional, via `--domain` flag)
- Base64 encoding of attachments and embedded files
- Comprehensive email field extraction (headers, body, addresses, etc.)
- Configurable timeouts and message size limits

## Development Commands

### Local Development
```bash
go mod vendor
go build
./smtp2http --webhook=http://localhost:8080/api/hook
```

### Docker Development
```bash
# Development build
go mod vendor
docker build -f Dockerfile.dev -t smtp2http-dev .
docker run -p 25:25 smtp2http-dev --timeout.read=50 --timeout.write=50 --webhook=http://some.hook/api

# Production build
docker build -t smtp2http .
docker run -p 25:25 smtp2http --webhook=http://some.hook/api
```

## Configuration Flags

- `--listen` - SMTP listen address (default: `:smtp`)
- `--webhook` - HTTP endpoint to POST email data (required)
- `--domain` - Restrict accepted emails to specific domain
- `--timeout.read` / `--timeout.write` - SMTP timeouts in seconds
- `--msglimit` - Maximum message size (default: 2MB)
- `--name` - Server banner name
- `--user` / `--pass` - SMTP authentication (if needed)

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

## Dependencies

- `github.com/alash3al/go-smtpsrv` - SMTP server implementation
- `github.com/go-resty/resty/v2` - HTTP client for webhook calls