package cipher

import "unicode"

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
	shiftDistance int32
}

type vigenere struct {
	key []int32
}

const min = 97
const max = 122

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 {
		return nil
	}
	if distance > 25 || distance < -25 {
		return nil
	}
	return shift{shiftDistance: int32(distance)}
}

func (c shift) Encode(input string) string {
	output := ""
	for _, v := range input {
		if unicode.IsLetter(v) {
			output = output + string(adjust(unicode.ToLower(v)+c.shiftDistance))
		}
	}
	return output
}

func (c shift) Decode(input string) string {
	output := ""
	for _, v := range input {
		if unicode.IsLetter(v) {
			output = output + string(adjust(unicode.ToLower(v)-c.shiftDistance))
		}
	}
	return output
}

func NewVigenere(key string) Cipher {
	var kV []int32
	allA := true
	for _, v := range key {
		if unicode.IsLetter(v) && unicode.IsLower(v) {
			t := v - 97
			kV = append(kV, t)
			if allA {
				if t != 0 {
					allA = false
				}
			}
		} else {
			return nil
		}
	}
	if allA {
		return nil
	}

	return vigenere{key: kV}
}

func (v vigenere) Encode(input string) string {
	output := ""
	index := 0
	for _, letter := range input {
		if unicode.IsLetter(letter) {
			output = output + string(adjust(unicode.ToLower(letter)+rune(v.key[index])))
			index++
			if index >= len(v.key) {
				index = 0
			}
		}
	}
	return output
}

func (v vigenere) Decode(input string) string {
	output := ""
	index := 0
	for _, letter := range input {
		if unicode.IsLetter(letter) {
			output = output + string(adjust(unicode.ToLower(letter)-rune(v.key[index])))
			index++
			if index >= len(v.key) {
				index = 0
			}
		}
	}
	return output
}

func adjust(i int32) int32 {
	if i < min {
		return i + 26
	} else if i > max {
		return i - 26
	}
	return i
}
