package main

import (
	"fmt"
	"math"
	"unicode"
n	"sort"
)

type stack []interface{}

func (s stack) Push(v interface{}) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, interface{}) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	max := 0
	for num := range set {
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

func isValid(s string) bool {
	st := stack{}
	maping := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, c := range s {
		if open, exists := maping[c]; exists {
			if len(st) > 0 {
				_, top := st.Pop()
				if top.(rune) != open {
					return false
				}
			} else {
				return false
			}
		} else {
			st.Push(c)
		}
	}

	return !(len(st) > 0)
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isAlpaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlpaNum(rune(s[l])) {
			r++
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlpaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		if (numbers[l] + numbers[r]) > target {
			r--
		} else if (numbers[l] + numbers[r]) < target {
			l++
		} else {
			return []int{l+1, r+1}
		}
	}

	return []int{}
}

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		target := nums[i]

		if target > 0 {
			break
        }
		
        if i > 0 && target == nums[i-1] {
            continue
        }
		
		l, r := i+1, len(nums) - 1
		for l < r {
			sum := target + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{target, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
                    l++
                }
			}
		}
	}

	return res
}

func maxArea(height []int) int {
    l, r := 0, len(height) - 1
	max := 0
	for l < r {
		var m int
		if height[l] < height[r] {
			m = height[l]
		} else {
			m = height[r]
		}
		
		area := (r - l) * m

		if area > max {
			max = area
		}
		
		if height[l] <= height[r] {
            l++
        } else {
            r--
        }
	}

	return max
}

func trap(height []int) int {
    l, r := 0, len(height)-1
	maxL, maxR := height[l], height[r]
	ans := 0

	for l < r {
		var val int
		if maxL > maxR {
			val = maxR - height[r]
			r--
			if height[r] > maxR {
				maxR = height[r]
			}
		} else {
			val = maxL - height[l]
			l++
			if height[l] > maxL {
				maxL = height[l]
			}
		}

		if val < 0 {
			val = 0
		}

		ans += val
	}

	return ans
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	return ans
}

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1

	for l <= r {
		m := l + (r-l)/2
		row, col := m / cols, m % cols

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	
	return false
}

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 1
	for _, p := range piles {
		if p > r {
			r = p
		}
	}
	ans := r

	for l <= r {
		mid := (l + r) / 2
		k := 0
		for _, p := range piles {
			k += int(math.Ceil(float64(p) / float64(mid)))
		}
		if k > h {
			l = mid + 1
		} else {
			ans = mid
			r = mid - 1
		}
	}
    return ans
}

func findMin(nums []int) int {
	ans := nums[0]
	l, r := 0, len(nums) - 1
	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < ans {
				ans = nums[l]
				break
			}
			
			mid := (r+l)/2
			if ans > nums[mid] {
				ans = nums[mid]
			}
			
			if nums[mid] >= nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return ans
}

func searchRotatedArray(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r)/2
		if target == nums[mid] {
			return mid
		}

		if nums[l] <= nums[mid] {
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
}
		}
	}
	return -1
}
