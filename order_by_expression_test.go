package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitOrderByExpression(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("A desc")
		result := mySqlParser.OrderByExpression().Accept(visitor)
		assert.EqualValues(t, OrderByExpression{
			Expression: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{
				FullColumnName: FullColumnName{
					Uid:       "A",
					DottedIds: nil,
				},
			}},
			Order: "desc",
		}, result)
	})
}
