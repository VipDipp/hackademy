package downcase

func Downcase(str string) (string, error) {
	s := []rune(str)
	for i, el := range str {
		if el >= 'A' && el <= 'Z' {
			s[i] = el - 'A' + 'a'
		}
	}
	return string(s), nil
}
