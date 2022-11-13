package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ AlterDatabase = (*AlterSimpleDatabase)(nil)
	_ AlterDatabase = (*AlterUpgradeName)(nil)
)

type (
	AlterDatabase interface {
		isAlterDatabase()
	}
	AlterSimpleDatabase struct {
		DatabaseName          string
		CreateDatabaseOptions []CreateDatabaseOption
	}
	AlterUpgradeName struct {
		DatabaseName string
	}
)

func (a AlterUpgradeName) isAlterDatabase() {
}

func (a AlterSimpleDatabase) isAlterDatabase() {
}

func (v *parseTreeVisitor) VisitAlterSimpleDatabase(ctx *parser.AlterSimpleDatabaseContext) interface{} {
	var databaseName string
	uidContext := ctx.Uid()
	if uidContext != nil {
		databaseName = uidContext.GetText()
	}

	allDatabaseOption := ctx.AllCreateDatabaseOption()
	databaseOptions := make([]CreateDatabaseOption, 0, len(allDatabaseOption))
	for _, databaseOptionContext := range allDatabaseOption {
		databaseOptions = append(databaseOptions, databaseOptionContext.Accept(v).(CreateDatabaseOption))
	}

	return AlterSimpleDatabase{
		DatabaseName:          databaseName,
		CreateDatabaseOptions: databaseOptions,
	}
}

func (v *parseTreeVisitor) VisitAlterUpgradeName(ctx *parser.AlterUpgradeNameContext) interface{} {
	return AlterUpgradeName{DatabaseName: ctx.Uid().GetText()}
}
