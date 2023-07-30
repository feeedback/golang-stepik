package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "conditionalConstructs_4"

	type args struct {
		a int
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{613244}, want: "YES"},
		{args: args{613243}, want: "NO"},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := conditionalConstructs_4(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
