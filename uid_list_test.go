package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitUidList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a, b")
		result := mySqlParser.UidList().Accept(visitor)
		assert.EqualValues(t,
			[]string{"a", "b"}, result)
	})
}
