package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(numsToSum ...[]int) (sums []int) {
	for _, num := range numsToSum {
		sums = append(sums, Sum(num))
	}
	return sums
}
