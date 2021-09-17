package main

import (
	"github.com/mailtokun/yutu/build/golang"
	"github.com/mailtokun/yutu/github"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	log.Info("# yutu started.")
	err := github.Clone()
	if err != nil {
		log.Error("Failed to clone source code.", err.Error())
		os.Exit(1)
	}
	first()
	ticker := time.NewTicker(time.Second * 20)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			watching()
		}
	}
	log.Info("shut down.")
}
func first() {
	err := golang.Build()
	if err != nil {
		log.Error("Failed to run make build.", err.Error())
		return
	}
	err = golang.Run()
	if err != nil {
		log.Error("Failed to run make run.", err.Error())
	}
}
func watching() {
	err := github.Pull()
	if err != nil {
		return
	}
	err = golang.Build()
	if err != nil {
		log.Error("Failed to run make build.", err.Error())
		return
	}
	err = golang.Run()
	if err != nil {
		log.Error("Failed to run make run.", err.Error())
		return
	}
}
