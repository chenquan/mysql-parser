package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ UpdateStatement = (*SingleUpdateStatement)(nil)
	_ UpdateStatement = (*MultipleUpdateStatement)(nil)
)

type (
	UpdateStatement interface {
		isUpdateStatement()
	}
	SingleUpdateStatement struct {
		LowPriority     bool
		Ignore          bool
		TableName       TableName
		Alias           string
		UpdatedElements []UpdatedElement
		WhereExpression Expression
		OrderByClause   OrderByClause
		LimitClause     LimitClause
	}
	MultipleUpdateStatement struct {
		LowPriority     bool
		Ignore          bool
		TableSources    []TableSource
		UpdatedElements []UpdatedElement
		WhereExpression Expression
	}
)

func (m MultipleUpdateStatement) isUpdateStatement() {}
func (s SingleUpdateStatement) isUpdateStatement()   {}

func (v *parseTreeVisitor) VisitSingleUpdateStatement(ctx *parser.SingleUpdateStatementContext) interface{} {
	var alias string
	uidContext := ctx.Uid()
	if uidContext != nil {
		alias = uidContext.GetText()
	}

	allUpdatedElement := ctx.AllUpdatedElement()
	updatedElements := make([]UpdatedElement, 0, len(allUpdatedElement))
	for _, updatedElementContext := range allUpdatedElement {
		updatedElements = append(updatedElements, updatedElementContext.Accept(v).(UpdatedElement))
	}

	var (
		whereExpression Expression
		orderByClause   OrderByClause
		limitClause     LimitClause
	)
	expressionContext := ctx.Expression()
	if expressionContext != nil {
		whereExpression = expressionContext.Accept(v).(Expression)
	}

	orderByClauseContext := ctx.OrderByClause()
	if orderByClauseContext != nil {
		orderByClause = orderByClauseContext.Accept(v).(OrderByClause)
	}

	limitClauseContext := ctx.LimitClause()
	if limitClauseContext != nil {
		limitClause = limitClauseContext.Accept(v).(LimitClause)
	}

	return SingleUpdateStatement{
		LowPriority:     ctx.GetPriority() != nil,
		Ignore:          ctx.IGNORE() != nil,
		TableName:       ctx.TableName().Accept(v).(TableName),
		Alias:           alias,
		UpdatedElements: updatedElements,
		WhereExpression: whereExpression,
		OrderByClause:   orderByClause,
		LimitClause:     limitClause,
	}
}

func (v *parseTreeVisitor) VisitMultipleUpdateStatement(ctx *parser.MultipleUpdateStatementContext) interface{} {
	allUpdatedElement := ctx.AllUpdatedElement()
	updatedElements := make([]UpdatedElement, 0, len(allUpdatedElement))
	for _, updatedElementContext := range allUpdatedElement {
		updatedElements = append(updatedElements, updatedElementContext.Accept(v).(UpdatedElement))
	}

	var whereExpression Expression
	expressionContext := ctx.Expression()
	if expressionContext != nil {
		whereExpression = expressionContext.Accept(v).(Expression)
	}

	return MultipleUpdateStatement{
		LowPriority:     ctx.GetPriority() != nil,
		Ignore:          ctx.IGNORE() != nil,
		TableSources:    ctx.TableSources().Accept(v).([]TableSource),
		UpdatedElements: updatedElements,
		WhereExpression: whereExpression,
	}
}
