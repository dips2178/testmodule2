package calc

// returns sum and product of two integers
func Add(numbers ...int) (result int) {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}
