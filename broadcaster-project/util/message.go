package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	httpClient = &http.Client{Timeout: time.Second * 10}

	webhookURL = os.Getenv("WEBHOOK_URL")
)

type payload struct {
	content string
}

// SendMessage ...
func SendMessage(msg string) error {
	if len(webhookURL) == 0 {
		return errors.New("webhookURL is empty")
	}

	p := &payload{
		content: msg,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	log.Println(string(b), webhookURL)

	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !isStatusCode2xx(resp.StatusCode) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("failed with statusCode: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

func isStatusCode2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
