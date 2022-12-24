package util

type OperationFunc func(int, int) int

func AdditionFunc(a, b int) int {
	return a + b
}

func SubtractionFunc(a, b int) int {
	return a - b
}

func MultiplicationFunc(a, b int) int {
	return a * b
}

func DivisionFunc(a, b int) int {
	return a / b
}
