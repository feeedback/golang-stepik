package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		args args
		want []int
	}{
		{args: args{[]int{30, 11, 7, 101}}, want: []int{30, 11}},
		{args: args{[]int{30, 11, 7, 100, 109}}, want: []int{30, 11, 100}},
	}

	const name = "cycles_6"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := cycles_6(tt.args.nums); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
