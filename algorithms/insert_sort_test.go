package algorithms_test

import (
	"testing"

	"github.com/joseluis8906/yummies/go-code/src/algorithms"
	"github.com/joseluis8906/yummies/go-code/src/pkg/cmp"
)

func TestInsertSort(t *testing.T) {
	t.Parallel()

	want := []int{1, 2, 3, 4, 5, 6}
	got := []int{5, 2, 4, 6, 1, 3}

	algorithms.InsertSort(got)

	if !cmp.Equal(got, want) {
		t.Errorf("InsertSort(%v) = %v; want %v\n%v", got, want, got, cmp.Diff(want, got))
	}

}

func BenchmarkInsertSort(b *testing.B) {
	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "Small",
			input: []int{9, 6, 2},
		},
		{
			name:  "Middle",
			input: []int{10, 70, 20, 5, 33, 22},
		},
		{
			name:  "Big",
			input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algorithms.InsertSort(tc.input)
			}
		})
	}
}
