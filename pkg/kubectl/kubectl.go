package kubectl

import (
	"fmt"
	"io"
	"os/exec"
)

// Exists returns true if kubectl binary is present in $PATH
func Exists() bool {
	if _, err := exec.LookPath("kubectl"); err != nil {
		return false
	}

	return true
}

// GenericExec takes cmds slice that needs to be run against `kubectl` binary
// and returns its stdout, stderr, error
func GenericExec(cmds []string) (string, string, error) {
	cmd := exec.Command("kubectl", cmds...)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", fmt.Errorf("failed to get stdout pipe")
	}

	er, err := cmd.StderrPipe()
	if err != nil {
		return "", "", fmt.Errorf("failed to get stderr pipe")
	}

	if err := cmd.Start(); err != nil {
		return "", "", err
	}

	outByt := []byte{}
	errByt := []byte{}

	io.ReadFull(out, outByt)
	io.ReadFull(er, errByt)

	cmd.Wait()

	return string(outByt), string(errByt), nil
}

func Apply(locs []string) (string, string, error) {
	args := []string{"apply"}

	for _, loc := range locs {
		args = append(args, "-f", loc)
	}

	return GenericExec(args)
}

func Delete(locs []string) (string, string, error) {
	args := []string{"delete"}

	for _, loc := range locs {
		args = append(args, "-f", loc)
	}

	return GenericExec(args)
}
