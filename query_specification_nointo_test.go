package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitQuerySpecificationNointo(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT * FROM U`)
		result := parser.QuerySpecificationNointo().Accept(visitor)
		assert.EqualValues(t, QuerySpecificationNointo{
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
		}, result)

	})

	t.Run("2", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT count(*) as cnt FROM U`)
		result := parser.QuerySpecificationNointo().Accept(visitor)
		assert.EqualValues(t, QuerySpecificationNointo{
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
		}, result)
	})

	t.Run("3", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT * FROM U`)
		result := parser.QuerySpecificationNointo().Accept(visitor)
		assert.EqualValues(t, QuerySpecificationNointo{
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
		}, result)
	})

	t.Run("4", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT * FROM U group by a desc, c asc`)
		result := parser.QuerySpecificationNointo().Accept(visitor)
		assert.EqualValues(t, QuerySpecificationNointo{
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
			GroupByClause: &GroupByClause{GroupByItems: []GroupByItem{
				{
					Expression: ExpressionAtomPredicate{
						ExpressionAtom: FullColumnNameExpressionAtom{
							FullColumnName: FullColumnName{Uid: "a"},
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
		}, result)
	})

	t.Run("5", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT name,avg(a) FROM U group by name having avg(a) > 1`)
		result := parser.QuerySpecificationNointo().Accept(visitor)
		assert.EqualValues(t, QuerySpecificationNointo{
			SelectSpecs: nil,
			SelectElements: SelectElements{
				SelectElements: []SelectElement{
					SelectColumnElement{FullColumnName: FullColumnName{Uid: "name"}},
					SelectFunctionElement{FunctionCall: AggregateWindowedFunction{
						Function:    "avg",
						StarArg:     false,
						Aggregator:  "",
						FunctionArg: &FunctionArg{F: FullColumnName{Uid: "a"}},
					}},
				},
			},
			FromClause: &FromClause{
				TableSources: &TableSources{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
				},
			},
			GroupByClause: &GroupByClause{
				GroupByItems: []GroupByItem{
					{
						Expression: ExpressionAtomPredicate{
							ExpressionAtom: FullColumnNameExpressionAtom{
								FullColumnName: FullColumnName{Uid: "name"},
							},
						},
					},
				},
			},
			HavingClause: &HavingClause{
				Expression: BinaryComparisonPredicate{
					LeftPredicate: ExpressionAtomPredicate{
						ExpressionAtom: FunctionCallExpressionAtom{FunctionCall: AggregateWindowedFunction{
							Function: "avg",
							FunctionArg: &FunctionArg{F: FullColumnName{
								Uid: "a",
							}},
						}},
					},
					ComparisonOperator: ">",
					RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				},
			},
		}, result)
	})

	t.Run("6", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT name FROM U order by a desc `)
		result := parser.QuerySpecificationNointo().Accept(visitor)
		assert.EqualValues(t, QuerySpecificationNointo{
			SelectSpecs: nil,
			SelectElements: SelectElements{
				SelectElements: []SelectElement{
					SelectColumnElement{FullColumnName: FullColumnName{Uid: "name"}},
				},
			},
			FromClause: &FromClause{
				TableSources: &TableSources{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
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

	t.Run("7", func(t *testing.T) {
		parser, visitor := createMySqlParser(`SELECT name FROM U LIMIT 2`)
		result := parser.QuerySpecificationNointo().Accept(visitor)
		assert.EqualValues(t, QuerySpecificationNointo{
			SelectSpecs: nil,
			SelectElements: SelectElements{
				SelectElements: []SelectElement{
					SelectColumnElement{FullColumnName: FullColumnName{Uid: "name"}},
				},
			},
			FromClause: &FromClause{
				TableSources: &TableSources{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
				},
			},
			LimitClause: &LimitClause{Limit: 2},
		}, result)

		parser, visitor = createMySqlParser(`SELECT name FROM U LIMIT 2,3 `)
		result = parser.QuerySpecification().Accept(visitor)
		assert.EqualValues(t, QuerySpecification{
			SelectSpecs: nil,
			SelectElements: SelectElements{
				SelectElements: []SelectElement{
					SelectColumnElement{FullColumnName: FullColumnName{Uid: "name"}},
				},
			},
			FromClause: &FromClause{
				TableSources: &TableSources{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
				},
			},
			LimitClause: &LimitClause{Limit: 3, Offset: 2},
		}, result)

		parser, visitor = createMySqlParser(`SELECT name FROM U LIMIT 3 OFFSET 2 `)
		result = parser.QuerySpecification().Accept(visitor)
		assert.EqualValues(t, QuerySpecification{
			SelectSpecs: nil,
			SelectElements: SelectElements{
				SelectElements: []SelectElement{
					SelectColumnElement{FullColumnName: FullColumnName{Uid: "name"}},
				},
			},
			FromClause: &FromClause{
				TableSources: &TableSources{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
				},
			},
			LimitClause: &LimitClause{Limit: 3, Offset: 2},
		}, result)
	})

}
