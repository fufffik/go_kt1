//go:build ignore

package main

import (
	"fmt"
	"unicode"
)

func isUnique(s string) bool {
	seen := make(map[rune]struct{})
	for _, r := range s {
		r = unicode.ToLower(r)
		if _, ok := seen[r]; ok {
			return false
		}
		seen[r] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println("abcd:", isUnique("abcd"))           // true
	fmt.Println("abCdefAaf:", isUnique("abCdefAaf")) // false
	fmt.Println("aabcd:", isUnique("aabcd"))         // false
}