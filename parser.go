package parser

import (
	"encoding/json"
	"fmt"

	"github.com/chenquan/mysql-parser/internal/parser"
)

type Result struct {
	CreateDatabaseNames []string `json:"createDatabaseNames"`
	AlterDatabaseNames  []string `json:"alterDatabaseNames"`
	DropDatabaseNames   []string `json:"dropDatabaseNames"`

	AlterTables []AlterTable
	CreatTables []CreatTable
}

type parseTreeVisitor struct {
	parser.BaseMySqlParserVisitor
	*Result
}

func (v *parseTreeVisitor) String() string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(data)
}

func (v *parseTreeVisitor) VisitRoot(ctx *parser.RootContext) interface{} {
	if ctx.SqlStatements() != nil {
		return ctx.SqlStatements().Accept(v)
	}

	return nil
}

// VisitSqlStatements visits a parse tree produced by MySqlParser#sqlStatements.
func (v *parseTreeVisitor) VisitSqlStatements(ctx *parser.SqlStatementsContext) interface{} {
	for _, e := range ctx.AllSqlStatement() {
		_ = e.Accept(v)
	}

	return nil
}

// VisitSqlStatement visits a parse tree produced by MySqlParser#sqlStatement.
func (v *parseTreeVisitor) VisitSqlStatement(ctx *parser.SqlStatementContext) interface{} {
	if ctx.DdlStatement() != nil {
		return ctx.DdlStatement().Accept(v)
	}

	return nil
}

func (v *parseTreeVisitor) VisitCreateDatabase(ctx *parser.CreateDatabaseContext) interface{} {
	//v.CreateDatabaseNames = append(v.CreateDatabaseNames, ctx.Uid().GetText())

	return nil
}

//func (v *parseTreeVisitor) VisitColumnCreateTable(ctx *parser.ColumnCreateTableContext) interface{} {
//	TableName := ctx.TableName().GetText()
//	TableName = strings.Trim(TableName, "`")
//	TableName = strings.Trim(TableName, "'")
//	replacer := strings.NewReplacer("\r", "", "\n", "")
//	TableName = replacer.Replace(TableName)
//	v.CreateTableNames = append(v.CreateTableNames, TableName)
//	return nil
//}

func (v *parseTreeVisitor) VisitDropDatabase(ctx *parser.DropDatabaseContext) interface{} {
	v.DropDatabaseNames = append(v.DropDatabaseNames, ctx.Uid().GetText())
	return nil
}

func (v *parseTreeVisitor) VisitAlterSimpleDatabase(ctx *parser.AlterSimpleDatabaseContext) interface{} {
	v.AlterDatabaseNames = append(v.AlterDatabaseNames, ctx.Uid().GetText())
	return nil
}

func (v *parseTreeVisitor) VisitAlterUpgradeName(ctx *parser.AlterUpgradeNameContext) interface{} {
	v.AlterDatabaseNames = append(v.AlterDatabaseNames, ctx.Uid().GetText())
	return nil
}

func (v *parseTreeVisitor) VisitDdlStatement(ctx *parser.DdlStatementContext) interface{} {
	if ctx.CreateTable() != nil {

		switch create := ctx.CreateTable().(type) {
		case *parser.CopyCreateTableContext:
			createTable := create.Accept(v).(CopyCreateTable)
			v.CreatTables = append(v.CreatTables, createTable)
		case *parser.QueryCreateTableContext:
			createTable := create.Accept(v).(QueryCreateTable)
			v.CreatTables = append(v.CreatTables, createTable)
		case *parser.ColumnCreateTableContext:
			createTable := create.Accept(v).(ColumnCreateTable)
			v.CreatTables = append(v.CreatTables, createTable)
		}

		return nil
	}

	if ctx.CreateDatabase() != nil {
		ctx.CreateDatabase().Accept(v)
		return nil
	}

	if ctx.DropDatabase() != nil {
		ctx.DropDatabase().Accept(v)
		return nil
	}

	if ctx.AlterDatabase() != nil {
		switch alter := ctx.AlterDatabase().(type) {
		case *parser.AlterSimpleDatabaseContext:
			alter.Accept(v)
		case *parser.AlterUpgradeNameContext:
			alter.Accept(v)
		}
		return nil
	}

	if ctx.AlterTable() != nil {
		v.AlterTables = append(v.AlterTables, ctx.AlterTable().Accept(v).(AlterTable))
		return nil
	}

	return nil
}

func (v *parseTreeVisitor) VisitDropTable(ctx *parser.DropTableContext) interface{} {

	return nil
}

func (v *parseTreeVisitor) VisitPartitionDefinitions(ctx *parser.PartitionDefinitionsContext) interface{} {
	for _, c := range ctx.AllPartitionDefinition() {
		fmt.Println(c.GetText())
	}

	return nil
}
