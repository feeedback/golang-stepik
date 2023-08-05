package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises3"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{59}, want: "It is 0 hours 0 minutes."},
		{args: args{60}, want: "It is 0 hours 1 minutes."},
		{args: args{90}, want: "It is 0 hours 1 minutes."},
		{args: args{3600}, want: "It is 1 hours 0 minutes."},
		{args: args{13257}, want: "It is 3 hours 40 minutes."},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises3(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
