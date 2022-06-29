package selection_sort

func Sort(s []int) []int {
	n := len(s)

	for i := 0; i < n; i++ {
		m := i // The lowest value's index

		for j := i; j < n; j++ {
			if s[j] < s[m] {
				m = j
			}
		}
		s[i], s[m] = s[m], s[i]
	}

	return s
}
