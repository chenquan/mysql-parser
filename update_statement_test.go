package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitSingleUpdateStatement(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("UPDATE LOW_PRIORITY IGNORE u as t SET a=1,b=2 WHERE a>1 AND b>2 ORDER BY a desc LIMIT 1")
		result := mySqlParser.SingleUpdateStatement().Accept(visitor)
		assert.EqualValues(t,
			SingleUpdateStatement{
				LowPriority: true,
				Ignore:      true,
				TableName: TableName{
					Uid:   "u",
					DotId: "",
				},
				Alias: "t",
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
				WhereExpression: LogicalExpression{
					LeftExpression: BinaryComparisonPredicate{
						LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}}},
						ComparisonOperator: ">",
						RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
					},
					LogicalOperator: "AND",
					RightExpression: BinaryComparisonPredicate{
						LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "b"}}},
						ComparisonOperator: ">",
						RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
					},
				},
				OrderByClause: OrderByClause{
					OrderByExpressions: []OrderByExpression{
						{
							Expression: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}}},
							Order:      "desc",
						},
					},
				},
				LimitClause: LimitClause{
					Offset: 0,
					Limit:  1,
				},
			}, result)
	})
}

func Test_parseTreeVisitor_VisitMultipleUpdateStatement(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("UPDATE LOW_PRIORITY IGNORE t1, t2 SET a=1,b=2 WHERE t1.a>1 AND t2.b>2")
		result := mySqlParser.SingleUpdateStatement().Accept(visitor)
		assert.EqualValues(t,
			MultipleUpdateStatement{
				LowPriority: true,
				Ignore:      true,
				TableSources: []TableSource{
					TableSourceBase{
						TableSourceItem: AtomTableItem{TableName: "t1"},
					},
					TableSourceBase{
						TableSourceItem: AtomTableItem{TableName: "t2"},
					},
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
				WhereExpression: LogicalExpression{
					LeftExpression: BinaryComparisonPredicate{
						LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "t1", DottedIds: []DottedId{{Uid: "a"}}}}},
						ComparisonOperator: ">",
						RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
					},
					LogicalOperator: "AND",
					RightExpression: BinaryComparisonPredicate{
						LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "t2", DottedIds: []DottedId{{Uid: "b"}}}}},
						ComparisonOperator: ">",
						RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
					},
				},
			}, result)
	})
}
