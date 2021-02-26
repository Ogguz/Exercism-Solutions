package etl

import "strings"

// Transform receives map with int key and string value and transforms it to map with string key and int value
func Transform(original map[int][]string) (transformed map[string]int) {
	transformed = make(map[string]int)

	for score, chars := range original {
		for _, char := range chars {
			transformed[strings.ToLower(char)] = score
		}
	}
	return transformed
}
