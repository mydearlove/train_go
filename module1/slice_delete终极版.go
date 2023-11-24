package main

import (
	"fmt"
)

// DeleteAtIndex 删除切片中指定下标的元素
func DeleteAtIndex[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		// 下标越界，直接返回原切片
		return slice
	}

	// 使用 append 将切片分为前后两部分，并且不包含要删除的元素

	copy(slice[index:], slice[index+1:])
	// 切片长度减一
	return slice[:len(slice)-1]
}

// RemoveElement 删除切片中的指定元素
func RemoveElement[T comparable](slice []T, element T) []T {
	index := -1
	for i, v := range slice {
		if v == element {
			index = i
			break
		}
	}
	if index == -1 {
		// 没有找到要删除的元素，直接返回原切片
		return slice
	}

	return DeleteAtIndex(slice, index)
}

// ShrinkToFit 缩容切片至其长度等于其容量
func ShrinkToFit[T any](slice []T) []T {
	return append([]T(nil), slice...)
}

func main() {
	// 示例用法
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", intSlice)

	// 删除指定下标的元素
	indexToDelete := 2
	intSlice = DeleteAtIndex(intSlice, indexToDelete)
	fmt.Println("删除下标为", indexToDelete, "的元素:", intSlice)

	// 删除指定元素
	elementToDelete := 4
	intSlice = RemoveElement(intSlice, elementToDelete)
	fmt.Println("删除元素", elementToDelete, ":", intSlice)

	// 缩容切片
	intSlice = ShrinkToFit(intSlice)
	fmt.Println("缩容后的切片:", intSlice)
}
