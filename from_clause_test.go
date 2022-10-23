package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitFromClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("FROM a as b")
		result := mySqlParser.FromClause().Accept(visitor)

		assert.EqualValues(t,
			FromClause{TableSources: &TableSources{TableSources: []TableSource{
				TableSourceBase{
					TableSourceItem: AtomTableItem{
						TableName: "a",
						Alias:     "b",
					},
				},
			}}}, result)

	})
}
