package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_AlterTableAddIndex(t *testing.T) {
	result := Parser("ALTER TABLE Persons\n ADD INDEX user_name_index (user_name);")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{{
			tableName:    "PERSONS",
			AddColumns:   nil,
			DeleteColumn: nil,
			AddIndexes: []TableAddIndex{{
				ifNotExists: false,
				indexName:   "USER_NAME_INDEX",
				indexType:   "",
				columns:     []string{"USER_NAME"},
			}},
		}},
	)
}

func TestParser_AlterTablePrimaryKey(t *testing.T) {
	//result := Parser("ALTER TABLE Persons\n ADD PRIMARY KEY USING HASH (user_name);")
	result := Parser("ALTER TABLE Persons\n ADD PRIMARY KEY  (user_name) USING HASH;")
	fmt.Println(result)
}

func TestParser_AlterTableUniqueKey(t *testing.T) {
	//result := Parser("ALTER TABLE Persons\n ADD PRIMARY KEY USING HASH (user_name);")
	result := Parser("ALTER TABLE Persons\n ADD UNIQUE user_name_index (user_name);")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{
			{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				AddUniqueKeys: []TableAddUniqueKey{{
					indexName: "USER_NAME_INDEX",
					indexType: "",
					columns:   []string{"USER_NAME"},
				}},
			},
		},
	)
}

func TestParser_AlterTableModifyColumn(t *testing.T) {
	result := Parser("ALTER TABLE Persons\n MODIFY COLUMN user_name varchar(100);")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{
			{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				ModifyColumns: []TableModifyColumn{{
					ifExists:         false,
					column:           "USER_NAME",
					columnDefinition: ColumnDefinition{dataType: "VARCHAR(100)"},
				}},
			},
		},
	)
}

func TestParser_AlterTableDropColumn(t *testing.T) {
	result := Parser("ALTER TABLE Persons\n DROP COLUMN user_name;")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{
			{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					ifExists: false,
					column:   "USER_NAME",
					restrict: false,
				}},
			},
		},
	)

	result = Parser("ALTER TABLE Persons\n DROP COLUMN user_name RESTRICT;")
	assert.EqualValues(
		t,
		result.AlterTables,
		[]AlterTable{
			{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					ifExists: false,
					column:   "USER_NAME",
					restrict: true,
				}},
			},
		},
	)

	result = Parser("ALTER TABLE Persons\n DROP COLUMN IF EXISTS user_name RESTRICT;")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{
			{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					ifExists: true,
					column:   "USER_NAME",
					restrict: true,
				}},
			},
		},
	)
}

func TestParser_AlterTableDropPrimaryKey(t *testing.T) {
	result := Parser("ALTER TABLE Persons\n DROP PRIMARY KEY;")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{{
			tableName:      "PERSONS",
			DropPrimaryKey: true,
		}},
	)
}

func TestParser_AlterTableRenameIndex(t *testing.T) {
	result := Parser("ALTER TABLE Persons\n RENAME INDEX A TO B;")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{{
			tableName: "PERSONS",
			RenameIndexes: []TableRenameIndex{
				{
					FromColumn: "A",
					ToColumn:   "B",
				},
			},
		}},
	)
}

func TestParser_AlterTableDropIndex(t *testing.T) {
	result := Parser("ALTER TABLE Persons\n DROP INDEX U")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{{
			tableName: "PERSONS",
			DropIndexes: []TableDropIndex{
				{
					ifExists: false,
					column:   "U",
				},
			},
		}},
	)

	result = Parser("ALTER TABLE Persons\n DROP INDEX IF EXISTS U")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{{
			tableName: "PERSONS",
			DropIndexes: []TableDropIndex{
				{
					ifExists: true,
					column:   "U",
				},
			},
		}},
	)
}

func TestParser_AlterTableRename(t *testing.T) {
	result := Parser("ALTER TABLE Persons\n RENAME A")
	assert.EqualValues(t,
		result.AlterTables,
		[]AlterTable{{
			tableName: "PERSONS",
			Renames:   []string{"A"},
		}},
	)
}
