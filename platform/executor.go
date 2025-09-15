package platform

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func query(query string, binaryPath string) ([]byte, error) {
	cmd := exec.Command(binaryPath, "--json", query)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return nil, fmt.Errorf("error executing binary: %v\nOutput: %s", err, output)
	}

	return output, nil
}

func queryWithFallbacks(queries []string, binaryPath string) ([]byte, error) {
	var lastErr error

	for i, queryStr := range queries {
		output, err := query(queryStr, binaryPath)
		if err != nil {
			lastErr = err
			continue
		}

		var result []map[string]interface{}
		if err := json.Unmarshal(output, &result); err != nil {
			lastErr = fmt.Errorf("invalid JSON from query %d: %v", i+1, err)
			continue
		}

		if len(result) > 0 {
			return output, nil
		}

		lastErr = fmt.Errorf("query %d returned no data", i+1)
	}

	return nil, fmt.Errorf("all queries failed, last error: %v", lastErr)
}

func Execute(binaryPath string) (Data, error) {
	softwareOutput, err := queryWithFallbacks(queryInstalledSoftwareFallbacks, binaryPath)
	if err != nil {
		return Data{}, fmt.Errorf("error querying software: %v", err)
	}

	var software []Software
	if err := json.Unmarshal(softwareOutput, &software); err != nil {
		return Data{}, fmt.Errorf("error parsing software JSON: %v", err)
	}

	systemOutput, err := query(querySystemInfo, binaryPath)
	if err != nil {
		return Data{}, fmt.Errorf("error querying system info: %v", err)
	}

	var systemInfoArray []System
	if err := json.Unmarshal(systemOutput, &systemInfoArray); err != nil {
		return Data{}, fmt.Errorf("error parsing system JSON: %v", err)
	}

	if len(systemInfoArray) == 0 {
		return Data{}, fmt.Errorf("no system information returned from query")
	}

	systemInfo := systemInfoArray[0]

	return Data{
		Software:   software,
		SystemInfo: systemInfo,
	}, nil
}
