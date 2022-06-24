package merge_sort

// TODO: TBD

// MergeSort cases:
// 1,2:4,5 -> *2<=4* 					== merge as it is: 1245 (1 op)
// 7,8:2,3 -> 7*>=*3 					== swap and merge: 2378 (1 op)
// 1,7:3,4 -> 1*<=3*, _7<_5        	    == place one by 1 (2 op)
// ^^ at this point we know the 2nd num from the 1st slice
// is bigger than 1st num from the 2nd slice,
// so we need to compare it with the 2nd number to decide where to put it:
// before the last element or after 3 (op)
// 3,7:2,4 -> |3*>=2*| |3,*>=_4| |_,7>=_4|:    == (3 op)

//func MergeSort(n []int) []int {
//	var (
//		l      = len(n)
//		h1, h2 []int
//	)
//
//	if l%2 != 0 {
//		h1, h2 = n[:l/2+1], n[l/2:]
//	}
//	h1, h2 = n[:l/2], n[l/2:]
//
//	return nil
//}

/*	return Merge(MergeSort(h1), MergeSort(h2))
}

func Merge(n1, n2 []int) []int {
	l1, l2 := len(n1), len(n2)
	if l1%2 == 0 {

	} else {

	}

	if l2%2 == 0 {

	} else {

	}

	return nil
}*/
