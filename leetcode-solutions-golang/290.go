func wordPattern(pattern string, s string) bool {
	charMap := make(map[string]string)
	wordMap := make(map[string]string)
	indexOfPattern := 0

	var makeWord string
	for i := 0; i < len(s); i++ {
		if s[i] == ' '{
			if indexOfPattern >= len(pattern) {
				return false
			}

			char, check1 := charMap[makeWord]
			word, check2 := wordMap[string(pattern[indexOfPattern])]

			if check1 != check2 {
				return false
			} else{
				if check1 == true && (word != makeWord || char != string(pattern[indexOfPattern])) {
					return false
				}
				if check1 == false {
				charMap[makeWord] = string(pattern[indexOfPattern])
				wordMap[string(pattern[indexOfPattern])] = makeWord
				}
			}
			makeWord = ""
			indexOfPattern++
		} else{
			makeWord += string(s[i])

			if i == len(s) - 1 {
				if indexOfPattern >= len(pattern) {
					return false
				}
				if(indexOfPattern != len(pattern) - 1) {
					return false
				}
				char, check1 := charMap[makeWord]
				word, check2 := wordMap[string(pattern[indexOfPattern])]

				if check1 != check2 {
					return false
				} else {
					if check1 == true && (word != makeWord || char != string(pattern[indexOfPattern])) {
						return false
					}
				}
			}

		}
	}
	return true
}