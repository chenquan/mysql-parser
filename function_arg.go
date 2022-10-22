package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type FunctionArg struct {
	F interface{}
}

func (v *parseTreeVisitor) VisitFunctionArgs(ctx *parser.FunctionArgsContext) interface{} {
	allFunctionArgCtx := ctx.AllFunctionArg()
	if len(allFunctionArgCtx) == 0 {
		return nil
	}

	functionArgs := make([]FunctionArg, 0, len(allFunctionArgCtx))
	for _, functionArgCtx := range allFunctionArgCtx {
		functionArgs = append(functionArgs, functionArgCtx.Accept(v).(FunctionArg))
	}

	return functionArgs
}

func (v *parseTreeVisitor) VisitFunctionArg(ctx *parser.FunctionArgContext) interface{} {
	constantContext := ctx.Constant()
	if constantContext != nil {
		return FunctionArg{
			F: constantContext.Accept(v),
		}
	}

	fullColumnNameContext := ctx.FullColumnName()
	if fullColumnNameContext != nil {
		return FunctionArg{F: fullColumnNameContext.Accept(v)}
	}

	functionCallContext := ctx.FunctionCall()
	if functionCallContext != nil {
		return FunctionArg{F: functionCallContext.Accept(v)}
	}

	expressionContext := ctx.Expression()
	if expressionContext != nil {
		return FunctionArg{F: expressionContext.Accept(v)}
	}

	return nil
}
