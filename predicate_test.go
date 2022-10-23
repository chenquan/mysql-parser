package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitIsNullPredicate(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("a is NOT NULL")
	result := mySqlParser.Predicate().Accept(visitor)
	assert.EqualValues(t, IsNullPredicate{
		Predicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
			Uid:       "a",
			DottedIds: nil,
		}}},
		NullNotnull: NullNotnull{
			Not:  true,
			Null: "NULL",
		},
	}, result)
}

func Test_parseTreeVisitor_VisitBinaryComparisonPredicate(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		parser, visitor := createMySqlParser("1!=1")
		result := parser.Predicate().Accept(visitor)
		assert.EqualValues(t, BinaryComparisonPredicate{
			LeftPredicate:      ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
			ComparisonOperator: "!=",
			RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		parser, visitor := createMySqlParser("a!=1")
		result := parser.Predicate().Accept(visitor)
		assert.EqualValues(t, BinaryComparisonPredicate{
			LeftPredicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
				Uid: "a",
			}}},
			ComparisonOperator: "!=",
			RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
		}, result)
	})

	t.Run("3", func(t *testing.T) {
		parser, visitor := createMySqlParser("a.b.c!=1")
		result := parser.Predicate().Accept(visitor)
		assert.EqualValues(t, BinaryComparisonPredicate{
			LeftPredicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
				Uid:       "a",
				DottedIds: []DottedId{{Uid: "b"}, {Uid: "c"}},
			}}},
			ComparisonOperator: "!=",
			RightPredicate:     ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}}},
		}, result)
	})

}

func Test_parseTreeVisitor_VisitInPredicate(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		parser, visitor := createMySqlParser("a in (1,23,3)")
		result := parser.Predicate().Accept(visitor)
		assert.EqualValues(t, InPredicate{
			Predicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}}},
			NotIn:     false,
			Expressions: []Expression{
				ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 1}},
				}, ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 23}},
				}, ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 3}},
				},
			},
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		parser, visitor := createMySqlParser(`a in ('1','2','3')`)
		result := parser.Predicate().Accept(visitor)
		assert.EqualValues(t, InPredicate{
			Predicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}}},
			NotIn:     false,
			Expressions: []Expression{
				ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantString{Val: "1"}},
				}, ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantString{Val: "2"}},
				}, ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantString{Val: "3"}},
				},
			},
		}, result)

		parser, visitor = createMySqlParser(`a in ("1","2","3")`)
		result = parser.Predicate().Accept(visitor)
		assert.EqualValues(t, InPredicate{
			Predicate: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}}},
			NotIn:     false,
			Expressions: []Expression{
				ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantString{Val: "1"}},
				}, ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantString{Val: "2"}},
				}, ExpressionAtomPredicate{
					ExpressionAtom: ConstantExpressionAtom{Constant: ConstantString{Val: "3"}},
				},
			},
		}, result)
	})
}
