//go:build darwin

package platform

import (
	"fmt"
)

func Run() (Data, error) {
	response, err := Execute("./binaries/darwin/osqueryi")
	if err != nil {
		return Data{}, fmt.Errorf("error has occurred %v", err)
	}

	return response, nil
}
