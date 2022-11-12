package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
)

func Parse(sql string) interface{} {
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

	accept := mysqlParser.Root().Accept(visitor)
	return accept
}
