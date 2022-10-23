package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitNotExpression(t *testing.T) {
	// TODO
	//mySqlParser, visitor := createMySqlParser("\"1!=1 AND a!=2\"")
	//result := mySqlParser.Expression().Accept(visitor)
	//assert.EqualValues(t, NotExpression, result)
}

func Test_parseTreeVisitor_VisitLogicalExpression(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1!=1 AND a!=2")
		result := mySqlParser.Expression().Accept(visitor)
		assert.EqualValues(t, LogicalExpression{
			LeftExpression: BinaryComparisonPredicate{
				LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				ComparisonOperator: "!=",
				RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
			},
			LogicalOperator: "AND",
			RightExpression: BinaryComparisonPredicate{
				LeftPredicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
					Uid: "a",
				}}},
				ComparisonOperator: "!=",
				RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
			},
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1!=1 AND a!=2 AND A!=3")
		result := mySqlParser.Expression().Accept(visitor)
		assert.EqualValues(t, LogicalExpression{
			LeftExpression: LogicalExpression{
				LeftExpression: BinaryComparisonPredicate{
					LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
					ComparisonOperator: "!=",
					RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				},
				LogicalOperator: "AND",
				RightExpression: BinaryComparisonPredicate{
					LeftPredicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
						Uid: "a",
					}}},
					ComparisonOperator: "!=",
					RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
				},
			},
			LogicalOperator: "AND",
			RightExpression: BinaryComparisonPredicate{
				LeftPredicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
					Uid: "A",
				}}},
				ComparisonOperator: "!=",
				RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 3}}},
			},
		}, result)
	})
}

func Test_parseTreeVisitor_VisitIsExpression(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1!=1 is true")
		result := mySqlParser.Expression().Accept(visitor)
		assert.EqualValues(t, IsExpression{
			predicateExpression: BinaryComparisonPredicate{
				LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				ComparisonOperator: "!=",
				RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
			},
			IsNot:     false,
			TestValue: "true",
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1!=1 IS NOT FALSE")
		result := mySqlParser.Expression().Accept(visitor)
		assert.EqualValues(t, IsExpression{
			predicateExpression: BinaryComparisonPredicate{
				LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
				ComparisonOperator: "!=",
				RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
			},
			IsNot:     true,
			TestValue: "FALSE",
		}, result)
	})

}

func Test_parseTreeVisitor_VisitPredicateExpression(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1!=1")
		result := mySqlParser.Expression().Accept(visitor)
		assert.EqualValues(t, BinaryComparisonPredicate{
			LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
			ComparisonOperator: "!=",
			RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
		}, result)
	})
}

func Test_parseTreeVisitor_VisitExpressions(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1,2")
		result := mySqlParser.Expressions().Accept(visitor)
		assert.EqualValues(t, []Expression{
			ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
			ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
		}, result)
	})
}
