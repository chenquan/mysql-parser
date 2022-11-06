package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type CurrentTimestamp struct {
	Current   string
	Precision int
}

func (v *parseTreeVisitor) VisitCurrentTimestamp(ctx *parser.CurrentTimestampContext) interface{} {
	text := ctx.GetChild(0).(interface {
		GetText() string
	}).GetText()

	var integer int
	integerLiteralContext := ctx.IntegerLiteral()
	if integerLiteralContext != nil {
		integer = integerLiteralContext.Accept(v).(int)
	}

	return CurrentTimestamp{
		Current:   text,
		Precision: integer,
	}
}

func (v *parseTreeVisitor) VisitDecimalLiteral(ctx *parser.DecimalLiteralContext) interface{} {
	return toDecimal(ctx.GetText())
}
