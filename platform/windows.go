//go:build windows

package platform

import (
	"fmt"
)

func Run() (Data, error) {
	response, err := Execute("./binaries/windows/osqueryi.exe")
	if err != nil {
		return Data{}, fmt.Errorf("error has occurred %v", err)
	}

	return response, nil
}
