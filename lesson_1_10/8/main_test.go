package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		args args
		want []string
	}{
		{args: args{564, 8954}, want: []string{"5", "4"}},
	}

	const name = "cycles_8"

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := getCommonDigits(tt.args.a, tt.args.b); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
