package platform

import (
	"fmt"
	"os"
	"os/exec"
)

var osqueryProcess *os.Process

func startOsqueryDaemon(binaryPath string, configPath string) error {
	token := os.Getenv("TOKEN")
	organization := os.Getenv("ORGANIZATION")

	if token == "" || organization == "" {
		return fmt.Errorf("TOKEN and ORGANIZATION environment variables must be set")
	}

	endpoint := fmt.Sprintf("29aeebc790eb.ngrok-free.app/agent?authorization=%s&organization=%s", token, organization)

	cmd := exec.Command("sudo", binaryPath,
		"--allow_unsafe=true",
		"--logger_plugin=tls",
		"--disable_enrollment=true",
		"--logger_tls_endpoint="+endpoint,
		"--tls_server_certs=./certs/certs.pem",
		"--database_path=/tmp/osquery.db",
		"--config_path="+configPath,
		"--disable_database=true",
		"--disable_caching=true",
		"--disable_events=true",
		"--logger_min_status=1",
		"--disable_audit=true",
		"--config_refresh=1",
		"--json",
	)

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start osqueryd daemon: %v", err)
	}

	osqueryProcess = cmd.Process
	return nil
}

func StopOsqueryDaemon() error {
	if osqueryProcess != nil {
		fmt.Printf("Killing osquery process PID: %d\n", osqueryProcess.Pid)
		err := osqueryProcess.Kill()
		osqueryProcess = nil
		return err
	}

	exec.Command("sudo", "pkill", "-f", "osqueryd").Run()
	return nil
}

func Execute(binaryPath string) error {
	err := startOsqueryDaemon(binaryPath, "./osquery.conf")
	if err != nil {
		return fmt.Errorf("error starting osqueryd daemon: %v", err)
	}

	return nil
}
