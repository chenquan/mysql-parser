package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type TableName = FullId

func (v *parseTreeVisitor) VisitTableName(ctx *parser.TableNameContext) interface{} {
	return ctx.FullId().Accept(v).(FullId)
}
