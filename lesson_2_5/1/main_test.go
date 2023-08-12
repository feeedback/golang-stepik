package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises1"

	type args struct {
		a string
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{"Прямоугольный."}, want: "Right"},
		{args: args{"G123."}, want: "Right"},
		{args: args{"E."}, want: "Right"},
		{args: args{"E"}, want: "Wrong"},
		{args: args{""}, want: "Wrong"},
		{args: args{"."}, want: "Wrong"},
		{args: args{"a."}, want: "Wrong"},
		{args: args{"aasd"}, want: "Wrong"},
		{args: args{"aE."}, want: "Wrong"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises1(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
