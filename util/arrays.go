package util

// ArrayRemove removes an element at index i from an array of strings a
func ArrayRemove(a []string, i int) ([]string, string) {
	a[len(a)-1], a[i] = a[i], a[len(a)-1]
	return a[:len(a)-1], a[len(a)-1]
}
