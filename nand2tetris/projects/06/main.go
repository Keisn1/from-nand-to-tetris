package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	target := os.Args[1]
	Assemble(target)
}

func Assemble(target string) {
	parser := NewParser(target)
	code := parser.Parse()
	FileSave(code, "add/Add.hack")
}

func FileSave(text, targetFP string) {
	os.WriteFile("add/Add.hack", []byte(text), 0644)
}

type Parser struct {
	cmds []string
}

func NewParser(fp string) *Parser {
	content, _ := os.ReadFile(fp)
	lines := strings.Split(string(content), "\n")
	var cmds []string

	for _, l := range lines {
		cmd := strings.TrimSpace(l)
		if isNotInstruction(cmd) {
			continue
		}
		cmds = append(cmds, cmd)
	}

	return &Parser{cmds: cmds}
}

func (p Parser) Parse() string {
	var hackCmds []string
	for _, cmd := range p.cmds {
		hackCmds = append(hackCmds, p.ParseCmd(cmd))
	}
	return strings.Join(hackCmds, "\n")
}

func isNotInstruction(cmd string) bool {
	if len(cmd) == 0 {
		return true
	}
	if cmd[0] == '/' {
		return true
	}
	return false
}

func parseAInstruction(cmd string) string {
	val := strings.Split(cmd, "@")[1]

	address, ok := symbolTable[val]
	if !ok {
		address, _ = strconv.Atoi(val)
	}

	binaryStr := strconv.FormatInt(int64(address), 2)
	missingZeros := 16 - len(binaryStr)
	for range missingZeros {
		binaryStr = "0" + binaryStr
	}
	return binaryStr
}

func (p Parser) jump(cmd string) string {
	if strings.Contains(cmd, ";") {
		jumpPart := strings.Split(cmd, ";")[1]
		return jumpTable[jumpPart]
	}
	return "000"
}

func (p Parser) dest(cmd string) string {
	if strings.Contains(cmd, "=") {
		destPart := strings.Split(cmd, "=")[0]
		return destTable[destPart]
	}
	return "000"
}

func (p Parser) comp(cmd string) string {
	if strings.Contains(cmd, "=") {
		compString := strings.Split(cmd, "=")[1]
		if strings.Contains(compString, "M") {
			a := "1"
			return a + compTableWithM[compString]
		} else {
			a := "0"
			return a + compTableWithOutM[compString]
		}
	}
	compString := strings.Split(cmd, ";")[0]
	if strings.Contains(compString, "M") {
		a := "1"
		return a + compTableWithM[compString]
	} else {
		a := "0"
		return a + compTableWithOutM[compString]
	}

}

func (p Parser) parseCInstruction(cmd string) string {
	return "111" + p.comp(cmd) + p.dest(cmd) + p.jump(cmd)
}

func (p Parser) ParseCmd(cmd string) string {
	switch cmd[0] {
	case '@':
		return parseAInstruction(cmd)
	default:
		return p.parseCInstruction(cmd)
	}
}
