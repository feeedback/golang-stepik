package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises1"

	type args struct {
		a float64
		b float64
	}

	tests := []struct {
		args args
		want float64
	}{
		{args: args{6, 8}, want: 10},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := calculateHypotenuse(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
