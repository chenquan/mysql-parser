package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	AlterTable struct {
		tableName      TableName
		AddColumns     []TableAddColumn
		DeleteColumn   []string
		AddIndexes     []TableAddIndex
		AddUniqueKeys  []TableAddUniqueKey
		AddPrimaryKeys []TableAddPrimaryKey
		DropColumns    []TableDropColumn
		ModifyColumns  []TableModifyColumn
		DropPrimaryKey bool
		DropIndexes    []TableDropIndex
		RenameIndexes  []TableRenameIndex
		Renames        []string
	}

	TableAddColumn struct {
		IfNotExists      bool
		Column           string
		ColumnDefinition ColumnDefinition
	}

	ColumnDefinition struct {
		DataType          string
		ColumnConstraints []ColumnConstraint
	}

	TableAddIndex struct {
		IfNotExists bool
		IndexName   string
		IndexType   IndexType
		Columns     []IndexColumnName
	}
	TableAddPrimaryKey struct {
		Index     string
		IndexType string
		Columns   []IndexColumnName
	}

	TableAddUniqueKey struct {
		IndexName string
		IndexType string
		Columns   []IndexColumnName
	}
	TableModifyColumn struct {
		IfExists         bool
		Column           string
		ColumnDefinition ColumnDefinition
	}
	TableDropColumn struct {
		IfExists bool
		Column   string
		Restrict bool
	}
	TableDropIndex struct {
		IfExists bool
		Column   string
	}
	TableRenameIndex struct {
		FromColumn string
		ToColumn   string
	}
)

func (a AlterTable) IsSqlStatement() {
}

func (a AlterTable) IsDdlStatement() {
}

func (v *parseTreeVisitor) VisitAlterTable(ctx *parser.AlterTableContext) interface{} {
	table := AlterTable{
		tableName: ctx.TableName().Accept(v).(TableName),
	}

	for _, alterSpecification := range ctx.AllAlterSpecification() {
		switch a := alterSpecification.(type) {
		case *parser.AlterByAddColumnContext:
			table.AddColumns = append(table.AddColumns, a.Accept(v).(TableAddColumn))
		case *parser.AlterByAddColumnsContext:
			table.AddColumns = append(table.AddColumns, a.Accept(v).([]TableAddColumn)...)
		case *parser.AlterByAddIndexContext:
			table.AddIndexes = append(table.AddIndexes, a.Accept(v).(TableAddIndex))
		case *parser.AlterByAddUniqueKeyContext:
			table.AddUniqueKeys = append(table.AddUniqueKeys, a.Accept(v).(TableAddUniqueKey))
		case *parser.AlterByAddPrimaryKeyContext:
			table.AddPrimaryKeys = append(table.AddPrimaryKeys, a.Accept(v).(TableAddPrimaryKey))
		case *parser.AlterByModifyColumnContext:
			table.ModifyColumns = append(table.ModifyColumns, a.Accept(v).(TableModifyColumn))
		case *parser.AlterByDropColumnContext:
			table.DropColumns = append(table.DropColumns, a.Accept(v).(TableDropColumn))
		case *parser.AlterByDropPrimaryKeyContext:
			table.DropPrimaryKey = true
		case *parser.AlterByDropIndexContext:
			table.DropIndexes = append(table.DropIndexes, a.Accept(v).(TableDropIndex))
		case *parser.AlterByRenameIndexContext:
			table.RenameIndexes = append(table.RenameIndexes, a.Accept(v).(TableRenameIndex))
		case *parser.AlterByRenameContext:
			table.Renames = append(table.Renames, a.Accept(v).(string))
		}
	}

	return table
}

func (v *parseTreeVisitor) VisitAlterByAddColumn(ctx *parser.AlterByAddColumnContext) interface{} {
	allUid := ctx.AllUid()
	if len(allUid) != 0 {
		acceptColumnDefinition := ctx.ColumnDefinition().Accept(v)
		return TableAddColumn{
			IfNotExists:      ctx.IfNotExists() != nil,
			Column:           allUid[0].GetText(),
			ColumnDefinition: acceptColumnDefinition.(ColumnDefinition),
		}
	}

	return nil
}

func (v *parseTreeVisitor) VisitAlterByAddColumns(ctx *parser.AlterByAddColumnsContext) interface{} {
	allUid := ctx.AllUid()
	definitions := ctx.AllColumnDefinition()
	if len(allUid) == len(definitions) {
		addColumns := make([]TableAddColumn, len(definitions))
		for i := 0; i < len(allUid); i++ {
			uid := allUid[i]
			definition := definitions[i]
			acceptColumnDefinition := definition.Accept(v)

			addColumns[i] = TableAddColumn{
				IfNotExists:      ctx.IfNotExists() != nil,
				Column:           uid.GetText(),
				ColumnDefinition: acceptColumnDefinition.(ColumnDefinition),
			}
		}
		return addColumns
	}

	return nil
}

func (v *parseTreeVisitor) VisitAlterByAddIndex(ctx *parser.AlterByAddIndexContext) interface{} {
	var indexName string
	uid := ctx.Uid()
	if uid != nil {
		indexName = uid.GetText()
	}

	var indexType IndexType
	indexTypeContext := ctx.IndexType()
	if indexTypeContext != nil {
		indexType = indexTypeContext.Accept(v).(IndexType)
	}

	var columns []IndexColumnName
	indexColumnNamesContext := ctx.IndexColumnNames()
	if indexColumnNamesContext != nil {
		indexColumnNamesValue := indexColumnNamesContext.Accept(v)
		if indexColumnNamesValue != nil {
			columns = indexColumnNamesValue.([]IndexColumnName)
		}
	}

	return TableAddIndex{
		IfNotExists: ctx.IfNotExists() != nil,
		IndexName:   indexName,
		IndexType:   indexType,
		Columns:     columns,
	}
}

func (v *parseTreeVisitor) VisitAlterByAddPrimaryKey(ctx *parser.AlterByAddPrimaryKeyContext) interface{} {
	var index string
	indexContext := ctx.GetIndex()
	if indexContext != nil {
		index = indexContext.GetText()
	}

	var indexType string
	indexTypeContext := ctx.IndexType()
	if indexTypeContext != nil {
		indexType = indexTypeContext.Accept(v).(string)
	}

	var columns []IndexColumnName
	indexColumnNamesContext := ctx.IndexColumnNames()
	if indexColumnNamesContext != nil {
		indexColumnNamesValue := indexColumnNamesContext.Accept(v)
		if indexColumnNamesValue != nil {
			columns = indexColumnNamesValue.([]IndexColumnName)
		}
	}
	allIndexOption := ctx.AllIndexOption()
	for _, optionContext := range allIndexOption {
		valContext := optionContext.Accept(v)
		switch val := valContext.(type) {
		case IndexType:
			if indexType == "" {
				indexType = string(val)
			}
		}
	}

	return TableAddPrimaryKey{
		Index:     index,
		IndexType: indexType,
		Columns:   columns,
	}
}

func (v *parseTreeVisitor) VisitAlterByAddUniqueKey(ctx *parser.AlterByAddUniqueKeyContext) interface{} {
	var indexName string
	indexNameContext := ctx.GetIndexName()
	if indexNameContext != nil {
		indexName = indexNameContext.GetText()
	}
	var indexType string
	indexTypeContext := ctx.IndexType()
	if indexTypeContext != nil {
		indexType = indexTypeContext.Accept(v).(string)
	}

	var indexColumnNames []IndexColumnName
	indexColumnNamesContext := ctx.IndexColumnNames()
	if indexColumnNamesContext != nil {
		indexColumnNames = indexColumnNamesContext.Accept(v).([]IndexColumnName)
	}

	allIndexOption := ctx.AllIndexOption()
	for _, optionContext := range allIndexOption {
		acceptVal := optionContext.Accept(v)
		if acceptVal != nil {
			switch val := acceptVal.(type) {
			case IndexType:
				if indexType == "" {
					indexType = string(val)
				}
			}
		}
	}

	return TableAddUniqueKey{
		IndexName: indexName,
		IndexType: indexType,
		Columns:   indexColumnNames,
	}
}

func (v *parseTreeVisitor) VisitAlterByModifyColumn(ctx *parser.AlterByModifyColumnContext) interface{} {
	return TableModifyColumn{
		IfExists:         ctx.IfExists() != nil,
		Column:           ctx.Uid(0).GetText(),
		ColumnDefinition: ctx.ColumnDefinition().Accept(v).(ColumnDefinition),
	}
}

func (v *parseTreeVisitor) VisitAlterByDropColumn(ctx *parser.AlterByDropColumnContext) interface{} {
	return TableDropColumn{
		IfExists: ctx.IfExists() != nil,
		Column:   ctx.Uid().GetText(),
		Restrict: ctx.RESTRICT() != nil,
	}
}

func (v *parseTreeVisitor) VisitAlterByDropIndex(ctx *parser.AlterByDropIndexContext) interface{} {
	return TableDropIndex{
		IfExists: ctx.IfExists() != nil,
		Column:   ctx.Uid().GetText(),
	}
}

func (v *parseTreeVisitor) VisitAlterByRenameIndex(ctx *parser.AlterByRenameIndexContext) interface{} {
	return TableRenameIndex{
		FromColumn: ctx.Uid(0).GetText(),
		ToColumn:   ctx.Uid(1).GetText(),
	}
}

func (v *parseTreeVisitor) VisitAlterByRename(ctx *parser.AlterByRenameContext) interface{} {
	return ctx.Uid().GetText()
}
