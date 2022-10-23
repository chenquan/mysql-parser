package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type GroupByClause struct {
	GroupByItems []GroupByItem
}

func (v *parseTreeVisitor) VisitGroupByClause(ctx *parser.GroupByClauseContext) interface{} {
	allGroupByItems := ctx.AllGroupByItem()
	groupByItems := make([]GroupByItem, 0, len(allGroupByItems))

	for _, groupByItemContext := range allGroupByItems {
		groupByItems = append(groupByItems, groupByItemContext.Accept(v).(GroupByItem))
	}

	return GroupByClause{GroupByItems: groupByItems}
}
