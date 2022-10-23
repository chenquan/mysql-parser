package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitOrderByClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("order by A desc")
		result := mySqlParser.OrderByClause().Accept(visitor)
		assert.EqualValues(t, OrderByClause{OrderByExpressions: []OrderByExpression{
			{
				Expression: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{
					FullColumnName: FullColumnName{
						Uid:       "A",
						DottedIds: nil,
					},
				}},
				Order: "desc",
			},
		}}, result)
	})

}
