package github

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	_const "github.com/mailtokun/yutu/const"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func Clone() error {
	log.Info("## git clone " + os.Getenv("GITHUB_REPO"))
	os.RemoveAll(_const.CODE_FOLDER)
	//var branch = plumbing.ReferenceName(os.Getenv("GITHUB_BRANCH"))
	var op = &git.CloneOptions{
		URL:      os.Getenv("GITHUB_REPO"),
		Auth:     &http.BasicAuth{Username: "anyuser", Password: os.Getenv("GITHUB_TOKEN")},
		Progress: os.Stdout,
	}
	var cloneErr error
	var maxRetry int = 0
	for maxRetry < 5 {
		maxRetry++
		_, cloneErr = git.PlainClone(_const.CODE_FOLDER, false, op)
		if cloneErr == nil {
			break
		} else {
			log.Warn(cloneErr)
			time.Sleep(time.Second * 5)
		}
	}
	return cloneErr
}
func Pull() error {
	log.Info("## git pull " + os.Getenv("GITHUB_REPO"))
	r, err := git.PlainOpen(_const.CODE_FOLDER)
	if err != nil {
		log.Error(err)
	}
	w, err := r.Worktree()
	if err != nil {
		log.Error(err)
	}
	var op = &git.PullOptions{
		RemoteName: "origin",
		Auth:       &http.BasicAuth{Username: "anyuser", Password: os.Getenv("GITHUB_TOKEN")},
		Progress:   os.Stdout,
	}
	var pullErr error
	var maxRetry int = 0
	for maxRetry < 5 {
		maxRetry++
		pullErr = w.Pull(op)
		if pullErr == nil || pullErr == git.NoErrAlreadyUpToDate {
			break
		} else {
			log.Warn(pullErr)
			time.Sleep(time.Second * 5)
		}
	}
	return pullErr
}
