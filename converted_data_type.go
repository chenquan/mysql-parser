package parser

import (
	"strconv"

	"github.com/chenquan/mysql-parser/internal/parser"
)

var _ LengthDimension = (*LengthOneDimension)(nil)

type ConvertedDataType struct {
	TypeName        string
	LengthDimension LengthDimension
}

type (
	LengthDimension interface {
		isLengthDimension()
	}
	LengthOneDimension struct {
		Dimension float64
	}
	LengthTwoDimension struct {
		Dimension1 float64
		Dimension2 float64
	}
)

func (l LengthTwoDimension) isLengthDimension() {}
func (l LengthOneDimension) isLengthDimension() {}

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
	float, err := strconv.ParseFloat(deci, 64)
	if err != nil {
		panic(err)
	}

	return LengthOneDimension{Dimension: float}
}

//func (v *parseTreeVisitor) VisitLengthTwoDimension(ctx *parser.LengthTwoDimensionContext) interface{} {
//	decimalLiteralContexts := ctx.AllDecimalLiteral()
//	return LengthTwoOptionalDimension{
//		Dimension1: toDecimal(decimalLiteralContexts[0].GetText()),
//		Dimension2: toDecimal(decimalLiteralContexts[1].GetText()),
//	}
//}

func (v *parseTreeVisitor) VisitLengthTwoOptionalDimension(ctx *parser.LengthTwoOptionalDimensionContext) interface{} {
	decimalLiteralContexts := ctx.AllDecimalLiteral()
	if len(decimalLiteralContexts) == 2 {
		return LengthTwoDimension{
			Dimension1: toDecimal(decimalLiteralContexts[0].GetText()),
			Dimension2: toDecimal(decimalLiteralContexts[1].GetText()),
		}
	}

	return LengthOneDimension{
		Dimension: toDecimal(decimalLiteralContexts[0].GetText()),
	}

}

func toDecimal(v string) float64 {
	float, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return float
}
