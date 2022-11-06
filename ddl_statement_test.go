package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitDdlStatement(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CREATE TABLE A(a int default 1);")
		result := mySqlParser.DdlStatement().Accept(visitor)

		assert.EqualValues(t,
			ColumnCreateTable{
				IfNotExists: false,
				Replace:     false,
				Temporary:   false,
				Table:       "A",
				CreateDefinitions: []CreateDefinition{
					ColumnDeclaration{
						Column:           "a",
						ColumnDefinition: ColumnDefinition{DataType: "int"},
					},
				},
			}, result)
	})
}
