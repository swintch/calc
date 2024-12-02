package calc

import "testing"

func TestAddition_Calculate(t1 *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Addition{}
			if got := t.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t1 *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Division{}
			if got := t.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication_Calculate(t1 *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Multiplication{}
			if got := t.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction_Calculate(t1 *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Subtraction{}
			if got := t.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
