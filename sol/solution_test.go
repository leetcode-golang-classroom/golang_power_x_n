package sol

import (
	"math"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	x := 2.10000
	n := 3
	for idx := 0; idx < b.N; idx++ {
		myPow(x, n)
	}
}
func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "x = 2.00000, n = 10",
			args: args{x: 2.00000, n: 10},
			want: 1024.00000,
		},
		{
			name: "x = 2.10000, n = 3",
			args: args{x: 2.10000, n: 3},
			want: 9.26100,
		},
		{
			name: "x = 2.00000, n = -2",
			args: args{x: 2.00000, n: -2},
			want: 0.25000,
		},
	}
	for _, tt := range tests {
		tolerance := 0.00001
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); math.Abs(got-tt.want) > tolerance {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
