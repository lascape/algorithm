package base

type InsertSort struct{}

func (s InsertSort) Sort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		e := arr[i]
		j := i
		for ; j > 0 && e < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
