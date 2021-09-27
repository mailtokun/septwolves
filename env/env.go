package env

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	_const "github.com/mailtokun/yutu/const"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func WatchEnv() {
	walkFile()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Chmod == fsnotify.Chmod {
					loadEnv(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Error(err)
			}
		}
	}()

	err = watcher.Add(_const.ROOT_FOLDER)
	if err != nil {
		log.Error(err)
	}
	<-done
}

func walkFile() {
	items, _ := ioutil.ReadDir(_const.ROOT_FOLDER)
	for _, item := range items {
		if !item.IsDir() {
			fn := item.Name()
			ext := filepath.Ext(fn)
			if strings.ToLower(ext) == ".env" {
				log.Info("loading environment variables in " + fn + ".")
				loadEnv(_const.ROOT_FOLDER + "/" + fn)
			}
		}
	}
}
func loadEnv(fp string) {
	if strings.ToLower(filepath.Ext(fp)) != ".env" {
		return
	}
	file, err := os.Open(fp)
	if err != nil {
		log.Error("failed to open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}
		items := strings.Split(line, "=")
		if len(items) < 2 {
			continue
		}
		if strings.HasPrefix(items[1], "#") {
			continue
		}
		os.Setenv(items[0], items[1])
	}
	file.Close()
}
