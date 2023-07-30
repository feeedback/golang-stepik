package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{20, 3, 5}, want: 3},
		{args: args{20, 5, 2}, want: 5},
		{args: args{20, 21, 2}, want: -1},
	}

	const name = "cycles_5"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := cycles_5(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
