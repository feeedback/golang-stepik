package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises6"

	type args struct {
		a, b int
	}

	tests := []struct {
		args args
		want float64
	}{
		{args: args{4, 5}, want: 4.5},
		{args: args{1, 1}, want: 1},
		{args: args{75, 5}, want: 40},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises6(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
