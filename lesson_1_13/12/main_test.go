package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		a int
	}

	tests := []struct {
		args args
		want []int
	}{
		{args: args{1}, want: []int{1}},
		{args: args{0}},
		{args: args{3}, want: []int{1, 2}},
		{args: args{50}, want: []int{1, 2, 4, 8, 16, 32}},
	}

	const name = "exercises12"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises12(tt.args.a); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
