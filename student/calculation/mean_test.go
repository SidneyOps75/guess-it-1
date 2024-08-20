package formulas

import "testing"

func TestCalculateAverage(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Empty slice",
			args: args{data: []float64{}},
			want: 0,
		},
		{
			name: "Single element",
			args: args{data: []float64{5}},
			want: 5,
		},
		{
			name: "Multiple elements",
			args: args{data: []float64{1, 2, 3, 4, 5}},
			want: 3,
		},
		{
			name: "Negative numbers",
			args: args{data: []float64{-1, -2, -3, -4, -5}},
			want: -3,
		},
		{
			name: "Mixed positive and negative",
			args: args{data: []float64{-10, 0, 10}},
			want: 0,
		},
		{
			name: "Fractional numbers",
			args: args{data: []float64{1.5, 2.5, 3.5}},
			want: 2.5,
		},
		{
			name: "Large numbers",
			args: args{data: []float64{1e6, 2e6, 3e6}},
			want: 2e6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateAverage(tt.args.data); got != tt.want {
				t.Errorf("CalculateAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
