package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ InsertStatement = (*InsertStatementIntoValue)(nil)
	_ InsertStatement = (*InsertStatementSetValue)(nil)
)

type (
	InsertStatementIntoValue struct {
		Priority             string
		Ignore               bool
		TableName            TableName
		Columns              []string
		InsertStatementValue InsertStatementValue
	}

	InsertStatementSetValue struct {
		Priority        string
		Ignore          bool
		TableName       TableName
		UpdatedElements []UpdatedElement
	}

	InsertStatement interface {
		DmlStatement
		isInsertStatement()
	}
)

func (i InsertStatementSetValue) isSqlStatement() {
}

func (i InsertStatementSetValue) isDmlStatement() {
}

func (i InsertStatementSetValue) isInsertStatement() {
}

func (i InsertStatementIntoValue) isInsertStatement() {
}

func (i InsertStatementIntoValue) isSqlStatement() {
}

func (i InsertStatementIntoValue) isDmlStatement() {
}

func (v *parseTreeVisitor) VisitInsertStatement(ctx *parser.InsertStatementContext) interface{} {
	var priority string
	priorityCtx := ctx.GetPriority()
	if priorityCtx != nil {
		priority = priorityCtx.GetText()
	}
	tableName := ctx.TableName().Accept(v).(TableName)

	insertStatementValueContext := ctx.InsertStatementValue()
	if insertStatementValueContext != nil {
		columnsCtx := ctx.GetColumns()
		var columns []string
		if columnsCtx != nil {
			columns = columnsCtx.Accept(v).([]string)
		}

		return InsertStatementIntoValue{
			Priority:             priority,
			Ignore:               ctx.IGNORE() != nil,
			TableName:            tableName,
			Columns:              columns,
			InsertStatementValue: insertStatementValueContext.Accept(v).(InsertStatementValue),
		}
	}

	if ctx.SET() != nil {
		allUpdatedElement := ctx.AllUpdatedElement()
		updatedElements := make([]UpdatedElement, 0, len(allUpdatedElement))
		for _, updatedElementContext := range allUpdatedElement {
			updatedElements = append(updatedElements, updatedElementContext.Accept(v).(UpdatedElement))
		}

		return InsertStatementSetValue{
			Priority:        priority,
			Ignore:          ctx.IGNORE() != nil,
			TableName:       tableName,
			UpdatedElements: updatedElements,
		}
	}

	return nil
}
