package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises5"

	type args struct {
		nums []int
	}

	tests := []struct {
		args args
		want struct{ a, b int }
	}{
		{args: args{[]int{1, 0}}, want: struct{ a, b int }{2, 1}},
		{args: args{[]int{1, 1, 10}}, want: struct{ a, b int }{3, 12}},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {
			a, b := sumInt(tt.args.nums...)
			if a != tt.want.a || b != tt.want.b {
				t.Errorf("= %v, but want %v", []int{a, b}, tt.want)
			}
		})
	}
}
