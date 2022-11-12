package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type UpdatedElement struct {
	FullColumnName FullColumnName
	Value          ExpressionOrDefault
}

func (v *parseTreeVisitor) VisitUpdatedElement(ctx *parser.UpdatedElementContext) interface{} {
	return UpdatedElement{
		FullColumnName: ctx.FullColumnName().Accept(v).(FullColumnName),
		Value:          ctx.ExpressionOrDefault().Accept(v).(ExpressionOrDefault),
	}
}
