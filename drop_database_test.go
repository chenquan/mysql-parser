package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitDropDatabase(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("DROP DATABASE A")
		result := mySqlParser.DropDatabase().Accept(visitor)

		assert.EqualValues(t, DropDatabase{DatabaseName: "A"}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("DROP DATABASE IF EXISTS A")
		result := mySqlParser.DropDatabase().Accept(visitor)

		assert.EqualValues(t, DropDatabase{DatabaseName: "A", IfExists: true}, result)
	})
}
