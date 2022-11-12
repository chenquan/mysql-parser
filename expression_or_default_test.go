package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitExpressionOrDefault(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1!=1 AND a!=2")
		result := mySqlParser.ExpressionOrDefault().Accept(visitor)
		assert.EqualValues(t, ExpressionOrDefault{Expression: LogicalExpression{
			LeftExpression: BinaryComparisonPredicate{
				LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				ComparisonOperator: "!=",
				RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
			},
			LogicalOperator: "AND",
			RightExpression: BinaryComparisonPredicate{
				LeftPredicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
					Uid: "a",
				}}},
				ComparisonOperator: "!=",
				RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
			},
		}}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("DEFAULT")
		result := mySqlParser.ExpressionOrDefault().Accept(visitor)
		assert.EqualValues(t, ExpressionOrDefault{Default: true}, result)
	})
}
