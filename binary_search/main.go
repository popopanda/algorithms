package main

import "fmt"

func main() {
	primeNums := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	binarySearch(primeNums, 3)
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
			fmt.Printf("Value %v is too high! \n", guess)
			max = mid - 1
		} else {
			fmt.Printf("Value: %v is too low! \n", guess)
			min = mid + 1
		}
	}

}
