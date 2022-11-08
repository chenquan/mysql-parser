package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type IndexColumnName struct {
	IndexColumnName   string
	IndexColumnLength int
	SortType          string
}

func (v *parseTreeVisitor) VisitIndexColumnName(ctx *parser.IndexColumnNameContext) interface{} {
	var indexColumnName string
	uidContext := ctx.Uid()
	if uidContext != nil {
		indexColumnName = uidContext.GetText()
	} else {
		indexColumnName = ctx.STRING_LITERAL().GetText()
	}

	var sortType string
	sortTypeCtx := ctx.GetSortType()
	if sortTypeCtx != nil {
		sortType = sortTypeCtx.GetText()
	}

	var indexColumnLength int
	integerLiteralContext := ctx.IntegerLiteral()
	if integerLiteralContext != nil {
		indexColumnLength = integerLiteralContext.Accept(v).(int)
	}

	return IndexColumnName{
		IndexColumnName:   indexColumnName,
		SortType:          sortType,
		IndexColumnLength: indexColumnLength,
	}
}

func (v *parseTreeVisitor) VisitIndexColumnNames(ctx *parser.IndexColumnNamesContext) interface{} {
	allIndexColumnName := ctx.AllIndexColumnName()
	var indexColumnNames []IndexColumnName
	for _, indexColumnNameContext := range allIndexColumnName {
		indexColumnNames = append(indexColumnNames, indexColumnNameContext.Accept(v).(IndexColumnName))
	}

	return indexColumnNames
}
