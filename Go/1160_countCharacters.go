package main

func countCharacters(words []string, chars string) int {
	chMap := make(map[byte]int)

	for i := range chars {
		chMap[chars[i]]++
	}

	count := 0
	for _, v := range words {
		if isContainer(v, chMap) {
			count += len(v)
		}
	}

	return count
}

func isContainer(pattern string, chMap map[byte]int) bool {
	for i := range pattern {
		chMap[pattern[i]]--
		if chMap[pattern[i]] < 0 {

			for j := i; j >= 0; j-- {
				chMap[pattern[j]]++
			}
			return false
		}
	}

	for i := range pattern {
		chMap[pattern[i]]++
	}
	return true
}

// func main() {
// 	words := []string{"cat", "bt", "hat", "tree"}
// 	chars := "atach"

// 	fmt.Println(countCharacters(words, chars))
// }
