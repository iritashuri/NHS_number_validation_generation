package nhshandling

import (
	"fmt"
	"math/rand"
	"strconv"
	"unicode"
)

type Service interface {
	GenerateNHS() string
	ValidateNHS(nhsNumber string) (bool, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GenerateNHS() string {
	// generate 9 random digits
	digits := make([]int, 10)
	for i := 0; i < 9; i++ {
		digits[i] = rand.Intn(10)
	}

	// calculate the sum of the multiply first 9 digits
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}

	// calculate the checkDigit
	digits[9] = 11 - (sum % 11)

	// if checkDigit is 10 then its invalid, generate new number
	if digits[9] == 10 {
		return s.GenerateNHS()
	}

	// if checkDigit the insert 0 instead in order to make the number valid
	if digits[9] == 11 {
		digits[9] = 0
	}

	// construct the NHS number
	nhsNumber := ""
	for _, digit := range digits {
		nhsNumber = fmt.Sprintf("%s%s", nhsNumber, strconv.Itoa(digit))
	}

	return nhsNumber
}

func (s *service) ValidateNHS(nhsNumber string) (bool, error) {
	if len(nhsNumber) != 10 {
		return false, nil
	}

	sum := 0
	// go over all characters
	for i, char := range nhsNumber {
		// validate each character is a valid digit (0-9)
		if !unicode.IsDigit(char) {
			return false, nil
		}

		// calculate the sum of the multiply first 9 digits
		if i != 9 {
			digit, err := strconv.Atoi(string(nhsNumber[i]))
			if err != nil {
				return false, err
			}

			sum += digit * (10 - i)
		}
	}

	// calculate the reminder from the multiply first 9 digits sum
	remainder := 11 - sum%11

	// get the checkDigit
	checkDigit, err := strconv.Atoi(string(nhsNumber[9]))
	if err != nil {
		return false, err
	}

	// if the reminder is 11 and checkDigit is 0 then its valid, else its invalid
	if remainder == 11 {
		return checkDigit == 0, nil
	}

	// if remainder = 10 (then it will not be equal to checkDigit which is 0-9) or if remainder != checkDigit then it's invalid
	return remainder == checkDigit, nil
}
