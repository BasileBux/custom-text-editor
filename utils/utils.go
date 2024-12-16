package utils

// utils package is a collection of missing stuff in golang

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func PushFront(x []int, y int) []int {
	x = append([]int{y}, x...)
	return x
}
