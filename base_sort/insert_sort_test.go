package base_sort

import "testing"

func TestInsertSort_Sort(t *testing.T) {
	tests := []struct {
		arr    []int
		output []int
	}{
		{
			arr:    []int{1, 8, 9, 5, 20, 15, 4, 2},
			output: []int{1, 2, 4, 5, 8, 9, 15, 20},
		},
		{
			arr:    []int{10, 8, 9, 5, 20, 15, 6, 2},
			output: []int{2, 5, 6, 8, 9, 10, 15, 20},
		},
	}
	for _, tt := range tests {
		InsertSort{}.Sort(tt.arr)
		for i, v := range tt.arr {
			if tt.output[i] != v {
				t.Errorf("want: %v,result: %v", tt.output, tt.arr)
				return
			}
		}
	}
}
