package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitQueryExpression(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("((SELECT name FROM U order by a desc))")
		result := mySqlParser.QueryExpression().Accept(visitor)
		assert.EqualValues(t, QuerySpecification{
			SelectSpecs: nil,
			SelectElements: SelectElements{
				SelectElements: []SelectElement{
					SelectColumnElement{FullColumnName: FullColumnName{Uid: "name"}},
				},
			},
			FromClause: &FromClause{
				TableSources: []TableSource{
					TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
				},
			},
			OrderByClause: &OrderByClause{
				OrderByExpressions: []OrderByExpression{
					{
						Expression: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{
							FullColumnName: FullColumnName{
								Uid:       "a",
								DottedIds: nil,
							},
						}},
						Order: "desc",
					},
				},
			},
		}, result)
	})

}
