package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
)

func Parser(sql string) *Result {
	if sql == "" {
		return nil
	}

	inputStream := antlr.NewInputStream(sql)
	lexer := parser.NewMySqlLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	mysqlParser := parser.NewMySqlParser(tokens)
	visitor := &parseTreeVisitor{
		Result: &Result{},
	}

	mysqlParser.Root().Accept(visitor)
	return visitor.Result
}
