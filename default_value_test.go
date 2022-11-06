package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitDefaultValue(t *testing.T) {
	t.Run("NULL", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("NULL")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValue{
			Type:        "NULL",
			NullLiteral: true,
		}, result)
	})

	t.Run("CAST", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CAST( a AS int)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValue{
			Type: "CAST",
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

		assert.EqualValues(t, DefaultValue{
			Type:     "CONSTANT",
			Constant: ConstantDecimal{Val: 2},
		}, result)

		mySqlParser, visitor = createMySqlParser("not 2")
		result = mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValue{
			Type:          "CONSTANT",
			UnaryOperator: "not",
			Constant:      ConstantDecimal{Val: 2},
		}, result)

	})

	t.Run("WRAP_EXPRESSION", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("(2)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValue{
			Type:       "WRAP_EXPRESSION",
			Expression: ExpressionAtomPredicate{ExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: 2}}},
		}, result)
	})

	t.Run("LASTVAL", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("LASTVAL (a)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValue{
			Type: "LASTVAL",
			FullId: FullId{
				Uid:   "a",
				DotId: "",
			},
		}, result)
	})

	t.Run("NEXTVAL", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("NEXTVAL (a)")
		result := mySqlParser.DefaultValue().Accept(visitor)

		assert.EqualValues(t, DefaultValue{
			Type: "NEXTVAL",
			FullId: FullId{
				Uid:   "a",
				DotId: "",
			},
		}, result)
	})

}
