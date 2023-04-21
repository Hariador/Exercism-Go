package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

type phone struct {
	areaCode string
	exchange string
	sub      string
}

func Number(phoneNumber string) (string, error) {
	p, err := toPhone(phoneNumber)
	if err != nil {
		return "", err
	}
	return p.areaCode + p.exchange + p.sub, nil
}

func AreaCode(phoneNumber string) (string, error) {
	p, err := toPhone(phoneNumber)
	if err != nil {
		return "", err
	}
	return p.areaCode, nil
}

func Format(phoneNumber string) (string, error) {
	p, err := toPhone(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%v) %v-%v", p.areaCode, p.exchange, p.sub), nil
}

func toPhone(p string) (phone, error) {

	area := ""
	exchange := ""
	sub := ""
	var scrubbed []rune
	for _, v := range p {
		if unicode.IsDigit(v) {
			scrubbed = append(scrubbed, v)
		}
	}
	if len(scrubbed) > 11 || len(scrubbed) < 10 {
		return phone{}, errors.New("invalid length")
	}
	if len(scrubbed) == 11 {
		if scrubbed[0] == '1' {
			scrubbed = scrubbed[1:]
		} else {
			return phone{}, errors.New("country code mush be 1")
		}
	}
	if scrubbed[0] == '1' || scrubbed[0] == '0' {
		return phone{}, errors.New("area code cannot start with 0 or 1")
	}
	area = string(scrubbed[0]) + string(scrubbed[1]) + string(scrubbed[2])
	if scrubbed[3] == '1' || scrubbed[3] == '0' {
		return phone{}, errors.New("exchange cannot start with 0 or 1")
	}

	exchange = string(scrubbed[3]) + string(scrubbed[4]) + string(scrubbed[5])
	sub = string(scrubbed[6]) + string(scrubbed[7]) + string(scrubbed[8]) + string(scrubbed[9])

	t := phone{areaCode: area, exchange: exchange, sub: sub}
	fmt.Println(fmt.Sprintf("(%v) %v-%v", t.areaCode, t.exchange, t.sub))

	return phone{areaCode: area, exchange: exchange, sub: sub}, nil
}
