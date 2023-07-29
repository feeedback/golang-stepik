package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "arithmeticCalculations_4"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{0}, want: 0},
		{args: args{9}, want: 9},
		{args: args{123}, want: 3},
		{args: args{12}, want: 2},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := arithmeticCalculations_4(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
