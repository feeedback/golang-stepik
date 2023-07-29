package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "conditionalConstructs"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{1}, want: "Число положительное"},
		{args: args{-1}, want: "Число отрицательное"},
		{args: args{0}, want: "Ноль"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := conditionalConstructs(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
