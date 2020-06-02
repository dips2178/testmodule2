package calc

// returns sum and product of two integers
func AddAndMul(numbers ...int) (sum int,prod int) {
	sum = 0
	prod = 1
	for _, num := range numbers {
		sum = sum + num
		prod = prod * num
	}
	return sum, prod
}
