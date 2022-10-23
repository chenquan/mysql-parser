package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitGroupByItem(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a")
		result := mySqlParser.GroupByItem().Accept(visitor)
		assert.EqualValues(t, GroupByItem{Expression: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}}}}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a.b desc")
		result := mySqlParser.GroupByItem().Accept(visitor)
		assert.EqualValues(t,
			GroupByItem{
				Expression: ExpressionAtomPredicate{
					ExpressionAtom: FullColumnNameExpressionAtom{
						FullColumnName: FullColumnName{Uid: "a", DottedIds: []DottedId{
							{Uid: "b"},
						}},
					},
				},
				Order: "desc",
			},
			result)
	})
}
