package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

type DefaultValue struct {
	Type              string
	NullLiteral       bool
	Expression        Expression
	ConvertedDataType ConvertedDataType
	UnaryOperator     string
	Constant          Constant
	CurrentTimestamp  CurrentTimestamp
	FullId            FullId
}

func (v *parseTreeVisitor) VisitDefaultValue(ctx *parser.DefaultValueContext) interface{} {
	nullLiteralCtx := ctx.NULL_LITERAL()
	if nullLiteralCtx != nil {
		nullLiteralCtx.GetText()
		return DefaultValue{Type: "NULL", NullLiteral: true}
	}
	castCtx := ctx.CAST()
	if castCtx != nil {
		return DefaultValue{
			Type:              "CAST",
			Expression:        ctx.Expression().Accept(v).(Expression),
			ConvertedDataType: ctx.ConvertedDataType().Accept(v).(ConvertedDataType)}
	}

	constantContext := ctx.Constant()
	if constantContext != nil {
		var unaryOperator string
		unaryOperatorContext := ctx.UnaryOperator()
		if unaryOperatorContext != nil {
			unaryOperator = unaryOperatorContext.GetText()
		}

		return DefaultValue{Type: "CONSTANT", Constant: constantContext.Accept(v).(Constant), UnaryOperator: unaryOperator}
	}

	if ctx.LR_BRACKET() != nil && ctx.RR_BRACKET() != nil {
		expressionContext := ctx.Expression()
		if expressionContext != nil {
			return DefaultValue{
				Type:       "WRAP_EXPRESSION",
				Expression: expressionContext.Accept(v).(Expression),
			}
		}
	}

	if ctx.LASTVAL() != nil {
		return DefaultValue{
			Type:   "LASTVAL",
			FullId: ctx.FullId().Accept(v).(FullId),
		}
	}

	if ctx.NEXTVAL() != nil {
		return DefaultValue{
			Type:   "NEXTVAL",
			FullId: ctx.NEXTVAL().Accept(v).(FullId),
		}
	}

	return DefaultValue{}
}
