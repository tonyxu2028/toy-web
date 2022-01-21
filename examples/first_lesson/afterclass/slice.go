package main

import "fmt"

func main() {

	s := []int{1, 2, 4, 7}

	fmt.Printf("Add :::::: \n")

	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)
	fmt.Printf("s: %v, len %d, cap: %d \n", s, len(s), cap(s))

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	fmt.Printf("s: %v, len %d, cap: %d \n", s, len(s), cap(s))

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)
	fmt.Printf("s: %v, len %d, cap: %d \n", s, len(s), cap(s))

	fmt.Printf("Delete :::::: \n")

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)
	fmt.Printf("s: %v, len %d, cap: %d \n", s, len(s), cap(s))

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)
	fmt.Printf("s: %v, len %d, cap: %d \n", s, len(s), cap(s))

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)
	fmt.Printf("s: %v, len %d, cap: %d \n", s, len(s), cap(s))

}

func Add(s []int, index int, value int) []int {
	s = append(s[:index], append([]int{value}, s[index:]...)...)
	return s
}

func Delete(s []int, index int) []int {
	s = append(s[:index], s[index+1:]...)
	return s
}
