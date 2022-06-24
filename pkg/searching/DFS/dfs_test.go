package DFS

import (
	"hunkevych-philip/algorithms-data-structures/pkg/data_generators"

	"testing"
)

func BenchmarkDFS(b *testing.B) {
	testingTree := data_generators.GetPopulatedTreeOfGivenLength(500)

	b.Run("DFS slice", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			GetDfsBySlicePreOrder(testingTree)
		}
	})

	testingTree = data_generators.GetPopulatedTreeOfGivenLength(16)
	b.Run("DFS slice recursion", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			GetDfsBySliceRPreOrder(testingTree)
		}
	})

	testingTree = data_generators.GetPopulatedTreeOfGivenLength(500)
	b.Run("DFS stack", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			GetDfsByStackPreOrder(testingTree)
		}
	})
}

func TestDFS(t *testing.T) {
	testingTree := data_generators.InitTestingTree()

	t.Run("DFS slice preorder ", func(t *testing.T) {
		actualDFS := GetDfsBySlicePreOrder(testingTree)
		if len(actualDFS) != len(data_generators.ExpectedDFSPreOrder) {
			t.Fatalf("Expected length %d, got %d. Expected: %+v, Actual: %+v", len(data_generators.ExpectedDFSPreOrder), len(actualDFS), data_generators.ExpectedDFSPreOrder, actualDFS)
		}

		for i, v := range actualDFS {
			if v != data_generators.ExpectedDFSPreOrder[i] {
				t.Fatalf("Expected %d, got %d", data_generators.ExpectedDFSPreOrder[i], v)
			}
		}
	})

	testingTree = data_generators.InitTestingTree()
	t.Run("DFS slice preorder recursion", func(t *testing.T) {
		actualDFS := GetDfsBySliceRPreOrder(testingTree)
		if len(actualDFS) != len(data_generators.ExpectedDFSPreOrder) {
			t.Fatalf("Expected length %d, got %d. Expected: %+v, Actual: %+v", len(data_generators.ExpectedDFSPreOrder), len(actualDFS), data_generators.ExpectedDFSPreOrder, actualDFS)
		}

		for i, v := range actualDFS {
			if v != data_generators.ExpectedDFSPreOrder[i] {
				t.Fatalf("Expected %d, got %d", data_generators.ExpectedDFSPreOrder[i], v)
			}
		}
	})

	testingTree = data_generators.InitTestingTree()
	t.Run("DFS slice inorder", func(t *testing.T) {
		actualDFS := GetDfsBySliceInOrder(testingTree)
		if len(actualDFS) != len(data_generators.ExpectedDFSInOrder) {
			t.Fatalf("Expected length %d, got %d. Expected: %+v, Actual: %+v", len(data_generators.ExpectedDFSPreOrder), len(actualDFS), data_generators.ExpectedDFSPreOrder, actualDFS)
		}

		for i, v := range actualDFS {
			if v != data_generators.ExpectedDFSInOrder[i] {
				t.Fatalf("Expected %d, got %d", data_generators.ExpectedDFSInOrder[i], v)
			}
		}
	})

	testingTree = data_generators.InitTestingTree()
	t.Run("DFS stack preorder", func(t *testing.T) {
		actualDFS := GetDfsByStackPreOrder(testingTree)
		if len(actualDFS) != len(data_generators.ExpectedDFSPreOrder) {
			t.Fatalf("Expected length %d, got %d. Expected: %+v, Actual: %+v", len(data_generators.ExpectedDFSPreOrder), len(actualDFS), data_generators.ExpectedDFSPreOrder, actualDFS)
		}

		for i, v := range actualDFS {
			if v != data_generators.ExpectedDFSPreOrder[i] {
				t.Fatalf("Expected %d, got %d", data_generators.ExpectedDFSPreOrder[i], v)
			}
		}
	})
}
