package helper

import (
	"math"
	"math/rand"
	"regexp"
	"time"
)

func ValidatePhoneNumber(phoneNumber string) (string, bool) {
	regex := `^(08|8|628|\+628)\d{5,12}$`
	match, _ := regexp.MatchString(regex, phoneNumber)

	if match {
		if phoneNumber[0:1] == "8" {
			phoneNumber = "62" + phoneNumber
		} else if phoneNumber[0:2] == "08" {
			phoneNumber = "62" + phoneNumber[1:]
		} else if phoneNumber[0:3] == "+62" {
			phoneNumber = "62" + phoneNumber[3:]
		}
	}

	return phoneNumber, match
}

func GenerateRandomNumber(digits int) int {
	min := int(math.Pow(10, float64(digits-1)))
	max := int(math.Pow(10, float64(digits))) - 1

	return rand.Intn(max-min+1) + min
}

func ParseDate(sdate string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", sdate)

	if err != nil {
		return parsedDate, err
	}
	return parsedDate, nil
}
