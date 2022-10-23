package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type OrderByClause struct {
	OrderByExpressions []OrderByExpression
}

func (v *parseTreeVisitor) VisitOrderByClause(ctx *parser.OrderByClauseContext) interface{} {
	allOrderByExpressions := ctx.AllOrderByExpression()
	orderByExpressions := make([]OrderByExpression, 0, len(allOrderByExpressions))

	for _, expressionContext := range allOrderByExpressions {
		orderByExpressions = append(orderByExpressions, expressionContext.Accept(v).(OrderByExpression))
	}

	return OrderByClause{OrderByExpressions: orderByExpressions}
}
