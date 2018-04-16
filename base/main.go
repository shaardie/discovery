package base

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"

	"github.com/shaardie/discovery/utils"
)

const (
	releaseFile = "/etc/os-release"
)

type osRelease struct{}

func (o osRelease) Description() string {
	return fmt.Sprintf("Operating System Release")
}

func (o osRelease) Run() (string, error) {
	out, err := ioutil.ReadFile(releaseFile)
	if err != nil {
		return fmt.Sprintf("failed, %v", err), nil
	}

	env, err := godotenv.Unmarshal(string(out))
	if err != nil {
		return "", fmt.Errorf("unable to source, %v", err)
	}

	return env["NAME"], nil
}

type initSystem struct{}

func (i initSystem) Description() string {
	return fmt.Sprintf("Init System")
}

func (i initSystem) Run() (string, error) {
	f := "/proc/1/comm"
	out, err := ioutil.ReadFile(f)
	if err != nil {
		return "", fmt.Errorf("unable to read '%v', %v", f, err)

	}

	return strings.TrimSuffix(string(out), "\n"), nil
}

type kernel struct{}

func (k kernel) Description() string {
	return fmt.Sprintf("Kernel")
}

func (k kernel) Run() (string, error) {
	out, err := exec.Command("uname", "--all").Output()
	if err != nil {
		return "", fmt.Errorf("unable to exec uname, %v", err)
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}

type logger struct{}

func (l logger) Description() string {
	return fmt.Sprintf("Logger")
}

func (l logger) Run() (string, error) {
	var result []string
	for _, logger := range []string{"systemd-journal", "rsyslogd", "syslog-ng"} {
		running, err := runningProcesses(logger)
		if err != nil {
			return "", fmt.Errorf("discovering Logger failed, %v", err)
		}
		if running {
			result = append(result, logger)
		}
	}
	return strings.Join(result, ","), nil
}

func discoverHostname() (string, error) {
	return os.Hostname()
}

func Tests() (tests []utils.Test) {
	tests = append(tests, osRelease{})
	tests = append(tests, initSystem{})
	tests = append(tests, kernel{})
	tests = append(tests, logger{})
	return
}
