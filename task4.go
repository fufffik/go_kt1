//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s) // делит по пробелам, лишние игнорирует
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("snow dog sun")) // sun dog snow
}