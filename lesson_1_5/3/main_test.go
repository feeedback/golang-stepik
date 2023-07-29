package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "arithmeticCalculations_3"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{0}, want: 0},
		{args: args{1}, want: 1},
		{args: args{3}, want: 9},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := arithmeticCalculations_3(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
