package util

import (
	"sort"
	"strings"
)

// WordMatchesFill determines if a word matches a "fill string". E.g. "style" matches "s y  " but not " sy  "
func WordMatchesFill(word, fill string) bool {
	word = strings.ToLower(word)
	fill = strings.ToLower(fill)

	if len(word) != len(fill) {
		return false
	}

	for i := 0; i < len(fill); i++ {
		if fill[i:i+1] != " " && word[i:i+1] != fill[i:i+1] {
			return false
		}
	}

	return true
}

// WordMatchesSet determines if a word can be built with a given set of characters
func WordMatchesSet(word, set string, reuse bool) bool {
	split := strings.Split(set, "")
	// fmt.Printf("word %v matches set %v reuse %v\n", word, set, reuse)
	// for each letter in the word
	for _, c := range word {
		// is the letter in the set?
		found := false
		for j := range split {
			if split[j] == string(c) {
				found = true
				if !reuse {
					split, _ = ArrayRemove(split, j)
				}
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// WordIsAnagram determines if two words are anagrams of each other
func WordIsAnagram(word, anagram string) bool {
	return NormalizeString(word) == NormalizeString(anagram)
}

// NormalizeString turns a string to lowercase and sorts it alphabetically
func NormalizeString(word string) string {
	letters := []string{}
	for _, letter := range word {
		letters = append(letters, strings.ToLower(string(letter)))
	}
	sort.Strings(letters)
	return strings.Join(letters, "")
}
