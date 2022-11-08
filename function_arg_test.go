package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitFunctionArgs(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("sum(a), avg(b)")
		result := mySqlParser.FunctionArgs().Accept(visitor)

		assert.EqualValues(t,
			[]FunctionArg{
				AggregateWindowedFunction{
					Function:     "sum",
					FunctionArgs: []FunctionArg{FullColumnName{Uid: "a"}},
				},
				AggregateWindowedFunction{
					Function:     "avg",
					FunctionArgs: []FunctionArg{FullColumnName{Uid: "b"}},
				},
			}, result)
	})

}

func Test_parseTreeVisitor_VisitFunctionArg(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("sum(a)")
		result := mySqlParser.FunctionArg().Accept(visitor)

		assert.EqualValues(t,
			AggregateWindowedFunction{
				Function:     "sum",
				FunctionArgs: []FunctionArg{FullColumnName{Uid: "a"}},
			}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1")
		result := mySqlParser.FunctionArg().Accept(visitor)

		assert.EqualValues(t,
			ConstantDecimal{Val: 1}, result)
	})
}
