package platform

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func query(query string, binaryPath string) ([]byte, error) {
	cmd := exec.Command(binaryPath, "--json", query)
	output, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("error executing binary: %v", err)
	}

	return output, nil
}

func queryWithFallbacks(queries []string, binaryPath string) ([]byte, error) {
	var allResults []map[string]any
	var hasSuccessfulQuery bool

	for _, queryStr := range queries {
		output, err := query(queryStr, binaryPath)
		if err != nil {
			continue
		}

		var result []map[string]any
		if err := json.Unmarshal(output, &result); err != nil {
			continue
		}

		if len(result) > 0 {
			allResults = append(allResults, result...)
			hasSuccessfulQuery = true
		}
	}

	if !hasSuccessfulQuery {
		return json.Marshal([]map[string]any{})
	}

	combinedOutput, err := json.Marshal(allResults)
	if err != nil {
		return nil, fmt.Errorf("error marshaling combined results: %v", err)
	}

	return combinedOutput, nil
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
