package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitSimpleFunctionCall(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("CURRENT_DATE")
	result := mySqlParser.SpecificFunction().Accept(visitor)
	assert.EqualValues(t, SimpleFunctionCall{function: "CURRENT_DATE"}, result)

	mySqlParser, visitor = createMySqlParser("CURRENT_DATE()")
	result = mySqlParser.SpecificFunction().Accept(visitor)
	assert.EqualValues(t, SimpleFunctionCall{function: "CURRENT_DATE"}, result)
}
