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
		{args: args{3456}, want: 9},
		{args: args{12}, want: 3},
		{args: args{1}, want: 1},
		{args: args{0}, want: 0},
	}

	const name = "exercises9"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := calculateDigitalRoot(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
