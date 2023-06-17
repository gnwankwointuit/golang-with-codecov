package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_abs(t *testing.T) {
	type args struct {
		a int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "positive number",
			args: args{a: 3},
			want: 3,
		},
		{
			name: "negative number",
			args: args{a: -3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, abs(tt.args.a), "abs(%v)", tt.args.a)
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "positive numbers",
			args: args{a: 3, b: 3},
			want: 9,
		},
		{
			name: "negative numbers",
			args: args{a: -3, b: -3},
			want: 9,
		},
		{
			name: "mixed positive/negative numbers",
			args: args{a: -3, b: 3},
			want: -9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, multiply(tt.args.a, tt.args.b), "multiply(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func Test_subtract(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "positive numbers",
			args: args{a: 3, b: 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, subtract(tt.args.a, tt.args.b), "subtract(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "positive numbers",
			args: args{a: 3, b: 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sum(tt.args.a, tt.args.b), "sum(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
