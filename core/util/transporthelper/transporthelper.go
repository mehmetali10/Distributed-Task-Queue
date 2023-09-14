package transporthelper

import "regexp"

func IsPhoneValid(phoneNumber string) bool {

	phonePattern := `^\+90\d{10}$`

	regex := regexp.MustCompile(phonePattern)

	return regex.MatchString(phoneNumber)

}
