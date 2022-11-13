package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ SelectIntoExpression = (*SelectIntoVariables)(nil)
)

type (
	SelectIntoExpression interface {
		isSelectIntoExpression()
	}
	SelectIntoVariables struct {
		AssignmentFields []AssignmentField
	}
)

func (s SelectIntoVariables) isSelectIntoExpression() {}

func (v *parseTreeVisitor) VisitSelectIntoVariables(ctx *parser.SelectIntoVariablesContext) interface{} {
	allAssignmentFields := ctx.AllAssignmentField()
	assignmentFields := make([]AssignmentField, 0, len(allAssignmentFields))
	for _, assignmentFieldContext := range allAssignmentFields {
		assignmentFields = append(assignmentFields, assignmentFieldContext.Accept(v).(AssignmentField))
	}

	return SelectIntoVariables{AssignmentFields: assignmentFields}
}
