package agent

import (
	"bytes"
	"os/exec"
	"time"
)

func RunTask(taskname string, timout time.Duration) (string, error) {
	cmd := exec.Command("sh", "-c", taskname)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Start()
	err, IsTimeout := timeOut(cmd, timout)
	errStr := stderr.String()
	if errStr != "" {
		ColorLog("[ERROR] 执行任务: %v 失败: %s\n", taskname, errStr)
	}
	if IsTimeout {
		if err == nil {
			ColorLog("[INFO] 执行任务超时并杀死: %v 成功 \n", taskname)
		}
		if err != nil {
			ColorLog("[ERROR] 执行任务超时并杀死: %v 失败: %s\n", taskname, err)
		}
		return "", err
	}
	return stdout.String(), err
}

func timeOut(command *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- command.Wait()
	}()
	var err error
	select {
	case <-time.After(timeout):
		ColorLog("[ERROR] 任务执行超时,进程: %s 将被杀掉\n", command.Path)
		go func() {
			<-done
		}()
		if err = command.Process.Kill(); err != nil {
			ColorLog("[ERROR] 杀掉任务失败: %s,错误是:%s\n", command.Path, err)
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
