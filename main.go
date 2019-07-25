package main

import (
	"./pkg/program"
)

func main() {
	prog := program.FromFile("assets/hello.bf")
	prog.Execute()
}
