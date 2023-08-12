package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises2"

	type args struct {
		a string
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{"abb"}, want: "Нет"},
		{args: args{"ab"}, want: "Нет"},
		{args: args{"AA"}, want: "Палиндром"},
		{args: args{"abba"}, want: "Палиндром"},
		{args: args{"ABBA"}, want: "Палиндром"},
		{args: args{"ABABA"}, want: "Палиндром"},
		{args: args{"бабаб"}, want: "Палиндром"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises2(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
