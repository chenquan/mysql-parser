package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ IndexOption = (*IndexType)(nil)
)

type IndexType string

func (i IndexType) isIndexOption() {}

func (v *parseTreeVisitor) VisitIndexType(ctx *parser.IndexTypeContext) interface{} {
	return IndexType(ctx.GetChild(1).(interface{ GetText() string }).GetText())
}
