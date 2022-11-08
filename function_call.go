package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var _ FunctionCall = (*SpecificFunctionCall)(nil)

type (
	FunctionCall interface {
		FunctionArg
		IsFunctionCall()
	}
	SpecificFunctionCall struct {
		SpecificFunction SpecificFunction
	}
)

func (s SpecificFunctionCall) IsFunctionArg() {
}

func (s SpecificFunctionCall) IsFunctionCall() {
}

func (v *parseTreeVisitor) VisitSpecificFunctionCall(ctx *parser.SpecificFunctionCallContext) interface{} {
	switch c := ctx.GetChild(0).(type) {
	case *parser.SimpleFunctionCallContext:
		return SpecificFunctionCall{SpecificFunction: c.Accept(v).(SpecificFunction)}
	}

	return nil
}

func (v *parseTreeVisitor) VisitAggregateFunctionCall(ctx *parser.AggregateFunctionCallContext) interface{} {
	switch c := ctx.GetChild(0).(type) {
	case *parser.AggregateWindowedFunctionContext:
		return c.Accept(v).(FunctionCall)
	}

	return nil
}
