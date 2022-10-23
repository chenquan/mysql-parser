package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type GroupByItem struct {
	Expression Expression
	Order      string
}

func (v *parseTreeVisitor) VisitGroupByItem(ctx *parser.GroupByItemContext) interface{} {
	var order string
	orderCtx := ctx.GetOrder()
	if orderCtx != nil {
		order = orderCtx.GetText()
	}
	return GroupByItem{
		Expression: ctx.Expression().Accept(v).(Expression),
		Order:      order,
	}
}
