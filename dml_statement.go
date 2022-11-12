package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

func (v *parseTreeVisitor) VisitDmlStatement(ctx *parser.DmlStatementContext) interface{} {
	selectStatementContext := ctx.SelectStatement()
	if selectStatementContext != nil {
		return selectStatementContext.Accept(v).(SelectStatement)
	}

	insertStatementContext := ctx.InsertStatement()
	if insertStatementContext != nil {
		return insertStatementContext.Accept(v).(InsertStatement)
	}

	return nil
}
