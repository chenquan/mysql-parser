package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitLimitClause(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("LIMIT 1")
		result := mySqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t,
			LimitClause{
				Offset: 0,
				Limit:  1,
			},
			result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("LIMIT 1,4")
		result := mySqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t,
			LimitClause{
				Offset: 1,
				Limit:  4,
			},
			result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("LIMIT 4 offset 1")
		result := mySqlParser.LimitClause().Accept(visitor)
		assert.EqualValues(t,
			LimitClause{
				Offset: 1,
				Limit:  4,
			},
			result)
	})

}
