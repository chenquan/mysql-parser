package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	CreateIndex struct {
		Replace          bool
		InTimeAction     string
		IndexCategory    string
		IndexType        IndexType
		IndexName        string
		TableName        TableName
		IndexColumnNames []IndexColumnName
		IndexOption      []IndexOption
		Algorithm        string
		Lock             string
	}
)

func (v *parseTreeVisitor) VisitCreateIndex(ctx *parser.CreateIndexContext) interface{} {
	var (
		inTimeAction, indexCategory, algorithm, lock string
		indexType                                    IndexType
	)

	inTimeActionCtx := ctx.GetIntimeAction()
	if inTimeActionCtx != nil {
		inTimeAction = inTimeActionCtx.GetText()
	}

	indexCategoryCtx := ctx.GetIndexCategory()
	if indexCategoryCtx != nil {
		indexCategory = indexCategoryCtx.GetText()
	}
	indexTypeContext := ctx.IndexType()
	if indexTypeContext != nil {
		indexType = indexTypeContext.Accept(v).(IndexType)
	}

	allIndexOption := ctx.AllIndexOption()
	var indexOptions []IndexOption
	if len(allIndexOption) != 0 {
		indexOptions = make([]IndexOption, 0, len(allIndexOption))
		for _, indexOptionContext := range allIndexOption {
			indexOptions = append(indexOptions, indexOptionContext.Accept(v).(IndexOption))
		}
	}

	algTypeCtx := ctx.GetAlgType()
	if algTypeCtx != nil {
		algorithm = algTypeCtx.GetText()
	}

	lockTypeCtx := ctx.GetLockType()
	if lockTypeCtx != nil {
		lock = lockTypeCtx.GetText()
	}

	return CreateIndex{
		Replace:          ctx.REPLACE() != nil,
		InTimeAction:     inTimeAction,
		IndexCategory:    indexCategory,
		IndexType:        indexType,
		IndexName:        ctx.Uid().GetText(),
		TableName:        ctx.TableName().Accept(v).(TableName),
		IndexColumnNames: ctx.IndexColumnNames().Accept(v).([]IndexColumnName),
		IndexOption:      indexOptions,
		Algorithm:        algorithm,
		Lock:             lock,
	}
}
