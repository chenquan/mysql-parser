package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser_CreateTableCopyCreateTable(t *testing.T) {
	result := Parse("CREATE TABLE A LIKE B;")
	assert.EqualValues(t, []SqlStatement{
		CopyCreateTable{
			IfNotExists:   false,
			Replace:       false,
			Temporary:     false,
			FromTableName: "A",
			ToTableName:   "B",
		},
	}, result)

	result = Parse("CREATE TABLE IF NOT EXISTS A LIKE B;")
	assert.EqualValues(t, []SqlStatement{CopyCreateTable{
		IfNotExists:   true,
		Replace:       false,
		Temporary:     false,
		FromTableName: "A",
		ToTableName:   "B",
	}}, result)

	result = Parse("CREATE OR REPLACE TABLE IF NOT EXISTS A LIKE B;")
	assert.EqualValues(t, []SqlStatement{CopyCreateTable{
		IfNotExists:   true,
		Replace:       true,
		Temporary:     false,
		FromTableName: "A",
		ToTableName:   "B",
	}}, result)

	result = Parse("CREATE OR REPLACE TEMPORARY TABLE IF NOT EXISTS A LIKE B;")
	assert.EqualValues(t, []SqlStatement{CopyCreateTable{
		IfNotExists:   true,
		Replace:       true,
		Temporary:     true,
		FromTableName: "A",
		ToTableName:   "B",
	}}, result)
}

func TestParser_CreateTableQueryCreateTable(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CREATE TEMPORARY TABLE a (a int,b int) as select * from u")
		result := mySqlParser.CreateTable().Accept(visitor)

		assert.EqualValues(t, QueryCreateTable{
			IfNotExists: false,
			Replace:     false,
			Temporary:   true,
			Table: FullId{
				Uid:   "a",
				DotId: "",
			},
			CreateDefinitions: []CreateDefinition{
				ColumnDeclaration{
					Column: "a",
					ColumnDefinition: ColumnDefinition{
						DataType:          "int",
						ColumnConstraints: nil,
					},
				},
				ColumnDeclaration{
					Column: "b",
					ColumnDefinition: ColumnDefinition{
						DataType:          "int",
						ColumnConstraints: nil,
					},
				},
			},
			SelectStatement: SimpleSelect{
				QuerySpecification: QuerySpecification{
					SelectSpecs: nil,
					SelectElements: SelectElements{
						All: true,
					},
					FromClause: &FromClause{TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "u"}},
					}},
				},
			},
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CREATE TEMPORARY TABLE a (a int,b int);")
		result := mySqlParser.CreateTable().Accept(visitor)

		assert.EqualValues(t, QueryCreateTable{
			IfNotExists: false,
			Replace:     false,
			Temporary:   true,
			Table: FullId{
				Uid:   "a",
				DotId: "",
			},
			CreateDefinitions: []CreateDefinition{
				ColumnDeclaration{
					Column: "a",
					ColumnDefinition: ColumnDefinition{
						DataType:          "int",
						ColumnConstraints: nil,
					},
				},
				ColumnDeclaration{
					Column: "b",
					ColumnDefinition: ColumnDefinition{
						DataType:          "int",
						ColumnConstraints: nil,
					},
				},
			},
		}, result)
	})
}

func TestParser_CreateTableColumnCreateTable(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CREATE TEMPORARY TABLE a (a int,b int);")
		result := mySqlParser.CreateTable().Accept(visitor)

		assert.EqualValues(t, ColumnCreateTable{
			IfNotExists: false,
			Replace:     false,
			Temporary:   true,
			Table: FullId{
				Uid:   "a",
				DotId: "",
			},
			CreateDefinitions: []CreateDefinition{
				ColumnDeclaration{
					Column: "a",
					ColumnDefinition: ColumnDefinition{
						DataType:          "int",
						ColumnConstraints: nil,
					},
				},
				ColumnDeclaration{
					Column: "b",
					ColumnDefinition: ColumnDefinition{
						DataType:          "int",
						ColumnConstraints: nil,
					},
				},
			},
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("CREATE OR REPLACE TEMPORARY TABLE IF NOT EXISTS a (a int,b int);")
		result := mySqlParser.CreateTable().Accept(visitor)

		assert.EqualValues(t, ColumnCreateTable{
			IfNotExists: true,
			Replace:     true,
			Temporary:   true,
			Table: FullId{
				Uid:   "a",
				DotId: "",
			},
			CreateDefinitions: []CreateDefinition{
				ColumnDeclaration{
					Column: "a",
					ColumnDefinition: ColumnDefinition{
						DataType:          "int",
						ColumnConstraints: nil,
					},
				},
				ColumnDeclaration{
					Column: "b",
					ColumnDefinition: ColumnDefinition{
						DataType:          "int",
						ColumnConstraints: nil,
					},
				},
			},
		}, result)
	})
}
