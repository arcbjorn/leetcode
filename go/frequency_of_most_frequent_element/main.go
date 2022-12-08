package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	// key = number, value = frequency (how many types the number is in array)
	// if need to know the actual number in some case
	// frequencyMap := make(map[int]int)

	maxFrequency := 1

	// descending order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	// go by each element
	// for each element increment next element by 1 until it equals the first (target) element
	// skip last element
	for i := 0; i < len(nums)-1; i++ {
		availableIncrements := k

		currentNumber := nums[i]

		// if need to know the actual number in some case
		// frequencyMap[currentNumber]++
		localFrequency := 1

		nextIndex := i + 1

		// go until the end of the array and until 0 availableIncrements
		for nextElementIndex := nextIndex; nextElementIndex < len(nums) && availableIncrements > 0; nextElementIndex++ {
			// exit if equal, it's a match
			for nums[nextElementIndex] < currentNumber && availableIncrements > 0 {
				nums[nextElementIndex]++

				// when numbers match, add frequency & go to the next target element
				if nums[nextElementIndex] == currentNumber {
					// if need to know the actual number in some case
					// frequencyMap[currentNumber]++
					localFrequency++
				}

				availableIncrements--
			}
		}

		if localFrequency > maxFrequency {
			maxFrequency = localFrequency
		}

	}

	return maxFrequency
}

func main() {
	sampleArray := []int{3, 9, 6}
	sampleInt := 2

	maxFreq := maxFrequency(sampleArray, sampleInt)
	fmt.Println(maxFreq)
}
