package brackets

func Bracket(str string) (bool, error) {

	if len(str) == 0 {
		return true, nil
	}

	var closed [100]bool
	var lit [100]rune

	for i, letter := range str {
		if letter == '(' || letter == '[' || letter == '{' {
			// Initialize bracket is open
			closed[i] = true
			lit[i] = letter

		} else if letter == ')' || letter == ']' || letter == '}' {

			lit[i] = letter
			closed[i] = true

			for k := i - 1; k >= 0; k-- {
				// Check if there is open bracket of the type
				if (letter-lit[k] == 1 || letter-lit[k] == 2) && closed[k] == true {
					// Check if anything in between is open, if so return false
					for j := k + 1; j < i; j++ {
						if closed[j] == true {
							return false, nil
						}
					}
					// Else close open bracket
					closed[k] = false
					closed[i] = false
					break
				}
			}
		}
	}
	// Check if there is unclosed brackets
	for i := 0; i < len(str); i++ {
		if closed[i] == true {
			return false, nil
		}
	}
	return true, nil
}

// shysh
