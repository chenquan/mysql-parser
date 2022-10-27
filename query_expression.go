package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type QueryExpression interface {
	IsQuerySpecification()
}

func (v *parseTreeVisitor) VisitQueryExpression(ctx *parser.QueryExpressionContext) interface{} {
	querySpecificationContext := ctx.QuerySpecification()
	if querySpecificationContext != nil {
		return querySpecificationContext.Accept(v).(QueryExpression)
	}

	return ctx.QueryExpression().Accept(v).(QueryExpression)
}
