package utils

import (
	"bytes"
	"log"
	"os/exec"
)

func ExecCommand(cmdstr string) (stdout string, stderr string, err error) {
	cmd := exec.Command(cmdstr)
	var stdoutBuffer, stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer
	err = cmd.Run()
	if err != nil {
		log.Printf("error when exec command %s\n", err.Error())
	}
	stdout, stderr = string(stdoutBuffer.Bytes()), string(stderrBuffer.Bytes())
	return
}
