package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises3"

	type args struct {
		a string
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{"1112221112"}, want: 2},
		{args: args{"3"}, want: 3},
		{args: args{"0"}, want: 0},
		{args: args{"091234"}, want: 9},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises3(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
