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
		{args: args{"9119"}, want: "811181"},
		{args: args{"1"}, want: "1"},
		{args: args{"2"}, want: "4"},
		{args: args{"9"}, want: "81"},
		{args: args{"0"}, want: "0"},
		{args: args{"34"}, want: "916"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises4(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
