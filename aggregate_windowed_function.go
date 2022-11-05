package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ FunctionCall = (*AggregateWindowedFunction)(nil)
)

type (
	AggregateWindowedFunction struct {
		Function     string
		StarArg      bool
		Aggregator   string
		FunctionArgs []FunctionArg
	}
)

func (a AggregateWindowedFunction) IsFunctionCall() {
}

func (v *parseTreeVisitor) VisitAggregateWindowedFunction(ctx *parser.AggregateWindowedFunctionContext) interface{} {
	var function string
	functionChild := ctx.GetChild(0)
	function = functionChild.(antlr.TerminalNode).GetText()

	var aggregator string
	aggregatorCtx := ctx.GetAggregator()
	if aggregatorCtx != nil {
		aggregator = aggregatorCtx.GetText()
	}

	var functionArgs []FunctionArg
	switch aggregator {
	case "ALL", "":
		functionArgCtx := ctx.FunctionArg()
		if functionArgCtx != nil {
			arg := functionArgCtx.Accept(v).(FunctionArg)
			functionArgs = append(functionArgs, arg)
		}
	case "DISTINCT":
		functionArgsCtx := ctx.FunctionArgs()
		if functionArgsCtx != nil {
			args := functionArgsCtx.Accept(v).([]FunctionArg)
			functionArgs = append(functionArgs, args...)
		}
	}

	return AggregateWindowedFunction{
		Function:     function,
		StarArg:      ctx.GetStarArg() != nil,
		Aggregator:   aggregator,
		FunctionArgs: functionArgs,
	}
}
