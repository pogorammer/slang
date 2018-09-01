package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/pogorammer/slang/eval"
	"github.com/pogorammer/slang/lexer"
	"github.com/pogorammer/slang/object"
	"github.com/pogorammer/slang/parser"
	"github.com/pogorammer/slang/repl"
)

func main() {
	// Start Slang REPL
	if len(os.Args) == 1 {
		fmt.Println("This is the Slang programming language!")
		fmt.Println("Feel free to type in commands")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// Run a Slang script
	if err := runProgram(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runProgram(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", filename, err)
	}

	p := parser.New(lexer.New(string(data)))
	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		return errors.New(p.Errors()[0])
	}

	env := object.NewEnvironment()
	result := eval.Eval(program, env)
	if _, ok := result.(*object.Nil); ok {
		return nil
	}

	_, err = io.WriteString(os.Stdout, result.Inspect()+"\n")
	return err
}
