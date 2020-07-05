package parsing

import (
	"regexp"
	"unicode"
)

// IsBinary checks if a string is the binary representation of a number.
func IsBinary(str string) bool {
	r, err := regexp.Compile(`^[01]+$`)
	if err != nil {
		return false
	}

	return r.MatchString(str)
}

// IsDecimal checks if a string is the decimal representation of a number.
func IsDecimal(str string) bool {
	r, err := regexp.Compile(`^\d+$`)
	if err != nil {
		return false
	}

	return r.MatchString(str)
}

// IsHex checks if a string is the hex representation of a number.
func IsHex(str string) bool {
	r, err := regexp.Compile(`^[\dabcdef]+$`)
	if err != nil {
		return false
	}

	return r.MatchString(str)
}

// IsBase64 checks if a string is an encoded base64 payload.
func IsBase64(str string) bool {
	r, err := regexp.Compile(`^[0-9a-zA-Z+/=]+$`)
	// TODO: check padding size
	if err != nil {
		return false
	}

	return r.MatchString(str)
}

// IsPrintable checks if a string is composed only of human readable characters.
func IsPrintable(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, r := range str {
		if !unicode.IsPrint(r) {
			return false
		}
	}

	return true
}
