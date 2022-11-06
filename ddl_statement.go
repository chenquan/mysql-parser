package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

func (v *parseTreeVisitor) VisitDdlStatement(ctx *parser.DdlStatementContext) interface{} {
	createTableContext := ctx.CreateTable()
	if createTableContext != nil {
		return createTableContext.Accept(v)
	}

	createDatabaseContext := ctx.CreateDatabase()
	if createDatabaseContext != nil {
		// TODO CreateDatabase
		return createDatabaseContext.Accept(v)
	}

	dropDatabaseContext := ctx.DropDatabase()
	if dropDatabaseContext != nil {
		return dropDatabaseContext.Accept(v)
	}

	alterDatabaseContext := ctx.AlterDatabase()
	if alterDatabaseContext != nil {
		switch alter := alterDatabaseContext.(type) {
		case *parser.AlterSimpleDatabaseContext, *parser.AlterUpgradeNameContext:
			// TODO AlterDatabase
			alter.Accept(v)
		}
		return nil
	}

	alterTableContext := ctx.AlterTable()
	if alterTableContext != nil {
		v.AlterTables = append(v.AlterTables, alterTableContext.Accept(v).(AlterTable))
		return nil
	}

	return nil
}
