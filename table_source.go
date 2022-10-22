package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var _ TableSource = (*TableSourceBase)(nil)

type (
	TableSource interface {
		IsTableSource()
	}

	TableSourceBase struct {
		TableSourceItem TableSourceItem
	}
)

func (t TableSourceBase) IsTableSource() {
}
func (v *parseTreeVisitor) VisitTableSourceBase(ctx *parser.TableSourceBaseContext) interface{} {
	return TableSourceBase{TableSourceItem: ctx.TableSourceItem().Accept(v).(TableSourceItem)}
}
