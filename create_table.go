package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ CreatTable = (*CopyCreateTable)(nil)
	_ CreatTable = (*QueryCreateTable)(nil)
	_ CreatTable = (*ColumnCreateTable)(nil)
)

type (
	CreatTable interface {
		IsCreatTable()
	}

	CopyCreateTable struct {
		IfNotExists   bool
		Replace       bool
		Temporary     bool
		FromTableName string
		ToTableName   string
	}
	QueryCreateTable struct {
		IfNotExists       bool
		Replace           bool
		Temporary         bool
		Table             string
		CreateDefinitions []CreateDefinition
	}
	ColumnCreateTable struct {
		IfNotExists       bool
		Replace           bool
		Temporary         bool
		Table             string
		CreateDefinitions []CreateDefinition
	}
	CreateDefinition interface {
		IsCreateDefinition()
	}
	ColumnDeclaration struct {
		Column           string
		ColumnDefinition ColumnDefinition
	}
	IndexDeclaration interface {
		IsIndexDeclaration()
		CreateDefinition
	}

	PrimaryKeyTableConstraint struct {
		Constraint bool
		Name       string
		Index      string
	}
)

func (c CopyCreateTable) IsSqlStatement() {
}

func (p PrimaryKeyTableConstraint) IsColumnConstraint() {
}

func (p PrimaryKeyTableConstraint) IsCreateDefinition() {
}

func (c ColumnDeclaration) IsCreateDefinition() {}

func (c ColumnCreateTable) IsCreatTable() {
}

func (q QueryCreateTable) IsCreatTable() {
}

func (c CopyCreateTable) IsCreatTable() {
}

func (v *parseTreeVisitor) VisitCopyCreateTable(ctx *parser.CopyCreateTableContext) interface{} {
	return CopyCreateTable{
		IfNotExists:   ctx.IfNotExists() != nil,
		Replace:       ctx.REPLACE() != nil,
		Temporary:     ctx.TEMPORARY() != nil,
		FromTableName: ctx.TableName(0).GetText(),
		ToTableName:   ctx.TableName(1).GetText(),
	}
}

func (v *parseTreeVisitor) VisitQueryCreateTable(ctx *parser.QueryCreateTableContext) interface{} {
	//ctx.OR() != nil && ctx.AllREPLACE()

	var createDefinitions []CreateDefinition
	createDefinitionsContext := ctx.CreateDefinitions()
	if createDefinitionsContext != nil {
		createDefinitions = createDefinitionsContext.Accept(v).([]CreateDefinition)
	}
	selectStatementContext := ctx.SelectStatement()
	switch selectStatement := selectStatementContext.(type) {
	case *parser.SimpleSelectContext:
		selectStatement.Accept(v)
	}

	return QueryCreateTable{
		IfNotExists: ctx.IfNotExists() != nil,
		//Replace:     ctx.REPLACE() != nil,
		Temporary:         ctx.TEMPORARY() != nil,
		CreateDefinitions: createDefinitions,
	}
}

func (v *parseTreeVisitor) VisitCreateDefinitions(ctx *parser.CreateDefinitionsContext) interface{} {
	allCreateDefinition := ctx.AllCreateDefinition()
	createDefinitions := make([]CreateDefinition, 0, len(allCreateDefinition))
	for _, createDefinitionContext := range allCreateDefinition {
		switch create := createDefinitionContext.(type) {
		case *parser.ColumnDeclarationContext, *parser.ConstraintDeclarationContext, *parser.IndexDeclarationContext:
			createDefinitions = append(createDefinitions, create.Accept(v).(CreateDefinition))
		}
	}

	return createDefinitions
}

func (v *parseTreeVisitor) VisitColumnDeclaration(ctx *parser.ColumnDeclarationContext) interface{} {
	return ColumnDeclaration{
		Column:           ctx.Uid().GetText(),
		ColumnDefinition: ctx.ColumnDefinition().Accept(v).(ColumnDefinition),
	}
}

func (v *parseTreeVisitor) VisitConstraintDeclaration(ctx *parser.ConstraintDeclarationContext) interface{} {
	switch tableConstraint := ctx.TableConstraint().(type) {
	case *parser.PrimaryKeyTableConstraintContext:
		return tableConstraint.Accept(v)
	}

	return nil
}

func (v *parseTreeVisitor) VisitPrimaryKeyTableConstraint(ctx *parser.PrimaryKeyTableConstraintContext) interface{} {

	name := ctx.GetName().GetText()
	index := ctx.GetIndex().GetText()

	return PrimaryKeyTableConstraint{
		Constraint: ctx.CONSTRAINT() != nil,
		Name:       name,
		Index:      index,
	}
}

func (v *parseTreeVisitor) VisitColumnCreateTable(ctx *parser.ColumnCreateTableContext) interface{} {

	ctx.CreateDefinitions()

	return ColumnCreateTable{
		IfNotExists:       ctx.IfNotExists() != nil,
		Replace:           ctx.REPLACE() != nil,
		Temporary:         ctx.TEMPORARY() != nil,
		Table:             ctx.TableName().GetText(),
		CreateDefinitions: ctx.CreateDefinitions().Accept(v).([]CreateDefinition),
	}
}
