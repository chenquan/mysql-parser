package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitCharSet(t *testing.T) {
	charSets := []string{
		"CHARSET",
		"CHARACTER SET",
		"CHAR SET",
	}

	for _, charSet := range charSets {
		t.Run(charSet, func(t *testing.T) {
			mySqlParser, visitor := createMySqlParser(charSet)
			result := mySqlParser.CharSet().Accept(visitor)

			assert.EqualValues(t, charSet, result)
		})
	}

}

func Test_parseTreeVisitor_VisitCreateDatabaseOption(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CHARACTER SET DEFAULT")
		result := mySqlParser.CreateDatabaseOption().Accept(visitor)

		assert.EqualValues(t, CreateDatabaseOptionCharSet{
			CharSet:     "CHARACTER SET",
			CharsetName: "DEFAULT",
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CHARACTER SET BINARY")
		result := mySqlParser.CreateDatabaseOption().Accept(visitor)

		assert.EqualValues(t, CreateDatabaseOptionCharSet{
			CharSet:     "CHARACTER SET",
			CharsetName: "BINARY",
		}, result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CHARACTER SET BIG5")
		result := mySqlParser.CreateDatabaseOption().Accept(visitor)

		assert.EqualValues(t, CreateDatabaseOptionCharSet{
			CharSet:     "CHARACTER SET",
			CharsetName: "BIG5",
		}, result)
	})

	t.Run("COLLATE", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("COLLATE a")
		result := mySqlParser.CreateDatabaseOption().Accept(visitor)

		assert.EqualValues(t, CreateDatabaseOptionCollate{
			CollationName: "a",
		}, result)
	})

	t.Run("ENCRYPTION", func(t *testing.T) {
		t.Run("1", func(t *testing.T) {
			mySqlParser, visitor := createMySqlParser("ENCRYPTION = 'a'")
			result := mySqlParser.CreateDatabaseOption().Accept(visitor)
			assert.EqualValues(t, CreateDatabaseOptionEncryption{
				Encryption: "a",
			}, result)
		})

		t.Run("2", func(t *testing.T) {
			mySqlParser, visitor := createMySqlParser("ENCRYPTION = `a`")
			result := mySqlParser.CreateDatabaseOption().Accept(visitor)
			assert.EqualValues(t, CreateDatabaseOptionEncryption{
				Encryption: "a",
			}, result)
		})

		t.Run("3", func(t *testing.T) {
			mySqlParser, visitor := createMySqlParser(`ENCRYPTION = "a"`)
			result := mySqlParser.CreateDatabaseOption().Accept(visitor)
			assert.EqualValues(t, CreateDatabaseOptionEncryption{
				Encryption: "a",
			}, result)
		})
	})

	t.Run("READ ONLY", func(t *testing.T) {
		t.Run("DEFAULT", func(t *testing.T) {
			mySqlParser, visitor := createMySqlParser(`READ ONLY DEFAULT`)
			result := mySqlParser.CreateDatabaseOption().Accept(visitor)
			assert.EqualValues(t, CreateDatabaseOptionReadonly{
				Default: true,
			}, result)
		})
		t.Run("1", func(t *testing.T) {
			mySqlParser, visitor := createMySqlParser(`READ ONLY 1`)
			result := mySqlParser.CreateDatabaseOption().Accept(visitor)
			assert.EqualValues(t, CreateDatabaseOptionReadonly{
				Default: false,
				Num:     1,
			}, result)
		})
		t.Run("0", func(t *testing.T) {
			mySqlParser, visitor := createMySqlParser(`READ ONLY 0`)
			result := mySqlParser.CreateDatabaseOption().Accept(visitor)
			assert.EqualValues(t, CreateDatabaseOptionReadonly{
				Default: false,
			}, result)
		})
	})

}
