package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	FunctionCall interface {
		IsFunctionCall()
	}
	SpecificFunctionCall struct {
		SpecificFunction SpecificFunction
	}
)

func (s SpecificFunctionCall) IsFunctionCall() {
}

func (v *parseTreeVisitor) VisitSpecificFunctionCall(ctx *parser.SpecificFunctionCallContext) interface{} {
	children := ctx.GetChildren()

	for _, child := range children {
		switch c := child.(type) {
		case *parser.SimpleFunctionCallContext:
			return SpecificFunctionCall{SpecificFunction: c.Accept(v).(SpecificFunction)}
		}
	}

	return nil
}
