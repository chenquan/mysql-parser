package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitCreateDatabase(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser(`CREATE DATABASE IF NOT EXISTS A READ ONLY 0 CHARACTER SET BIG5`)
		result := mySqlParser.CreateDatabase().Accept(visitor)
		assert.EqualValues(t, CreateDatabase{
			IfNotExists:  true,
			DatabaseName: "A",
			CreateDatabaseOptions: []CreateDatabaseOption{
				CreateDatabaseOptionReadonly{
					Default: false,
					Num:     0,
				},
				CreateDatabaseOptionCharSet{
					CharSet:     "CHARACTER SET",
					CharsetName: "BIG5",
				},
			},
		}, result)
	})

	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser(`CREATE DATABASE IF NOT EXISTS A `)
		result := mySqlParser.CreateDatabase().Accept(visitor)
		assert.EqualValues(t, CreateDatabase{
			IfNotExists:  true,
			DatabaseName: "A",
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser(`CREATE DATABASE IF NOT EXISTS A COLLATE A`)
		result := mySqlParser.CreateDatabase().Accept(visitor)
		assert.EqualValues(t, CreateDatabase{
			IfNotExists:  true,
			DatabaseName: "A",
			CreateDatabaseOptions: []CreateDatabaseOption{
				CreateDatabaseOptionCollate{CollationName: "A"},
			},
		}, result)
	})
}
