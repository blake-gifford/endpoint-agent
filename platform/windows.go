//go:build windows

package platform

import (
	"fmt"
)

func Run() error {
	err := Execute("./binaries/windows/osqueryd.exe")
	if err != nil {
		return fmt.Errorf("error has occurred %v", err)
	}
	return nil
}
