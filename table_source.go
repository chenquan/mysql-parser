package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var _ TableSource = (*TableSourceBase)(nil)

type (
	TableSource interface {
		isTableSource()
	}

	TableSourceBase struct {
		TableSourceItem TableSourceItem
	}
)

func (t TableSourceBase) isTableSource() {
}
func (v *parseTreeVisitor) VisitTableSourceBase(ctx *parser.TableSourceBaseContext) interface{} {
	return TableSourceBase{TableSourceItem: ctx.TableSourceItem().Accept(v).(TableSourceItem)}
}
