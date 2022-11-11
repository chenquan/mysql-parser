package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitCreateIndex(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CREATE OR REPLACE INDEX  index_name USING HASH ON user (a)")
		result := mySqlParser.CreateIndex().Accept(visitor)

		assert.EqualValues(t, CreateIndex{
			Replace:       true,
			InTimeAction:  "",
			IndexCategory: "",
			IndexType:     "HASH",
			IndexName:     "index_name",
			TableName: TableName{
				Uid:   "user",
				DotId: "",
			},
			IndexColumnNames: []IndexColumnName{
				{
					IndexColumnName:   "a",
					IndexColumnLength: 0,
					SortType:          "",
				},
			},
			IndexOption: nil,
			Algorithm:   "",
			Lock:        "",
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CREATE OR REPLACE ONLINE UNIQUE  INDEX  index_name USING HASH ON user (a(1) DESC) KEY_BLOCK_SIZE 1 ALGORITHM=DEFAULT LOCK=DEFAULT")
		result := mySqlParser.CreateIndex().Accept(visitor)

		assert.EqualValues(t, CreateIndex{
			Replace:       true,
			InTimeAction:  "ONLINE",
			IndexCategory: "UNIQUE",
			IndexType:     "HASH",
			IndexName:     "index_name",
			TableName: TableName{
				Uid:   "user",
				DotId: "",
			},
			IndexColumnNames: []IndexColumnName{
				{
					IndexColumnName:   "a",
					IndexColumnLength: 1,
					SortType:          "DESC",
				},
			},
			IndexOption: []IndexOption{IndexOptionKeyBlockSize{BlockSize: "1"}},
			Algorithm:   "DEFAULT",
			Lock:        "DEFAULT",
		}, result)
	})
}
