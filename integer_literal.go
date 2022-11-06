package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

func (v *parseTreeVisitor) VisitIntegerLiteral(ctx *parser.IntegerLiteralContext) interface{} {
	return toInter(ctx.GetText())
}
