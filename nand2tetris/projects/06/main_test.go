package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hackAsm"
)

func TestParser(t *testing.T) {
	type testCase struct {
		cmd  string
		want string
	}

	testCases := []testCase{
		{
			cmd:  "M=D",
			want: "1110001100001000",
		},
		{
			cmd:  "D=D+A",
			want: "1110000010010000",
		},
		{
			cmd:  "D=A",
			want: "1110110000010000",
		},
		{
			cmd:  "@0",
			want: "0000000000000000",
		},
		{
			cmd:  "@1",
			want: "0000000000000001",
		},
		{
			cmd:  "@2",
			want: "0000000000000010",
		},
		{
			cmd:  "@3",
			want: "0000000000000011",
		},
		{
			cmd:  "@4",
			want: "0000000000000100",
		},
	}

	for _, tc := range testCases {
		got := main.ParseCmd(tc.cmd)
		assert.Equal(t, tc.want, got)
	}
}
