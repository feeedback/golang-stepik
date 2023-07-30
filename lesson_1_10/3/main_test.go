package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		n    int
		nums []int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{5, []int{38, 24, 800, 8, 16}}, want: 40},
	}

	const name = "cycles_3"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := cycles_3(tt.args.n, tt.args.nums); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
