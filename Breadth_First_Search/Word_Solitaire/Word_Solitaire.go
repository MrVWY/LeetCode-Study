package Word_Solitaire

//127. 单词接龙
//BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dictionary := make(map[string]bool,len(wordList))
	for i := 0 ; i < len(wordList) ; i++ {
		dictionary[wordList[i]] = true
	}

	//判断endword是否在dictionary中
	if _, ok := dictionary[endWord] ; !ok {
		return 0
	}

	//开始BFS
	queue := []string{}
	queue = append(queue,beginWord)

	step := 0
	length := len(beginWord)
	for len(queue) > 0 {
		step++
		size := len(queue)
		for i := size ; i >0 ; i-- {
			vlaue := queue[0]
			queue = queue[1:]
			chs := []rune(vlaue)
			//开始替换
			for i := 0 ; i < length ; i++ {
				a := chs[i]
				for c := 'a'; c <= 'z'; c++ {
					if c == a {
						continue
					}
					chs[i] = c
					b := string(chs)
					if b == endWord{
						return step+1
					}
					if _, ok := dictionary[b]; !ok{
						continue
					}
					delete(dictionary,b)
					queue = append(queue,b)
				}
				//单词串中单词复位才进行下一位替换
				chs[i] = a
			}
		}
	}
	return 0
}


//双向BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dictionary := make(map[string]bool)
	var flag bool
	for i := 0 ; i < len(wordList) ; i++ {
		if wordList[i] == endWord{
			flag = true
			break
		}
	}

	//判断endword是否在dictionary中
	if !flag {
		return 0
	}

	//开始BFS
	queue_1 := []string{beginWord}
	queue_2 := []string{endWord}
	step := 0


	for len(queue_1) != 0 {
		step++
		tmp := []string{}
		length := len(queue_1)
		for length > 0 {
			value := queue_1[0]
			queue_1 = queue_1[1:]

			for i := 0 ; i < len(queue_2) ; i++ {
				if C(value, queue_2[i]) {
					return step + 1
				}
			}

			for i := 0 ; i < len(wordList) ; i++ {
				if _, ok := dictionary[wordList[i]]; !ok && C(value,wordList[i]) {
					tmp = append(tmp,wordList[i])
					dictionary[wordList[i]] = true
				}
			}
			length--
		}
		queue_1 = tmp
		if len(queue_1) > len(queue_2) {
			queue_1, queue_2 = queue_1, queue_2
		}
	}

	return 0
}

func C(s1,s2 string) bool {
	var count int
	for i := 0 ; i < len(s1) ; i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count == 1
}