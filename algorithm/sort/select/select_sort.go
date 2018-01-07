package _select

//Simple Selection Sort
func SelectSort(data []int) {
	len := len(data)
	for i := 1; i < len; i++ {
		minKey := chooseMinKey(data, i)
		if data[minKey] < data[i-1] {
			data[minKey], data[i-1] = data[i-1], data[minKey]
		}
	}
}

func chooseMinKey(data []int, start int) int {
	min := start
	for i := start; i < len(data); i++ {
		if data[min] > data[i] {
			min = i
		}
	}
	return min
}
