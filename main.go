package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var (
	host = "example.com"
)

func topic(s string) {
	t := fmt.Sprintf("%v:", s)
	underline := strings.Repeat("=", utf8.RuneCountInString(t))
	fmt.Printf("\n\n%v\n%v\n\n", t, underline)
}

func line(k string, v string, err error) {
	if err != nil {
		fmt.Printf("%v: failed (%v)\n", k, err)
		return
	}

	if v == "" {
		v = "success"
	}

	fmt.Printf("%v: %v\n", k, v)
}

func lines(k string, vs []string, err error) {
	if vs == nil {
		return
	}
	if err != nil {
		line(k, "", err)
		return
	}

	if len(vs) == 1 {
		line(k, vs[0], err)
		return
	}

	fmt.Printf("%v:\n", k)
	for _, v := range vs {
		fmt.Printf(" - %v\n", v)
	}
}

func main() {

	topic("Operating System")

	os, err := discoverOS()
	line("OS", os, err)

	init, err := discoverInit()
	line("Init", init, err)

	kernel, err := discoverKernel()
	line("Kernel", kernel, err)

	logger, err := discoverLogger()
	lines("Logger", logger, err)

	hostname, err := discoverHostname()
	line("Hostname", hostname, err)

	topic("Network")

	interfaces, err := discoverInterfaces()
	lines("Network Interfaces", interfaces, err)

	err = discoverDNS(host)
	line("DNS", "", err)

	IPs, err := discoverAddresses()
	lines("IPs", IPs, err)

	line("Internet", "", discoverInternet(host))
}
