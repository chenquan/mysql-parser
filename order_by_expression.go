package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type OrderByExpression struct {
	Expression Expression
	Order      string
}

func (v *parseTreeVisitor) VisitOrderByExpression(ctx *parser.OrderByExpressionContext) interface{} {
	var order string
	orderCtx := ctx.GetOrder()
	if orderCtx != nil {
		order = orderCtx.GetText()
	}
	return OrderByExpression{
		Expression: ctx.Expression().Accept(v).(Expression),
		Order:      order,
	}
}
