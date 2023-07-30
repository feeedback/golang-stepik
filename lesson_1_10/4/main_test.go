package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		args args
		want int
	}{
		{args: args{[]int{1, 3, 3, 1, 0}}, want: 2},
		{args: args{[]int{1, 3, 3, 1, 1}}, want: 2},
		{args: args{[]int{3, 3, 3, 0}}, want: 3},
		{args: args{[]int{}}, want: 0},
	}

	const name = "cycles_4"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := countMaxElements(tt.args.nums); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
