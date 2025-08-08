package phonenumber

import (
	"errors"
	"strconv"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	cleaned := cleanNumber(phoneNumber)
	_, err := strconv.Atoi(cleaned)
	if err != nil {
		return "", errors.New("number contains invalid characters")
	}
	if len(cleaned) < 10 || len(cleaned) > 11 {
		return "", errors.New("invalid length")
	}
	if len(cleaned) == 11 {
		if string(cleaned[0]) != "1" {
			return "", errors.New("invalid country code")
		}
		cleaned = cleaned[1:]
	}
	if string(cleaned[0]) < "2" {
		return "", errors.New("invalid area code")
	}
	if string(cleaned[3]) < "2" {
		return "", errors.New("invalid local number")
	}
	return cleaned, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + number[0:3] + ") " + number[3:6] + "-" + number[6:], nil
}

// remove redundant characters
func cleanNumber(phoneNumber string) string {
	r := strings.NewReplacer(" ", "", "+", "", ".", "", "-", "", "(", "", ")", "")
	return r.Replace(phoneNumber)
}
