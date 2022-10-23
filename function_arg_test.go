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
				{
					F: AggregateWindowedFunction{
						Function:    "sum",
						FunctionArg: &FunctionArg{F: FullColumnName{Uid: "a"}},
					},
				},
				{
					F: AggregateWindowedFunction{
						Function:    "avg",
						FunctionArg: &FunctionArg{F: FullColumnName{Uid: "b"}},
					},
				},
			}, result)
	})

}

func Test_parseTreeVisitor_VisitFunctionArg(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("sum(a)")
		result := mySqlParser.FunctionArg().Accept(visitor)

		assert.EqualValues(t,
			FunctionArg{
				F: AggregateWindowedFunction{
					Function:    "sum",
					FunctionArg: &FunctionArg{F: FullColumnName{Uid: "a"}},
				},
			}, result)
	})
}
