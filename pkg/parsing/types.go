package parsing

// StringType defines the meaning of a string.
type StringType string

// Following constants define types of strings that bear a specific meaning.
const (
	MAC         StringType = "mac"
	DNSRecord   StringType = "dns"
	IP          StringType = "ip"
	URL         StringType = "url"
	EMail       StringType = "e-mail"
	Binary      StringType = "binary"
	Hexadecimal StringType = "hex"
	Base64      StringType = "base64"
	Decimal     StringType = "decimal"
	Printable   StringType = "printable"
	Unknown     StringType = "unknown"
)

// GetType returns the type of a string.
func GetType(str string) StringType {
	if IsMAC(str) {
		return MAC
	} else if IsDNS(str) {
		return DNSRecord
	} else if IsIP(str) {
		return IP
	} else if IsURL(str) {
		return URL
	} else if IsEmail(str) {
		return EMail
	} else if IsBinary(str) {
		return Binary
	} else if IsDecimal(str) {
		return Decimal
	} else if IsHex(str) {
		return Hexadecimal
	} else if IsBase64(str) {
		return Base64
	} else if IsPrintable(str) {
		return Printable
	} else {
		return Unknown
	}
}
