package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func discoverOS() (string, error) {
	osFile := "/etc/os-release"
	out, err := ioutil.ReadFile(osFile)
	if err != nil {
		return "", fmt.Errorf("unable to read '%v', %v", osFile, err)
	}

	env, err := godotenv.Unmarshal(string(out))
	if err != nil {
		return "", fmt.Errorf("unable to source '%v', %v", osFile, err)
	}

	return env["NAME"], nil
}

func discoverInit() (string, error) {
	f := "/proc/1/comm"
	out, err := ioutil.ReadFile(f)
	if err != nil {
		return "", fmt.Errorf("unable to read '%v', %v", f, err)

	}

	return strings.TrimSuffix(string(out), "\n"), nil
}

func discoverKernel() (string, error) {
	out, err := exec.Command("uname", "--all").Output()
	if err != nil {
		return "", fmt.Errorf("unable to exec uname, %v", err)
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}

func discoverLogger() ([]string, error) {
	var result []string
	for _, logger := range []string{"systemd-journal", "rsyslogd", "syslog-ng"} {
		running, err := runningProcesses(logger)
		if err != nil {
			return nil, fmt.Errorf("discovering Logger failed, %v", err)
		}
		if running {
			result = append(result, logger)
		}
	}
	return result, nil
}

func discoverHostname() (string, error) {
	return os.Hostname()
}
