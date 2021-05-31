package util

import (
  "os"
  "net/http"
  "time"
)

var (
  webhookURL = os.Getenv("WEBHOOK_URL")

  client = &http.Client{Timeout: time.Second * 10}
)

// SendMessage ...
func SendMessage(msg string) error {
  return nil
}
