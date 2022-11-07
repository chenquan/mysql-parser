package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ TableSourceItem = (*AtomTableItem)(nil)
)

type TableSourceItem interface {
	IsTableSourceItem()
}
type (
	AtomTableItem struct {
		TableName string
		Alias     string
	}
	SubqueryTableItem struct {
		SelectStatement SelectStatement
	}
	TableSourcesItem struct {
		TableSources []TableSource
	}
)

func (a AtomTableItem) IsTableSourceItem() {
}

func (v *parseTreeVisitor) VisitAtomTableItem(ctx *parser.AtomTableItemContext) interface{} {
	var alias string
	aliasCtx := ctx.GetAlias()
	if aliasCtx != nil {
		alias = aliasCtx.GetText()
	}

	return AtomTableItem{
		TableName: ctx.TableName().GetText(),
		Alias:     alias,
	}
}

func (v *parseTreeVisitor) VisitSubqueryTableItem(ctx *parser.SubqueryTableItemContext) interface{} {
	return SubqueryTableItem{SelectStatement: ctx.SelectStatement().Accept(v).(SelectStatement)}
}

func (v *parseTreeVisitor) VisitTableSourcesItem(ctx *parser.TableSourcesItemContext) interface{} {
	return TableSourcesItem{TableSources: ctx.TableSources().Accept(v).([]TableSource)}
}
