//+build aix

package osversion

import (
	"fmt"
	"os/exec"
	"strings"
)

func Get() (string, error) {
	oslevel, err := exec.Command("/usr/bin/oslevel").Output()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("AIX %s", strings.TrimSuffix(string(oslevel), "\n")), nil
}
