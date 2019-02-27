package osversion

import (
	"fmt"
	"os/exec"
)

func Get() (string, error) {
	cmd := exec.Command("defaults", "read", "loginwindow", "SystemVersionStampAsString")
	b, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not read SystemVersion from defaults: %s", err)
	}
	return "MacOS " + string(b), nil
}
