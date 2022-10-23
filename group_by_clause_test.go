package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitGroupByClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("group by a.b desc, c asc")
		result := mySqlParser.GroupByClause().Accept(visitor)
		assert.EqualValues(t,
			GroupByClause{GroupByItems: []GroupByItem{
				{
					Expression: ExpressionAtomPredicate{
						ExpressionAtom: FullColumnNameExpressionAtom{
							FullColumnName: FullColumnName{Uid: "a", DottedIds: []DottedId{
								{Uid: "b"},
							}},
						},
					},
					Order: "desc",
				},
				{
					Expression: ExpressionAtomPredicate{
						ExpressionAtom: FullColumnNameExpressionAtom{
							FullColumnName: FullColumnName{Uid: "c"},
						},
					},
					Order: "asc",
				},
			}},
			result)
	})
}
