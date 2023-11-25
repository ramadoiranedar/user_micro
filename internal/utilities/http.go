package utilities

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

func GetClientIPv4(request *http.Request) (ip string, err error) {
	ip = request.Header.Get("X-Real-IP") // some proxies
	if ip == "" {
		ip = request.Header.Get("X-Forwarded-For") // some proxies
	}

	if ip == "" {
		ip = request.RemoteAddr
	}

	ip, _, err = net.SplitHostPort(ip)
	if err != nil {
		return "", err
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil || parsedIP.To4() == nil {
		return "", fmt.Errorf("not a valid IPv4 address: %s", ip)
	}

	return
}

func ReadRequestBodyApplicationJson(r *http.Request) (bodyString string) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		bodyString = string(body)
	}

	return
}
