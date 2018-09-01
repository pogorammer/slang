package ast

import (
	"testing"

	"github.com/pogorammer/slang/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&FuckStatement{
				Token: token.Token{Type: token.FUCK, Literal: "fuck"},
				Name: &Ident{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Ident{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "fuck myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%T", program.String())
	}
}
