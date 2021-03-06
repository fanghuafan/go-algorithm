package insert

// insertion sort [front to back]
func InsertSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		key := data[i]
		for j := 0; j < i; j++ {
			if key < data[j] {
				// move back
				for k := i; k > j; k-- {
					data[k] = data[k-1];
				}
				// key stance
				data[j] = key;
				break;
			}
		}
	}
	return data
}
