package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	QuerySpecification struct {
		SelectSpecs          []string
		SelectElements       SelectElements
		FromClause           *FromClause
		SelectIntoExpression SelectIntoExpression
		GroupByClause        *GroupByClause
		HavingClause         *HavingClause
	}
)

func (v *parseTreeVisitor) VisitQuerySpecification(ctx *parser.QuerySpecificationContext) interface{} {
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

	var selectIntoExpression SelectIntoExpression
	selectIntoExpressionContext := ctx.SelectIntoExpression()
	if selectIntoExpressionContext != nil {
		selectIntoExpression = selectIntoExpressionContext.Accept(v).(SelectIntoExpression)
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

	return QuerySpecification{
		SelectSpecs:          selectSpecs,
		SelectElements:       ctx.SelectElements().Accept(v).(SelectElements),
		FromClause:           fromClause,
		SelectIntoExpression: selectIntoExpression,
		GroupByClause:        groupByClause,
		HavingClause:         havingClause,
	}
}
