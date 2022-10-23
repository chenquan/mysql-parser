package parser

import (
	"strconv"

	"github.com/chenquan/mysql-parser/internal/parser"
)

type LimitClause struct {
	Offset int
	Limit  int
}

func (v *parseTreeVisitor) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	var offset int
	offsetCtx := ctx.GetOffset()
	if offsetCtx != nil {
		offset = toInter(offsetCtx.GetText())
	}

	limit := toInter(ctx.GetLimit().GetText())

	return LimitClause{
		Offset: offset,
		Limit:  limit,
	}
}

func toInter(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
