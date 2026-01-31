package main

import (
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
