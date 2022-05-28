package base_sort

type MergeSort struct{}

func (m MergeSort) Sort(arr []int) {
	helper := make([]int, len(arr))
	m.mergeSort(arr, helper, 0, len(arr)-1)
}

func (m MergeSort) mergeSort(arr []int, helper []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		m.mergeSort(arr, helper, left, mid)
		m.mergeSort(arr, helper, mid+1, right)
		m.mergeSort2(arr, helper, left, mid, right)
	}
}

func (m MergeSort) mergeSort2(arr []int, helper []int, left, mid, right int) {
	for i := left; i <= right; i++ {
		helper[i] = arr[i]
	}
	helperLeft := left
	helperRight := mid + 1
	current := left

	for helperLeft <= mid && helperRight <= right {
		if helper[helperLeft] <= helper[helperRight] {
			arr[current] = helper[helperLeft]
			helperLeft++
		} else {
			arr[current] = helper[helperRight]
			helperRight++
		}
		current++
	}
	remaining := mid - helperLeft
	for i := 0; i <= remaining; i++ {
		arr[current+i] = helper[helperLeft+i]
	}
}
