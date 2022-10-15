package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_CreateTableCopyCreateTable(t *testing.T) {
	result := Parser("CREATE TABLE A LIKE B;")
	assert.EqualValues(t, []CreatTable{CopyCreateTable{
		IfNotExists:   false,
		Replace:       false,
		Temporary:     false,
		FromTableName: "A",
		ToTableName:   "B",
	}}, result.CreatTables)

	result = Parser("CREATE TABLE IF NOT EXISTS A LIKE B;")
	assert.EqualValues(t, []CreatTable{CopyCreateTable{
		IfNotExists:   true,
		Replace:       false,
		Temporary:     false,
		FromTableName: "A",
		ToTableName:   "B",
	}}, result.CreatTables)

	result = Parser("CREATE OR REPLACE TABLE IF NOT EXISTS A LIKE B;")
	assert.EqualValues(t, []CreatTable{CopyCreateTable{
		IfNotExists:   true,
		Replace:       true,
		Temporary:     false,
		FromTableName: "A",
		ToTableName:   "B",
	}}, result.CreatTables)

	result = Parser("CREATE OR REPLACE TEMPORARY TABLE IF NOT EXISTS A LIKE B;")
	assert.EqualValues(t, []CreatTable{CopyCreateTable{
		IfNotExists:   true,
		Replace:       true,
		Temporary:     true,
		FromTableName: "A",
		ToTableName:   "B",
	}}, result.CreatTables)
}

func TestParser_CreateTableQueryCreateTable(t *testing.T) {
	//result := Parser("CREATE TABLE A LIKE B;")

}
