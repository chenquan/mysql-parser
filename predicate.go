package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	Predicate interface {
		IsPredicate()
	}
	InPredicate struct {
		Predicate       Predicate
		NotIn           bool
		SelectStatement SelectStatement
		Expressions     Expressions
	}
	IsNullPredicate struct {
		Predicate   Predicate
		NullNotnull NullNotnull
	}
)

func (i InPredicate) IsPredicate() {
}

func (v *parseTreeVisitor) VisitInPredicate(ctx *parser.InPredicateContext) interface{} {
	var selectStatement SelectStatement
	var expressions Expressions
	selectStatementContext := ctx.SelectStatement()
	if selectStatementContext != nil {
		selectStatement = selectStatementContext.Accept(v).(SelectStatement)
	} else {
		expressions = ctx.Expressions().Accept(v).(Expressions)
	}

	return InPredicate{
		Predicate:       ctx.Predicate().Accept(v).(Predicate),
		NotIn:           ctx.NOT() != nil,
		SelectStatement: selectStatement,
		Expressions:     expressions,
	}
}

func (v *parseTreeVisitor) VisitIsNullPredicate(ctx *parser.IsNullPredicateContext) interface{} {
	return IsNullPredicate{
		Predicate:   ctx.Predicate().Accept(v).(Predicate),
		NullNotnull: ctx.NullNotnull().Accept(v).(NullNotnull),
	}
}
