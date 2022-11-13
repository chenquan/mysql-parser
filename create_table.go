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
		Table             FullId
		CreateDefinitions []CreateDefinition
		SelectStatement   SelectStatement
	}
	ColumnCreateTable struct {
		IfNotExists       bool
		Replace           bool
		Temporary         bool
		Table             FullId
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
	var createDefinitions []CreateDefinition
	createDefinitionsContext := ctx.CreateDefinitions()
	if createDefinitionsContext != nil {
		createDefinitions = createDefinitionsContext.Accept(v).([]CreateDefinition)
	}

	var selectStatement SelectStatement
	selectStatementContext := ctx.SelectStatement()
	if selectStatementContext != nil {
		selectStatement = selectStatementContext.Accept(v).(SelectStatement)
	}

	return QueryCreateTable{
		IfNotExists:       ctx.IfNotExists() != nil,
		Replace:           ctx.OR() != nil && ctx.REPLACE(0) != nil,
		Temporary:         ctx.TEMPORARY() != nil,
		Table:             ctx.TableName().Accept(v).(FullId),
		CreateDefinitions: createDefinitions,
		SelectStatement:   selectStatement,
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
	return PrimaryKeyTableConstraint{
		Constraint: ctx.CONSTRAINT() != nil,
		Name:       ctx.GetName().GetText(),
		Index:      ctx.GetIndex().GetText(),
	}
}

func (v *parseTreeVisitor) VisitColumnCreateTable(ctx *parser.ColumnCreateTableContext) interface{} {
	return ColumnCreateTable{
		IfNotExists:       ctx.IfNotExists() != nil,
		Replace:           ctx.REPLACE() != nil,
		Temporary:         ctx.TEMPORARY() != nil,
		Table:             ctx.TableName().Accept(v).(FullId),
		CreateDefinitions: ctx.CreateDefinitions().Accept(v).([]CreateDefinition),
	}
}
