package data_generators

import (
	"hunkevych-philip/algorithms-data-structures/pkg/types/rand"
	"time"
)

func GetPopulatedSliceOfGivenLength(n, limit int) []int {
	time.Sleep(time.Second)

	var (
		res = make([]int, n)
		r   = rand.R()
	)

	for i := 0; i < n; i++ {
		res[i] = r.Intn(limit)
	}

	return res
}
