package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	SelectStatement interface {
		IsSelectStatement()
	}
	SimpleSelect struct {
		QuerySpecification QuerySpecification
	}
)

func (s SimpleSelect) IsSelectStatement() {
}

func (v *parseTreeVisitor) VisitSimpleSelect(ctx *parser.SimpleSelectContext) interface{} {
	return SimpleSelect{QuerySpecification: ctx.QuerySpecification().Accept(v).(QuerySpecification)}
}
