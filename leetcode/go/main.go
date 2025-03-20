package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}

	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	max := 0
	for num := range set{
		if !set[num-1] {
            c := num
			len := 1
			for set[c+1] {
                c++
                len++
			}

			if max < len {
				max = len
			}
		}
	}
	return max
}
