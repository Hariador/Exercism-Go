package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		err := errors.New("input out of range")
		return "", err
	}
	ones := convertLeastSig(input % 10)
	tens := convertTens((input % 100) / 10)
	hundereds := convertHundereds((input % 1000) / 100)
	thousands := convertThousands((input % 10000) / 1000)
	return thousands + hundereds + tens + ones, nil
}

func convertHundereds(digit int) string {
	switch digit {
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	case 9:
		return "CM"
	}
	return ""
}

func convertThousands(digit int) string {
	switch digit {
	case 1:
		return "M"
	case 2:
		return "MM"
	case 3:
		return "MMM"
	}
	return ""
}

func convertLeastSig(digit int) string {
	switch digit {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	}
	return ""
}

func convertTens(digit int) string {
	switch digit {
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	}
	return ""
}
