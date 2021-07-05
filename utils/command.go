package utils

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ExecCommand(cmdstr string) (stdout string, stderr string, err error) {
	path, _ := os.Executable()
	dir := filepath.Dir(path)
	stdout, stderr, err = ExecCommandInPath(cmdstr, dir)
	return
}
func ExecCommandInPath(cmdstr string, cmdPath string) (stdout string, stderr string, err error) {
	cmd := exec.Command(cmdstr)
	cmd.Dir = cmdPath
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
