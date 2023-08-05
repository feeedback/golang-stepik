package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		a, b int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{100, 500}, want: 497},
		{args: args{1, 8}, want: 7},
	}

	const name = "exercises10"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises10(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
