package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{1, 5}, want: 15},
		{args: args{1, 1}, want: 1},
		{args: args{1, 2}, want: 3},
	}

	const name = "cycles_2"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := cycles_2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
