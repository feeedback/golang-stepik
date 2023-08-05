package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises1"

	type args struct {
		a int16
	}

	tests := []struct {
		args args
		want int16
	}{
		{args: args{321}, want: 6},
		{args: args{102}, want: 3},
		{args: args{322}, want: 7},
		{args: args{222}, want: 6},
		{args: args{232}, want: 7},
		{args: args{223}, want: 7},
		{args: args{0}, want: 0},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises1(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
