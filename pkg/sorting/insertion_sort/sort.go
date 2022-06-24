package insertion_sort

func SortOne(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i-1] < s[i] {
			continue
		}
		for j := i; j >= 1; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			} else {
				break
			}
		}
	}
	return s
}

func SortTwo(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		currElem := arr[i]
		prevIndx := i - 1

		for prevIndx >= 0 && currElem < arr[prevIndx] {
			arr[prevIndx+1] = arr[prevIndx]
			prevIndx--
		}
		arr[prevIndx+1] = currElem
	}
	return arr
}
