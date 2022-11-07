package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitDefaultValue(t *testing.T) {
	t.Run("NULL", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("NULL")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValueNullLiteral{}, result)
	})

	t.Run("CAST", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CAST( a AS int)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValueCast{
			Expression: ExpressionAtomPredicate{
				ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}},
			},
			ConvertedDataType: ConvertedDataType{
				TypeName:        "int",
				LengthDimension: nil,
			},
		}, result)

	})

	t.Run("CONSTANT", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("2")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValueConstant{
			Constant: ConstantDecimal{Val: 2},
		}, result)

		mySqlParser, visitor = createMySqlParser("not 2")
		result = mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValueConstant{
			UnaryOperator: "not",
			Constant:      ConstantDecimal{Val: 2},
		}, result)

	})

	t.Run("WRAP_EXPRESSION", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("(2)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValueExpression{
			HasBracket: true,
			Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
		}, result)
	})

	t.Run("LASTVAL", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("LASTVAL (a)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValueLastval{
			Val: FullId{
				Uid:   "a",
				DotId: "",
			},
		}, result)
	})

	t.Run("NEXTVAL", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("NEXTVAL (a)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValueNextval{
			Val: FullId{
				Uid:   "a",
				DotId: "",
			},
		}, result)
	})

	t.Run("PREVIOUS", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("(PREVIOUS VALUE FOR A)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValuePreviousValue{
			Val: FullId{Uid: "A"},
		}, result)
	})

	t.Run("NEXT", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("(NEXT VALUE FOR A)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValuePreviousValue{
			Val: FullId{Uid: "A"},
		}, result)
	})

	t.Run("EXPRESSION", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a+2")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValueExpression{
			Expression: ExpressionAtomPredicate{ExpressionAtom: MathExpressionAtom{
				LeftExpressionAtom:  FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}},
				MathOperator:        "+",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}},
			}},
		}, result)
	})

}
