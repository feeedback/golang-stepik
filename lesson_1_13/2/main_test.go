package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises2"

	type args struct {
		a int16
	}

	tests := []struct {
		args args
		want int16
	}{
		{args: args{102}, want: 201},
		{args: args{321}, want: 123},
		{args: args{222}, want: 222},
		{args: args{221}, want: 122},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises2(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
