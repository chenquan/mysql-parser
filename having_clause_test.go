package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitHavingClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("HAVING AVG(A) > 2")
		result := mySqlParser.HavingClause().Accept(visitor)
		assert.EqualValues(t,
			HavingClause{
				Expression: BinaryComparisonPredicate{
					LeftPredicate: ExpressionAtomPredicate{
						ExpressionAtom: FunctionCallExpressionAtom{FunctionCall: AggregateWindowedFunction{
							Function:   "AVG",
							Aggregator: "",
							FunctionArgs: []FunctionArg{{F: FullColumnName{
								Uid: "A",
							}}},
						}},
					},
					ComparisonOperator: ">",
					RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
				},
			},
			result)
	})
}
