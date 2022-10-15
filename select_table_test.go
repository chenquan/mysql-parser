package parser

import (
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
	"github.com/stretchr/testify/assert"
)

func createMySqlParser(sql string) (*parser.MySqlParser, *parseTreeVisitor) {
	sql = strings.ToUpper(sql)
	inputStream := antlr.NewInputStream(sql)
	lexer := parser.NewMySqlLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	mysqlParser := parser.NewMySqlParser(tokens)
	visitor := &parseTreeVisitor{
		Result: &Result{},
	}

	return mysqlParser, visitor
}

func TestParseTreeVisitor_VisitSelectElements(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("*")
	result := mySqlParser.SelectElements().Accept(visitor)
	assert.EqualValues(t, SelectElements{
		All:            true,
		SelectElements: nil,
	}, result)
}

func Test_parseTreeVisitor_VisitFullId(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a.*")
	result := mySqlParser.FullId().Accept(visitor)
	assert.EqualValues(t, FullId{
		Uid:   "A",
		DotId: "",
	}, result)

	mySqlParser, visitor = createMySqlParser("db_name.a.*")
	result = mySqlParser.FullId().Accept(visitor)
	assert.EqualValues(t, FullId{
		Uid:   "DB_NAME",
		DotId: "A",
	}, result)
}

func Test_parseTreeVisitor_VisitSelectStarElement(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a.*")
	result := mySqlParser.SelectElement().Accept(visitor)
	assert.EqualValues(t, SelectStarElement{FullId: FullId{
		Uid:   "A",
		DotId: "",
	}}, result)

	mySqlParser, visitor = createMySqlParser("db_name.a.*")
	result = mySqlParser.SelectElement().Accept(visitor)
	assert.EqualValues(t, SelectStarElement{FullId: FullId{
		Uid:   "DB_NAME",
		DotId: "A",
	}}, result)
}

func Test_parseTreeVisitor_VisitDottedId(t *testing.T) {
	mySqlParser, visitor := createMySqlParser(".name")
	result := mySqlParser.DottedId().Accept(visitor)
	assert.EqualValues(t, DottedId{Uid: "NAME"}, result)
}

func Test_parseTreeVisitor_VisitFullColumnName(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a.c.name")
	result := mySqlParser.FullColumnName().Accept(visitor)
	assert.EqualValues(t, FullColumnName{
		Uid:       "A",
		DottedIds: []DottedId{{Uid: "C"}, {Uid: "NAME"}},
	}, result)

	mySqlParser, visitor = createMySqlParser("a.c")
	result = mySqlParser.FullColumnName().Accept(visitor)
	assert.EqualValues(t, FullColumnName{
		Uid:       "A",
		DottedIds: []DottedId{{Uid: "C"}},
	}, result)

}

func Test_parseTreeVisitor_VisitSelectColumnElement(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a as b")
	result := mySqlParser.SelectElement().Accept(visitor)
	assert.EqualValues(t, SelectColumnElement{Alias: "B",
		FullColumnName: FullColumnName{
			Uid:       "A",
			DottedIds: []DottedId{},
		}}, result)
}

func Test_parseTreeVisitor_VisitSimpleFunctionCall(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("CURRENT_DATE")
	result := mySqlParser.SpecificFunction().Accept(visitor)
	assert.EqualValues(t, SimpleFunctionCall{function: "CURRENT_DATE"}, result)

	mySqlParser, visitor = createMySqlParser("CURRENT_DATE()")
	result = mySqlParser.SpecificFunction().Accept(visitor)
	assert.EqualValues(t, SimpleFunctionCall{function: "CURRENT_DATE"}, result)
}
