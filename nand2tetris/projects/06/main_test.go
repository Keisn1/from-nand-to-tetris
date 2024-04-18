package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	cmd := "@2"
	want := "0000000000000010"
	got := parseCmd(cmd)
	assert.Equal(t, want, got)
}
