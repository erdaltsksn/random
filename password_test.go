package random_test

import (
	"testing"

	"github.com/erdaltsksn/random"
)

func TestPassword(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 Character", args{length: 1}, 1},
		{"2 Characters", args{length: 2}, 2},
		{"10 Characters", args{length: 10}, 10},
		{"99 Characters", args{length: 99}, 99},
	}
	for _, tt := range tests {
		for i := 0; i < 1000; i++ {
			got := random.Password(tt.args.length)
			t.Run(tt.name, func(t *testing.T) {
				if len(got) != tt.want {
					t.Errorf("Got: %s (length:%d) Want: %d", got, len(got), tt.want)
				}
			})
		}
	}
}
