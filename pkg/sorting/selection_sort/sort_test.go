package selection_sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "valid",
			args: args{
				s: []int{10, 5, 6, 2, 7},
			},
			want: []int{2, 5, 6, 7, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
