package main

import "fmt"

func main() {

	s := "aaaaaaabpbppleee"
	dict := map[int]string{
		0: "able",
		1: "ale",
		2: "apple",
		3: "bale",
		4: "kangaroo",
	}

	longestWord := findLongestSubsequence(s, dict)
	fmt.Println(longestWord)
}

func findLongestSubsequence(s string, dict map[int]string) string {
	longestWord := ""

	for i := 0; i < len(dict); i++ {
		index := 0
		for x := 0; x < len(dict[i]); x++ {
			for y := index; y < len(s); y++ {
				if dict[i][x] == s[y] {
					index++
					break
				} else {
					index++
				}
			}
			if index == len(s) {
				break
			}
		}
		if index < len(s) {
			longestWord = dict[i]
		}
		index = 0
	}

	return longestWord
}
