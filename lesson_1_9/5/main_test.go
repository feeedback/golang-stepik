package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "conditionalConstructs_5"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{2000}, want: "YES"},
		{args: args{2004}, want: "YES"},
		{args: args{1996}, want: "YES"},
		{args: args{2001}, want: "NO"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := conditionalConstructs_5(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
