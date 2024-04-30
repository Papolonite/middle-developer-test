package lib

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	emailRegex, err := regexp.Compile(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`)

	if err != nil {
		return err
	}

	isEmail := emailRegex.MatchString(email)
	if !isEmail {
		return errors.New("invalid email format")
	}
	return nil
}
