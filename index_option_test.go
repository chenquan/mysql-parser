package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitIndexOption(t *testing.T) {
	t.Run("KEY_BLOCK_SIZE", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("KEY_BLOCK_SIZE 1")
		result := mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionKeyBlockSize{BlockSize: "1"},
			result)

		mySqlParser, visitor = createMySqlParser("KEY_BLOCK_SIZE 1M")
		result = mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionKeyBlockSize{BlockSize: "1M"},
			result)
	})

	t.Run("indexType", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("USING HASH")
		result := mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexType("HASH"),
			result)

		mySqlParser, visitor = createMySqlParser("USING BTREE")
		result = mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexType("BTREE"),
			result)
	})

	t.Run("COMMENT", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("COMMENT 'xxx'")
		result := mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionComment{Comment: "xxx"},
			result)

		mySqlParser, visitor = createMySqlParser("COMMENT `xxx`")
		result = mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionComment{Comment: "xxx"},
			result)

		mySqlParser, visitor = createMySqlParser(`COMMENT "xxx"`)
		result = mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionComment{Comment: "xxx"},
			result)
	})

	t.Run("VISIBLE", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("VISIBLE")
		result := mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionVisible{Visible: true},
			result)
	})

	t.Run("INVISIBLE", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("INVISIBLE")
		result := mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionVisible{},
			result)
	})

	t.Run("ENGINE_ATTRIBUTE", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("ENGINE_ATTRIBUTE 'innodb'")
		result := mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionEngineAttribute{EngineAttribute: "innodb"},
			result)
	})
	t.Run("SECONDARY_ENGINE_ATTRIBUTE", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("SECONDARY_ENGINE_ATTRIBUTE 'innodb'")
		result := mySqlParser.IndexOption().Accept(visitor)
		assert.EqualValues(t,
			IndexOptionSecondaryEngineAttribute{EngineAttribute: "innodb"},
			result)
	})

}
