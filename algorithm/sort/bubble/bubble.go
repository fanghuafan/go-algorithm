package bubble

// bubble sort
func Sort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i];
			}
		}
	}
	return data
}
