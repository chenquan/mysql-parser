package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type HavingClause struct {
	Expression Expression
}

func (v *parseTreeVisitor) VisitHavingClause(ctx *parser.HavingClauseContext) interface{} {
	return HavingClause{
		Expression: ctx.Expression().Accept(v).(Expression),
	}
}
