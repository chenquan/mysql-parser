package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

func (v *parseTreeVisitor) VisitTableSources(ctx *parser.TableSourcesContext) interface{} {
	allTableSources := ctx.AllTableSource()
	tableSources := make([]TableSource, 0, len(allTableSources))
	for _, sourceContext := range allTableSources {
		tableSources = append(tableSources, sourceContext.Accept(v).(TableSource))
	}

	return tableSources
}
