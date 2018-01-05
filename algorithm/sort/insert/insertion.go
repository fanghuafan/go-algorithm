package insert

// 插入排序
func InsertSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		key := data[i]
		for j := 0; j < i; j++ {
			if key < data[j] {
				// 将数据往后移动
				for k := i; k > j; k-- {
					data[k] = data[k-1];
				}
				// key 站位
				data[j] = key;
				break;
			}
		}
	}
	return data
}
