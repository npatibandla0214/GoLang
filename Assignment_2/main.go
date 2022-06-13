package main

import "fmt"

func main() {
	s := []int{20, 30, 40, 50, 60, 70, 80}
	fmt.Printf("Length of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))

	s = s[:0]
	fmt.Println(s)
	fmt.Printf("Length of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))

	s = s[1:4]
	fmt.Println(s)
	fmt.Printf("Length of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))

	s = s[2:]
	fmt.Println(s)
	fmt.Printf("Length of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))
}
