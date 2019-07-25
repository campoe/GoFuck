# GoFuck
A brainfuck interpreter written in Go

## Example
### _hello.bf_
```
++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.
```

### _main.go_
```
package main

import (
	"./pkg/program"
)

func main() {
	prog := program.FromFile("assets/hello.bf")
	prog.Execute()
}
```

### Output
`Hello World!`
