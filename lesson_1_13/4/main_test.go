package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises4"

	type args struct {
		a, b, c float64
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{6, 8, 10}, want: "Прямоугольный"},
		{args: args{6, 8, 11}, want: "Непрямоугольный"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises4(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
