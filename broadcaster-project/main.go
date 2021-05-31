package main

import (
	"log"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
  "github.com/dhanusaputra/k8s-exercises/broadcaster-project/util"
)

const defaultNatsURL = "my-nats:4222"

func main() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = defaultNatsURL
	}

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := nc.Subscribe("todo", func(m *nats.Msg) {
		defer wg.Done()
		log.Printf("Received a message: %s\n", string(m.Data))
    if err := util.SendMessage("test"); err != nil {
      log.Fatal(err)
    }
	}); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
