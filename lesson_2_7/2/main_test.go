package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	const name = "exercises2"

	type args struct {
		a string
	}

	tests := []struct {
		args args
		want string
	}{
		{args: args{"LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO"}, want: "L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O"},
		{args: args{"ab"}, want: "a*b"},
		{args: args{"a"}, want: "a"},
		{args: args{""}, want: ""},
	}

	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v args:%v", name, tt.args), func(t *testing.T) {

			if got := exercises2(tt.args.a); got != tt.want {
				t.Errorf("= %v, but want %v", got, tt.want)
			}
		})
	}
}
