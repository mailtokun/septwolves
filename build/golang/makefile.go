package golang

import (
	"github.com/mailtokun/yutu/cmd"
	_const "github.com/mailtokun/yutu/const"
	log "github.com/sirupsen/logrus"
	"time"
)

func Build() error {
	log.Info("## make build")
	var buildError error
	var maxRetry int = 0
	for maxRetry < 5 {
		maxRetry++
		buildError = cmd.StreamingBashCommand(time.Hour, "cd "+_const.CODE_FOLDER+" && make build")
		if buildError == nil {
			break
		} else {
			log.Warn(buildError)
			time.Sleep(time.Second * 5)
		}
	}
	return buildError
}

func Run() error {
	log.Info("## make run")
	var runError error
	var maxRetry int = 0
	for maxRetry < 5 {
		maxRetry++
		runError = cmd.StreamingBashCommand(time.Hour, "cd "+_const.CODE_FOLDER+" && make run")
		if runError == nil {
			break
		} else {
			log.Warn(runError)
			time.Sleep(time.Second * 5)
		}
	}
	return runError
}
