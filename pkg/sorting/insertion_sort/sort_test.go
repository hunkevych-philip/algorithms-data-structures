package insertion_sort

import (
	"hunkevych-philip/algorithms-data-structures/pkg/data_generators"
	"testing"
)

func BenchmarkSortOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortOne(data_generators.GetPopulatedSliceOfGivenLength(99999, 99999))
	}
}

func BenchmarkSortTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortTwo(data_generators.GetPopulatedSliceOfGivenLength(99999, 9999))
	}
}
