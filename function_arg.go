package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type FunctionArg interface {
	isFunctionArg()
}

func (v *parseTreeVisitor) VisitFunctionArgs(ctx *parser.FunctionArgsContext) interface{} {
	allFunctionArgCtx := ctx.AllFunctionArg()
	functionArgs := make([]FunctionArg, 0, len(allFunctionArgCtx))
	for _, functionArgCtx := range allFunctionArgCtx {
		functionArgs = append(functionArgs, functionArgCtx.Accept(v).(FunctionArg))
	}

	return functionArgs
}

func (v *parseTreeVisitor) VisitFunctionArg(ctx *parser.FunctionArgContext) interface{} {
	constantContext := ctx.Constant()
	if constantContext != nil {
		return constantContext.Accept(v).(Constant)
	}

	fullColumnNameContext := ctx.FullColumnName()
	if fullColumnNameContext != nil {
		return fullColumnNameContext.Accept(v).(FullColumnName)
	}

	functionCallContext := ctx.FunctionCall()
	if functionCallContext != nil {
		return functionCallContext.Accept(v).(FunctionCall)
	}

	expressionContext := ctx.Expression()
	if expressionContext != nil {
		return expressionContext.Accept(v).(Expression)
	}

	return nil
}
