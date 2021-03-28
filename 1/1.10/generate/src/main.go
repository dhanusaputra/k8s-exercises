package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	random := md5.Sum([]byte(time.Now().String()))
	for {
		p := "./shared/test"
		if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
			log.Fatal(err)
		}
		f, err := os.Create(p)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		msg := fmt.Sprintf("%s: %x\n", time.Now().Format(time.RFC3339), random)
		_, err = f.WriteString(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(msg)
		f.Sync()
		time.Sleep(5 * time.Second)
	}
}
