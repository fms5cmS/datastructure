package string_match

func BF(main, pattern string) int {
	if len(main) < len(pattern) || len(main) == 0 || len(pattern) == 0 {
		return -1
	}
	for i := 0; i <= len(main)-len(pattern); i++ {
		str := main[i : i+len(pattern)]
		if pattern == str {
			return i
		}
	}
	return -1
}
