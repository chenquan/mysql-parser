package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ Constant = (*ConstantBool)(nil)
	_ Constant = (*ConstantString)(nil)
	_ Constant = (*ConstantNull)(nil)
	_ Constant = (*ConstantDecimal)(nil)
	_ Constant = (*ConstantHexadecimal)(nil)
	_ Constant = (*ConstantReal)(nil)
	_ Constant = (*ConstantBitString)(nil)
)

type (
	Constant interface {
		IsConstant()
	}
	ConstantBool struct {
		Val bool
	}
	ConstantString struct {
		Val string
	}
	ConstantNull struct {
		Val string
	}
	ConstantDecimal struct {
		Val string
	}
	ConstantHexadecimal struct {
		Val string
	}
	ConstantReal struct {
		Val string
	}
	ConstantBitString struct {
		Val string
	}
)

func (c ConstantBitString) IsConstant() {
}

func (c ConstantReal) IsConstant() {
}

func (c ConstantHexadecimal) IsConstant() {
}

func (c ConstantDecimal) IsConstant() {
}

func (c ConstantNull) IsConstant() {
}

func (c ConstantBool) IsConstant() {
}

func (c ConstantString) IsConstant() {
}

func (v *parseTreeVisitor) VisitConstant(ctx *parser.ConstantContext) interface{} {

	stringLiteralContext := ctx.StringLiteral()
	if stringLiteralContext != nil {
		return ConstantString{Val: stringLiteralContext.GetText()}
	}

	decimalContext := ctx.DecimalLiteral()
	if decimalContext != nil {
		var decimal string
		minusCtx := ctx.MINUS()
		if minusCtx != nil {
			decimal = "-"
		}
		decimal += decimalContext.GetText()

		return ConstantDecimal{Val: decimal}
	}

	hexadecimalLiteralContext := ctx.HexadecimalLiteral()
	if hexadecimalLiteralContext != nil {
		return ConstantHexadecimal{Val: hexadecimalLiteralContext.GetText()}
	}

	booleanLiteralContext := ctx.BooleanLiteral()
	if booleanLiteralContext != nil {
		b := booleanLiteralContext.GetText()
		return ConstantBool{Val: b == "TRUE"}
	}

	realLiteralContext := ctx.REAL_LITERAL()
	if realLiteralContext != nil {
		return ConstantReal{Val: realLiteralContext.GetText()}
	}

	bitStringCtx := ctx.BIT_STRING()
	if bitStringCtx != nil {
		return ConstantBitString{Val: bitStringCtx.GetText()}
	}

	nullLiteralContext := ctx.GetNullLiteral()
	if nullLiteralContext != nil {
		notContext := ctx.NOT()
		var constant string
		if notContext != nil {
			constant = notContext.GetText()
		}
		if constant != "" {
			constant += " "
		}
		constant += nullLiteralContext.GetText()

		return ConstantNull{Val: constant}
	}

	return ctx.GetText()
}
