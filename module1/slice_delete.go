package main

import "fmt"

// DeleteAtIndex 删除切片中指定下标的元素'
func deleteAtIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		// 下标越界，直接返回原切片
		return slice
	}

	// 使用 append 将切片分为前后两部分，并且不包含要删除的元素
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// 示例用法
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", intSlice)

	// 删除指定下标的元素
	indexToDelete := 2
	intSlice = deleteAtIndex(intSlice, indexToDelete)
	fmt.Println("删除下标为", indexToDelete, "的元素:", intSlice)
}
