package main

import "fmt"

/**
 * 切片处理方式，切片和数组的区别是切片不需要声明长度，但是数组需要
 */
func main() {
	s1 := []int{1, 2, 3, 4} // 直接初始化了 4 个元素的切片
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) // 创建了一个包含三个元素，容量为4的切片
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// s2 目前 [0, 0, 0], append（追加）一个元素，变成什么？
	s2 = append(s2, 7) // 后边添加一个元素，没有超出容量限制，不会发生扩容
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 8) // 后边添加了一个元素，触发扩容
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4) // 只传入一个参数，表示创建一个含有四个元素，容量也为四个元素的
	// 等价于 s3 := make([]int, 4, 4)
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	// 按下标索引
	fmt.Printf("s3[2]: %d \n", s3[2])

	// 超出下标范围，直接崩溃
	// runtime error: index out of range [99] with length 4
	// fmt.Printf("s3[99]: %d", s3[99])
	// SubSlice()
	//shareArr()
	// shareSlice
	fmt.Printf("shareSlice ::: \n")
	shareSlice()
}

func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3]
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4: %v, len %d, cap: %d \n", s4, len(s4), cap(s4))
}

func shareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// 1.这里在s2没有进行append之前，由于s2是由s1派生处理的，所以此时两者共享数据，因此会两者同时修改数据
	s2[0] = 99
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// 2.这里在s2没有进行append之后，新的s2是根据之前派生出的s2进行追加的，此时不在于母体s1共享数据了，
	// 故导致之后s2[1]=1999，没有修改之前s1的数据
	s2 = append(s2, 199)
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// 这里没有修改s1的数据详见2
	s2[1] = 1999
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
}
