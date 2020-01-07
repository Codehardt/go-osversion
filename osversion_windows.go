//+build windows

package osversion

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func Get() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "Software\\Microsoft\\Windows NT\\CurrentVersion", registry.QUERY_VALUE)
	if err != nil {
		return "", fmt.Errorf("could not open registry key for CurrentVersion: %s", err)
	}
	defer k.Close()
	productName, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return "", fmt.Errorf("could not get ProductName in CurrentVersion registry key: %s", err)
	}
	return productName, nil
}
