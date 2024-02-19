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
		currentIsDigit := errCurrent == nil

		_, errNext := strconv.Atoi(string(next))
		nextIsDigit := errNext == nil

		if currentIsDigit && nextIsDigit {
			return true
		}
	}

	return false
}

func Unpack(packedString string) (string, error) {
	result := strings.Builder{}
	runes := []rune(packedString)

	if len(packedString) == 0 {
		return result.String(), nil
	}

	if isNumInString(packedString) {
		return result.String(), ErrInvalidString
	}

	for i := 0; i < len(runes)-1; i++ {
		currentRune := runes[i]
		nextRune := runes[i+1]

		_, errCurrent := strconv.Atoi(string(currentRune))
		currentRuneIsString := errCurrent != nil

		if i == 0 && !currentRuneIsString {
			return result.String(), ErrInvalidString
		}

		nextDigit, errNext := strconv.Atoi(string(nextRune))
		nextRuneIsDigit := errNext == nil

		if currentRuneIsString && nextRuneIsDigit {
			result.WriteString(strings.Repeat(string(currentRune), nextDigit))
		}

		if currentRuneIsString && !nextRuneIsDigit {
			result.WriteString(string(currentRune))
		}
	}

	_, err := strconv.Atoi(string(runes[len(runes)-1]))
	lastRuneIsString := err != nil
	if lastRuneIsString {
		result.WriteString(string(runes[len(runes)-1]))
	}

	return result.String(), nil
}
