package util

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

const defaultNatsURL = "my-nats:4222"

// InitNats ...
func InitNats() *nats.Conn {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = defaultNatsURL
	}

	var err error
	nc, err = nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	return nc
}

// PublishNats ...
func PublishNats(subject string, v interface{}) error {
	if nc == nil {
		return errors.New("nats is nil")
	}

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return nc.Publish(subject, b)
}
