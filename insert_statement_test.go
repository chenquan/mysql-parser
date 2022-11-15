package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitInsertStatement(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("INSERT INTO A VALUES (1,2),(2,3)")
		result := mySqlParser.InsertStatement().Accept(visitor)
		assert.EqualValues(t,
			InsertStatementIntoValue{
				Priority: "",
				Ignore:   false,
				TableName: TableName{
					Uid:   "A",
					DotId: "",
				},
				Columns: nil,
				InsertStatementValue: InsertStatementValuedExpressionsWithDefaults{ExpressionsWithDefaultsList: []ExpressionsWithDefaults{
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
			},
			result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("INSERT INTO A(a,b) VALUES (1,2),(2,3)")
		result := mySqlParser.InsertStatement().Accept(visitor)
		assert.EqualValues(t,
			InsertStatementIntoValue{
				Priority: "",
				Ignore:   false,
				TableName: TableName{
					Uid:   "A",
					DotId: "",
				},
				Columns: []string{"a", "b"},
				InsertStatementValue: InsertStatementValuedExpressionsWithDefaults{ExpressionsWithDefaultsList: []ExpressionsWithDefaults{
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
			},
			result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("INSERT DELAYED Ignore INTO A SET a=1,b=2")
		result := mySqlParser.InsertStatement().Accept(visitor)
		assert.EqualValues(t,
			InsertStatementSetValue{
				Priority: "DELAYED",
				Ignore:   true,
				TableName: TableName{
					Uid:   "A",
					DotId: "",
				},
				UpdatedElements: []UpdatedElement{
					{
						FullColumnName: FullColumnName{
							Uid:       "a",
							DottedIds: nil,
						},
						Value: ExpressionOrDefault{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
						},
					},

					{
						FullColumnName: FullColumnName{
							Uid:       "b",
							DottedIds: nil,
						},
						Value: ExpressionOrDefault{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
						},
					},
				},
			},
			result)
	})

	t.Run("4", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("INSERT DELAYED Ignore INTO A SET a=1,b=2 ON DUPLICATE KEY UPDATE c=1")
		result := mySqlParser.InsertStatement().Accept(visitor)
		assert.EqualValues(t,
			InsertStatementSetValue{
				Priority: "DELAYED",
				Ignore:   true,
				TableName: TableName{
					Uid:   "A",
					DotId: "",
				},
				UpdatedElements: []UpdatedElement{
					{
						FullColumnName: FullColumnName{
							Uid:       "a",
							DottedIds: nil,
						},
						Value: ExpressionOrDefault{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
						},
					},

					{
						FullColumnName: FullColumnName{
							Uid:       "b",
							DottedIds: nil,
						},
						Value: ExpressionOrDefault{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
						},
					},
				},
				DuplicatedUpdatedElements: []UpdatedElement{
					{
						FullColumnName: FullColumnName{
							Uid:       "c",
							DottedIds: nil,
						},
						Value: ExpressionOrDefault{
							Default:    false,
							Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
						},
					},
				},
			},
			result)
	})

}
