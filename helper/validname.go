package helper

func IsValidNameFormat(name string) bool {
	if name == "" {
		return false
	}

	first := rune(name[0])

	if first >= 'a' && first <= 'z' {
		return false
	}

	return true
}
