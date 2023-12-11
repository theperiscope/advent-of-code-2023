package lib

import (
	"cmp"
)

// Computes minimum from specified values; can be used with slices with "sliceName..." syntax
func Min[T cmp.Ordered](values ...T) T {
	var min T

	for i, value := range values {
		if i == 0 || value < min {
			min = value
		}
	}

	return min
}

// Computes maximum from specified values; can be used with slices with "sliceName..." syntax
func Max[T cmp.Ordered](values ...T) T {
	var max T

	for i, value := range values {
		if i == 0 || value > max {
			max = value
		}
	}

	return max
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Computes sum from specified values; can be used with slices with "sliceName..." syntax
func Sum[T Number](values ...T) T {
	var sum T

	for _, value := range values {
		sum += value
	}

	return sum
}

// Computes sum where value is adjusted by a provided function
func SumWithAdjustment[T Number](f func(int, T) T, values ...T) T {
	var sum T

	for index, value := range values {
		sum += f(index, value)
	}

	return sum
}

// returns absolute value of v
func Abs[T Number](v T) T {
	if v < 0 {
		return -v
	}

	return v
}

// returns absolute difference between the two values
func AbsDiff[T Number](x, y T) T {
	if x < y {
		return y - x
	}
	return x - y
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
