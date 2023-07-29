package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "arithmeticCalculations_6"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{90}, want: "It is 3 hours 0 minutes."},
		{args: args{195}, want: "It is 6 hours 30 minutes."},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := arithmeticCalculations_6(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
