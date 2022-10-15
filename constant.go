package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type Constant struct {
	Constant string
}

func (v *parseTreeVisitor) VisitConstant(ctx *parser.ConstantContext) interface{} {
	return Constant{Constant: ctx.GetText()}
}
