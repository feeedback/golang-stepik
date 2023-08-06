package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		num, index int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{38012732, 3}, want: 801272},
		{args: args{2221222, 2}, want: 1},
		{args: args{123, 4}, want: 123},
		{args: args{123, 2}, want: 13},
	}

	const name = "exercises15"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises15(tt.args.num, tt.args.index); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
