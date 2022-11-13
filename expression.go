package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ Expression = (*LogicalExpression)(nil)
	_ Expression = (*IsExpression)(nil)
	_ Expression = (PredicateExpression)(nil)
)

type (
	Expression interface {
		FunctionArg
		isExpression()
	}

	NotExpression Expression

	LogicalExpression struct {
		LeftExpression  Expression
		LogicalOperator string
		RightExpression Expression
	}
	IsExpression struct {
		predicateExpression PredicateExpression
		IsNot               bool
		TestValue           string
	}
	PredicateExpression Predicate
)

func (i IsExpression) isFunctionArg() {
}

func (l LogicalExpression) isFunctionArg() {
}

func (i IsExpression) isExpression() {
}

func (l LogicalExpression) isExpression() {
}

func (v *parseTreeVisitor) VisitExpressions(ctx *parser.ExpressionsContext) interface{} {
	allExpression := ctx.AllExpression()
	expressions := make([]Expression, 0, len(allExpression))
	for _, expressionContext := range allExpression {
		expressions = append(expressions, expressionContext.Accept(v).(Expression))
	}

	return expressions
}

func (v *parseTreeVisitor) VisitNotExpression(ctx *parser.NotExpressionContext) interface{} {
	return NotExpression(ctx.Expression().Accept(v).(Expression))
}

func (v *parseTreeVisitor) VisitLogicalExpression(ctx *parser.LogicalExpressionContext) interface{} {
	expressionContexts := ctx.AllExpression()
	return LogicalExpression{
		LeftExpression:  expressionContexts[0].Accept(v).(Expression),
		LogicalOperator: ctx.LogicalOperator().GetText(),
		RightExpression: expressionContexts[1].Accept(v).(Expression),
	}
}

func (v *parseTreeVisitor) VisitIsExpression(ctx *parser.IsExpressionContext) interface{} {
	return IsExpression{
		predicateExpression: ctx.Predicate().Accept(v).(Predicate),
		IsNot:               ctx.NOT() != nil,
		TestValue:           ctx.GetTestValue().GetText(),
	}
}

func (v *parseTreeVisitor) VisitPredicateExpression(ctx *parser.PredicateExpressionContext) interface{} {
	return ctx.Predicate().Accept(v)
}
