package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"main/platform"
	"net/http"
	"os"
)

type Data = platform.Data

type Headers struct {
	OrganizationId int    `json:"organizationId"`
	Authorization  string `json:"authorization"`
}

type Request struct {
	Data    Data    `json:"data"`
	Headers Headers `json:"headers"`
}

var URL string

func send(data Data) error {
	url := URL
	if url == "" {
		return fmt.Errorf("URL is required but not set")
	}

	client := &http.Client{}

	request := Request{
		Data: data,
		Headers: Headers{
			OrganizationId: 1,
			Authorization:  "Bearer your-token",
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP request failed with status %d", resp.StatusCode)
	}

	return nil
}

func main() {
	data, err := platform.Run()
	if err != nil {
		log.Printf("Failed to execute platform: %v", err)
		os.Exit(1)
	}

	if err := send(data); err != nil {
		log.Printf("Failed to send data: %v", err)
		os.Exit(1)
	}

	log.Printf("Data sent successfully")
}
