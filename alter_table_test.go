package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_AlterTableAddIndex(t *testing.T) {
	result := Parse("ALTER TABLE PERSONS\n ADD INDEX user_name_index using HASH (user_name);")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName:    TableName{Uid: "PERSONS"},
				AddColumns:   nil,
				DeleteColumn: nil,
				AddIndexes: []TableAddIndex{{
					IfNotExists: false,
					IndexName:   "user_name_index",
					IndexType:   "HASH",
					Columns: []IndexColumnName{
						{
							IndexColumnName:   "user_name",
							IndexColumnLength: 0,
							SortType:          "",
						},
					},
				}},
			},
		},
	)
}

func TestParser_AlterTablePrimaryKey(t *testing.T) {
	result := Parse("ALTER TABLE Persons\n ADD PRIMARY KEY  (user_name) USING HASH;")
	assert.EqualValues(t, []SqlStatement{AlterTable{
		tableName:    TableName{Uid: "Persons"},
		AddColumns:   nil,
		DeleteColumn: nil,
		AddPrimaryKeys: []TableAddPrimaryKey{
			{
				Index:     "",
				IndexType: "HASH",
				Columns: []IndexColumnName{
					{
						IndexColumnName:   "user_name",
						IndexColumnLength: 0,
						SortType:          "",
					},
				},
			},
		},
	}}, result)
}

func TestParser_AlterTableUniqueKey(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		result := Parse("ALTER TABLE PERSONS\n ADD UNIQUE user_name_index (user_name);")
		assert.EqualValues(t,
			[]SqlStatement{
				AlterTable{
					tableName:    TableName{Uid: "PERSONS"},
					AddColumns:   nil,
					DeleteColumn: nil,
					AddUniqueKeys: []TableAddUniqueKey{{
						IndexName: "user_name_index",
						IndexType: "",
						Columns: []IndexColumnName{
							{
								IndexColumnName:   "user_name",
								IndexColumnLength: 0,
								SortType:          "",
							},
						},
					}},
				},
			},
			result,
		)
	})
}

func TestParser_AlterTableModifyColumn(t *testing.T) {
	result := Parse("ALTER TABLE PERSONS\n MODIFY COLUMN user_name varchar(100);")
	assert.EqualValues(t,
		[]SqlStatement{
			AlterTable{
				tableName:    TableName{Uid: "PERSONS"},
				AddColumns:   nil,
				DeleteColumn: nil,
				ModifyColumns: []TableModifyColumn{{
					IfExists:         false,
					Column:           "user_name",
					ColumnDefinition: ColumnDefinition{DataType: "varchar(100)"},
				}},
			},
		},
		result,
	)
}

func TestParser_AlterTableDropColumn(t *testing.T) {
	result := Parse("ALTER TABLE PERSONS\n DROP COLUMN user_name;")
	assert.EqualValues(t,
		[]SqlStatement{
			AlterTable{
				tableName:    TableName{Uid: "PERSONS"},
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					IfExists: false,
					Column:   "user_name",
					Restrict: false,
				}},
			},
		},
		result,
	)

	result = Parse("ALTER TABLE PERSONS\n DROP COLUMN user_name RESTRICT;")
	assert.EqualValues(
		t,
		[]SqlStatement{
			AlterTable{
				tableName:    TableName{Uid: "PERSONS"},
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					IfExists: false,
					Column:   "user_name",
					Restrict: true,
				}},
			},
		},
		result,
	)

	result = Parse("ALTER TABLE PERSONS\n DROP COLUMN IF EXISTS user_name RESTRICT;")
	assert.EqualValues(t,
		[]SqlStatement{
			AlterTable{
				tableName:    TableName{Uid: "PERSONS"},
				AddColumns:   nil,
				DeleteColumn: nil,
				DropColumns: []TableDropColumn{{
					IfExists: true,
					Column:   "user_name",
					Restrict: true,
				}},
			},
		},
		result,
	)
}

func TestParser_AlterTableDropPrimaryKey(t *testing.T) {
	result := Parse("ALTER TABLE PERSONS\n DROP PRIMARY KEY;")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName:      TableName{Uid: "PERSONS"},
				DropPrimaryKey: true,
			},
		},
	)
}

func TestParser_AlterTableRenameIndex(t *testing.T) {
	result := Parse("ALTER TABLE PERSONS\n RENAME INDEX A TO B;")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName: TableName{Uid: "PERSONS"},
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
	result := Parse("ALTER TABLE PERSONS\n DROP INDEX U")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName: TableName{Uid: "PERSONS"},
				DropIndexes: []TableDropIndex{
					{
						IfExists: false,
						Column:   "U",
					},
				},
			},
		},
	)

	result = Parse("ALTER TABLE PERSONS\n DROP INDEX IF EXISTS U")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName: TableName{Uid: "PERSONS"},
				DropIndexes: []TableDropIndex{
					{
						IfExists: true,
						Column:   "U",
					},
				},
			},
		},
	)
}

func TestParser_AlterTableRename(t *testing.T) {
	result := Parse("ALTER TABLE PERSONS\n RENAME A")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName: TableName{Uid: "PERSONS"},
				Renames:   []string{"A"},
			},
		},
	)
}

func Test_parseTreeVisitor_VisitAlterByAddColumn(t *testing.T) {
	result := Parse("ALTER TABLE PERSONS\n ADD column a int;")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName:    TableName{Uid: "PERSONS"},
				DeleteColumn: nil,

				AddColumns: []TableAddColumn{
					{
						IfNotExists: false,
						Column:      "a",
						ColumnDefinition: ColumnDefinition{
							DataType:          "int",
							ColumnConstraints: nil,
						},
					},
				},
			},
		},
	)
}

func Test_parseTreeVisitor_VisitAlterByAddColumns(t *testing.T) {
	result := Parse("ALTER TABLE PERSONS\n ADD column (a int, b int);")
	assert.EqualValues(t,
		result,
		[]SqlStatement{
			AlterTable{
				tableName:    TableName{Uid: "PERSONS"},
				DeleteColumn: nil,
				AddColumns: []TableAddColumn{
					{
						IfNotExists: false,
						Column:      "a",
						ColumnDefinition: ColumnDefinition{
							DataType:          "int",
							ColumnConstraints: nil,
						},
					},
					{
						IfNotExists: false,
						Column:      "b",
						ColumnDefinition: ColumnDefinition{
							DataType:          "int",
							ColumnConstraints: nil,
						},
					},
				},
			},
		},
	)
}

func Test_parseTreeVisitor_VisitAlterByAddUniqueKey(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("Add UNIQUE u (a,b) USING HASH")
		result := mySqlParser.AlterSpecification().Accept(visitor)
		assert.EqualValues(t, TableAddUniqueKey{
			IndexName: "u",
			IndexType: "",
			Columns: []IndexColumnName{
				{
					IndexColumnName:   "a",
					IndexColumnLength: 0,
					SortType:          "",
				},
				{
					IndexColumnName:   "b",
					IndexColumnLength: 0,
					SortType:          "",
				},
			},
			IndexOptions: []IndexOption{
				IndexType("HASH"),
			},
		}, result)
	})

	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("Add UNIQUE u USING HASH (a,b) ")
		result := mySqlParser.AlterSpecification().Accept(visitor)
		assert.EqualValues(t, TableAddUniqueKey{
			IndexName: "u",
			IndexType: IndexType("HASH"),
			Columns: []IndexColumnName{
				{
					IndexColumnName:   "a",
					IndexColumnLength: 0,
					SortType:          "",
				},
				{
					IndexColumnName:   "b",
					IndexColumnLength: 0,
					SortType:          "",
				},
			},
		}, result)
	})
}
