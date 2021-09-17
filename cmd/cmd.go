package cmd

import (
	"bufio"
	"bytes"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"time"
)

func OSStreamingCommand(timeout time.Duration, name string, args ...string) (cmderr error) {
	/* #nosec */
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	c1 := make(chan error, 1)
	go func() {
		c1 <- cmd.Run()
	}()

	select {
	case err := <-c1:
		cmderr = err
		if cmderr != nil {
			return
		}
	case <-time.After(timeout):
		cmd.Process.Kill()
		cmderr = errors.New("timeout")
		return
	}
	return
}
func StreamingCommand(timeout time.Duration, name string, args ...string) (stdout string, stderr string, cmderr error) {
	/* #nosec */
	cmd := exec.Command(name, args...)
	var stdErrBuffer bytes.Buffer
	cmd.Stderr = &stdErrBuffer
	//cmd.Stdout = os.Stdout
	stdoutPipe, cmderr := cmd.StdoutPipe()
	if cmderr != nil {
		log.Error(cmderr)
		return
	}
	cmderr = cmd.Start()
	if cmderr != nil {
		log.Error(cmderr)
		return
	}
	reader := bufio.NewReader(stdoutPipe)
	line, _, cmderr := reader.ReadLine()
	c1 := make(chan error, 1)
	go func() {
		for cmderr == nil && line != nil && len(line) > 0 {
			stringLine := string(line)
			stdout += stringLine + "\n"
			log.Info(stringLine)
			line, _, cmderr = reader.ReadLine()
		}
		if cmderr == io.EOF {
			c1 <- nil
		} else if cmderr != nil {
			c1 <- cmderr
		}
	}()
	select {
	case err := <-c1:
		cmderr = err
		if cmderr != nil {
			return
		}
	case <-time.After(timeout):
		cmd.Process.Kill()
		cmderr = errors.New("timeout")
		return
	}
	//cmderr = cmd.Wait()
	//if cmderr != nil {
	//	return
	//}
	stderr = stdErrBuffer.String()
	return
}
func StreamingBashCommand(timeout time.Duration, bashCmd string) (cmderr error) {
	/* #nosec */
	cmd := exec.Command("bash", "-c", bashCmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	c1 := make(chan error, 1)
	go func() {
		c1 <- cmd.Run()
	}()

	select {
	case err := <-c1:
		cmderr = err
		if cmderr != nil {
			return
		}
	case <-time.After(timeout):
		cmd.Process.Kill()
		cmderr = errors.New("timeout")
		return
	}
	return
}
