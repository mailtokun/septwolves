package github

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	_const "github.com/mailtokun/yutu/const"
	"github.com/mailtokun/yutu/models"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func Clone(proj models.Project) error {
	log.Info("## git clone " + proj.GithubRepo)
	os.RemoveAll(_const.CODE_FOLDER + "/" + models.GetMD5Hash(proj))
	//var branch = plumbing.ReferenceName(os.Getenv("GITHUB_BRANCH"))
	var op = &git.CloneOptions{
		URL:      proj.GithubRepo,
		Auth:     &http.BasicAuth{Username: "anyuser", Password: proj.GithubToken},
		Progress: os.Stdout,
	}
	var cloneErr error
	var maxRetry int = 0
	for maxRetry < 5 {
		maxRetry++
		_, cloneErr = git.PlainClone(_const.CODE_FOLDER+"/"+models.GetMD5Hash(proj), false, op)
		if cloneErr == nil {
			break
		} else {
			log.Warn(cloneErr)
			time.Sleep(time.Second * 5)
		}
	}
	return cloneErr
}
func Pull(proj models.Project) error {
	log.Info("## git pull " + proj.GithubRepo)
	r, err := git.PlainOpen(_const.CODE_FOLDER + "/" + models.GetMD5Hash(proj))
	if err != nil {
		log.Error(err)
	}
	w, err := r.Worktree()
	if err != nil {
		log.Error(err)
	}
	var op = &git.PullOptions{
		RemoteName: "origin",
		Auth:       &http.BasicAuth{Username: "anyuser", Password: proj.GithubToken},
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
