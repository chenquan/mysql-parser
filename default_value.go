package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ DefaultValue = (*DefaultValueNullLiteral)(nil)
	_ DefaultValue = (*DefaultValueExpression)(nil)
	_ DefaultValue = (*DefaultValueCast)(nil)
	_ DefaultValue = (*DefaultValueConstant)(nil)
	_ DefaultValue = (*DefaultValueLastval)(nil)
	_ DefaultValue = (*DefaultValueNextval)(nil)
)

type (
	DefaultValue interface {
		IsDefaultValue()
	}

	DefaultValueNullLiteral struct {
	}

	DefaultValueExpression struct {
		Expression Expression
		HasBracket bool
	}
	DefaultValueCast struct {
		Expression        Expression
		ConvertedDataType ConvertedDataType
	}
	DefaultValueConstant struct {
		UnaryOperator string
		Constant      Constant
	}
	DefaultValueLastval struct {
		Val FullId
	}
	DefaultValueNextval struct {
		Val FullId
	}
	DefaultValuePreviousValue struct {
		Val FullId
	}
	DefaultValueNextValue struct {
		Val FullId
	}
)

func (d DefaultValueNextValue) IsDefaultValue() {
}

func (d DefaultValuePreviousValue) IsDefaultValue() {
}

func (d DefaultValueNextval) IsDefaultValue() {
}

func (d DefaultValueLastval) IsDefaultValue() {
}

func (d DefaultValueConstant) IsDefaultValue() {
}

func (d DefaultValueExpression) IsDefaultValue() {
}

func (d DefaultValueCast) IsDefaultValue() {
}

func (d DefaultValueNullLiteral) IsDefaultValue() {
}

func (v *parseTreeVisitor) VisitDefaultValue(ctx *parser.DefaultValueContext) interface{} {
	nullLiteralCtx := ctx.NULL_LITERAL()
	if nullLiteralCtx != nil {
		nullLiteralCtx.GetText()
		return DefaultValueNullLiteral{}
	}

	castCtx := ctx.CAST()
	if castCtx != nil {
		return DefaultValueCast{
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

		return DefaultValueConstant{Constant: constantContext.Accept(v).(Constant), UnaryOperator: unaryOperator}
	}

	if ctx.LR_BRACKET() != nil && ctx.RR_BRACKET() != nil {
		expressionContext := ctx.Expression()
		if expressionContext != nil {
			return DefaultValueExpression{
				HasBracket: true,
				Expression: expressionContext.Accept(v).(Expression),
			}
		}

		if ctx.LASTVAL() != nil {
			return DefaultValueLastval{
				Val: ctx.FullId().Accept(v).(FullId),
			}
		}

		if ctx.NEXTVAL() != nil {
			return DefaultValueNextval{
				Val: ctx.FullId().Accept(v).(FullId),
			}
		}

		if ctx.PREVIOUS() != nil {
			return DefaultValuePreviousValue{
				Val: ctx.FullId().Accept(v).(FullId),
			}
		}
		if ctx.NEXT() != nil {
			return DefaultValueNextValue{
				Val: ctx.FullId().Accept(v).(FullId),
			}
		}

	}

	expressionCtx := ctx.Expression()
	if expressionCtx != nil {
		return DefaultValueExpression{
			Expression: expressionCtx.Accept(v).(Expression),
			HasBracket: false,
		}
	}

	return nil
}
