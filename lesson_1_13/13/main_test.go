package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		a int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{8}, want: 6},
		{args: args{2}, want: 3},
		{args: args{4}, want: -1},
		{args: args{3}, want: 4},
	}

	const name = "exercises13"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises13(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
