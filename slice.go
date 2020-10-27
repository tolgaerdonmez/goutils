package github.com/tolgaerdonmez/goutils

// HasItem checks for item in given slice
func HasItem(list []string, items ...string) bool {
	v := 0
	for _, item := range items {
		for _, b := range list {
			if b == item {
				v++
			}
		}
	}
	if v > 0 {
		return true
	}
	return false
}
