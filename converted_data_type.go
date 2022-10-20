package parser

import (
	"strconv"

	"github.com/chenquan/mysql-parser/internal/parser"
)

type ConvertedDataType struct {
	TypeName        string
	LengthDimension LengthDimension
}

type LengthDimension interface {
	IsLengthDimension()
}

var _ LengthDimension = (*LengthOneDimension)(nil)

type LengthOneDimension struct {
	Dimension float64
}
type LengthTowDimension struct {
	Dimension1 float64
	Dimension2 float64
}

func (l LengthOneDimension) IsLengthDimension() {
}

func (v *parseTreeVisitor) VisitConvertedDataType(ctx *parser.ConvertedDataTypeContext) interface{} {
	var typeName string
	typeNameCtx := ctx.GetTypeName()
	if typeNameCtx != nil {
		typeName = typeNameCtx.GetText()
	} else {
		typeName = ctx.GetText()
	}

	var lengthDimension LengthDimension
	lengthOneDimensionCtx := ctx.LengthOneDimension()
	if lengthOneDimensionCtx != nil {
		lengthDimension = lengthOneDimensionCtx.Accept(v).(LengthDimension)
	}

	lengthTwoOptionalDimensionCtx := ctx.LengthTwoOptionalDimension()
	if lengthTwoOptionalDimensionCtx != nil {
		lengthDimension = lengthTwoOptionalDimensionCtx.Accept(v).(LengthDimension)

	}

	return ConvertedDataType{TypeName: typeName, LengthDimension: lengthDimension}
}

func (v *parseTreeVisitor) VisitLengthOneDimension(ctx *parser.LengthOneDimensionContext) interface{} {
	deci := ctx.DecimalLiteral().GetText()
	float, err := strconv.ParseFloat(deci, 10)
	if err != nil {
		panic(err)
	}

	return LengthOneDimension{Dimension: float}
}

func (v *parseTreeVisitor) VisitLengthTwoDimension(ctx *parser.LengthTwoDimensionContext) interface{} {
	decimalLiteralContexts := ctx.AllDecimalLiteral()
	return LengthTowDimension{
		Dimension1: toDecimal(decimalLiteralContexts[0].GetText()),
		Dimension2: toDecimal(decimalLiteralContexts[1].GetText()),
	}
}

func toDecimal(v string) float64 {
	float, err := strconv.ParseFloat(v, 10)
	if err != nil {
		panic(err)
	}
	return float
}
