package main

import "fmt"

// DeleteAtIndex 删除切片中指定下标的元素（泛型版本）
func deleteAtIndex3[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		// 下标越界，直接返回原切片
		return slice
	}

	// 使用 copy 函数将后面的元素覆盖前面的元素，达到删除指定下标元素的效果
	copy(slice[index:], slice[index+1:])
	// 切片长度减一
	return slice[:len(slice)-1]
}

func main() {
	// 示例用法
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", intSlice)

	// 删除指定下标的元素
	indexToDelete := 2
	intSlice = deleteAtIndex3(intSlice, indexToDelete)
	fmt.Println("删除下标为", indexToDelete, "的元素:", intSlice)

	// 泛型版本的示例用法
	stringSlice := []string{"apple", "banana", "cherry"}
	fmt.Println("原始切片:", stringSlice)

	// 删除指定下标的元素
	indexToDeleteStr := 1
	stringSlice = DeleteAtIndex(stringSlice, indexToDeleteStr)
	fmt.Println("删除下标为", indexToDeleteStr, "的元素:", stringSlice)
}
