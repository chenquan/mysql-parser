package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	SqlStatement interface {
		isSqlStatement()
	}
	DdlStatement interface {
		SqlStatement
		isDdlStatement()
	}
	DmlStatement interface {
		SqlStatement
		isDmlStatement()
	}
)

func (v *parseTreeVisitor) VisitSqlStatements(ctx *parser.SqlStatementsContext) interface{} {
	allSqlStatement := ctx.AllSqlStatement()
	sqlStatements := make([]SqlStatement, 0, len(allSqlStatement))
	for _, e := range allSqlStatement {
		sqlStatements = append(sqlStatements, e.Accept(v).(SqlStatement))
	}

	return sqlStatements
}

func (v *parseTreeVisitor) VisitSqlStatement(ctx *parser.SqlStatementContext) interface{} {
	ddlStatement := ctx.DdlStatement()
	if ddlStatement != nil {
		return ddlStatement.Accept(v)
	}

	dmlStatement := ctx.DmlStatement()
	if dmlStatement != nil {
		return dmlStatement.Accept(v)
	}

	return nil
}
