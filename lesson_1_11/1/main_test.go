package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "stringFormat_1"

	type args struct {
		a float64
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{-000012.2123}, want: "число -12.21 не подходит"},
		{args: args{1000001}, want: "1.000001e+06"},
		{args: args{12.12345678}, want: "146.9782"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := stringFormat_1(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
