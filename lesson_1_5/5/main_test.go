package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "arithmeticCalculations_5"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{2010}, want: 1},
		{args: args{235}, want: 3},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := arithmeticCalculations_5(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
