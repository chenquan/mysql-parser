package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

func (v *parseTreeVisitor) VisitExpressionsWithDefaults(ctx *parser.ExpressionsWithDefaultsContext) interface{} {
	allExpressionOrDefault := ctx.AllExpressionOrDefault()
	expressionOrDefaults := make([]ExpressionOrDefault, 0, len(allExpressionOrDefault))
	for _, expressionOrDefaultContext := range allExpressionOrDefault {
		expressionOrDefaults = append(expressionOrDefaults, expressionOrDefaultContext.Accept(v).(ExpressionOrDefault))
	}

	return expressionOrDefaults
}
