package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitConstant(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("-323")
	result := mySqlParser.Constant().Accept(visitor)
	assert.EqualValues(t, Constant{
		Constant: "-323",
	}, result)

	mySqlParser, visitor = createMySqlParser("NULL")
	result = mySqlParser.Constant().Accept(visitor)
	assert.EqualValues(t, Constant{
		Constant: "NULL",
	}, result)
}
