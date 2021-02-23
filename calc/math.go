package calc

import "fmt"

func Add(numbers ...int) (error, int) {
	var sum int

	if len(numbers) < 2 {
		return fmt.Errorf("provide more than 2 number"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return nil, sum
	}
}
