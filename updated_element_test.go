package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitUpdatedElement(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a=1")
		result := mySqlParser.UpdatedElement().Accept(visitor)
		assert.EqualValues(t,
			UpdatedElement{
				FullColumnName: FullColumnName{
					Uid: "a",
				},
				Value: ExpressionOrDefault{
					Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				},
			}, result)
	})
}
