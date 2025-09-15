//go:build linux

package platform

import (
	"fmt"
)

func Run() (Data, error) {
	response, err := Execute("./binaries/linux/osqueryi")
	if err != nil {
		return Data{}, fmt.Errorf("error has occurred %v", err)
	}

	return response, nil
}
