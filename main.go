package main

import (
	"errors"
	"fmt"
	"math"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	if n == 2 {
		return true, nil
	}
	if n%2 == 0 {
		return false, nil
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	if exponent == 0 {
		return 1, nil
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

func main() {
	fmt.Println("=== Math Operations ===")
	f5, _ := Factorial(5)
	fmt.Printf("Factorial(5) = %d\n", f5)

	p17, _ := IsPrime(17)
	fmt.Printf("IsPrime(17) = %v\n", p17)

	pow, _ := Power(2, 8)
	fmt.Printf("Power(2, 8) = %d\n", pow)
	fmt.Println("===")
}
