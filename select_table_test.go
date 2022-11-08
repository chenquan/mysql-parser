package parser

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
	"github.com/stretchr/testify/assert"
)

func createMySqlParser(sql string) (*parser.MySqlParser, *parseTreeVisitor) {
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
	t.Run("all", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("*")
		result := mySqlParser.SelectElements().Accept(visitor)
		assert.EqualValues(t, SelectElements{
			All:            true,
			SelectElements: nil,
		}, result)
	})

	t.Run("SelectColumnElement", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a,b as b_t")
		result := mySqlParser.SelectElements().Accept(visitor)
		assert.EqualValues(t, SelectElements{
			SelectElements: []SelectElement{
				SelectColumnElement{
					FullColumnName: FullColumnName{
						Uid:       "a",
						DottedIds: nil,
					},
					Alias: "",
				},
				SelectColumnElement{
					FullColumnName: FullColumnName{
						Uid:       "b",
						DottedIds: nil,
					},
					Alias: "b_t",
				},
			},
		}, result)
	})

	t.Run("SelectStarElement", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a.*,b.*")
		result := mySqlParser.SelectElements().Accept(visitor)
		assert.EqualValues(t, SelectElements{
			SelectElements: []SelectElement{
				SelectStarElement{
					TableName: FullId{
						Uid: "a",
					},
				},
				SelectStarElement{
					TableName: FullId{
						Uid: "b",
					},
				},
			},
		}, result)
	})
	t.Run("SelectStarElement&SelectColumnElement", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a.*,b.*,c as c_t")
		result := mySqlParser.SelectElements().Accept(visitor)
		assert.EqualValues(t, SelectElements{
			SelectElements: []SelectElement{
				SelectStarElement{
					TableName: FullId{
						Uid: "a",
					},
				},
				SelectStarElement{
					TableName: FullId{
						Uid: "b",
					},
				},
				SelectColumnElement{
					FullColumnName: FullColumnName{
						Uid:       "c",
						DottedIds: nil,
					},
					Alias: "c_t",
				},
			},
		}, result)
	})

	t.Run("SelectFunctionElement", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("SUM(a) as c, count(*) as cnt")
		result := mySqlParser.SelectElements().Accept(visitor)
		assert.EqualValues(t, SelectElements{
			SelectElements: []SelectElement{
				SelectFunctionElement{
					FunctionCall: AggregateWindowedFunction{
						Function:     "SUM",
						StarArg:      false,
						Aggregator:   "",
						FunctionArgs: []FunctionArg{FullColumnName{Uid: "a"}},
					},
					Alias: "c",
				},
				SelectFunctionElement{
					FunctionCall: AggregateWindowedFunction{
						Function: "count",
						StarArg:  true,
					},
					Alias: "cnt",
				},
			},
		}, result)

	})
}

func Test_parseTreeVisitor_VisitFullId(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a.*")
	result := mySqlParser.FullId().Accept(visitor)
	assert.EqualValues(t, FullId{
		Uid:   "a",
		DotId: "",
	}, result)

	mySqlParser, visitor = createMySqlParser("db_name.a.*")
	result = mySqlParser.FullId().Accept(visitor)
	assert.EqualValues(t, FullId{
		Uid:   "db_name",
		DotId: "a",
	}, result)
}

func Test_parseTreeVisitor_VisitSelectStarElement(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a.*")
	result := mySqlParser.SelectElement().Accept(visitor)
	assert.EqualValues(t, SelectStarElement{TableName: FullId{
		Uid:   "a",
		DotId: "",
	}}, result)

	mySqlParser, visitor = createMySqlParser("db_name.a.*")
	result = mySqlParser.SelectElement().Accept(visitor)
	assert.EqualValues(t, SelectStarElement{TableName: FullId{
		Uid:   "db_name",
		DotId: "a",
	}}, result)
}

func Test_parseTreeVisitor_VisitDottedId(t *testing.T) {
	mySqlParser, visitor := createMySqlParser(".name")
	result := mySqlParser.DottedId().Accept(visitor)
	assert.EqualValues(t, DottedId{Uid: "name"}, result)
}

func Test_parseTreeVisitor_VisitFullColumnName(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a.c.name")
	result := mySqlParser.FullColumnName().Accept(visitor)
	assert.EqualValues(t, FullColumnName{
		Uid:       "a",
		DottedIds: []DottedId{{Uid: "c"}, {Uid: "name"}},
	}, result)

	mySqlParser, visitor = createMySqlParser("a.c")
	result = mySqlParser.FullColumnName().Accept(visitor)
	assert.EqualValues(t, FullColumnName{
		Uid:       "a",
		DottedIds: []DottedId{{Uid: "c"}},
	}, result)

}

func Test_parseTreeVisitor_VisitSelectColumnElement(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a as b")
	result := mySqlParser.SelectElement().Accept(visitor)
	assert.EqualValues(t, SelectColumnElement{Alias: "b",
		FullColumnName: FullColumnName{
			Uid: "a",
		}}, result)
}
