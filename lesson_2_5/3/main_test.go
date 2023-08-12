package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises3"

	type args struct {
		a string
		b string
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{"awesome", "es"}, want: 2},
		{args: args{"awesome", "a"}, want: 0},
		{args: args{"awesome", "b"}, want: -1},
		{args: args{"awesome", "b"}, want: -1},
		{args: args{"awesome", "awesome"}, want: 0},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises3(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
