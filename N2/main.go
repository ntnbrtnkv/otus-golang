package n2

import (
	"errors"
	"strings"
)

func Unpack(s string) (string, error) {
	runes := []rune(s)
	runes = append(runes, rune('a'))
	l := len(runes)
	var res strings.Builder
	var num int = 0
	var last rune = -1

	for i := 0; i < l; i++ {
		c := runes[i]

		if c >= '0' && c <= '9' {
			if num == -1 {
				num = 0
			}
			num = num*10 + int(c) - '0'
		} else {
			if num == -1 {
				if last == -1 {
					return "", errors.New("not correct string")
				}
				res.WriteRune(last)
			} else if num != 0 {
				if last == -1 {
					return "", errors.New("not correct string")
				}
				lastString := string(last)
				res.WriteString(strings.Repeat(lastString, num))
			}

			last = c
			num = -1

			if c == '\\' {
				if i+1 >= l {
					return "", errors.New("not correct string")
				}
				i++
				last = runes[i]
			}
		}
	}

	return res.String(), nil
}
