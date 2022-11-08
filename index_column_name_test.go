package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitIndexColumnName(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("(a)")
		result := mySqlParser.IndexColumnNames().Accept(visitor)
		assert.EqualValues(t,
			[]IndexColumnName{
				{
					IndexColumnName:   "a",
					IndexColumnLength: 0,
					SortType:          "",
				},
			},
			result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("('a' DESC,c(1))")
		result := mySqlParser.IndexColumnNames().Accept(visitor)
		assert.EqualValues(t,
			[]IndexColumnName{
				{
					IndexColumnName:   "'a'",
					IndexColumnLength: 0,
					SortType:          "DESC",
				},
				{
					IndexColumnName:   "c",
					IndexColumnLength: 1,
					SortType:          "",
				},
			},
			result)
	})

}
