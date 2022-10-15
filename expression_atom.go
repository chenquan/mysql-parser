package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	ExpressionAtom interface {
		IsExpressionAtom()
	}
	ConstantExpressionAtom struct {
		Constant Constant
	}
	FullColumnNameExpressionAtom struct {
		FullColumnName FullColumnName
	}
)

func (f FullColumnNameExpressionAtom) IsExpressionAtom() {
}

func (c ConstantExpressionAtom) IsExpressionAtom() {
}

func (v *parseTreeVisitor) VisitConstantExpressionAtom(ctx *parser.ConstantExpressionAtomContext) interface{} {
	return ConstantExpressionAtom{Constant: ctx.Constant().Accept(v).(Constant)}
}
