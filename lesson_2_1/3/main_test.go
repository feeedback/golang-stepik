package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises3"

	type args struct {
		a, b, c int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{0, 0, 1}, want: 0},
		{args: args{0, 0, 0}, want: 0},
		{args: args{1, 1, 1}, want: 1},
		{args: args{1, 1, 0}, want: 1},
		{args: args{1, 0, 1}, want: 1},
		{args: args{1, 0, 1}, want: 1},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := vote(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
