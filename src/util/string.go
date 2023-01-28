package util

func Contains(str string, items []string) bool {
	for _, v := range items {
		if v == str {
			return true
		}
	}

	return false
}
