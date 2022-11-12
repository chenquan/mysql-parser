package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitExpressionsWithDefaults(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1!=1 AND a!=2")
		result := mySqlParser.ExpressionsWithDefaults().Accept(visitor)
		assert.EqualValues(t, ExpressionsWithDefaults{ExpressionOrDefaults: []ExpressionOrDefault{
			{Expression: LogicalExpression{
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
			}},
		}}, result)
	})

}
