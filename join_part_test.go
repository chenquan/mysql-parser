package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitInnerJoin(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("INNER JOIN a ON a.name = 1")
		result := mySqlParser.JoinPart().Accept(visitor)
		assert.EqualValues(t,
			InnerJoin{
				Join:            "INNER",
				TableSourceItem: TableSourceItem(AtomTableItem{TableName: "a"}),
				OnExpression: Expression(BinaryComparisonPredicate{
					LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a", DottedIds: []DottedId{{Uid: "name"}}}}},
					ComparisonOperator: "=",
					RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				}),
				UsingUidList: nil,
			},
			result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CROSS JOIN a USING (a,b,c)")
		result := mySqlParser.JoinPart().Accept(visitor)
		assert.EqualValues(t,
			InnerJoin{
				Join:            "CROSS",
				TableSourceItem: TableSourceItem(AtomTableItem{TableName: "a"}),
				UsingUidList:    []string{"a", "b", "c"},
			},
			result)
	})
}

func Test_parseTreeVisitor_VisitStraightJoin(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("STRAIGHT_JOIN a ON a.name = 1")
		result := mySqlParser.JoinPart().Accept(visitor)
		assert.EqualValues(t,
			StraightJoin{
				TableSourceItem: TableSourceItem(AtomTableItem{TableName: "a"}),
				OnExpression: Expression(BinaryComparisonPredicate{
					LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a", DottedIds: []DottedId{{Uid: "name"}}}}},
					ComparisonOperator: "=",
					RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				}),
			},
			result)
	})
}

func Test_parseTreeVisitor_VisitOuterJoin(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("LEFT OUTER a ON a.name = 1")
		result := mySqlParser.JoinPart().Accept(visitor)
		assert.EqualValues(t,
			OuterJoin{
				Join:            "LEFT",
				TableSourceItem: TableSourceItem(AtomTableItem{TableName: "a"}),
				OnExpression: Expression(BinaryComparisonPredicate{
					LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a", DottedIds: []DottedId{{Uid: "name"}}}}},
					ComparisonOperator: "=",
					RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				}),
				UsingUidList: nil,
			},
			result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("LEFT JOIN a USING (a,b,c)")
		result := mySqlParser.JoinPart().Accept(visitor)
		assert.EqualValues(t,
			OuterJoin{
				Join:            "LEFT",
				TableSourceItem: TableSourceItem(AtomTableItem{TableName: "a"}),
				UsingUidList:    []string{"a", "b", "c"},
			},
			result)
	})
}

func Test_parseTreeVisitor_VisitNaturalJoin(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("NATURAL LEFT JOIN a")
		result := mySqlParser.JoinPart().Accept(visitor)
		assert.EqualValues(t,
			NaturalJoin{
				Join:            "LEFT",
				TableSourceItem: TableSourceItem(AtomTableItem{TableName: "a"}),
			},
			result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("NATURAL RIGHT OUTER JOIN a")
		result := mySqlParser.JoinPart().Accept(visitor)
		assert.EqualValues(t,
			NaturalJoin{
				Join:            "RIGHT",
				TableSourceItem: TableSourceItem(AtomTableItem{TableName: "a"}),
			},
			result)
	})

}
