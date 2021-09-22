package main

import (
	"encoding/json"
	"github.com/mailtokun/yutu/build/golang"
	_const "github.com/mailtokun/yutu/const"
	"github.com/mailtokun/yutu/github"
	"github.com/mailtokun/yutu/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var projects []models.Project

func main() {
	projPath := filepath.Clean(_const.PROJECTS_PATH)
	if _, err := os.Stat(projPath); os.IsNotExist(err) {
		// load project from OS environment variable
		var project models.Project
		project.GithubRepo = os.Getenv("GITHUB_REPO")
		project.GithubBranch = os.Getenv("GITHUB_BRANCH")
		project.GithubToken = os.Getenv("GITHUB_TOKEN")
		projects = append(projects, project)
		return
	} else {
		// load project(s) from file
		content, readErr := ioutil.ReadFile(projPath)
		if readErr != nil {
			log.Error(readErr)
			return
		}
		unmarshalErr := json.Unmarshal(content, &projects)
		if unmarshalErr != nil {
			log.Error(unmarshalErr)
			return
		}
	}
	log.Info("# yutu started.")
	for i := 0; i < len(projects); i++ {
		log.Info("## " + strconv.Itoa(i) + "/" + strconv.Itoa(len(projects)))
		err := github.Clone(projects[i])
		if err != nil {
			log.Error("Failed to clone source code.", err.Error())
			os.Exit(1)
		}
		first(projects[i])
	}

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
func first(proj models.Project) {
	err := golang.Build(proj)
	if err != nil {
		log.Error("Failed to run make build.", err.Error())
		return
	}
	err = golang.Run(proj)
	if err != nil {
		log.Error("Failed to run make run.", err.Error())
	}
}
func watching() {
	for i := 0; i < len(projects); i++ {
		proj := projects[i]
		err := github.Pull(proj)
		if err != nil {
			continue
		}
		err = golang.Build(proj)
		if err != nil {
			log.Error("Failed to run make build.", err.Error())
			continue
		}
		err = golang.Run(proj)
		if err != nil {
			log.Error("Failed to run make run.", err.Error())
			continue
		}
	}
}
