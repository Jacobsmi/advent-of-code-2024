package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func convertStringsToNumbers(numStrings []string) []int {
	nums := []int{}
	for _, value := range numStrings {
		numValue, _ := strconv.Atoi(value)
		nums = append(nums, numValue)
	}
	return nums
}

func checkValidDiffs(nums []int) bool {
	validDifferences := true
	for i := 0; i < len(nums)-1; i++ {
		difference := math.Abs(float64(nums[i] - nums[i+1]))
		if difference < 1 || difference > 3 {
			validDifferences = false
		}
	}
	return validDifferences
}

// Check if the array is sorted or equal to the reverse of the sorted array
func isSortedOrReverse(nums []int) bool {
	// Check if the array is sorted in ascending order
	if sort.IntsAreSorted(nums) {
		return true
	}

	// Create a copy of the array and reverse it
	copyArr := make([]int, len(nums))
	copy(copyArr, nums)

	reversed := sort.Reverse(sort.IntSlice(copyArr))

	// Check if the reversed array is sorted in ascending order
	return sort.IsSorted(reversed)
}

func main() {
	start := time.Now()
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	safeReports := 0
	for scanner.Scan() {
		// Read in values by line
		numStrings := strings.Split(scanner.Text(), " ")
		nums := convertStringsToNumbers(numStrings)

		// Check for sort (increase and decrease) and valid differences
		if isSortedOrReverse(nums) && checkValidDiffs(nums) {
			safeReports++
		} else {
			for i := 0; i < len(nums); i++ {
				newArr := make([]int, 0, len(nums)-1)
				newArr = append(newArr, nums[:i]...)
				newArr = append(newArr, nums[i+1:]...)
				if isSortedOrReverse(newArr) && checkValidDiffs(newArr) {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Printf("Number of safe reports %v\nFunction took %v\n", safeReports, time.Since(start))
}
