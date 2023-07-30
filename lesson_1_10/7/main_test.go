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
		{args: args{100, 10, 200}, want: 8},
	}

	const name = "cycles_7"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := cycles_7(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
