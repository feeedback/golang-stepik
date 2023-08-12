package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises6"

	type args struct {
		a string
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{"fdsghdfgjsdDD1"}, want: "Ok"},
		{args: args{"12345"}, want: "Ok"},
		{args: args{"1234"}, want: "Wrong password"},
		{args: args{"1234-"}, want: "Wrong password"},
		{args: args{""}, want: "Wrong password"},
		{args: args{"пппппппп99"}, want: "Wrong password"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises6(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
