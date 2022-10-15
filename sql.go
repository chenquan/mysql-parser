package parser

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
)

type Sql interface {
	Err() error
}

func Parser(sql string) *Result {
	if sql == "" {
		return nil
	}

	sql = strings.ToUpper(sql)
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

func parserSingleSql(singleSql string) Sql {

	//visitor.CreateTableNames
	return nil
}
