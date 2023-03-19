package iteration

// Contains reports whether character is within text
func Contains(text, character string) bool {
	for i := 0; i < len(text); i++ {
		if string(text[i]) == character {
			return true
		}
	}
	return false
}
