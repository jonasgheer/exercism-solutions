package romannumerals

import "errors"

type numeral struct {
	decimal int
	roman   string
}

var decimalToRoman = [13]numeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input >= 4000 || input <= 0 {
		return "", errors.New("input cannot be represented in roman numerals")
	}
	output := ""
	for _, numeral := range decimalToRoman {
		for input >= numeral.decimal {
			output += numeral.roman
			input -= numeral.decimal
		}
	}
	return output, nil
}
