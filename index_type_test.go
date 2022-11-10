package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitIndexType(t *testing.T) {
	t.Run("HASH", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("USING HASH")
		result := mySqlParser.IndexType().Accept(visitor)
		assert.EqualValues(t,
			IndexType("HASH"),
			result)
	})

	t.Run("BTREE", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("USING BTREE")
		result := mySqlParser.IndexType().Accept(visitor)
		assert.EqualValues(t,
			IndexType("BTREE"),
			result)
	})

}
