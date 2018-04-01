package main

import (
	"fmt"
	"net"
	"net/http"
)

func discoverInternet(host string) error {
	url := fmt.Sprintf("http://%v", host)
	_, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to GET '%v', %v", host, err)
	}
	return nil
}

func discoverAddresses() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, fmt.Errorf("unable to get IPs, %v", err)
	}
	var result []string
	for _, addr := range addrs {
		result = append(result, addr.String())
	}
	return result, nil

}

func discoverInterfaces() ([]string, error) {

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("unable to get interfaces, %v", err)
	}

	result := make([]string, len(interfaces))
	for index, ifi := range interfaces {
		result[index] = ifi.Name
		if ifi.HardwareAddr != nil {
			result[index] = fmt.Sprintf("%v, %v", result[index], ifi.HardwareAddr)
		}
		result[index] = fmt.Sprintf("%v, %v", result[index], ifi.Flags)
	}
	return result, nil
}

func discoverDNS(host string) error {
	_, err := net.LookupHost(host)
	if err != nil {
		return fmt.Errorf("unable to resolve %v, %v", host, err)
	}
	return nil
}