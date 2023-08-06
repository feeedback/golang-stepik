package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises_2_1"

	type args struct {
		a, b, c, d int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{6, 8, 1, 10}, want: 1},
		{args: args{6, 8, 1, 0}, want: 0},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises_2_1(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
