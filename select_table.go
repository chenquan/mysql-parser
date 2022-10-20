package parser

import (
	"strings"

	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	SelectElements struct {
		All            bool
		SelectElements []SelectElement
	}
	SelectElement interface {
		IsSelectElement()
	}
	SelectStarElement struct {
		FullId FullId
	}
	SelectColumnElement struct {
		FullColumnName FullColumnName
		Alias          string
	}

	DottedId struct {
		Uid string
	}
	FullId struct {
		Uid   string
		DotId string
	}
	FullColumnName struct {
		Uid       string
		DottedIds []DottedId
	}
)

type (
	SelectFunctionElement struct {
	}
)

func (s SelectStarElement) IsSelectElement() {
}

func (v *parseTreeVisitor) VisitSelectElements(ctx *parser.SelectElementsContext) interface{} {
	star := ctx.GetStar() != nil
	if star {
		return SelectElements{All: true}
	}

	elementContexts := ctx.AllSelectElement()
	selectElements := make([]SelectElement, 0, len(elementContexts))
	for _, selectElementContext := range elementContexts {
		switch selectElement := selectElementContext.(type) {
		case *parser.SelectStarElementContext, *parser.SelectColumnElementContext:
			selectElements = append(selectElements, selectElement.Accept(v).(SelectStarElement))
		}
	}

	return SelectElements{
		SelectElements: selectElements,
	}
}

func (v *parseTreeVisitor) VisitSelectStarElement(ctx *parser.SelectStarElementContext) interface{} {
	return SelectStarElement{FullId: ctx.FullId().Accept(v).(FullId)}
}

func (v *parseTreeVisitor) VisitFullId(ctx *parser.FullIdContext) interface{} {
	uid := ctx.Uid(0)
	var dotId string
	dotIdContext := ctx.DOT_ID()
	if dotIdContext != nil {
		dotId = strings.ReplaceAll(dotIdContext.GetText(), ".", "")
	}
	return FullId{
		Uid:   uid.GetText(),
		DotId: dotId,
	}
}

func (v *parseTreeVisitor) VisitSelectColumnElement(ctx *parser.SelectColumnElementContext) interface{} {
	var uid string
	uidContext := ctx.Uid()
	if uidContext != nil {
		uid = uidContext.GetText()
	}
	return SelectColumnElement{
		FullColumnName: ctx.FullColumnName().Accept(v).(FullColumnName),
		Alias:          uid,
	}
}

func (v *parseTreeVisitor) VisitFullColumnName(ctx *parser.FullColumnNameContext) interface{} {
	uid := "."
	uidContext := ctx.Uid()
	if uidContext != nil {
		uid = uidContext.GetText()
	}

	allDottedIdContext := ctx.AllDottedId()
	var dottedIds []DottedId
	if len(allDottedIdContext) != 0 {
		dottedIds = make([]DottedId, 0, len(allDottedIdContext))
		for _, context := range allDottedIdContext {
			dottedIds = append(dottedIds, context.Accept(v).(DottedId))
		}
	}

	return FullColumnName{
		Uid:       uid,
		DottedIds: dottedIds,
	}
}

func (v *parseTreeVisitor) VisitDottedId(ctx *parser.DottedIdContext) interface{} {
	var uid string
	dotIdContext := ctx.DOT_ID()
	if dotIdContext != nil {
		uid = strings.ReplaceAll(dotIdContext.GetText(), ".", "")
	} else {
		uidContext := ctx.Uid()
		if uidContext != nil {
			uid = uidContext.GetText()
		}
	}

	return DottedId{Uid: uid}
}

func (v *parseTreeVisitor) VisitSelectFunctionElement(ctx *parser.SelectFunctionElementContext) interface{} {

	return SelectFunctionElement{}

}
