package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type ExpressionOrDefault struct {
	Default    bool
	Expression Expression
}

func (v *parseTreeVisitor) VisitExpressionOrDefault(ctx *parser.ExpressionOrDefaultContext) interface{} {
	if ctx.DEFAULT() != nil {
		return ExpressionOrDefault{Default: true}
	}

	return ExpressionOrDefault{Expression: ctx.Expression().Accept(v).(Expression)}
}
