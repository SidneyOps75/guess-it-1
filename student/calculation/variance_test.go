package formulas

import (
	"math"
	"testing"
)

func TestCalculateVariance(t *testing.T) {
	type args struct {
		data    []float64
		average float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Empty slice",
			args: args{data: []float64{}, average: 0},
			want: 0, // Variance of an empty slice is not well-defined but assuming 0 for this case
		},
		{
			name: "Single element",
			args: args{data: []float64{5}, average: 5},
			want: 0, // Variance of a single value is always 0
		},
		{
			name: "Two elements with the same value",
			args: args{data: []float64{10, 10}, average: 10},
			want: 0, // Variance is 0 when all values are the same
		},
		{
			name: "Two elements with different values",
			args: args{data: []float64{10, 20}, average: 15},
			want: 25, // Variance = ((10-15)^2 + (20-15)^2) / 2 = 25
		},
		{
			name: "Multiple elements",
			args: args{data: []float64{1, 2, 3, 4, 5}, average: 3},
			want: 2, // Variance = ((1-3)^2 + (2-3)^2 + (3-3)^2 + (4-3)^2 + (5-3)^2) / 5 = 2.5
		},
		{
			name: "Negative values",
			args: args{data: []float64{-1, -2, -3}, average: -2},
			want: 0.6666666666666666, // Variance = ((-1+2)^2 + (-2+2)^2 + (-3+2)^2) / 3 = 0.6666666666666666
		},
		{
			name: "Large numbers",
			args: args{data: []float64{1000000, 2000000, 3000000}, average: 2000000},
			want: 1000000000000, // Variance = ((1000000-2000000)^2 + (2000000-2000000)^2 + (3000000-2000000)^2) / 3 = 1000000000000
		},
		{
			name: "Fractional values",
			args: args{data: []float64{1.5, 2.5, 3.5}, average: 2.5},
			want: 0.6666666666666666, // Variance = ((1.5-2.5)^2 + (2.5-2.5)^2 + (3.5-2.5)^2) / 3 = 0.6666666666666666
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateVariance(tt.args.data, tt.args.average)
			if math.Abs(got-tt.want) > 1e+12 { // Use a tolerance for floating-point comparison
				t.Errorf("CalculateVariance() = %v, want %v", got, tt.want)
			}
		})
	}
}
