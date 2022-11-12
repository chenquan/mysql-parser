package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ InsertStatementValue = (*InsertStatementValueSelectStatement)(nil)
	_ InsertStatementValue = (*InsertStatementValuedExpressionsWithDefaults)(nil)
)

type (
	InsertStatementValue interface {
		IsInsertStatementValue()
	}
	InsertStatementValueSelectStatement struct {
		SelectStatement SelectStatement
	}
	InsertStatementValuedExpressionsWithDefaults struct {
		ExpressionsWithDefaultsList []ExpressionsWithDefaults
	}
)

func (i InsertStatementValuedExpressionsWithDefaults) IsInsertStatementValue() {
}

func (i InsertStatementValueSelectStatement) IsInsertStatementValue() {
}

func (v *parseTreeVisitor) VisitInsertStatementValue(ctx *parser.InsertStatementValueContext) interface{} {
	selectStatementContext := ctx.SelectStatement()
	if selectStatementContext != nil {
		return InsertStatementValueSelectStatement{SelectStatement: selectStatementContext.Accept(v).(SelectStatement)}
	}
	allExpressionsWithDefaults := ctx.AllExpressionsWithDefaults()
	expressionOrDefaultsList := make([]ExpressionsWithDefaults, 0, len(allExpressionsWithDefaults))
	for _, expressionsWithDefault := range allExpressionsWithDefaults {
		expressionOrDefaultsList = append(expressionOrDefaultsList, expressionsWithDefault.Accept(v).(ExpressionsWithDefaults))
	}

	return InsertStatementValuedExpressionsWithDefaults{ExpressionsWithDefaultsList: expressionOrDefaultsList}
}
