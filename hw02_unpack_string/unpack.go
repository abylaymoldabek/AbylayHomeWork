package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const (
	backslash = '\\'
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	concatResult := strings.Builder{}

	symbols := []rune(str)

	if len(symbols) == 0 {
		return "", nil
	}

	for i := 0; i < len(symbols); i++ {
		if unicode.IsDigit(symbols[i]) {
			return "", ErrInvalidString
		}
		if i+1 < len(symbols) && symbols[i] == backslash {
			i++
			if !unicode.IsDigit(symbols[i]) && symbols[i] != backslash {
				return "", ErrInvalidString
			}
		}
		if i+1 < len(symbols) && unicode.IsDigit(symbols[i+1]) {
			count, _ := strconv.Atoi(string(symbols[i+1]))
			if count == 0 {
				concatResult.WriteString(strings.TrimRight(string(symbols[i]), string(symbols[i])))
			}
			concatResult.WriteString(strings.Repeat(string(symbols[i]), count))
			i++
		} else {
			concatResult.WriteString(string(symbols[i]))
		}
	}
	return concatResult.String(), nil
}
