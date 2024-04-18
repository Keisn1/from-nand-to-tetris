package main_test

import (
	"os"
	"testing"

	"hackAsm"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	type testCase struct {
		targetFP string
		want     string
	}
	testCases := []testCase{
		{
			targetFP: "add/Add.asm",
			want:     testTables["add"],
		},
		{
			targetFP: "max/MaxL.asm",
			want:     testTables["maxL"],
		},
	}

	for _, tc := range testCases {

		parser := main.NewParser(tc.targetFP)
		got := parser.Parse()
		assert.Equal(t, tc.want, got)
	}

}

func TestParser(t *testing.T) {
	t.Run("Test cmds dependent on file", func(t *testing.T) {
		parser := main.NewParser("max/Max.asm")
		type testCase struct {
			cmd  string
			want string
		}

		testCases := []testCase{
			{
				cmd:  "@INFINITE_LOOP",
				want: "0000000000001110",
			},
			{
				cmd:  "@OUTPUT_D",
				want: "0000000000001100",
			},
			{
				cmd:  "@OUTPUT_FIRST",
				want: "0000000000001010",
			},
		}

		for _, tc := range testCases {
			got := parser.ParseCmd(tc.cmd)
			assert.Equal(t, tc.want, got)
		}
	})
	t.Run("Test cmds that are independent of file", func(t *testing.T) {

		parser := main.Parser{}
		type testCase struct {
			cmd  string
			want string
		}

		testCases := []testCase{
			{
				cmd:  "@R1",
				want: "0000000000000001",
			},
			{
				cmd:  "0;JMP",
				want: "1110101010000111",
			},
			{
				cmd:  "D;JGT",
				want: "1110001100000001",
			},
			{
				cmd:  "@10",
				want: "0000000000001010",
			},
			{
				cmd:  "D=D-M",
				want: "1111010011010000",
			},
			{
				cmd:  "D=M",
				want: "1111110000010000",
			},
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
			got := parser.ParseCmd(tc.cmd)
			assert.Equal(t, tc.want, got)
		}
	})
}

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
