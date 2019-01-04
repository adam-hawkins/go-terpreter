package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/literallystan/gopiler/evaluator"
	"github.com/literallystan/gopiler/lexer"
	"github.com/literallystan/gopiler/object"
	"github.com/literallystan/gopiler/parser"
	"github.com/literallystan/gopiler/repl"
)

func main() {
	out := os.Stdout
	if len(os.Args) >= 3 {
		fmt.Printf("Wanted 1 arg, got {%d}\n", len(os.Args)-1)
		return
	} else if len(os.Args) == 2 {
		env := object.NewEnvironment()
		p, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}

		l := lexer.New(string(p))
		parse := parser.New(l)

		program := parse.ParseProgram()
		if len(parse.Errors()) != 0 {
			printParserErrors(out, parse.Errors())
			return
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	} else {
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
