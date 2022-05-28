package base

import "testing"

func TestBubbleSort_Sort(t *testing.T) {
	tests := []struct {
		array  []int
		output []int
	}{
		{
			array:  []int{1, 8, 9, 5, 20, 15, 4, 2},
			output: []int{1, 2, 4, 5, 8, 9, 15, 20},
		},
		{
			array:  []int{10, 8, 9, 5, 20, 15, 6, 2},
			output: []int{2, 5, 6, 8, 9, 10, 15, 20},
		},
	}
	for _, tt := range tests {
		BubbleSort{}.Sort(tt.array)
		for i, v := range tt.array {
			if tt.output[i] != v {
				t.Errorf("want: %v,result: %v", tt.output, tt.array)
				return
			}
		}
	}
}
