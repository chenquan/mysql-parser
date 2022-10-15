package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
)

type NullNotnull struct {
	Not  bool
	Null string // NULL_LITERAL | NULL_SPEC_LITERAL
}

func (v *parseTreeVisitor) VisitNullNotnull(ctx *parser.NullNotnullContext) interface{} {
	return NullNotnull{
		Not:  ctx.NOT() != nil,
		Null: ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText(),
	}
}
