package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
	s := []string{"00:00","23:59","00:00"}
	fmt.Println(findMinDifference(s))
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
			return []int{l + 1, r + 1}
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

		l, r := i+1, len(nums)-1
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
	l, r := 0, len(height)-1
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
		row, col := m/cols, m%cols

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
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < ans {
				ans = nums[l]
				break
			}

			mid := (r + l) / 2
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
		mid := (l + r) / 2
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
	return -1
}

type TimeMap struct {
	m map[string][]pair
}

type pair struct {
	time  int
	value string
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], pair{value: value, time: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, e := this.m[key]; !e {
		return ""
	}

	pairs := this.m[key]

	l, r := 0, len(pairs)-1

	for l <= r {
		mid := (l + r) / 2
		if pairs[mid].time <= timestamp {
			if mid == len(pairs)-1 || pairs[mid+1].time > timestamp {
				return pairs[mid].value
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

func nthSuperUglyNumber(n int, primes []int) int {
	return 0
}

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = rev
		rev = head
		head = temp
	}
	return rev
}

func putMarbles(weights []int, k int) int64 {
	return 0
}

func mostPoints(questions [][]int) int64 {
	length := len(questions)
	dp := make([]int64, length)
	dp[length-1] = int64(questions[length-1][0])
	for i := length - 2; i >= 0; i-- {
		dp[i] = solve(questions, i, dp)
	}

	var max int64
	for _, p := range dp {
		if p > max {
			max = p
		}
	}

	return max
}

func solve(questions [][]int, start int, dp []int64) int64 {
	i := start
	var points int64
	points = int64(questions[i][0])
	if questions[i][1]+1+i < len(questions) {
		points = points + dp[questions[i][1]+1+i]
	}
	return max(points, dp[start+1])
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maximumTripletValue(nums []int) int64 {	
	n := len(nums)
    if n < 3 {
        return 0
    }
    
    maxVal := 0
    maxLeft := nums[0]
    maxDiff := 0
    
    for k := 2; k < n; k++ {
        if maxLeft - nums[k-1] > maxDiff {
            maxDiff = maxLeft - nums[k-1]
        }
        
        if nums[k-1] > maxLeft {
            maxLeft = nums[k-1]
        }
        
        current := maxDiff * nums[k]
        if current > maxVal {
            maxVal = current
        }
    }
    
    if maxVal < 0 {
        return 0
    }
    return int64(maxVal)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	list := dummy
	
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			list.Next = list2
			list2 = list2.Next
		} else {
			list.Next = list1
			list1 = list1.Next
		}

		list = list.Next
	}
	list.Next = list1
	if list1 == nil {
		list.Next = list2
	}

	return dummy.Next
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
    return false
}

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil
	var prev *ListNode
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	first := head
	second = prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	N := 0
	curr := head
	for curr != nil {
		N++
		curr = curr.Next
	}

	index := N - n
	if index == 0 {
		return head.Next
	}

	curr = head
	for i := 0; i < N - 1; i++ {
		if (i + 1) == index {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return head
}

func findMinDifference(timePoints []string) int {
	count := int(1e4)
	times := make([]int, 0)

	for _, timeStr := range timePoints {
		hourStr := strings.Split(timeStr, ":")[0]
		minuteStr := strings.Split(timeStr, ":")[1]

		hour, _ := strconv.Atoi(hourStr)
		minute, _ := strconv.Atoi(minuteStr)

		if hour == 0 {
			hour = 24
		}

		time := hour * 60 + minute
		times = append(times, time)
	}
	
	for i := 0; i < len(times); i++ {
		for j := 0; j < len(times); j++ {
			diff := int(math.Abs(float64(times[i] - times[j])))
			fmt.Println(count, diff)
			if count > diff && i != j{
				count = diff
			}
		}
	}

	return count
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (int, *TreeNode)
	dfs = func(node *TreeNode) (int, *TreeNode) {
		if node == nil {
			return 0, nil
		} 
		if node.Left == nil && node.Right == nil {
			return 1, node
		}

		deepestLevelLeft, lcaLeft := dfs(node.Left)
		deepestLevelRight, lcaRight := dfs(node.Right)
		if deepestLevelLeft == deepestLevelRight {
			return deepestLevelLeft + 1, node
		}
		if deepestLevelLeft > deepestLevelRight {
			return deepestLevelLeft + 1, lcaLeft
		}

		return deepestLevelRight+1, lcaRight
	}

	_, rs := dfs(root)
	return rs 
}

func countSymmetricIntegers(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if isSymmetric(i) {
			count++
		}
	}
	return count
}

func isSymmetric(num int) bool {
	numStr := strconv.Itoa(num)
	if len(numStr) % 2 != 0 {
		return false
	}
	count1 := 0
	count2 := 0
	for i, dStr := range numStr {
		if i > len(numStr)/2 {
			d, _ := strconv.Atoi(string(dStr))
			count1 += d
		} else {
			d, _ := strconv.Atoi(string(dStr))
			count2 += d
		}
	}

	if count1 != count2 {
		return false
	}

	return true
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	result := 0
	for i := 0; i < len(arr) - 2; i++ {
        for j := i + 1; j < len(arr) - 1; j++ {
            if mod(arr[i] - arr[j]) > a {
                continue
            }

            for k := j + 1; k < len(arr); k++ {
                if mod(arr[j] - arr[k]) <= b && mod(arr[i] - arr[k]) <= c {
                    result++
                }
            }
        }
    }

    return result
}

func mod(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
