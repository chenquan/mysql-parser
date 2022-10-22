package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ ExpressionAtom = (*MathExpressionAtom)(nil)
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
	MathExpressionAtom struct {
		LeftExpressionAtom  ExpressionAtom
		MathOperator        string
		RightExpressionAtom ExpressionAtom
	}
	SubqueryExpressionAtom struct {
		SelectStatement SelectStatement
	}
	NestedExpressionAtom struct {
		Expressions []Expression
	}
	ExistsExpressionAtom struct {
		SelectStatement SelectStatement
	}
)

func (n NestedExpressionAtom) IsExpressionAtom() {
}

func (m MathExpressionAtom) IsExpressionAtom() {
}

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

func (v *parseTreeVisitor) VisitMathExpressionAtom(ctx *parser.MathExpressionAtomContext) interface{} {
	return MathExpressionAtom{
		LeftExpressionAtom:  ctx.GetLeft().Accept(v).(ExpressionAtom),
		MathOperator:        ctx.MathOperator().GetText(),
		RightExpressionAtom: ctx.GetRight().Accept(v).(ExpressionAtom),
	}
}

func (v *parseTreeVisitor) VisitSubqueryExpressionAtom(ctx *parser.SubqueryExpressionAtomContext) interface{} {
	return SubqueryExpressionAtom{
		SelectStatement: ctx.SelectStatement().Accept(v).(SelectStatement),
	}
}

func (v *parseTreeVisitor) VisitNestedExpressionAtom(ctx *parser.NestedExpressionAtomContext) interface{} {
	allExpressions := ctx.AllExpression()

	expressions := make([]Expression, 0, len(allExpressions))
	for _, expressionCtx := range allExpressions {
		expressions = append(expressions, expressionCtx.Accept(v).(Expression))
	}

	return NestedExpressionAtom{Expressions: expressions}
}

func (v *parseTreeVisitor) VisitExistsExpressionAtom(ctx *parser.ExistsExpressionAtomContext) interface{} {
	return ExistsExpressionAtom{SelectStatement: ctx.SelectStatement().Accept(v).(SelectStatement)}
}
