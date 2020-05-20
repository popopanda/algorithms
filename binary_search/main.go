package main

import "fmt"

func main() {
	primeNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	binarySearch(primeNums, 9)
}

func binarySearch(sliceVar []int, masterNum int) {
	min := 0
	max := len(sliceVar) - 1

	for min <= max {
		mid := (min + max) / 2
		guess := sliceVar[mid]

		if guess == masterNum {
			fmt.Printf("Guess right! value: %v\n", guess)
			return
		}
		if guess > masterNum {
			fmt.Printf("Value: %v not correct\n", guess)
			max = mid - 1
		} else {
			fmt.Printf("Value: %v not correct\n", guess)
			min = mid + 1
		}
	}

}
