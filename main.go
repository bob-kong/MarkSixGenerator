package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// Seed the random number generator to ensure different results each time.
	rand.Seed(time.Now().UnixNano())

	// Create a slice to store the random digits.
	randomDigits := make([]int, 6)

	for {
		// Generate 6 unique random digits between 1 and 49.
		for i := 0; i < 6; {
			// Generate a random number between 1 and 49.
			randomDigit := rand.Intn(49) + 1

			// Check if the generated digit is unique.
			unique := true
			for j := 0; j < i; j++ {
				if randomDigit == randomDigits[j] {
					unique = false
					break
				}
			}

			// If it's unique, add it to the slice.
			if unique {
				randomDigits[i] = randomDigit
				i++
			}
		}

		sort.Ints(randomDigits)

		//Condition Checking
		if ContainsMixOfOddAndEven(randomDigits) &&
			!ContainsConsecutiveNumbers(randomDigits) &&
			IsSumInRange(randomDigits) &&
			IsMinValueGreaterThanOrEqualTo(randomDigits) &&
			IsLargestValueInRange(randomDigits) &&
			OnlyOneDoorMissing(randomDigits) {
			break
		}
	}

	// Print the sorted slice of random digits.
	fmt.Println("Generated numbers meet all conditions:", randomDigits)
}
