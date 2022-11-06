package parser

import (
	"encoding/json"
	"fmt"

	"github.com/chenquan/mysql-parser/internal/parser"
)

type Result struct {
	DropDatabaseNames []string `json:"dropDatabaseNames"`

	AlterTables      []AlterTable
	CreatTables      []CreatTable
	SelectStatements []SelectStatement
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
	ddlStatement := ctx.DdlStatement()
	if ddlStatement != nil {
		return ddlStatement.Accept(v)
	}

	dmlStatement := ctx.DmlStatement()
	if dmlStatement != nil {
		return dmlStatement.Accept(v)
	}

	return nil
}

func (v *parseTreeVisitor) VisitCreateDatabase(ctx *parser.CreateDatabaseContext) interface{} {
	// TODO CreateDatabase

	return nil
}

func (v *parseTreeVisitor) VisitDropDatabase(ctx *parser.DropDatabaseContext) interface{} {
	return ctx.Uid().GetText()
}

func (v *parseTreeVisitor) VisitAlterSimpleDatabase(ctx *parser.AlterSimpleDatabaseContext) interface{} {
	// TODO AlterSimpleDatabase
	return nil
}

func (v *parseTreeVisitor) VisitAlterUpgradeName(ctx *parser.AlterUpgradeNameContext) interface{} {
	// TODO AlterUpgradeName
	return nil
}

func (v *parseTreeVisitor) VisitDropTable(ctx *parser.DropTableContext) interface{} {

	return nil
}

func (v *parseTreeVisitor) VisitPartitionDefinitions(ctx *parser.PartitionDefinitionsContext) interface{} {
	for _, c := range ctx.AllPartitionDefinition() {
		// TODO PartitionDefinitions
		fmt.Println(c.GetText())
	}

	return nil
}

func (v *parseTreeVisitor) VisitDmlStatement(ctx *parser.DmlStatementContext) interface{} {
	selectStatementContext := ctx.SelectStatement()
	if selectStatementContext != nil {
		v.SelectStatements = append(v.SelectStatements, selectStatementContext.Accept(v).(SelectStatement))
		return nil
	}

	return nil
}
