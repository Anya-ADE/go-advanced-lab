package main

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "negative input", input: -1, want: 0, wantErr: true},
		{name: "negative large", input: -10, want: 0, wantErr: true},
		{name: "factorial of 3", input: 3, want: 6, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{"prime 2", 2, true, false},
		{"prime 17", 17, true, false},
		{"composite 20", 20, false, false},
		{"composite 25", 25, false, false},
		{"edge case 1", 1, false, true},
		{"negative number", -5, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{"2^8", 2, 8, 256, false},
		{"5^3", 5, 3, 125, false},
		{"base^0", 10, 0, 1, false},
		{"0^n", 0, 5, 0, false},
		{"negative exponent", 2, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeCounter(t *testing.T) {
	t.Run("counters are independent", func(t *testing.T) {
		c1 := MakeCounter(0)
		c2 := MakeCounter(10)

		if res := c1(); res != 1 {
			t.Errorf("c1 expected 1, got %d", res)
		}
		if res := c2(); res != 11 {
			t.Errorf("c2 expected 11, got %d", res)
		}
		if res := c1(); res != 2 {
			t.Errorf("c1 expected 2, got %d", res)
		}
	})
}

func TestMakeMultiplier(t *testing.T) {
	{
		double := MakeMultiplier(2)
		triple := MakeMultiplier(3)

		if res := double(5); res != 10 {
			t.Errorf("double expected 10, got %d", res)
		}
		if res := triple(5); res != 15 {
			t.Errorf("triple expected 15, got %d", res)
		}
		if res := double(10); res != 20 {
			t.Errorf("double expected 20, got %d", res)
		}
	}
}

func TestMakeAccumulator(t *testing.T) {
	add, sub, get := MakeAccumulator(100)
	add(50)
	if get() != 150 {
		t.Errorf("expected 150, got %d", get())
	}
	sub(30)
	if get() != 120 {
		t.Errorf("expected 120, got %d", get())
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		want      []int
	}{
		{
			name:      "square numbers",
			nums:      []int{1, 2, 3, 4},
			operation: func(x int) int { return x * x },
			want:      []int{1, 4, 9, 16},
		},
		{
			name:      "negate numbers",
			nums:      []int{10, -5, 0},
			operation: func(x int) int { return -x },
			want:      []int{-10, 5, 0},
		},
		{
			name:      "double numbers",
			nums:      []int{1, 2, 3},
			operation: func(x int) int { return x * 2 },
			want:      []int{2, 4, 6},
		},
		{
			name:      "empty slice",
			nums:      []int{},
			operation: func(x int) int { return x + 1 },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.nums, tt.operation)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "filter even numbers",
			nums:      []int{1, 2, 3, 4, 5, 6},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int{2, 4, 6},
		},
		{
			name:      "filter numbers greater than 10",
			nums:      []int{5, 12, 8, 20, 3},
			predicate: func(x int) bool { return x > 10 },
			want:      []int{12, 20},
		},
		{
			name:      "filter positive numbers",
			nums:      []int{-2, -1, 0, 1, 2},
			predicate: func(x int) bool { return x > 0 },
			want:      []int{1, 2},
		},
		{
			name:      "filter no matches",
			nums:      []int{1, 3, 5},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.predicate)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(int, int) int
		want      int
	}{
		{
			name:      "sum of numbers",
			nums:      []int{1, 2, 3, 4},
			initial:   0,
			operation: func(acc, curr int) int { return acc + curr },
			want:      10,
		},
		{
			name:      "product of numbers",
			nums:      []int{1, 2, 3, 4},
			initial:   1,
			operation: func(acc, curr int) int { return acc * curr },
			want:      24,
		},
		{
			name:    "find maximum",
			nums:    []int{5, 12, 3, 8, 2},
			initial: 0,
			operation: func(acc, curr int) int {
				if curr > acc {
					return curr
				}
				return acc
			},
			want: 12,
		},
		{
			name:      "empty slice returns initial",
			nums:      []int{},
			initial:   100,
			operation: func(acc, curr int) int { return acc + curr },
			want:      100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.nums, tt.initial, tt.operation)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	addOne := func(x int) int { return x + 1 }
	double := func(x int) int { return x * 2 }
	square := func(x int) int { return x * x }

	tests := []struct {
		name  string
		f     func(int) int
		g     func(int) int
		input int
		want  int
	}{
		{
			name:  "double then add one",
			f:     addOne,
			g:     double,
			input: 5,
			want:  11,
		},
		{
			name:  "add one then double",
			f:     double,
			g:     addOne,
			input: 5,
			want:  12,
		},
		{
			name:  "square then double",
			f:     double,
			g:     square,
			input: 3,
			want:  18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composedFunc := Compose(tt.f, tt.g)
			got := composedFunc(tt.input)
			if got != tt.want {
				t.Errorf("Compose() result = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleValue(t *testing.T) {
	t.Run("Value should not change original", func(t *testing.T) {
		val := 10
		DoubleValue(val)
		if val != 10 {
			t.Errorf("DoubleValue modified the original! Got %d, want 10", val)
		}
	})
}

func TestDoublePointer(t *testing.T) {
	t.Run("Value should be doubled", func(t *testing.T) {
		val := 15
		DoublePointer(&val)
		if val != 30 {
			t.Errorf("DoublePointer failed! Got %d, want 30", val)
		}
	})
}

func TestSwapValues(t *testing.T) {
	tests := []struct {
		name         string
		a, b         int
		wantA, wantB int
	}{
		{"swap positives", 5, 10, 10, 5},
		{"swap with zero", 0, 100, 100, 0},
		{"swap negatives", -1, -5, -5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapValues(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("SwapValues(%d, %d) = %d, %d; want %d, %d", tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestSwapPointers(t *testing.T) {
	tests := []struct {
		name         string
		a, b         int
		wantA, wantB int
	}{
		{"pointer swap 1", 10, 20, 20, 10},
		{"pointer swap 2", 100, 200, 200, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localA, localB := tt.a, tt.b
			SwapPointers(&localA, &localB)
			if localA != tt.wantA || localB != tt.wantB {
				t.Errorf("SwapPointers() failed: a=%d, b=%d; want a=%d, b=%d", localA, localB, tt.wantA, tt.wantB)
			}
		})
	}
}
