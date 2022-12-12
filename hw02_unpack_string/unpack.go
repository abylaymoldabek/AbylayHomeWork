package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var symbol rune
	str := ""
	checkFistSymbol := true
	checkSecondNumber := false
	checkException := false
	exeption := `\`
	for _, v := range s {
		switch {
		case checkFistSymbol:
			if 48 <= v && 57 >= v {
				return "", ErrInvalidString
			}
			symbol = v
			str += string(symbol)
			checkFistSymbol = false
		case string(v) == exeption && !checkException:
			checkException = true
			continue
		case 48 < v && 57 >= v && !checkException:
			if checkSecondNumber {
				return "", ErrInvalidString
			}
			checkSecondNumber = true
			for j := 0; j < int(v-48)-1; j++ {
				str += string(symbol)
			}
		case 48 < v && v <= 57 && checkException:
			symbol = v
			str += string(symbol)
			checkException = false
		case v == 48:
			if checkSecondNumber {
				return "", ErrInvalidString
			}
			str = str[:len(str)-1]
		case string(v) == exeption && checkException:
			checkException = false
			symbol = v
			str += string(symbol)
		default:
			symbol = v
			str += string(symbol)
			checkSecondNumber = false
		}
	}
	return str, nil
}
