package main

func main() {
	Sum([]int{1, 2, 3, 4, 5})
}

// Sum adds the numbers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll sums all lists
func SumAll(numbers ...[]int) []int {
	var result []int
	for _, number := range numbers {
		result = append(result, Sum(number))
	}
	return result
}
