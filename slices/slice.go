package main

import (
	"fmt"
	"slices"
)

func main() {
	var slice1 []int = []int{10,20}

	fmt.Println(slice1)

	var slice2 []int  = make([]int,5,10)

	fmt.Println(slice2 == nil) // true, because slice2 is not initialized
	length := len(slice2) // initializing slice2 with zero length
	fmt.Println(length)
	// slice2[2] = 100 // This will cause a runtime panic: index out of range
	fmt.Println(slice2) // This line will not be executed due to the panic
	fmt.Println( cap(slice2)) // This will not be executed due to the panic
	slice2 = append(slice2,100)
	fmt.Println(slice2) // Now slice2 has been initialized and can be used
	fmt.Println(len(slice2)) // Length of slice2 after appending
	fmt.Println(cap(slice2)) // Capacity of slice2 after appending
	var nums []int = []int{10,50}
	var nums2 []int = []int{20,50,60}
	nums[1] = 500
	// nums2... // This is a variadic parameter, which allows us to pass a slice as individual elements
	nums = append(nums,nums2...) // Appending nums2 to nums
	fmt.Println(nums) // Output: [10 50 20 50 60]
	fmt.Println(cap(nums)) // Output: 4, because the capacity is doubled when appending

	var nums3 []int = make([]int, len((nums)))
	copy(nums3, nums) 

	fmt.Println(nums3) // This will cause a runtime panic: invalid memory address or nil pointer dereference
	fmt.Println(nums2) // Length of nums3 after copying
	fmt.Println(nums2[1:3]) // Slicing nums2 from index 1 to 3

	fmt.Println(nums)
	fmt.Println(nums3)
	fmt.Println(slices.Equal(nums, nums3)) // Checking if nums and nums3 are equal
	var nums4 []int = nums3
	fmt.Println("Capacity of nums3:", cap(nums3)) // Capacity of nums3
	fmt.Println(nums4) // Output: [10 500 20 50 60]
	nums4[0] = 1000
	fmt.Println((nums4))
	fmt.Println(nums4, cap(nums4)) // Output: [1000 500 20 50 60 5000]
	nums4 = append(nums4, 5000)
	fmt.Println(nums4, cap(nums4)) // Output: [1000 500 20 50 60 5000]
	var nums5 = [][]int{{10,20,30}, {40,50,60}}
	fmt.Println(nums5) // Output: [[10 20 30] [40 50 60]]
	fmt.Println(nums5[0][1])
}