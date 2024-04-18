package main

import (
	"strconv"
	"strings"
)

func ParseCmd(cmd string) string {
	if cmd[0] == '@' {
		nbr, _ := strconv.Atoi(strings.Split(cmd, "@")[1])

		binaryStr := strconv.FormatInt(int64(nbr), 2)
		missingZeros := 16 - len(binaryStr)
		for range missingZeros {
			binaryStr = "0" + binaryStr
		}
		return binaryStr
	}

	var opCode, fill, a, comp, dest, jump string
	opCode = "1"
	fill = "11"
	a = "0"
	jump = "000"

	parts := strings.Split(cmd, "=")
	switch parts[0] {
	case "D":
		dest = "010"
	case "M":
		dest = "001"
	}

	switch parts[1] {
	case "A":
		comp = "110000"
	case "D+A":
		comp = "000010"
	case "D":
		comp = "001100"
	}

	return opCode + fill + a + comp + dest + jump
}
