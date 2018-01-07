package insert

// shell sort implements [back to front]
func ShellInsertSort(data []int, dk int) {
	for i := dk; i < len(data); i++ {
		if data[i] < data[i-dk] {
			j := i - dk
			x := data[i]
			data[i] = data[i-dk]
			for j >= 0 && x < data[j] {
				data[j+dk] = data[j]
				j -= dk
			}
			data[j+dk] = x
		}
	}
}

/*
shell sort
 */
func ShellSort(data []int) {
	dk := len(data) / 2;
	for dk >= 1 {
		ShellInsertSort(data, dk);
		dk = dk / 2;
	}
}
