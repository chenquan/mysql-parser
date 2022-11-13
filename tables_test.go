package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitTables(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a, b")
		result := mySqlParser.Tables().Accept(visitor)
		assert.EqualValues(t,
			[]TableName{
				{
					Uid:   "a",
					DotId: "",
				}, {
					Uid:   "b",
					DotId: "",
				},
			}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a.c, b.c")
		result := mySqlParser.Tables().Accept(visitor)
		assert.EqualValues(t,
			[]TableName{
				{
					Uid:   "a",
					DotId: "c",
				}, {
					Uid:   "b",
					DotId: "c",
				},
			}, result)
	})
}
