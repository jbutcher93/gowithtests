package main

func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	newSlice := []int{}
	for _, slice := range slices {
		newSlice = append(newSlice, Sum(slice))
	}
	return newSlice
}

func SumAllTails(slices ...[]int) []int {
	newSlice := []int{}
	for _, slice := range slices {
		if len(slice) == 0 {
			newSlice = append(newSlice, 0)
		} else {
			newSlice = append(newSlice, Sum(slice[1:]))
		}
	}
	return newSlice
}
