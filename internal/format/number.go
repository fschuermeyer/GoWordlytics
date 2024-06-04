package format

import (
	"strconv"
)

func InsertThousandSeparator(number int, separator rune) string {
	var formatedNumber []rune
	stringNumber := strconv.Itoa(number)

	for i, char := range reverseString(stringNumber) {
		formatedNumber = append([]rune{char}, formatedNumber...)

		if (i%3) == 2 && len(stringNumber)-1 > i {
			formatedNumber = append([]rune{separator}, formatedNumber...)
		}
	}

	return string(formatedNumber)
}

func reverseString(str string) string {
	var rstr string

	for _, char := range str {
		rstr = string(char) + rstr
	}

	return rstr
}
