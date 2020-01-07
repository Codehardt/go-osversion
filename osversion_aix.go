//+build aix

package osversion

import (
	"fmt"
	"os/exec"
	"strings"
)

func Get() (string, error) {
	version, err := exec.Command("uname -v").Output()
	if err != nil {
		return "", err
	}
	release, err := exec.Command("uname -r").Output()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(string(version), "\n"), strings.TrimSuffix(string(release), "\n")), nil
}
