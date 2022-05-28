package base_sort

type QuickSort struct{}

func (s QuickSort) Sort(arr []int) {
	s.quickSort(arr, 0, len(arr)-1)
}

func (s QuickSort) quickSort(arr []int, left, right int) {
	index := s.partition(arr, left, right)
	if left < index-1 {
		s.quickSort(arr, left, index-1)
	}
	if index < right {
		s.quickSort(arr, index, right)
	}
}

func (s QuickSort) partition(arr []int, left, right int) int {
	pivot := arr[(left+right)/2]
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}
