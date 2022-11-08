package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type CreateDatabase struct {
	IfNotExists           bool
	DatabaseName          string
	CreateDatabaseOptions []CreateDatabaseOption
}

func (v *parseTreeVisitor) VisitCreateDatabase(ctx *parser.CreateDatabaseContext) interface{} {
	allCreateDatabaseOption := ctx.AllCreateDatabaseOption()
	var createDatabaseOptions []CreateDatabaseOption
	if len(allCreateDatabaseOption) != 0 {
		createDatabaseOptions = make([]CreateDatabaseOption, 0, len(allCreateDatabaseOption))
	}

	for _, createDatabaseOptionContext := range allCreateDatabaseOption {
		createDatabaseOptions = append(createDatabaseOptions, createDatabaseOptionContext.Accept(v).(CreateDatabaseOption))
	}

	return CreateDatabase{
		IfNotExists:           ctx.IfNotExists() != nil,
		DatabaseName:          ctx.Uid().GetText(),
		CreateDatabaseOptions: createDatabaseOptions,
	}
}
