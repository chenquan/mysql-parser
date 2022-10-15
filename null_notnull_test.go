package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitNullNotnull(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("NOT NULL")
	result := mySqlParser.NullNotnull().Accept(visitor)
	assert.EqualValues(t, NullNotnull{
		Not:  true,
		Null: "NULL",
	}, result)
}
