package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	fmt.Println("///")
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			watching()
		}
	}
	log.Info("shut down.")
}
func watching() {
	fmt.Println("watching")
	time.Sleep(time.Second * 5)
	fmt.Println("watching done")
}
