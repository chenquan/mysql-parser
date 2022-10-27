package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type QuerySpecificationNointo struct {
	SelectSpecs    []string
	SelectElements SelectElements
	FromClause     *FromClause
	GroupByClause  *GroupByClause
	HavingClause   *HavingClause
	OrderByClause  *OrderByClause
	LimitClause    *LimitClause
}

func (v *parseTreeVisitor) VisitQuerySpecificationNointo(ctx *parser.QuerySpecificationNointoContext) interface{} {
	selectSpecContexts := ctx.AllSelectSpec()
	var selectSpecs []string

	if len(selectSpecContexts) != 0 {
		selectSpecs = make([]string, len(selectSpecContexts))
	}

	for i, selectSpecContext := range selectSpecContexts {
		selectSpecs[i] = selectSpecContext.GetText()
	}

	var fromClause *FromClause
	clauseCtx := ctx.FromClause()
	if clauseCtx != nil {
		clause := clauseCtx.Accept(v).(FromClause)
		fromClause = &clause
	}

	var groupByClause *GroupByClause
	groupByClauseContext := ctx.GroupByClause()
	if groupByClauseContext != nil {
		clause := groupByClauseContext.Accept(v).(GroupByClause)
		groupByClause = &clause
	}

	var havingClause *HavingClause
	havingClauseContext := ctx.HavingClause()
	if havingClauseContext != nil {
		clause := havingClauseContext.Accept(v).(HavingClause)
		havingClause = &clause
	}

	var orderByClause *OrderByClause
	orderByClauseContext := ctx.OrderByClause()
	if orderByClauseContext != nil {
		clause := orderByClauseContext.Accept(v).(OrderByClause)
		orderByClause = &clause
	}

	var limitClause *LimitClause
	limitClauseContext := ctx.LimitClause()
	if limitClauseContext != nil {
		clause := limitClauseContext.Accept(v).(LimitClause)
		limitClause = &clause
	}

	return QuerySpecificationNointo{
		SelectSpecs:    selectSpecs,
		SelectElements: ctx.SelectElements().Accept(v).(SelectElements),
		FromClause:     fromClause,
		GroupByClause:  groupByClause,
		HavingClause:   havingClause,
		OrderByClause:  orderByClause,
		LimitClause:    limitClause,
	}
}
