package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	SpecificFunction interface {
		isSpecificFunction()
	}
	SimpleFunctionCall struct {
		Function string
	}
	DataTypeFunctionCall struct {
		Function          string
		Expression        Expression
		ConvertedDataType ConvertedDataType
	}
)

func (s SimpleFunctionCall) isSpecificFunction() {}
func (d DataTypeFunctionCall) isFunctionCall()   {}

func (v *parseTreeVisitor) VisitSimpleFunctionCall(ctx *parser.SimpleFunctionCallContext) interface{} {
	child := ctx.GetChild(0)
	return SimpleFunctionCall{Function: child.(*antlr.TerminalNodeImpl).GetText()}
}

func (v *parseTreeVisitor) VisitDataTypeFunctionCall(ctx *parser.DataTypeFunctionCallContext) interface{} {
	var function string
	convertContext := ctx.CONVERT()
	if convertContext != nil {
		function = convertContext.GetText()
	} else {
		castContext := ctx.CONVERT()
		if castContext != nil {
			function = castContext.GetText()
		}
	}
	return DataTypeFunctionCall{
		Function:          function,
		Expression:        ctx.Expression().Accept(v).(Expression),
		ConvertedDataType: ctx.ConvertedDataType().Accept(v).(ConvertedDataType),
	}
}
