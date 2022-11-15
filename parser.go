package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type Result struct {
	DropDatabaseNames []string `json:"dropDatabaseNames"`

	AlterTables      []AlterTable
	CreatTables      []CreatTable
	SelectStatements []SelectStatement
}

type parseTreeVisitor struct {
	parser.BaseMySqlParserVisitor
}

func (v *parseTreeVisitor) VisitRoot(ctx *parser.RootContext) interface{} {
	if ctx.SqlStatements() != nil {
		return ctx.SqlStatements().Accept(v)
	}

	return nil
}
