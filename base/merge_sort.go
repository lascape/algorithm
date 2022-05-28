package base

type MergeSort struct{}

func (m MergeSort) Sort(array []int) {
	helper := make([]int, len(array))
	m.mergeSort(array, helper, 0, len(array)-1)
}

func (m MergeSort) mergeSort(array []int, helper []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		m.mergeSort(array, helper, left, mid)
		m.mergeSort(array, helper, mid+1, right)
		m.mergeSort2(array, helper, left, mid, right)
	}
}

func (m MergeSort) mergeSort2(array []int, helper []int, left, mid, right int) {
	for i := left; i <= right; i++ {
		helper[i] = array[i]
	}
	helperLeft := left
	helperRight := mid + 1
	current := left

	for helperLeft <= mid && helperRight <= right {
		if helper[helperLeft] <= helper[helperRight] {
			array[current] = helper[helperLeft]
			helperLeft++
		} else {
			array[current] = helper[helperRight]
			helperRight++
		}
		current++
	}
	remaining := mid - helperLeft
	for i := 0; i <= remaining; i++ {
		array[current+i] = helper[helperLeft+i]
	}
}
