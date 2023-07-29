package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "arithmeticCalculations_2"

	type args struct {
		a int
		b int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{0, 0}, want: 0},
		{args: args{2, 2}, want: 8},
		{args: args{2, 3}, want: 13},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := arithmeticCalculations_2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
