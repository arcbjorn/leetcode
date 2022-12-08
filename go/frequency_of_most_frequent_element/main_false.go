package main

import (
	"fmt"
	"sort"
)

func maxFrequency_false(nums []int, k int) int {
	// key = number, value = frequency (how many types the number is in array)
	// if need to know the actual number in some case
	// frequencyMap := make(map[int]int)

	maxFrequency := 1

	// descending order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	fmt.Println(nums)
	fmt.Println("")

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
		for nextElementIndex := nextIndex; nextElementIndex < len(nums)-1 && availableIncrements > 0; nextElementIndex++ {
			// exit if equal, it's a match
			for nums[nextElementIndex] < currentNumber && availableIncrements > 0 {
				nums[nextElementIndex]++

				fmt.Println("nums[nextElementIndex]")
				fmt.Println(nums[nextElementIndex])
				fmt.Println("")

				// when numbers match, add frequency & go to the next target element
				if nums[nextElementIndex] == currentNumber {
					// if need to know the actual number in some case
					// frequencyMap[currentNumber]++
					localFrequency++
				}

				availableIncrements--
				fmt.Println("availableIncrements")
				fmt.Println(availableIncrements)
				fmt.Println("")
			}
		}

		if localFrequency > maxFrequency {
			maxFrequency = localFrequency
		}

		fmt.Println(maxFrequency)
		fmt.Println("")
	}

	return maxFrequency
}

func main_false() {
	sampleArray := []int{9930, 9923, 9983, 9997, 9934, 9952, 9945, 9914, 9985, 9982, 9970, 9932, 9985, 9902, 9975, 9990, 9922, 9990, 9994, 9937, 9996, 9964, 9943, 9963, 9911, 9925, 9935, 9945, 9933, 9916, 9930, 9938, 10000, 9916, 9911, 9959, 9957, 9907, 9913, 9916, 9993, 9930, 9975, 9924, 9988, 9923, 9910, 9925, 9977, 9981, 9927, 9930, 9927, 9925, 9923, 9904, 9928, 9928, 9986, 9903, 9985, 9954, 9938, 9911, 9952, 9974, 9926, 9920, 9972, 9983, 9973, 9917, 9995, 9973, 9977, 9947, 9936, 9975, 9954, 9932, 9964, 9972, 9935, 9946, 9966}
	sampleInt := 3056

	maxFreq := maxFrequency_false(sampleArray, sampleInt)
	fmt.Println(maxFreq)
}
