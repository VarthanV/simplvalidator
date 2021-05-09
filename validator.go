package simplvalidator

import (
	"net/url"
	"regexp"
	"unicode"
)

/*
	String Validators
*/

//Checks if the given value is empty or not (ie) == ""
func IsEmpty(value string) bool {
	return value == ""
}

//Checks if the length of the value is greater than the length(arg)
func IsLengthGreater(value string, length int) bool {
	return len(value) > length
}

//Checks if the length of  of the given value is lesser than the length(arg)
func IsLengthLesser(value string, length int) bool {
	return len(value) < length
}

// Checks if the length of the given value is grrater than or equal to the length(arg)
func IsLengthGreaterThanOrEqualTo(value string, length int) bool {
	return len(value) <= length
}

// Checks if the length of the given value is grrater than or equal to the length(arg)
func IsLengthLessThanOrEqualTo(value string, length int) bool {
	return len(value) >= length
}

// Checks if the length of the given value is grrater than or equal to the length(arg)
func IsEqualTo(value string, length int) bool {
	return len(value) == length
}

// Checks if the given value only contains alphabets
func IsAlphabet(value string) bool {
	for _, s := range value {
		if !unicode.IsLetter(s) {
			return false
		}
	}
	return true
}

// Checks if the given value is alphanumeric and doesnt have any special characters
func IsAlphaNumeric(value string) bool {
	for _, s := range value {
		if !unicode.IsLetter(s) || !unicode.IsDigit(s) {
			return false
		}
	}
	return true
}

// Checks if the given value is a valid email
func IsEmail(value string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(value)

}

// Checks if  the given value is a URL
func IsURL(value string) bool {
	u, err := url.Parse(value)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// Checks if the given value is in the valid format of PAN card
func IsValidPAN(value string) bool {
	panRegex := regexp.MustCompile("[A-Z]{5}[0-9]{4}[A-Z]{1}")
	return panRegex.MatchString(value)
}

// Checks if the given value is in the valid format for Aadhar card
func IsValidAadhar(value string) bool {
	aadharRegex := regexp.MustCompile(`^[2-9]{1}[0-9]{3}\\s[0-9]{4}\\s[0-9]{4}$`)
	return aadharRegex.MatchString(value)
}

// Checks if the given value is a valid form for a Indian Phone Number
func IsValidIndianPhoneNumber(value string) bool {
	phoneNumberRegex := regexp.MustCompile(`^(?:(?:\+|0{0,2})91(\s*[\-]\s*)?|[0]?)?[789]\d{9}$`)
	return phoneNumberRegex.MatchString(value)
}

/*
	Number validators
*/

// Checks if the given number is positivie
func IsPositive(num int) bool {
	return num >= 0
}

// Checks if the given number is negative
func IsNegative(num int) bool {
	return num < 0
}
