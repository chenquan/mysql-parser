package parser

import (
	"encoding/json"
	"fmt"

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
	*Result
}

func (v *parseTreeVisitor) String() string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(data)
}

func (v *parseTreeVisitor) VisitRoot(ctx *parser.RootContext) interface{} {
	if ctx.SqlStatements() != nil {
		return ctx.SqlStatements().Accept(v)
	}

	return nil
}

func (v *parseTreeVisitor) VisitDropTable(ctx *parser.DropTableContext) interface{} {

	return nil
}

func (v *parseTreeVisitor) VisitPartitionDefinitions(ctx *parser.PartitionDefinitionsContext) interface{} {
	for _, c := range ctx.AllPartitionDefinition() {
		// TODO PartitionDefinitions
		fmt.Println(c.GetText())
	}

	return nil
}

func (v *parseTreeVisitor) VisitDmlStatement(ctx *parser.DmlStatementContext) interface{} {
	selectStatementContext := ctx.SelectStatement()
	if selectStatementContext != nil {
		return selectStatementContext.Accept(v).(SelectStatement)
	}

	return nil
}
