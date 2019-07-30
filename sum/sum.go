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
