package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	CreateIndex struct {
		Replace          bool
		InTimeAction     string
		IndexCategory    string
		IndexType        string
		IndexName        string
		TableName        TableName
		IndexColumnNames []IndexColumnName
	}
)

func (v *parseTreeVisitor) VisitCreateIndex(ctx *parser.CreateIndexContext) interface{} {
	var (
		inTimeAction, indexCategory, indexType string
	)

	inTimeActionCtx := ctx.GetIntimeAction()
	if inTimeActionCtx != nil {
		inTimeAction = inTimeActionCtx.GetText()
	}

	indexCategoryCtx := ctx.GetIndexCategory()
	if indexCategoryCtx != nil {
		indexCategory = indexCategoryCtx.GetText()
	}

	return CreateIndex{
		Replace:          ctx.REPLACE() != nil,
		InTimeAction:     inTimeAction,
		IndexCategory:    indexCategory,
		IndexName:        ctx.Uid().GetText(),
		IndexType:        indexType,
		TableName:        ctx.TableName().Accept(v).(TableName),
		IndexColumnNames: ctx.IndexColumnNames().Accept(v).([]IndexColumnName),
	}
}
