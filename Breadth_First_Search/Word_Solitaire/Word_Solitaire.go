package Word_Solitaire

//127. 单词接龙
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
