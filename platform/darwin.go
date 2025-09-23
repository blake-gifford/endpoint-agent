//go:build darwin

package platform

import (
	"fmt"
)

func Run() error {
	err := Execute("./binaries/darwin/lib/osquery.app/Contents/MacOS/osqueryd")
	if err != nil {
		return fmt.Errorf("error has occurred %v", err)
	}
	return nil
}
