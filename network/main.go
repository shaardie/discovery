package network

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/shaardie/discovery/utils"
)

const (
	url       = "http://example.com"
	host      = "example.com"
	dnsFile   = "/etc/resolv.conf"
	dnsPrefix = "nameserver "
)

type internet struct {
	url string
}

func (i internet) Description() string {
	return fmt.Sprintf("Connect to %v to verify internet connection", i.url)
}

func (i internet) Run() (string, error) {
	_, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("failed, %v", err), nil
	}
	return "ok", nil
}

type resolve struct {
	host string
}

func (r resolve) Description() string {
	return fmt.Sprintf("Resolve IP from %v", r.host)
}

func (r resolve) Run() (string, error) {
	_, err := net.LookupHost(host)
	if err != nil {
		return fmt.Sprintf("failed, %v", err), nil
	}
	return "ok", nil
}

type dns struct{}

func (d dns) Description() string {
	return fmt.Sprintf("Look for DNS entries")
}

func (d dns) Run() (string, error) {
	out, err := ioutil.ReadFile(dnsFile)
	if err != nil {
		return fmt.Sprintf("failed, %v", err), nil
	}
	var dns []string
	for _, entry := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(entry, dnsPrefix) {
			entry = strings.TrimRight(entry, "\n")
			entry = strings.TrimPrefix(entry, dnsPrefix)
			dns = append(dns, entry)
		}
	}
	return strings.Join(dns, ","), nil
}

func Tests() (tests []utils.Test) {
	tests = append(tests, dns{})
	tests = append(tests, resolve{host})
	tests = append(tests, internet{url})
	return
}
