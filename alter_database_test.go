package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitAlterSimpleDatabase(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("ALTER DATABASE a CHARSET DEFAULT")
		result := mySqlParser.AlterDatabase().Accept(visitor)
		assert.EqualValues(t, AlterSimpleDatabase{
			DatabaseName:          "a",
			CreateDatabaseOptions: []CreateDatabaseOption{CreateDatabaseOptionCharSet{CharSet: "CHARSET", CharsetName: "DEFAULT"}},
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("ALTER DATABASE DEFAULT CHARSET DEFAULT COLLATE a")
		result := mySqlParser.AlterDatabase().Accept(visitor)
		assert.EqualValues(t, AlterSimpleDatabase{
			DatabaseName: "",
			CreateDatabaseOptions: []CreateDatabaseOption{
				CreateDatabaseOptionCharSet{CharSet: "CHARSET", CharsetName: "DEFAULT"},
				CreateDatabaseOptionCollate{CollationName: "a"},
			},
		}, result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("ALTER DATABASE a UPGRADE DATA DIRECTORY NAME")
		result := mySqlParser.AlterDatabase().Accept(visitor)
		assert.EqualValues(t, AlterUpgradeName{DatabaseName: "a"}, result)
	})
}
