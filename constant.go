package parser

import (
	"strings"

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
		FunctionArg
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
		Val float64
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

func (c ConstantBitString) IsFunctionArg() {
}

func (c ConstantReal) IsFunctionArg() {
}

func (c ConstantHexadecimal) IsFunctionArg() {
}

func (c ConstantDecimal) IsFunctionArg() {
}

func (c ConstantNull) IsFunctionArg() {
}

func (c ConstantString) IsFunctionArg() {
}

func (c ConstantBool) IsFunctionArg() {
}

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
		return ConstantString{Val: stringLiteralContext.Accept(v).(string)}
	}

	decimalContext := ctx.DecimalLiteral()
	if decimalContext != nil {
		var decimal string
		minusCtx := ctx.MINUS()
		if minusCtx != nil {
			decimal = "-"
		}
		decimal += decimalContext.GetText()

		return ConstantDecimal{Val: toDecimal(decimal)}
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

var quotaReplacer = strings.NewReplacer("'", "", `"`, "", `\'`, "")

func (v *parseTreeVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	text := ctx.STRING_LITERAL(0).GetText()
	// TODO
	return quotaReplacer.Replace(text)
}
