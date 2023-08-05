package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises5"

	type args struct {
		a, b, c int
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{4, 5, 6}, want: "Существует"},
		{args: args{1, 1, 0}, want: "Не существует"},
		{args: args{5, 2, 1}, want: "Не существует"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises5(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
