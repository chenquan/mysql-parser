package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitSelectIntoVariables(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		parser, visitor := createMySqlParser(`INTO A,B,c`)
		result := parser.SelectIntoExpression().Accept(visitor)
		assert.EqualValues(t, SelectIntoVariables{AssignmentFields: []AssignmentField{
			{
				Val: "A",
			},
			{
				Val: "B",
			},
			{
				Val: "c",
			},
		}}, result)

	})
}
