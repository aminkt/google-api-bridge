package lib

import (
	"errors"
	"net"
	"net/http"
	"strings"
)

func VerifyIsAllowedIPs(r *http.Request) (bool, error) {
	allowedIPsStr := ReadEnvironmentVariables().AllowedIps
	if allowedIPsStr == "" {
		return false, errors.New("allowed IPs not configured")
	}

	allowedIPs := strings.Split(allowedIPsStr, ",")
	clientIp, _ := extractIP(r.RemoteAddr)

	if clientIp == "" {
		return false, errors.New("Failed to get client's IP address")
	}

	for _, ip := range allowedIPs {
		if clientIp == ip {
			return true, nil
		}
	}
	return false, errors.New("Access Denied for " + r.RemoteAddr)
}

func extractIP(remoteAddr string) (string, error) {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return "", err
	}

	// Strip square brackets if present (IPv6)
	if len(host) > 0 && host[0] == '[' && host[len(host)-1] == ']' {
		host = host[1 : len(host)-1]
	}

	return host, nil
}
