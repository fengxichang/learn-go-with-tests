package list

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func SumALl(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}