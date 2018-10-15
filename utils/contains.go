package utils

// n がスライスに含まれているか
func Contains(n string, xs []string) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}
