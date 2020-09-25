package types

import (
	"regexp"
	"strings"
)

var cellphoneReges = regexp.MustCompile(`^09\d{9}$`)

// Cellphone is the data-type for cellphone numbers.
type Cellphone string

// IsValid returns true if p is a valid cellphone number.
func (p Cellphone) IsValid() bool {
	return cellphoneReges.MatchString(p.Standard().String())
}

// Standard converts p to a standard format.
func (p Cellphone) Standard() Cellphone {
	if strings.HasPrefix(p.String(), "+") {
		if strings.HasPrefix(p.String(), "+98") {
			return "0" + p[3:]
		}
		return "00" + p[1:]
	}
	if strings.HasPrefix(p.String(), "0098") {
		return "0" + p[4:]
	}
	return p
}

// International converts p to an international format.
func (p Cellphone) International() Cellphone {
	if strings.HasPrefix(p.String(), "0") {
		if strings.HasPrefix(p.String(), "00") {
			return "+" + p[2:]
		}
		return "+98" + p[1:]
	}
	return p
}

func (p Cellphone) String() string {
	return string(p)
}
