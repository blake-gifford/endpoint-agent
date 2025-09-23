//go:build linux

package platform

import (
	"fmt"
)

func Run() error {
	err := Execute("./binaries/linux/osquery/bin/osqueryd")
	if err != nil {
		return fmt.Errorf("error has occurred %v", err)
	}
	return nil
}
