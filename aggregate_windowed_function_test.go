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
			FunctionArgs: []FunctionArg{FunctionArg{F: FullColumnName{
				Uid:       "a",
				DottedIds: nil,
			}}},
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("SUM(a+2)")
		result := mySqlParser.AggregateWindowedFunction().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function: "SUM",
			StarArg:  false,
			FunctionArgs: []FunctionArg{FunctionArg{F: ExpressionAtomPredicate{ExpressionAtom: MathExpressionAtom{
				LeftExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
					Uid: "a",
				}},
				MathOperator:        "+",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}},
			}}}},
		}, result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("SUM(ALL 2 MOD 2)")
		result := mySqlParser.AggregateWindowedFunction().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function:   "SUM",
			StarArg:    false,
			Aggregator: "ALL",
			FunctionArgs: []FunctionArg{FunctionArg{F: ExpressionAtomPredicate{ExpressionAtom: MathExpressionAtom{
				LeftExpressionAtom:  ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}},
				MathOperator:        "MOD",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}},
			}}}},
		}, result)
	})

	t.Run("5", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("count(DISTINCT id)")
		result := mySqlParser.AggregateWindowedFunction().Accept(visitor)
		assert.EqualValues(t, AggregateWindowedFunction{
			Function:   "count",
			Aggregator: "DISTINCT",
			FunctionArgs: []FunctionArg{FunctionArg{
				F: FullColumnName{
					Uid: "id",
				},
			}},
		}, result)
	})
}
