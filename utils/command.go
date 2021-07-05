package utils 
import (
	"exec"
	"bytes"
)

func ExecCommand(cmd string) (stdout string,stderr string, err error) {
	cmd := exec.Command("ls", "-lah")
	var stdoutBuffer, strerrBuffer bytes.Buffer
	cmd.
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Println("error when exec command %s\n", err)
	}
	return 
}