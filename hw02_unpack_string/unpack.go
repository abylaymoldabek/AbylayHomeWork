package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var concatResult strings.Builder
	symbol := []rune(str)
	exeption := '\\'

	if len(symbol) == 0 {
		return "", nil
	}
	if !unicode.IsLetter(symbol[0]) {
		return "", ErrInvalidString
	}
	for i := 0; i < len(symbol); i++ {
		if unicode.IsDigit(symbol[i]) {
			return "", ErrInvalidString
		}
		if i+1 < len(symbol) && symbol[i] == exeption {
			i++
			if !unicode.IsDigit(symbol[i]) && symbol[i] != exeption {
				return "", ErrInvalidString
			}
		}
		if i+1 < len(symbol) && unicode.IsDigit(symbol[i+1]) {
			count, _ := strconv.Atoi(string(symbol[i+1]))
			if count == 0 {
				concatResult.WriteString(strings.TrimRight(string(symbol[i]), string(symbol[i])))
			}
			concatResult.WriteString(strings.Repeat(string(symbol[i]), count))
			i++
		} else {
			concatResult.WriteString(string(symbol[i]))
		}
	}
	return concatResult.String(), nil
}
