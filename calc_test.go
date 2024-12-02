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
		{name: "1+2=3", args: args{a: 1, b: 2}, want: 3},
		{name: "0+0=0", args: args{a: 0, b: 0}, want: 0},
		{name: "-2+2=0", args: args{a: -2, b: 2}, want: 0},
		{name: "2+-2=0", args: args{a: 2, b: -2}, want: 0},
		{name: "-2+-2=-4", args: args{a: -2, b: -2}, want: -4},
		{name: "1+0=1", args: args{a: 1, b: 0}, want: 1},
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
		{name: "6/2=3", args: args{a: 6, b: 2}, want: 3},
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
		{name: "2*2=4", args: args{a: 2, b: 2}, want: 4},
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
		{name: "2-2=0", args: args{a: 2, b: 2}, want: 0},
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
