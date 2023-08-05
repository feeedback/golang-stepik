package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises11"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{1}, want: "1 korova"},
		{args: args{2}, want: "2 korovy"},
		{args: args{10}, want: "10 korov"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises11(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
