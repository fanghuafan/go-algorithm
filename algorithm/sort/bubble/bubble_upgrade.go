package bubble

// bubble upgrade sort
func SortUpgrade(data []int) []int {
	for i := 0; i < len(data); i++ {
		isMove := false;
		for j := i; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i];
				isMove = true;
			}
		}
		if (!isMove) {
			break;
		}
	}
	return data
}
