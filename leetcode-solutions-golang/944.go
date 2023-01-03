func minDeletionSize(strs []string) int {
	Ans := 0
	numberOfCol := len(strs[0])
	numberOfRow := len(strs)
	for j := 0; j < numberOfCol; j++ {
		var tmpString string
		var originalString string
		for i := 0; i < numberOfRow; i++ {
			tmpString += string(strs[i][j])
			originalString += string(strs[i][j])
		}
		//sort.Strings(tmpString) --> gives error
		sortedString := sortString(tmpString)
		//fmt.Println(tmpString)
		if sortedString != originalString {
			Ans++
		}
	}
	return Ans
}

func sortString(input string) string{
	runeArray := []rune(input)                   // convert in Rune [act like an array]
	sort.Sort(sortRuneString(runeArray))
	fmt.Println(string(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]            // swap the values of s[i] and s[j]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]                 // sorting in ascending
}

func (s sortRuneString) Len() int {
	return len(s)                      // return the length
}


/*
Focus:
	Sort a string. First make a "rune" array. Then send them in sort.Sort() properties.
*/

/*
Signature of sort.Sort() function: func Sort(data Interface)
Definition of Interface :

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

So whatever we want to sort using the sort.Sort function then that needs to implement above three functions
1. Swap(i, j) int
2. Less(i, j int) bool
3. Len()
*/

/*

*/