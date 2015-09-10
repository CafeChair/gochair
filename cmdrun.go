package x

import (
	"log"
	"time"
	"errors"
	"os/exec"
	"bytes"
)

func Run(command string) (string,error) {
	cmd := exec.Command("sh","-c", command)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Start()
	cmdDone := make(chan error, 1)
	go func(){
		cmdDone <- cmd.Wait()
	}()

	select{
	case <- time.After(time.Duration(60) * time.Second):
		KillCmd(cmd)
		<- cmdDone
		return "", errors.New("Command timeout")
	case err := <- cmdDone:
		if err != nil {
			log.Println(stderr.String())
		}
		return out.String(), err
	}
}
func KillCmd(cmd *exec.Cmd) {
	if err := cmd.Process.Kill(); err != nil {
		log.Printf("Failed to kill command: %v", err)
	}
}
