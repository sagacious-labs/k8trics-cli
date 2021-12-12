package kubectl

import (
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
	outByt, err := cmd.Output()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			return "", string(exitErr.Stderr), err
		}

		return "", "", err
	}

	return string(outByt), "", nil
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
