package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises5"

	type args struct {
		a string
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{"zaabcbd"}, want: "zcd"},
		{args: args{"abcabc"}, want: ""},
		{args: args{"11"}, want: ""},
		{args: args{"211"}, want: "2"},
		{args: args{""}, want: ""},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises5(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
