package calc

func Add(args ...int) {
	var sum int
	for _, val := range args {
		sum = sum + val
	}
}
