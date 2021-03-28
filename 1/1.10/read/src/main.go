package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	for {
		b, err := ioutil.ReadFile("./shared/test")
		if err != nil {
			log.Println(err)
		}
		fmt.Print(string(b))
		time.Sleep(5 * time.Second)
	}
}
