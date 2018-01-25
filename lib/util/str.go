package util

func CheckContains(key string, set []string) bool {
	for i := 0; i < len(set); i++ {
		if key == set[i] {
			return true
		}
	}
	return false
}
