package base

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func runningProcesses(p string) (bool, error) {
	p = fmt.Sprintf("Name:\t%v", p)
	matches, err := filepath.Glob("/proc/[1-9]*")
	if err != nil {
		return false, fmt.Errorf("runningProcesses failed, %v", err)
	}

	for _, match := range matches {
		out, err := ioutil.ReadFile(filepath.Join(match, "status"))
		if err != nil {
			// maybe handle this error
			continue
		}
		for _, line := range strings.Split(string(out), "\n") {
			if p == line {
				return true, nil
			}
		}
	}
	return false, nil
}
