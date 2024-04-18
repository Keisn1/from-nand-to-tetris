package main_test

import (
	"os"
	"testing"

	"hackAsm"

	"github.com/stretchr/testify/assert"
)

func TestAssemble(t *testing.T) {
	target := "add/Add.asm"
	want := `0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1110001100001000`
	main.Assemble(target)

	out := "add/Add.hack"
	assert.FileExists(t, out)

	got, err := os.ReadFile(out)
	assert.NoError(t, err)
	assert.Equal(t, want, string(got))

	os.Remove(out)
	assert.NoFileExists(t, out)
}

func TestSaveResult(t *testing.T) {
	text := "text"
	targetFP := "add/Add.asm"
	want := "add/Add.hack"
	main.FileSave(text, targetFP)

	assert.FileExists(t, want)
	got, err := os.ReadFile(want)
	assert.NoError(t, err)
	assert.Equal(t, text, string(got))

	os.Remove(want)
	assert.NoFileExists(t, want)
}

func TestParse(t *testing.T) {
	targetFP := "add/Add.asm"
	parser := main.NewParser(targetFP)
	want := `0000000000000010
1110110000010000
0000000000000011
1110000010010000
0000000000000000
1110001100001000`
	got := parser.Parse()
	assert.Equal(t, want, got)
}

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
		parser := main.Parser{}
		got := parser.ParseCmd(tc.cmd)
		assert.Equal(t, tc.want, got)
	}
}
