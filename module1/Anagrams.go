package module1

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func Anagrams() {
	name1 := "Listen"
	name1 = strings.ToLower(name1)
	name2 := "silent"
	name2 = strings.ToLower(name2)
	name1slice := []rune(name1)
	name2slice := []rune(name2)
	slices.Sort(name1slice)
	sort.Slice(name2slice, func(i, j int) bool { return name2slice[i] < name2slice[j] })
	if string(name1slice) == string(name2slice) {
		fmt.Println("true")

	}
}
