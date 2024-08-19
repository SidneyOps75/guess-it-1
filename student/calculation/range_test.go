package calc

import "testing"

func TestRange(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "Empty slice",
			args:  args{data: []float64{}},
			want:  -9223372036854775808, // Expected values need to be determined based on implementation
			want1: -9223372036854775808,
		},
		{
			name:  "Single element",
			args:  args{data: []float64{10}},
			want:  10,
			want1: 10,
		},
		{
			name:  "Multiple positive elements",
			args:  args{data: []float64{10, 20, 30, 40}},
			want:  2, // Example values, adjust according to actual calculation
			want1: 48,
		},
		{
			name:  "Multiple negative elements",
			args:  args{data: []float64{-10, -20, -30, -40}},
			want:  -48,
			want1: -2,
		},
		{
			name:  "Mixed positive and negative elements",
			args:  args{data: []float64{-10, 20, -30, 40}},
			want:  -52, // Example values, adjust according to actual calculation
			want1: 62,
		},
		{
			name:  "Fractional elements",
			args:  args{data: []float64{1.5, 2.5, 3.5}},
			want:  1, // Example values, adjust according to actual calculation
			want1: 4,
		},
		{
			name:  "Large numbers",
			args:  args{data: []float64{1000000, 2000000, 3000000}},
			want:  285357,
			want1: 3714643,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Range(tt.args.data)
			if got != tt.want {
				t.Errorf("Range() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Range() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
