package main

import (
	"os"
	"strconv"
	"strings"
)

const varOffset = 16

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
	cmds       []string
	varCounter int
}

func removeWhiteSpace(line string) string {
	line = strings.TrimSpace(line)
	splits := strings.Split(line, " ")
	var elems []string
	for _, s := range splits {
		elems = append(elems, strings.TrimSpace(s))
	}
	return strings.Join(elems, "")
}

func removeComments(line string) string {
	line = strings.Split(line, "//")[0]
	return line
}

func NewParser(fp string) *Parser {
	content, _ := os.ReadFile(fp)
	var lines []string
	for _, line := range strings.Split(string(content), "\n") {
		line = removeWhiteSpace(line)
		line = removeComments(line)
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	var cmds []string
	cmdCounter := 0
	for _, l := range lines {
		cmd := strings.TrimSpace(l)
		if isNotInstruction(cmd) {
			continue
		}
		if isLabelSymbol(cmd) {
			addToSymbolTable(cmd, cmdCounter)
			continue
		}

		cmds = append(cmds, cmd)
		cmdCounter++
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

func (p *Parser) parseAInstruction(cmd string) string {
	val := strings.Split(cmd, "@")[1]

	address, ok := symbolTable[val]
	var err error
	if !ok {
		address, err = strconv.Atoi(val)
		if err != nil {
			symbolTable[val] = varOffset + p.varCounter
			p.varCounter++
			address = symbolTable[val]
		}
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
		return p.parseAInstruction(cmd)
	default:
		return p.parseCInstruction(cmd)
	}
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

func addToSymbolTable(cmd string, cmdCounter int) {
	symbol := cmd[1 : len(cmd)-1]
	symbolTable[symbol] = cmdCounter
}

func isLabelSymbol(cmd string) bool {
	return cmd[0] == '('
}
