package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	httpClient = &http.Client{Timeout: time.Second * 10}

	webhookURL = os.Getenv("WEBHOOK_URL")
  podName = os.Getenv("POD_NAME")
)

type payload struct {
  Content string `json:"content"`
}

// SendMessage ...
func SendMessage(msg []byte) error {
	if len(webhookURL) == 0 {
		return errors.New("webhookURL is empty")
	}

  var prettyMsg bytes.Buffer
  if err := json.Indent(&prettyMsg, msg, "", "\t"); err != nil {
    return err
  }

  msgFooter := fmt.Sprintf("\nbroadcasted by %s", podName)

	p := &payload{
		Content: string(prettyMsg.Bytes())+msgFooter,
	}

	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

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
