package BFS

import (
	"hunkevych-philip/algorithms-data-structures/pkg/data_generators"
	"testing"
)

func BenchmarkBFS(b *testing.B) {
	testingTree := data_generators.GetPopulatedTreeOfGivenLength(500)

	b.Run("bfs slice", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			GetBfsBySlicePreOrder(testingTree)
		}
	})

	testingTree = data_generators.GetPopulatedTreeOfGivenLength(500)
	b.Run("bfs queue", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			GetBfsByQueuePreOrder(testingTree)
		}
	})

	testingTree = data_generators.GetPopulatedTreeOfGivenLength(500)
	b.Run("bfs queue recursive", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			GetBfsByQueuePreOrderRecursion(testingTree)
		}
	})

	//Benchmark/bfs_slice-12           					 125880            10136 ns/op           18208  B/op        111 allocs/op
	//Benchmark/bfs_queue-12            				 59527             28927 ns/op           20208  B/op        511 allocs/op
	//Benchmark/bfs_queue_recursive-12                   37404             36095 ns/op  		 67880  B/op        622 allocs/op
}

func TestBFS(t *testing.T) {
	testingTree := data_generators.InitTestingTree()

	t.Run("BFS slice", func(t *testing.T) {
		actualBFS := GetBfsBySlicePreOrder(testingTree)
		if len(actualBFS) != len(data_generators.ExpectedBFSPreOrder) {
			t.Fatalf("Expected length %d, got %d. Expected: %+v, Actual: %+v", len(data_generators.ExpectedBFSPreOrder), len(actualBFS), data_generators.ExpectedBFSPreOrder, actualBFS)
		}

		for i, v := range actualBFS {
			if v != data_generators.ExpectedBFSPreOrder[i] {
				t.Fatalf("Expected %d, got %d", data_generators.ExpectedBFSPreOrder[i], v)
			}
		}
	})

	testingTree = data_generators.InitTestingTree()

	t.Run("BFS queue", func(t *testing.T) {
		actualBFS := GetBfsByQueuePreOrder(testingTree)
		if len(actualBFS) != len(data_generators.ExpectedBFSPreOrder) {
			t.Fatalf("Expected length %d, got %d. Expected: %+v, Actual: %+v", len(data_generators.ExpectedBFSPreOrder), len(actualBFS), data_generators.ExpectedBFSPreOrder, actualBFS)
		}

		for i, v := range actualBFS {
			if v != data_generators.ExpectedBFSPreOrder[i] {
				t.Fatalf("Expected %d, got %d", data_generators.ExpectedBFSPreOrder[i], v)
			}
		}
	})

	testingTree = data_generators.InitTestingTree()

	t.Run("BFS queue recursive", func(t *testing.T) {
		actualBFS := GetBfsByQueuePreOrderRecursion(testingTree)
		if len(actualBFS) != len(data_generators.ExpectedBFSPreOrder) {
			t.Fatalf("Expected length %d, got %d. Expected: %+v, Actual: %+v", len(data_generators.ExpectedBFSPreOrder), len(actualBFS), data_generators.ExpectedBFSPreOrder, actualBFS)
		}

		for i, v := range actualBFS {
			if v != data_generators.ExpectedBFSPreOrder[i] {
				t.Fatalf("Expected %d, got %d", data_generators.ExpectedBFSPreOrder[i], v)
			}
		}
	})
}
