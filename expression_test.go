package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitNotExpression(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("CURRENT_DATE")
	result := mySqlParser.Expression().Accept(visitor)
	assert.EqualValues(t, IsNullPredicate{
		Predicate:   nil,
		NullNotnull: NullNotnull{},
	}, result)
	_ = result
}
