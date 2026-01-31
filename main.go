package main

import (
	"errors"
	"fmt"
	"math"
	"os"
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

func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	accumulator := initial
	add = func(n int) {
		accumulator += n
	}
	subtract = func(n int) {
		accumulator -= n
	}
	get = func() int {
		return accumulator
	}
	return add, subtract, get
}

func Apply(num []int, operation func(int) int) []int {
	result := make([]int, len(num))
	for i, n := range num {
		result[i] = operation(n)
	}
	return result
}

func Filter(nums []int, predict func(int) bool) []int {
	result := []int{}
	for _, n := range nums {
		if predict(n) {
			result = append(result, n)
		}
	}
	return result
}

func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	accumulator := initial
	for _, n := range nums {
		accumulator = operation(accumulator, n)
	}
	return accumulator
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ExploreProcess() {
	fmt.Println("=== Process Information ===")
	// A Process ID (PID) is a unique numerical identifier assigned by the
	// operating system to every active process
	fmt.Printf("Current Process ID: %d\n", os.Getpid())
	fmt.Printf("Parent Process ID: %d\n", os.Getppid())
	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice data: %v\n", data)
	fmt.Printf("Memory address of slice header: %p\n", &data)
	fmt.Printf("Memory address of first element: %p\n", &data[0])
	// Process isolation is a security feature that prevents one process from
	// accessing the memory of another process, ensuring stability and privacy.
	fmt.Println("\nNote: Other processes cannot access these memory addresses due to process isolation.")
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

	fmt.Println("=== Counter Function ===")
	counter1 := MakeCounter(0)
	fmt.Println(counter1())
	fmt.Println(counter1())

	counter2 := MakeCounter(10)
	fmt.Println(counter2())
	fmt.Println(counter1())
	fmt.Println("===")

	fmt.Println("=== Multiplier Function ===")
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)
	fmt.Println(double(5))
	fmt.Println(triple(5))
	fmt.Println("===")

	fmt.Println("=== Accumulator Function ===")
	add, sub, get := MakeAccumulator(100)
	add(50)
	fmt.Println(get())
	sub(30)
	fmt.Println(get())
	fmt.Println("===")

	fmt.Println("=== Higher-Order Functions ===")
	nums := []int{1, 2, 3, 4}
	squared := Apply(nums, func(n int) int { return n * n })
	fmt.Printf("Apply (square): %v\n", squared)

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Filter (even): %v\n", evens)

	sum := Reduce(nums, 0, func(acc, curr int) int { return acc + curr })
	fmt.Printf("Reduce (sum): %d\n", sum)

	product := Reduce(nums, 1, func(acc, curr int) int { return acc * curr })
	fmt.Printf("Reduce (product): %d\n", product)

	addTwo := func(x int) int { return x + 2 }
	double = func(x int) int { return x * 2 }

	doubleTHenAddTwo := Compose(addTwo, double)
	result := doubleTHenAddTwo(5)
	fmt.Printf("Compose (double then add two) on 5: %d\n", result)
	fmt.Println("===")

	ExploreProcess()
}
