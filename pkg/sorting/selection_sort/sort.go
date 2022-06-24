package selection_sort

func Sort(s []int) []int {
	maxInt := 2147483647  // max int
	iSmallest := 0        // where to remove value
	valSmallest := maxInt // flag
	for i := range s {
		for j := i; j < len(s); j++ {
			if s[j] < valSmallest {
				iSmallest = j
				valSmallest = s[j]
			}
		}
		s[i], s[iSmallest] = s[iSmallest], s[i]
		valSmallest = maxInt
	}

	return s
}
