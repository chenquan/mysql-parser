package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type FromClause struct {
	TableSources *TableSources
}

func (v *parseTreeVisitor) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	var tableSources *TableSources
	tableSourcesCtx := ctx.TableSources()
	if tableSourcesCtx != nil {
		sources := tableSourcesCtx.Accept(v).(TableSources)
		tableSources = &sources
	}

	return FromClause{TableSources: tableSources}
}
