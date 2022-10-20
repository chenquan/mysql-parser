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
	FunctionCallExpressionAtom struct {
		FunctionCall FunctionCall
	}
)

func (f FullColumnNameExpressionAtom) IsExpressionAtom() {
}

func (c ConstantExpressionAtom) IsExpressionAtom() {
}

func (v *parseTreeVisitor) VisitConstantExpressionAtom(ctx *parser.ConstantExpressionAtomContext) interface{} {
	return ConstantExpressionAtom{Constant: ctx.Constant().Accept(v).(Constant)}
}

func (v *parseTreeVisitor) VisitFullColumnNameExpressionAtom(ctx *parser.FullColumnNameExpressionAtomContext) interface{} {
	return FullColumnNameExpressionAtom{FullColumnName: ctx.FullColumnName().Accept(v).(FullColumnName)}
}

func (v *parseTreeVisitor) VisitFunctionCallExpressionAtom(ctx *parser.FunctionCallExpressionAtomContext) interface{} {
	return FunctionCallExpressionAtom{
		FunctionCall: ctx.FunctionCall().Accept(v).(FunctionCall),
	}
}

func (v *parseTreeVisitor) VisitExpressionAtomPredicate(ctx *parser.ExpressionAtomPredicateContext) interface{} {
	return ExpressionAtomPredicate{ExpressionAtom: ctx.ExpressionAtom().Accept(v).(ExpressionAtom)}
}
