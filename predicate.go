package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ Predicate = (*InPredicate)(nil)
	_ Predicate = (*IsNullPredicate)(nil)
	_ Predicate = (*BinaryComparisonPredicate)(nil)
	_ Predicate = (*ExpressionAtomPredicate)(nil)
)

type (
	Predicate interface {
		Expression
		isPredicate()
	}

	InPredicate struct {
		Predicate       Predicate
		NotIn           bool
		SelectStatement SelectStatement
		Expressions     []Expression
	}

	IsNullPredicate struct {
		Predicate   Predicate
		NullNotnull NullNotnull
	}
	BinaryComparisonPredicate struct {
		LeftPredicate      Predicate
		ComparisonOperator string
		RightPredicate     Predicate
	}

	ExpressionAtomPredicate struct {
		ExpressionAtom ExpressionAtom
	}
)

func (e ExpressionAtomPredicate) isFunctionArg()   {}
func (b BinaryComparisonPredicate) isFunctionArg() {}
func (i IsNullPredicate) isFunctionArg()           {}
func (i InPredicate) isFunctionArg()               {}
func (e ExpressionAtomPredicate) isExpression()    {}
func (b BinaryComparisonPredicate) isExpression()  {}
func (b BinaryComparisonPredicate) isPredicate()   {}
func (i IsNullPredicate) isExpression()            {}
func (i IsNullPredicate) isPredicate()             {}
func (i InPredicate) isExpression()                {}
func (e ExpressionAtomPredicate) isPredicate()     {}
func (i InPredicate) isPredicate()                 {}

func (v *parseTreeVisitor) VisitInPredicate(ctx *parser.InPredicateContext) interface{} {
	var selectStatement SelectStatement
	var expressions []Expression
	selectStatementContext := ctx.SelectStatement()
	if selectStatementContext != nil {
		selectStatement = selectStatementContext.Accept(v).(SelectStatement)
	} else {
		expressions = ctx.Expressions().Accept(v).([]Expression)
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

func (v *parseTreeVisitor) VisitBinaryComparisonPredicate(ctx *parser.BinaryComparisonPredicateContext) interface{} {
	allPredicate := ctx.AllPredicate()
	return BinaryComparisonPredicate{
		LeftPredicate:      toPredicate(allPredicate[0].Accept(v)),
		ComparisonOperator: ctx.ComparisonOperator().GetText(),
		RightPredicate:     toPredicate(allPredicate[1].Accept(v)),
	}
}

func toPredicate(v interface{}) Predicate {
	if v != nil {
		return v.(Predicate)
	}
	return nil
}
