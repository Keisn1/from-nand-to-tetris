package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(os.Args)
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
	nbr, _ := strconv.Atoi(strings.Split(cmd, "@")[1])

	binaryStr := strconv.FormatInt(int64(nbr), 2)
	missingZeros := 16 - len(binaryStr)
	for range missingZeros {
		binaryStr = "0" + binaryStr
	}
	return binaryStr
}

var (
	destTable = map[string]string{
		"":    "000",
		"M":   "001",
		"D":   "010",
		"MD":  "011",
		"A":   "100",
		"AM":  "101",
		"AD":  "110",
		"AMD": "111",
	}
	compTableWithOutA = map[string]string{
		"0":   "101010",
		"1":   "111111",
		"-1":  "111010",
		"D":   "001100",
		"A":   "110000",
		"!D":  "001101",
		"!A":  "110001",
		"-D":  "001111",
		"-A":  "110011",
		"D+1": "011111",
		"A+1": "110111",
		"D-1": "001110",
		"A-1": "110010",
		"D+A": "000010",
		"D-A": "010011",
		"A-D": "000111",
		"D&A": "000000",
		"D|A": "010101",
	}
	compTableWithA = map[string]string{
		"M":   "110000",
		"!M":  "110001",
		"-M":  "110011",
		"M+1": "110111",
		"M-1": "110010",
		"D+M": "000010",
		"D-M": "010011",
		"M-D": "000111",
		"D&M": "000000",
		"D|M": "010101",
	}
)

func (p Parser) jump() string {
	return "000"
}

func (p Parser) dest(cmd string) string {
	parts := strings.Split(cmd, "=")
	return destTable[parts[0]]
}

func (p Parser) comp(cmd string) string {
	parts := strings.Split(cmd, "=")
	return compTableWithOutA[parts[1]]
}

func (p Parser) parseCInstruction(cmd string) string {
	return "111" + "0" + p.comp(cmd) + p.dest(cmd) + p.jump()
}

func (p Parser) ParseCmd(cmd string) string {
	switch cmd[0] {
	case '@':
		return parseAInstruction(cmd)
	default:
		return p.parseCInstruction(cmd)
	}
}
