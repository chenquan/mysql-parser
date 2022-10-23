package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitSimpleSelect(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT * FROM U`)
		result := parser.SelectStatement().Accept(visitor)
		assert.EqualValues(t, SimpleSelect{QuerySpecification: QuerySpecification{
			SelectSpecs: nil,
			SelectElements: SelectElements{
				All: true,
			},
			FromClause: &FromClause{
				TableSources: &TableSources{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
				},
			},
		}}, result)

	})
	t.Run("2", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT count(*) as cnt FROM U`)
		result := parser.SelectStatement().Accept(visitor)
		assert.EqualValues(t, SimpleSelect{QuerySpecification: QuerySpecification{
			SelectSpecs: nil,
			SelectElements: SelectElements{
				SelectElements: []SelectElement{
					SelectFunctionElement{
						FunctionCall: AggregateWindowedFunction{
							Function: "count",
							StarArg:  true,
						},
						Alias: "cnt",
					},
				},
			},
			FromClause: &FromClause{
				TableSources: &TableSources{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
				},
			},
		}}, result)

	})
}
