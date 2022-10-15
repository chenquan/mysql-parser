package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitIsNullPredicate(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("NOT NULL")
	result := mySqlParser.Predicate().Accept(visitor)
	assert.EqualValues(t, IsNullPredicate{
		Predicate:   nil,
		NullNotnull: NullNotnull{},
	}, result)
}
