package main

import "fmt"

// 10. 你能说出切片的扩容逻辑吗？
// 通过 append 扩容的时候，指针会发生变化，原切片不受影响。在 1024 大小之前每次扩容会翻倍，之后会按照 1.25 倍扩容。
func main() {
	fmt.Println("深浅拷贝")
	sliceA := []int{1, 2}
	sliceB := make([]int, 2) // 空的切片
	sliceC := sliceA
	fmt.Printf("slice A %p\n", sliceA)
	fmt.Printf("slice B %p\n", sliceB)
	fmt.Printf("slice C %p\n", sliceC)

	fmt.Println("切片扩容地址测试")
	slice1 := []int{2}
	fmt.Printf("[slice1] ptr:%p, content:%o, len:%d, cap:%d\n", &slice1, slice1, len(slice1), cap(slice1))
	slice2 := append(slice1, 3)
	fmt.Printf("[slice2] ptr:%p, content:%o, len:%d, cap:%d\n", &slice2, slice2, len(slice2), cap(slice2))
	fmt.Printf("[slice1] ptr:%p, content:%o, len:%d, cap:%d\n", &slice1, slice1, len(slice1), cap(slice1))

	fmt.Println("切片扩容容量变化测试")
	sliceZero := []int{1}
	length, capicity := len(sliceZero), cap(sliceZero)
	fmt.Printf("[sliceZero] ptr:%p, content:%o, len:%d, cap:%d\n", *&sliceZero, sliceZero, length, capicity)
	for i := 0; i < 4098; i++ {
		sliceZero = append(sliceZero, i)
		if capicity != cap(sliceZero) {
			length = len(sliceZero)
			capicity = cap(sliceZero)
			// 这里打印的是 sliceZero 指针所指向的地址，应该在不断变化
			fmt.Printf("[sliceZero] ptr:%p, len:%d, cap:%d\n", *&sliceZero, length, capicity)
		}
	}

}
