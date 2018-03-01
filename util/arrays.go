package util

// ArrayRemove removes an element at index i from an array of strings a
func ArrayRemove(a []string, i int) ([]string, string) {
	a[len(a)-1], a[i] = a[i], a[len(a)-1]
	return a[:len(a)-1], a[len(a)-1]
}

// ArrayUnique returns an array with only unique values in it
func ArrayUnique(a []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
