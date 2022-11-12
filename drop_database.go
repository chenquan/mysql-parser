package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type DropDatabase struct {
	IfExists     bool
	DatabaseName string
}

func (v *parseTreeVisitor) VisitDropDatabase(ctx *parser.DropDatabaseContext) interface{} {
	return DropDatabase{
		IfExists:     ctx.IfExists() != nil,
		DatabaseName: ctx.Uid().GetText(),
	}
}
