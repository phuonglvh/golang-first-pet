package utils

import (
	"net"
	"net/http"
)

// GetMyIP return current machine ip
func GetMyIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetScheme return scheme of the current request
func GetScheme(r *http.Request) string {
	if r.TLS == nil {
		return "http"
	}
	return "https"
}
