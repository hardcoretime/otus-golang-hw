package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

// isNumInString checks 'str' contains a number.
func isNumInString(str string) bool {
	runes := []rune(str)

	for i := 0; i < len(runes)-1; i++ {
		current := runes[i]
		next := runes[i+1]
		_, errCurrent := strconv.Atoi(string(current))
		_, errNext := strconv.Atoi(string(next))

		if errCurrent == nil && errNext == nil {
			return true
		}
	}

	return false
}

// stringStartWithDigit returns true if 'str' starts with digit.
func stringStartWithDigit(str string) bool {
	if len(str) == 0 {
		return false
	}

	runes := []rune(str)

	if _, err := strconv.Atoi(string(runes[0])); err == nil {
		return true
	}

	return false
}

func Unpack(packedString string) (string, error) {
	result := strings.Builder{}

	if len(packedString) == 0 {
		return result.String(), nil
	}

	if stringStartWithDigit(packedString) {
		return result.String(), ErrInvalidString
	}

	if isNumInString(packedString) {
		return result.String(), ErrInvalidString
	}

	runes := []rune(packedString)
	for i := 0; i < len(runes)-1; i++ {
		current := runes[i]
		next := runes[i+1]

		_, errCurrent := strconv.Atoi(string(current))
		nextInt, errNext := strconv.Atoi(string(next))

		if errCurrent != nil && errNext == nil {
			result.WriteString(strings.Repeat(string(current), nextInt))
		}

		if errCurrent != nil && errNext != nil {
			result.WriteString(string(current))
		}
	}

	lastRune := runes[len(runes)-1]
	if _, err := strconv.Atoi(string(lastRune)); err != nil {
		result.WriteString(string(lastRune))
	}

	return result.String(), nil
}
