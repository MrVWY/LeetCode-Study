package main

import "fmt"

func prefixTable(pattern string, prefix []int, size int) {
	prefix[0] = 0
	lens := 0 //最长前缀和
	index := 1
	for index < size {
		if pattern[index] == pattern[lens]{
			lens++
			prefix[index] = lens
			index++
		}else{
			if lens > 0 {
				lens = prefix[lens-1] // 当前位置的最长前缀没找到, 把prefix中原len指向len-1
			}else {
				//说明lens = 0
				prefix[index] = lens
				index++
			}
		}
	}
	moverPrefixtable(prefix, size)
}

func moverPrefixtable(prefix []int, size int)  {
	for i := size - 1 ; i>0 ; i-- {
		prefix[i] = prefix[i-1]
	}
	prefix[0] = -1
}

func kmpSearch(text string, pattern string) {
	size := len(pattern)
	prefix := make([]int, size)
	prefixTable(pattern, prefix, size)

	textIndex, patternIndex := 0, 0
	for textIndex < len(text) {
		if patternIndex == size - 1 && pattern[patternIndex] == text[textIndex] {
			fmt.Println("found pattern in ", textIndex-patternIndex)
			patternIndex = prefix[patternIndex] //or break
		}
		if text[textIndex] == pattern[patternIndex] {
			textIndex++
			patternIndex++
		}else {
			patternIndex = prefix[patternIndex]
			if patternIndex == -1 {
				textIndex++
				patternIndex++
			}
		}
	}
}

func main() {
	pattern := "ABABCABAA"
	text    := "ABABABCABAAABABCABAA"
	kmpSearch(text, pattern)
}