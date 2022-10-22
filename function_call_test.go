package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitSpecificFunctionCall(t *testing.T) {
	t.Run("SimpleFunctionCall", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CURRENT_DATE")
		result := mySqlParser.SpecificFunction().Accept(visitor)
		assert.EqualValues(t, SimpleFunctionCall{Function: "CURRENT_DATE"}, result)
	})

	t.Run("SimpleFunctionCall", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CONVERT( a , CHAR(2) )")
		result := mySqlParser.SpecificFunction().Accept(visitor)
		assert.EqualValues(t, DataTypeFunctionCall{
			Function:   "CONVERT",
			Expression: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}}},
			ConvertedDataType: ConvertedDataType{
				TypeName:        "CHAR",
				LengthDimension: LengthOneDimension{Dimension: 2},
			},
		}, result)

	})

	t.Run("CONVERT_DATETIME", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CONVERT(a,DATETIME)")
		result := mySqlParser.SpecificFunction().Accept(visitor)
		assert.EqualValues(t, DataTypeFunctionCall{
			Function:   "CONVERT",
			Expression: ExpressionAtomPredicate{ExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{Uid: "a"}}},
			ConvertedDataType: ConvertedDataType{
				TypeName: "DATETIME",
			},
		}, result)
	})
}
