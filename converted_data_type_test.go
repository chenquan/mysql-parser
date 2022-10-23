package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitConvertedDataType(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("char(2)")
		result := mySqlParser.ConvertedDataType().Accept(visitor)

		assert.EqualValues(t,
			ConvertedDataType{
				TypeName:        "char",
				LengthDimension: LengthOneDimension{Dimension: 2},
			}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("DECIMAL(2,3)")
		result := mySqlParser.ConvertedDataType().Accept(visitor)

		assert.EqualValues(t,
			ConvertedDataType{
				TypeName: "DECIMAL",
				LengthDimension: LengthTwoDimension{
					Dimension1: 2,
					Dimension2: 3,
				},
			}, result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("DECIMAL(2)")
		result := mySqlParser.ConvertedDataType().Accept(visitor)

		assert.EqualValues(t,
			ConvertedDataType{
				TypeName: "DECIMAL",
				LengthDimension: LengthOneDimension{
					Dimension: 2,
				},
			}, result)
	})

}
