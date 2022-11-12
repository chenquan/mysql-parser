package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var _ SqlStatement = (*DropTable)(nil)

type DropTable struct {
	IfExists   bool
	Temporary  bool
	TableNames []TableName
	DropType   string
}

func (d DropTable) IsSqlStatement() {
}

func (v *parseTreeVisitor) VisitDropTable(ctx *parser.DropTableContext) interface{} {
	var dropType string
	dropTypeCtx := ctx.GetDropType()
	if dropTypeCtx != nil {
		dropType = dropTypeCtx.GetText()
	}

	return DropTable{
		IfExists:   ctx.IfExists() != nil,
		Temporary:  ctx.TEMPORARY() != nil,
		TableNames: ctx.Tables().Accept(v).([]TableName),
		DropType:   dropType,
	}
}
