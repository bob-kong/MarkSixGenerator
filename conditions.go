package main

func ContainsMixOfOddAndEven(numbers []int) bool {
	hasOdd := false
	hasEven := false

	for _, num := range numbers {
		if num%2 == 0 {
			hasEven = true
		} else {
			hasOdd = true
		}

		// If both odd and even numbers are found, return true.
		if hasOdd && hasEven {
			return true
		}
	}

	// If we reach this point, the slice does not contain a mix of odd and even numbers.
	return false
}

func ContainsConsecutiveNumbers(numbers []int) bool {
	consecutiveCount := 1

	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1]+1 {
			consecutiveCount++
		} else {
			consecutiveCount = 1
		}

		if consecutiveCount >= 3 {
			return true
		}
	}

	return false
}

func IsSumInRange(numbers []int) bool {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum >= 91 && sum <= 210
}

func IsMinValueGreaterThanOrEqualTo(numbers []int) bool {
	for _, num := range numbers {
		if num < 20 {
			return true
		}
	}
	return false
}

func IsLargestValueInRange(numbers []int) bool {
	max := numbers[len(numbers)-1]
	return max >= 40 && max <= 49
}

func OnlyOneDoorMissing(numbers []int) bool {
	doorCount := [5]int{0, 0, 0, 0, 0} // Initialize door counts for each range.

	for _, num := range numbers {
		if num >= 1 && num <= 9 {
			doorCount[0]++
		} else if num >= 10 && num <= 19 {
			doorCount[1]++
		} else if num >= 20 && num <= 29 {
			doorCount[2]++
		} else if num >= 30 && num <= 39 {
			doorCount[3]++
		} else if num >= 40 && num <= 49 {
			doorCount[4]++
		}
	}

	missingDoors := 0
	for _, count := range doorCount {
		if count == 0 {
			missingDoors++
		}
	}

	return missingDoors == 1
}
