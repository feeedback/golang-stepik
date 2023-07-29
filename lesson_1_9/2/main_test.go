package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "conditionalConstructs_2"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{321}, want: "YES"},
		{args: args{102}, want: "YES"},
		{args: args{322}, want: "NO"},
		{args: args{222}, want: "NO"},
		{args: args{232}, want: "NO"},
		{args: args{223}, want: "NO"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := conditionalConstructs_2(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
