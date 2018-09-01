package lexer

import (
	"testing"

	"github.com/pogorammer/slang/token"
)

func TestNextToken(t *testing.T) {
	input := `
	fuck five = 5;
	fuck ten = 10;
	fuck add = fucks(x, y) {
		x + y;
	};
	fuck result = add(five, ten);
	!-/*0;
	2 < 10 > 7;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;

	"foobar";
	"foo bar";

	[1, 2];

	{"foo": "bar"};

	// comment
	fuck a = 1; // inline comment

	fuck b = 123.45;
	fuck c = 0.678;
	fuck d = 9.0;
	`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.FUCK, "fuck"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.FUCK, "fuck"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUCK, "fuck"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fucks"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.FUCK, "fuck"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTARISK, "*"},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.INT, "2"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "7"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foo bar"},
		{token.SEMICOLON, ";"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.FUCK, "fuck"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.FUCK, "fuck"},
		{token.IDENT, "b"},
		{token.ASSIGN, "="},
		{token.FLOAT, "123.45"},
		{token.SEMICOLON, ";"},
		{token.FUCK, "fuck"},
		{token.IDENT, "c"},
		{token.ASSIGN, "="},
		{token.FLOAT, "0.678"},
		{token.SEMICOLON, ";"},
		{token.FUCK, "fuck"},
		{token.IDENT, "d"},
		{token.ASSIGN, "="},
		{token.FLOAT, "9.0"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Logf("tests[%d] - tok: %#v", i, tok)
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Logf("tests[%d] - tok: %#v", i, tok)
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
