package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type AssignmentField struct {
	Val string
}

func (v *parseTreeVisitor) VisitAssignmentField(ctx *parser.AssignmentFieldContext) interface{} {
	return AssignmentField{Val: ctx.GetChild(0).(interface {
		GetText() string
	}).GetText()}
}
