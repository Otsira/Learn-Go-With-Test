package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(numsToSum ...[]int) []int {
	lenofNums := len(numsToSum)
	sums := make([]int, lenofNums)
	for i, num := range numsToSum {
		sums[i] = Sum(num)
	}
	return sums
}
