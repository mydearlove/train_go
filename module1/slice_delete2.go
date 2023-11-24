package main

import "fmt"

// DeleteAtIndex 删除切片中指定下标的元素
func deleteAtIndex2(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		// 下标越界，直接返回原切片
		return slice
	}

	// 使用 copy 函数将后面的元素覆盖前面的元素，达到删除指定下标元素的效果
	//在Go语言中，使用copy函数删除元素的方式相对于append具有性能上的优势，尤其是在大型切片上。这是因为copy函数是按字节复制的，而append可能会导致内存重新分配，将元素一个个复制到新的内存空间中。
	//
	//在删除切片中的元素时，使用copy的主要优势在于避免了创建新切片和内存分配。copy直接在原有的切片上覆盖需要删除的元素，然后通过调整切片的长度来完成删除操作。这样做的好处是：
	//
	//避免内存重新分配： append 有可能会触发内存重新分配，导致更多的内存管理开销。copy直接在原切片上进行操作，不需要额外的内存分配，因此更高效。
	//
	//减少内存使用： append 可能会在新的内存块中留下未使用的空间，而 copy 在原切片上直接操作，不会产生额外的未使用内存。
	//
	//更快的执行速度： 由于 copy 是按字节复制的，而 append 则可能需要更多的逻辑来处理扩容等问题，因此在某些情况下，copy 可能会更快。
	//
	//需要注意的是，对于小型切片或者切片中的元素数量相对较小时，两者的性能差异可能不太明显。因此，性能优化应该根据具体的使用场景和数据规模来进行权衡。
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
	intSlice = deleteAtIndex2(intSlice, indexToDelete)
	fmt.Println("删除下标为", indexToDelete, "的元素:", intSlice)
}
