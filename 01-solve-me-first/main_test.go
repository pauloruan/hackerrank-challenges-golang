package main

import (
	"testing"
)

func Test_SolveMeFirst(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{1, 2}, 3},
		{"Test 2", args{1, -2}, -1},
		{"Test 3", args{-1, 0}, -1},
		{"Test 4", args{-1, -2}, -3},
		{"Test 5", args{-1, 2}, 1},
		{"Test 6", args{1, 0}, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := SolveMeFirst(test.args.a, test.args.b); got != test.want {
				t.Errorf("%s - Expected: %v, Received: %v", test.name, test.want, got)
			}
		})
	}
}
