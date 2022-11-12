package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

func (v *parseTreeVisitor) VisitTables(ctx *parser.TablesContext) interface{} {
	allTableName := ctx.AllTableName()
	var tables []TableName
	if len(allTableName) != 0 {
		tables = make([]TableName, 0, len(allTableName))
		for _, tableNameContext := range allTableName {
			tables = append(tables, tableNameContext.Accept(v).(TableName))
		}
	}

	return tables
}
