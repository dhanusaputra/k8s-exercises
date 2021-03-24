package main

import (
  "crypto/md5"
  "fmt"
  "time"
)

func main() {
  curTime :=  time.Now().String()
  for {
    fmt.Printf("%s: %x\n", time.Now().Format(time.RFC3339), md5.Sum([]byte(curTime)))
    time.Sleep(5 * time.Second)
  }
}
