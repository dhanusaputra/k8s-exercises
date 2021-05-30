package util

import (
	"errors"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

var (
	ec *nats.EncodedConn
)

const defaultNatsURL = "my-nats:4222"

// InitNats ...
func InitNats() *nats.EncodedConn {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = defaultNatsURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	ec, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("nats successfully init")

	return ec
}

// PublishNats ...
func PublishNats(subj string, v interface{}) error {
	if ec == nil {
		return errors.New("ec is nil")
	}

	return ec.Publish(subj, v)
}
