package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitAggregateWindowedFunction(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("SUM(a)")
		result := mySqlParser.AggregateWindowedFunction().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function: "SUM",
			StarArg:  false,
			FunctionArg: &FunctionArg{F: FullColumnName{
				Uid:       "a",
				DottedIds: nil,
			}},
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("SUM(a+2)")
		result := mySqlParser.AggregateWindowedFunction().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function: "SUM",
			StarArg:  false,
			FunctionArg: &FunctionArg{F: ExpressionAtomPredicate{ExpressionAtom: MathExpressionAtom{
				LeftExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
					Uid: "a",
				}},
				MathOperator:        "+",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: "2"}},
			}}},
		}, result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("SUM(2 MOD 2)")
		result := mySqlParser.AggregateWindowedFunction().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function: "SUM",
			StarArg:  false,
			FunctionArg: &FunctionArg{F: ExpressionAtomPredicate{ExpressionAtom: MathExpressionAtom{
				LeftExpressionAtom:  ConstantExpressionAtom{Constant: ConstantDecimal{Val: "2"}},
				MathOperator:        "MOD",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: "2"}},
			}}},
		}, result)
	})

	t.Run("4", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("count(*)")
		result := mySqlParser.AggregateWindowedFunction().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function: "count",
			StarArg:  true,
		}, result)
	})
}
