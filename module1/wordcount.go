package module1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordCount() {
	fmt.Println("Enter a runes:")
	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n') // Reads the full input line
	fmt.Println("", name)
	name = strings.TrimSpace(name) // Remove trailing newline
	fmt.Println("", name)
	list := strings.Fields(name)
	fmt.Println("list:", list, "length:", len(list))
}
func wordcountandfrequency() {
	line := "Compares the first and last characters, then moves inward."
	slice1 := strings.Fields(line)
	fmt.Println("", len(slice1))
	fmt.Println("", slice1)
	map1 := make(map[string]int)
	for _, word := range slice1 {
		map1[word]++
	}
	fmt.Println("", map1)
}
