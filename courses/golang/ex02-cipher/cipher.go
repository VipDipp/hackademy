package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type caesar struct{}

func NewCaesar() Cipher { return caesar{} }

type shift struct {
	num int
}

func NewShift(num int) Cipher {
	if num == 0 || num > 25 || num < (-25) {
		return nil
	}
	return shift{num: num}
}

type vigenere struct {
	str string
}

func NewVigenere(str string) Cipher {

	if str == "" {
		return nil
	}
	temp := str
	str = ""
	if temp == "a" || temp == "aa" {
		return nil
	}
	for _, i := range temp {

		if i > 96 && i < 123 {
			str += string(i)
		}
	}
	if temp == str {
		return vigenere{str: str}
	}
	return nil
}

func (c caesar) Encode(str string) (output string) {
	var tosimple string
	for _, i := range str {
		if i >= 65 && i <= 90 {
			tosimple += string(i + 32)
		} else if i >= 97 && i <= 122 {
			tosimple += string(i)
		}
	}
	output = ""
	for _, i := range tosimple {
		if i+3 > 'z' {
			output += string(i + 3 - 26)
		} else {
			output += string(i + 3)
		}
	}

	return output
}

func (c caesar) Decode(str string) (output string) {
	for _, i := range str {
		if i-3 >= 'a' {
			output += string(i - 3)
		} else {
			output += string(i - 3 + 26)
		}
	}
	return output
}

func (s shift) Encode(str string) (output string) {
	var tosimple string
	for _, i := range str {
		if i >= 65 && i <= 90 {
			tosimple += string(i + 32)
		} else if i >= 97 && i <= 122 {
			tosimple += string(i)
		}
	}
	output = ""
	for _, i := range tosimple {
		if i+rune(s.num) > 'z' {
			output += string(i + rune(s.num) - 26)
		} else if i+rune(s.num) < 'a' {
			output += string(i + rune(s.num) + 26)
		} else {
			output += string(i + rune(s.num))
		}
	}

	return output
}

func (s shift) Decode(str string) (output string) {
	for _, i := range str {
		if i-rune(s.num) < 'a' {
			output += string(i - rune(s.num) + 26)
		} else if i-rune(s.num) > 'z' {
			output += string(i - rune(s.num) - 26)
		} else {
			output += string(i - rune(s.num))
		}
	}
	return output
}

func (v vigenere) Encode(str string) (output string) {
	for _, i := range str {
		if i >= 65 && i <= 90 {
			output += string(i + 32)
		} else if i > 96 && i < 123 {
			output += string(i)
		}
	}
	length := len(v.str)
	temp := output
	output = ""

	for k, i := range temp {
		key := string(rune(((int(v.str[k%length]) + int(i) - 64) % 26) + 97))
		output += key

	}
	return output
}

func (v vigenere) Decode(str string) (output string) {
	length := len(v.str)
	output = ""
	for i := 0; i < len(str); i++ {
		key := string(rune((int(str[i])+26-int(v.str[i%length]))%26 + 97))
		output += key
	}
	return output
}
