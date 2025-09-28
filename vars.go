package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	flagServerName     = flag.String("name", getEnvOrDefault("SMTP_SERVER_NAME", "smtp2http"), "the server name")
	flagListenAddr     = flag.String("listen", getEnvOrDefault("SMTP_LISTEN_ADDR", ":smtp"), "the smtp address to listen on")
	flagWebhook        = flag.String("webhook", getWebhookDefault(), "the webhook to send the data to")
	flagMaxMessageSize = flag.Int64("msglimit", getEnvOrDefaultInt64("MAX_MESSAGE_SIZE", 1024*1024*2), "maximum incoming message size")
	flagReadTimeout    = flag.Int("timeout.read", getEnvOrDefaultInt("READ_TIMEOUT", 5), "the read timeout in seconds")
	flagWriteTimeout   = flag.Int("timeout.write", getEnvOrDefaultInt("WRITE_TIMEOUT", 5), "the write timeout in seconds")
	flagAuthUSER       = flag.String("user", getEnvOrDefault("SMTP_USER", ""), "user for smtp client")
	flagAuthPASS       = flag.String("pass", getEnvOrDefault("SMTP_PASS", ""), "pass for smtp client")
	flagDomain         = flag.String("domain", getEnvOrDefault("ALLOWED_DOMAIN", ""), "domain for recieving mails")
	flagPostmark       = flag.Bool("postmark", getEnvOrDefaultBool("USE_POSTMARK_FORMAT", false), "use Postmark webhook format instead of default format")
	flagPostmarkToken  = flag.String("postmark-token", getEnvOrDefault("POSTMARK_SERVER_TOKEN", ""), "Postmark server token to send in X-Postmark-Server-Token header")
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvOrDefaultInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvOrDefaultBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getWebhookDefault() string {
	if value := os.Getenv("WEBHOOK_URL"); value != "" {
		return value
	}

	if getEnvOrDefaultBool("USE_POSTMARK_FORMAT", false) {
		return "https://api.postmarkapp.com/email"
	}

	return "http://localhost:8080/my/webhook"
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables and defaults")
	}
	flag.Parse()
}
