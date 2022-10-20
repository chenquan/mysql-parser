package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitConstantExpressionAtom(t *testing.T) {
	// FIXME
	sql := "-132.3"
	sql = strings.ToUpper(sql)
	inputStream := antlr.NewInputStream(sql)
	lexer := parser.NewMySqlLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	mySqlParser := parser.NewMySqlParser(tokens)
	visitor := &parseTreeVisitor{
		Result: &Result{},
	}
	fmt.Println(mySqlParser.ExpressionAtom().GetText())
	result := mySqlParser.ExpressionAtom().Accept(visitor)
	_ = result
}

func Test_parseTreeVisitor_VisitFullColumnNameExpressionAtom(t *testing.T) {
	sql := "a.c.name"
	sql = strings.ToUpper(sql)
	inputStream := antlr.NewInputStream(sql)
	lexer := parser.NewMySqlLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	mySqlParser := parser.NewMySqlParser(tokens)
	visitor := &parseTreeVisitor{
		Result: &Result{},
	}

	result := mySqlParser.ExpressionAtom().Accept(visitor)
	assert.EqualValues(t, result, FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
		Uid: "A",
		DottedIds: []DottedId{
			{Uid: "C"},
			{Uid: "NAME"},
		},
	}})
}

func Test_parseTreeVisitor_VisitFunctionCallExpressionAtom(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("CURRENT_DATE()")
	result := mySqlParser.ExpressionAtom().Accept(visitor)
	assert.EqualValues(t, result, FunctionCallExpressionAtom{FunctionCall: SpecificFunctionCall{SpecificFunction: SimpleFunctionCall{function: "CURRENT_DATE"}}})
}
