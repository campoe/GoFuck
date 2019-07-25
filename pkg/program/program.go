package program

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Program struct {
	instructionPointer int
	Commands           string
	loopStack          []int
	dataPointer        int
	data               [30000]byte
}

func FromFile(filename string) Program {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return Program{
		Commands: string(buf),
	}
}

func (p *Program) next() *Program {
	p.instructionPointer += 1
	return p
}

func (p *Program) increase() *Program {
	p.data[p.dataPointer] += 1
	return p.next()
}

func (p *Program) decrease() *Program {
	p.data[p.dataPointer] -= 1
	return p.next()
}

func (p *Program) right() *Program {
	p.dataPointer = (p.dataPointer + 1) % len(p.data)
	return p.next()
}

func (p *Program) left() *Program {
	p.dataPointer = (p.dataPointer - 1) % len(p.data)
	return p.next()
}

func (p *Program) out() *Program {
	fmt.Printf("%c", p.data[p.dataPointer])
	return p.next()
}

func (p *Program) in() *Program {
	// TODO: take input
	return p.next()
}

func (p *Program) jmpForward() *Program {
	ptr := p.instructionPointer + 1
	if p.data[p.dataPointer] == 0 {
		ptr = loopClosingIndex(ptr, 1, p.Commands)
	}
	p.loopStack = append(p.loopStack, p.instructionPointer)
	p.instructionPointer = ptr
	return p
}

func (p *Program) jmpBack() *Program {
	ptr := p.instructionPointer + 1
	if len(p.loopStack) > 0 {
		if p.data[p.dataPointer] != 0 {
			ptr = p.loopStack[len(p.loopStack)-1]
		}
		p.loopStack = p.loopStack[:len(p.loopStack)-1]
	}
	p.instructionPointer = ptr
	return p
}

func loopClosingIndex(pointer int, brackets int, commands string) int {
	if brackets == 0 {
		return pointer - 1
	} else {
		switch commands[pointer] {
		case '[':
			return loopClosingIndex(pointer+1, brackets+1, commands)
		case ']':
			return loopClosingIndex(pointer+1, brackets-1, commands)
		default:
			return loopClosingIndex(pointer+1, brackets, commands)
		}
	}
}

func (p *Program) executeCommand() *Program {
	if p.instructionPointer < 0 || p.instructionPointer >= len(p.Commands) {
		return p
	}
	switch p.Commands[p.instructionPointer] {
	case '>':
		p.right()
	case '<':
		p.left()
	case '+':
		p.increase()
	case '-':
		p.decrease()
	case '.':
		p.out()
	case ',':
		p.in()
	case '[':
		p.jmpForward()
	case ']':
		p.jmpBack()
	default:
		p.next()
	}
	return p
}

func (p *Program) reset() *Program {
	p.instructionPointer = 0
	p.dataPointer = 0
	p.loopStack = []int{}
	for i := 0; i < len(p.data); i++ {
		p.data[i] = 0
	}
	return p
}

func (p Program) Execute() Program {
	p.reset()
	for p.instructionPointer < len(p.Commands) {
		p.executeCommand()
	}
	return *p.reset()
}
