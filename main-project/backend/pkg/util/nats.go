package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var (
	nc *nats.Conn
)

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

	log.Println("nats successfully init")

	return nc
}

// PublishNats ...
func PublishNats(subj, action string, v interface{}) error {
	if nc == nil {
		return errors.New("nc is nil")
	}

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var prettyV bytes.Buffer
	if err := json.Indent(&prettyV, b, "", "\t"); err != nil {
		return err
	}

	msgB := append([]byte(fmt.Sprintf("A task was %sd\n", action)), prettyV.Bytes()...)

	return nc.Publish(subj, msgB)
}
