package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitInsertStatementValue(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("SELECT * FROM U")
		result := mySqlParser.InsertStatementValue().Accept(visitor)
		assert.EqualValues(t,
			InsertStatementValueSelectStatement{SelectStatement: SimpleSelect{QuerySpecification: QuerySpecification{
				SelectSpecs: nil,
				SelectElements: SelectElements{
					All: true,
				},
				FromClause: &FromClause{
					TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "U"}},
					},
				},
			}}},
			result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("VALUES (1,2),(2,3)")
		result := mySqlParser.InsertStatementValue().Accept(visitor)
		assert.EqualValues(t,
			InsertStatementValuedExpressionsWithDefaults{ExpressionsWithDefaultsList: []ExpressionsWithDefaults{
				{
					ExpressionOrDefaults: []ExpressionOrDefault{
						{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
						},
						{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
						},
					},
				},
				{
					ExpressionOrDefaults: []ExpressionOrDefault{
						{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
						},
						{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 3}}},
						},
					},
				},
			}},
			result)
	})
}
