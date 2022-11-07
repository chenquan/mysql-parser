package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_AlterTableAddIndex(t *testing.T) {
	result := Parser("ALTER TABLE PERSONS\n ADD INDEX user_name_index (user_name);")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				AddIndexes: []TableAddIndex{{
					ifNotExists: false,
					indexName:   "user_name_index",
					indexType:   "",
					columns:     []string{"user_name"},
				}},
			},
		},
	)
}

func TestParser_AlterTablePrimaryKey(t *testing.T) {
	result := Parser("ALTER TABLE Persons\n ADD PRIMARY KEY  (user_name) USING HASH;")
	assert.EqualValues(t, []SqlStatement{AlterTable{
		tableName:    "Persons",
		AddColumns:   nil,
		DeleteColumn: nil,
		AddPrimaryKeys: []TableAddPrimaryKey{
			{
				index:     "",
				indexType: "HASH",
				columns:   []string{"user_name"},
			},
		},
	}}, result)
}

func TestParser_AlterTableUniqueKey(t *testing.T) {
	result := Parser("ALTER TABLE PERSONS\n ADD UNIQUE user_name_index (user_name);")
	assert.EqualValues(t,
		[]SqlStatement{
			AlterTable{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				AddUniqueKeys: []TableAddUniqueKey{{
					indexName: "user_name_index",
					indexType: "",
					columns:   []string{"user_name"},
				}},
			},
		},
		result,
	)
}

func TestParser_AlterTableModifyColumn(t *testing.T) {
	result := Parser("ALTER TABLE PERSONS\n MODIFY COLUMN user_name varchar(100);")
	assert.EqualValues(t,
		[]SqlStatement{
			AlterTable{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				ModifyColumns: []TableModifyColumn{{
					ifExists:         false,
					column:           "user_name",
					columnDefinition: ColumnDefinition{DataType: "varchar(100)"},
				}},
			},
		},
		result,
	)
}

func TestParser_AlterTableDropColumn(t *testing.T) {
	result := Parser("ALTER TABLE PERSONS\n DROP COLUMN user_name;")
	assert.EqualValues(t,
		[]SqlStatement{
			AlterTable{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					ifExists: false,
					column:   "user_name",
					restrict: false,
				}},
			},
		},
		result,
	)

	result = Parser("ALTER TABLE PERSONS\n DROP COLUMN user_name RESTRICT;")
	assert.EqualValues(
		t,
		[]SqlStatement{
			AlterTable{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					ifExists: false,
					column:   "user_name",
					restrict: true,
				}},
			},
		},
		result,
	)

	result = Parser("ALTER TABLE PERSONS\n DROP COLUMN IF EXISTS user_name RESTRICT;")
	assert.EqualValues(t,
		[]SqlStatement{
			AlterTable{
				tableName:    "PERSONS",
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					ifExists: true,
					column:   "user_name",
					restrict: true,
				}},
			},
		},
		result,
	)
}

func TestParser_AlterTableDropPrimaryKey(t *testing.T) {
	result := Parser("ALTER TABLE PERSONS\n DROP PRIMARY KEY;")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName:      "PERSONS",
				DropPrimaryKey: true,
			},
		},
	)
}

func TestParser_AlterTableRenameIndex(t *testing.T) {
	result := Parser("ALTER TABLE PERSONS\n RENAME INDEX A TO B;")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName: "PERSONS",
				RenameIndexes: []TableRenameIndex{
					{
						FromColumn: "A",
						ToColumn:   "B",
					},
				},
			},
		},
	)
}

func TestParser_AlterTableDropIndex(t *testing.T) {
	result := Parser("ALTER TABLE PERSONS\n DROP INDEX U")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName: "PERSONS",
				DropIndexes: []TableDropIndex{
					{
						ifExists: false,
						column:   "U",
					},
				},
			},
		},
	)

	result = Parser("ALTER TABLE PERSONS\n DROP INDEX IF EXISTS U")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName: "PERSONS",
				DropIndexes: []TableDropIndex{
					{
						ifExists: true,
						column:   "U",
					},
				},
			},
		},
	)
}

func TestParser_AlterTableRename(t *testing.T) {
	result := Parser("ALTER TABLE PERSONS\n RENAME A")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName: "PERSONS",
				Renames:   []string{"A"},
			},
		},
	)
}
