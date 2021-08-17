package data

type Uint64List []uint64

// 定义类型长度
func (my64 Uint64List) Len() int {
	return len(my64)
}
// 交换数值
func (my64 Uint64List) Swap(i, j int) {
	my64[i], my64[j] = my64[j], my64[i]
}
// 判断大小
func (my64 Uint64List) Less(i, j int) bool {
	return my64[i] < my64[j]
}
