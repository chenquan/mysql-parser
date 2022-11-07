package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type FromClause struct {
	TableSources []TableSource
}

func (v *parseTreeVisitor) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	tableSourcesCtx := ctx.TableSources()
	if tableSourcesCtx != nil {
		return FromClause{TableSources: tableSourcesCtx.Accept(v).([]TableSource)}
	}

	return nil
}
