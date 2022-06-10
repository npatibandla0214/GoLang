package app

func xAdd(x int, y int) int {
	var sum int // Alternate sum := x+y
	sum = x + y
	return sum
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Multiplication(x, y int) int {
	return x * y
}

func Division(x, y int) (int, int) {
	return x / y, x % y
}
