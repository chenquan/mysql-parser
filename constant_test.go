package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitConstant(t *testing.T) {
	t.Run("ConstantDecimal", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("-323")
		result := mySqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantDecimal{Val: -323}, result)

		mySqlParser, visitor = createMySqlParser("323")
		result = mySqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantDecimal{Val: 323}, result)
	})

	t.Run("ConstantNull", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("NULL")
		result := mySqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantNull{Val: "NULL"}, result)

		mySqlParser, visitor = createMySqlParser("NOT NULL")
		result = mySqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantNull{Val: "NOT NULL"}, result)
	})

	t.Run("ConstantString", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("'xxx'")
		result := mySqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantString{Val: "xxx"}, result)
	})

	t.Run("ConstantBool", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("TRUE")
		result := mySqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantBool{Val: true}, result)
	})
	t.Run("ConstantBool", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("x'DFFF'")
		result := mySqlParser.Constant().Accept(visitor)
		assert.EqualValues(t, ConstantHexadecimal{Val: "x'DFFF'"}, result)
	})

}
