package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "conditionalConstructs_3"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{1234}, want: 1},
		{args: args{10000}, want: 1},
		{args: args{2000}, want: 2},
		{args: args{300}, want: 3},
		{args: args{40}, want: 4},
		{args: args{5}, want: 5},
		{args: args{0}, want: 0},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := conditionalConstructs_3(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
