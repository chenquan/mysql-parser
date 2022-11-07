package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	ColumnConstraint interface {
		IsColumnConstraint()
		CreateDefinition
	}
	DefaultColumnConstraint struct {
		DefaultValue DefaultValue
	}
)

func (d DefaultColumnConstraint) IsColumnConstraint() {
}

func (d DefaultColumnConstraint) IsCreateDefinition() {
}

func (v *parseTreeVisitor) VisitColumnDefinition(ctx *parser.ColumnDefinitionContext) interface{} {
	allColumnConstraint := ctx.AllColumnConstraint()
	var columnConstraints []ColumnConstraint
	if len(allColumnConstraint) != 0 {
		columnConstraints = make([]ColumnConstraint, 0, len(allColumnConstraint))
		for _, columnConstraintContext := range allColumnConstraint {
			columnConstraints = append(columnConstraints, columnConstraintContext.Accept(v).(ColumnConstraint))
		}
	}

	return ColumnDefinition{
		DataType:          ctx.DataType().GetText(),
		ColumnConstraints: columnConstraints,
	}
}

func (v *parseTreeVisitor) VisitDefaultColumnConstraint(ctx *parser.DefaultColumnConstraintContext) interface{} {

	return DefaultColumnConstraint{
		DefaultValue: ctx.DefaultValue().Accept(v).(DefaultValue),
	}
}
