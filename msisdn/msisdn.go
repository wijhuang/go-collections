package msisdn

import (
	"bytes"
)

func cleanSpaceAndChar(s string) string {
	data := bytes.Buffer{}
	for i := 0; i < len(s); i++ {
		b := s[i]
		if '0' <= b && b <= '9' {
			data.WriteByte(b)
		}
	}
	return data.String()

}

// make format msisdn 628123456789
func NormalizeMsisdn(msisdn string) string {
	s := cleanSpaceAndChar(msisdn)

	if len(s) < 2 {
		return ""
	}

	switch s[0] {
	case '0':
		return "62" + s[1:]
	case '2', '8':
		return "62" + s
	}
	return s
}
