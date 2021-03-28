package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	random := md5.Sum([]byte(time.Now().String()))
	for {
		fmt.Printf("%s: %x\n", time.Now().Format(time.RFC3339), random)
		time.Sleep(5 * time.Second)
	}
}
