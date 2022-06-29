package selection_sort

import "math"

func Sort(s []int) []int {
	iSmallest, vSmallest := 0, math.MaxInt

	for i := range s {
		for j := i; j < len(s); j++ {
			if s[j] < vSmallest {
				iSmallest, vSmallest = j, s[j]
			}
		}
		s[i], s[iSmallest] = s[iSmallest], s[i]
		vSmallest = math.MaxInt
	}

	return s
}
