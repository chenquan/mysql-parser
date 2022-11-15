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
		Priority                  string
		Ignore                    bool
		TableName                 TableName
		Columns                   []string
		InsertStatementValue      InsertStatementValue
		DuplicatedUpdatedElements []UpdatedElement
	}

	InsertStatementSetValue struct {
		Priority                  string
		Ignore                    bool
		TableName                 TableName
		UpdatedElements           []UpdatedElement
		DuplicatedUpdatedElements []UpdatedElement
	}

	InsertStatement interface {
		DmlStatement
		isInsertStatement()
	}
)

func (i InsertStatementSetValue) isSqlStatement()     {}
func (i InsertStatementSetValue) isDmlStatement()     {}
func (i InsertStatementSetValue) isInsertStatement()  {}
func (i InsertStatementIntoValue) isInsertStatement() {}
func (i InsertStatementIntoValue) isSqlStatement()    {}
func (i InsertStatementIntoValue) isDmlStatement()    {}

func (v *parseTreeVisitor) VisitInsertStatement(ctx *parser.InsertStatementContext) interface{} {
	var priority string
	priorityCtx := ctx.GetPriority()
	if priorityCtx != nil {
		priority = priorityCtx.GetText()
	}
	tableName := ctx.TableName().Accept(v).(TableName)

	var duplicatedUpdatedElements []UpdatedElement
	if ctx.DUPLICATE() != nil && ctx.KEY() != nil {
		allDuplicatedUpdatedElement := ctx.GetDuplicatedElements()
		duplicatedUpdatedElements = make([]UpdatedElement, 0, len(allDuplicatedUpdatedElement)+1)
		duplicatedUpdatedElements = append(duplicatedUpdatedElements, ctx.GetDuplicatedFirst().Accept(v).(UpdatedElement))
		for _, updatedElementContext := range allDuplicatedUpdatedElement {
			duplicatedUpdatedElements = append(duplicatedUpdatedElements, updatedElementContext.Accept(v).(UpdatedElement))
		}

	}

	insertStatementValueContext := ctx.InsertStatementValue()
	if insertStatementValueContext != nil {
		columnsCtx := ctx.GetColumns()
		var columns []string
		if columnsCtx != nil {
			columns = columnsCtx.Accept(v).([]string)
		}

		return InsertStatementIntoValue{
			Priority:                  priority,
			Ignore:                    ctx.IGNORE() != nil,
			TableName:                 tableName,
			Columns:                   columns,
			InsertStatementValue:      insertStatementValueContext.Accept(v).(InsertStatementValue),
			DuplicatedUpdatedElements: duplicatedUpdatedElements,
		}
	}

	if ctx.SET() != nil {
		allUpdatedElement := ctx.GetSetElements()
		updatedElements := make([]UpdatedElement, 0, len(allUpdatedElement)+1)
		updatedElements = append(updatedElements, ctx.GetSetFirst().Accept(v).(UpdatedElement))

		for _, updatedElementContext := range allUpdatedElement {
			updatedElements = append(updatedElements, updatedElementContext.Accept(v).(UpdatedElement))
		}

		return InsertStatementSetValue{
			Priority:                  priority,
			Ignore:                    ctx.IGNORE() != nil,
			TableName:                 tableName,
			UpdatedElements:           updatedElements,
			DuplicatedUpdatedElements: duplicatedUpdatedElements,
		}
	}

	return nil
}
