package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises4"

	type args struct {
		a string
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{"ihgewlqlkot"}, want: "hello"},
		{args: args{"i"}, want: ""},
		{args: args{"i0"}, want: "0"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises4(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
