//+build linux

package osversion

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Get() (string, error) {
	if version := getFromOSRelease(); version != "" {
		return version, nil
	}
	if version := getFromLSB(); version != "" {
		return version, nil
	}
	if version := getFromDebianVersion(); version != "" {
		return version, nil
	}
	if version := getFromRedhatRelease(); version != "" {
		return version, nil
	}
	if version := getFromSuSeRelease(); version != "" {
		return version, nil
	}
	return getFromUname()
}

func getFromOSRelease() string {
	b, err := readFileSafe("/etc/os-release")
	if err != nil {
		return ""
	}
	r := bufio.NewReader(bytes.NewReader(b))
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return ""
		}
		// PRETTY_NAME="Debian GNU/Linux 9 (stretch)"
		if strings.HasPrefix(line, `PRETTY_NAME="`) && len(line) >= 15 {
			return line[13 : len(line)-2]
		}
	}
	return ""
}

func getFromDebianVersion() string {
	b, err := readFileSafe("/etc/debian_version")
	if err != nil {
		return ""
	}
	r := bufio.NewReader(bytes.NewReader(b))
	line, err := r.ReadString('\n')
	if err != nil {
		return ""
	}
	return "Debian " + strings.TrimSuffix(line, "\n")
}

func getFromRedhatRelease() string {
	b, err := readFileSafe("/etc/redhat-release")
	if err != nil {
		return ""
	}
	r := bufio.NewReader(bytes.NewReader(b))
	line, err := r.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSuffix(line, "\n")
}

func getFromSuSeRelease() string {
	b, err := readFileSafe("/etc/SuSe-release")
	if err != nil {
		return ""
	}
	r := bufio.NewReader(bytes.NewReader(b))
	line, err := r.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSuffix(line, "\n")
}

func getFromLSB() string {
	oscmd := exec.Command("lsb_release", "-si")
	vercmd := exec.Command("lsb_release", "-sr")
	os, err := oscmd.Output()
	if err != nil {
		return ""
	}
	ver, err := vercmd.Output()
	if err != nil {
		return ""
	}
	return string(os) + " " + string(ver)
}

func getFromUname() (string, error) {
	oscmd := exec.Command("uname", "-s")
	vercmd := exec.Command("uname", "-r")
	os, err := oscmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not execute uname: %s", err)
	}
	ver, err := vercmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not execute uname: %s", err)
	}
	return string(os) + " " + string(ver), nil
}

func readFileSafe(path string) ([]byte, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := make([]byte, 1000000)
	n, err := f.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}
