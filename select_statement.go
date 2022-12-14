package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ SelectStatement = (*SimpleSelect)(nil)
)

type (
	SelectStatement interface {
		DmlStatement
		isSelectStatement()
	}
	SimpleSelect struct {
		QuerySpecification QuerySpecification
	}
	ParenthesisSelect struct {
		QueryExpression QueryExpression
	}
)

func (s SimpleSelect) isSqlStatement()    {}
func (s SimpleSelect) isDmlStatement()    {}
func (s SimpleSelect) isSelectStatement() {}

func (v *parseTreeVisitor) VisitSimpleSelect(ctx *parser.SimpleSelectContext) interface{} {
	return SimpleSelect{QuerySpecification: ctx.QuerySpecification().Accept(v).(QuerySpecification)}
}

func (v *parseTreeVisitor) VisitParenthesisSelect(ctx *parser.ParenthesisSelectContext) interface{} {
	return ParenthesisSelect{QueryExpression: ctx.QueryExpression().Accept(v).(QueryExpression)}
}
