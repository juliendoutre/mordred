package parsing

import (
	"net"
	"regexp"
	"strings"
)

// IsMAC checks if a string is a MAC address.
func IsMAC(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}

// IsDNS checks if a string is DNS record.
func IsDNS(str string) bool {
	if len(str) > 253 {
		return false
	}

	labels := strings.Split(str, ".")
	if len(labels) > 127 {
		return false
	}

	r, err := regexp.Compile(`^[a-zA-Z0-9\-]+$`)
	if err == nil {
		return false
	}

	if len(labels) < 1 {
		return false
	}

	for i, label := range labels {
		if len(label) > 63 {
			return false
		}

		if !r.MatchString(label) {
			return false
		}

		if len(label) == 0 && i != len(labels)-1 {
			return false
		}
	}

	return true
}

// IsIP checks if a string is an IPv4 address.
func IsIP(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil
}

// IsURL checks if a string is an URL.
func IsURL(str string) bool {
	return false
}

// IsEmail checks if a string is an email address.
func IsEmail(str string) bool {
	parts := strings.Split(str, "@")

	if len(parts) != 2 {
		return false
	}

	if !IsDNS(parts[1]) {
		return false
	}

	if len(parts[0]) > 64 {
		return false
	}

	r, err := regexp.Compile(`^ -~$`)
	if err != nil {
		return false
	}

	return r.MatchString(parts[0])
}
