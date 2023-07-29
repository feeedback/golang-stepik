package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "arithmeticCalculations"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{0}, want: 100},
		{args: args{1}, want: 102},
		{args: args{22}, want: 144},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := arithmeticCalculations(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
