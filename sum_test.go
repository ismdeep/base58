package base58

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		v []int64
		h int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				v: []int64{13, 12, 4, 11, 15, 7, 8, 2, 0, 0, 0, 0},
				h: 16,
			},
			want: 679457997,
		},
		{
			name: "",
			args: args{
				v: []int64{3, 11, 23, 2, 2, 1, 0, 0, 0},
				h: 58,
			},
			want: 679457997,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.v, tt.args.h); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
