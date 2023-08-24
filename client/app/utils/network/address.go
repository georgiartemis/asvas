package network

import (
	"io/ioutil"
	"net/http"
	"net"
)

func GetMacAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var address []string
	for _, i := range interfaces {
		a := i.HardwareAddr.String()
		if a != "" {
			address = append(address, a)
		}
	}
	if len(address) == 0 {
		return "", nil
	}
	return address[0], nil
}

func GetLocalIP() (string, error) {
	resp, err := http.Get("https://icanhazip.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ipBytes), nil
}
func Pablikip() (net.IP, error) {
	conn, err := net.Dial(`udp`, `8.8.8.8:80`)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}