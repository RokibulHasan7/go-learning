func backward(cnt *int, words []string, target string, idx1 int, idx2 int) bool{
	for i := idx1; i > idx2; i-- {
		*cnt += 1
		if(words[i] == target){
			return true;
		}
	}
	return false
}
func forward(cnt *int, words []string, target string, idx1 int, idx2 int) bool{
	for i := idx1; i < idx2; i++ {
		*cnt += 1
		if(words[i] == target){
			return true;
		}
	}
	return false
}
func closetTarget(words []string, target string, startIndex int) int {
	if(words[startIndex] == target){
		return 0
	}

	left := 0
	right := 0
	firstCheck := false
	sz := len(words)

	for i := 0; i < sz; i++ {
		if(words[i] == target){
			firstCheck = true
			break
		}
	}
	if firstCheck == false {
		return -1
	}

	check := false
	check = backward(&left, words, target, startIndex - 1, -1)
	if check == false {
		check = backward(&left, words, target, sz - 1, startIndex)
	}

	check = false
	check = forward(&right, words, target, startIndex + 1, sz)
	if check == false {
		check = forward(&right, words, target, 0, startIndex)
	}

	if(left < right){
		return left
	} else{
		return right;
	}
}