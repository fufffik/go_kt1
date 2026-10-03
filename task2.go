//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)
	for _, t := range temps {
		key := int(t/10) * 10 // int() отбрасывает дробную часть к нулю
		groups[key] = append(groups[key], t)
	}

	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys) // порядок ключей только для красивого вывода

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}