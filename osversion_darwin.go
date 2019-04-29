package osversion

import (
	"fmt"
	"os/exec"
)

func Get() (string, error) {
	cmd := exec.Command("sw_vers", "-productVersion")
	b, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not read productVersion from sw_vers: %s", err)
	}
	return "MacOS " + string(b), nil
}
