package calc
import "fmt"

// returns sum and product of two integers
func Add(numbers ...int) (sum int) {
	sum = 0
	prod := 1
	for _, num := range numbers {
		sum = sum + num
		prod = prod * num
	}
	fmt.Println("Prod=",prod)
	return sum
}
