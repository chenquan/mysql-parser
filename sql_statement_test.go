package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitSqlStatements(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT * FROM U;`)
		result := parser.SqlStatements().Accept(visitor)
		assert.EqualValues(t, []SqlStatement{
			SimpleSelect{QuerySpecification: QuerySpecification{
				SelectSpecs: nil,
				SelectElements: SelectElements{
					All: true,
				},
				FromClause: &FromClause{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
				},
			}},
		}, result)
	})

}
