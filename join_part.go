package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ JoinPart = (*InnerJoin)(nil)
	_ JoinPart = (*StraightJoin)(nil)
	_ JoinPart = (*OuterJoin)(nil)
	_ JoinPart = (*NaturalJoin)(nil)
)

type (
	JoinPart interface {
		isJoinPart()
	}
	InnerJoin struct {
		Join            string
		TableSourceItem TableSourceItem
		OnExpression    Expression
		UsingUidList    []string
	}
	StraightJoin struct {
		TableSourceItem TableSourceItem
		OnExpression    Expression
	}
	OuterJoin struct {
		Join            string
		TableSourceItem TableSourceItem
		OnExpression    Expression
		UsingUidList    []string
	}

	NaturalJoin struct {
		Join            string
		TableSourceItem TableSourceItem
	}
)

func (n NaturalJoin) isJoinPart()  {}
func (o OuterJoin) isJoinPart()    {}
func (s StraightJoin) isJoinPart() {}
func (i InnerJoin) isJoinPart()    {}

func (v *parseTreeVisitor) VisitInnerJoin(ctx *parser.InnerJoinContext) interface{} {
	var (
		join         string
		expression   Expression
		usingUidList []string
	)

	if ctx.INNER() != nil {
		join = ctx.INNER().GetText()
	}

	if ctx.CROSS() != nil {
		join = ctx.CROSS().GetText()
	}

	expressionContext := ctx.Expression()
	if expressionContext != nil {
		expression = expressionContext.Accept(v).(Expression)
	} else {
		usingUidList = ctx.UidList().Accept(v).([]string)
	}

	return InnerJoin{
		Join:            join,
		TableSourceItem: ctx.TableSourceItem().Accept(v).(TableSourceItem),
		OnExpression:    expression,
		UsingUidList:    usingUidList,
	}
}

func (v *parseTreeVisitor) VisitStraightJoin(ctx *parser.StraightJoinContext) interface{} {
	var expression Expression
	expressionContext := ctx.Expression()
	if expressionContext != nil {
		expression = expressionContext.Accept(v).(Expression)
	}

	return StraightJoin{
		TableSourceItem: ctx.TableSourceItem().Accept(v).(TableSourceItem),
		OnExpression:    expression,
	}
}

func (v *parseTreeVisitor) VisitOuterJoin(ctx *parser.OuterJoinContext) interface{} {
	var (
		join         string
		expression   Expression
		usingUidList []string
	)

	if ctx.LEFT() != nil {
		join = ctx.LEFT().GetText()
	} else {
		join = ctx.RIGHT().GetText()
	}

	expressionContext := ctx.Expression()
	if expressionContext != nil {
		expression = expressionContext.Accept(v).(Expression)
	} else {
		usingUidList = ctx.UidList().Accept(v).([]string)
	}

	return OuterJoin{
		Join:            join,
		TableSourceItem: ctx.TableSourceItem().Accept(v).(TableSourceItem),
		OnExpression:    expression,
		UsingUidList:    usingUidList,
	}
}

func (v *parseTreeVisitor) VisitNaturalJoin(ctx *parser.NaturalJoinContext) interface{} {
	var join string
	if ctx.LEFT() != nil {
		join = ctx.LEFT().GetText()
	} else {
		join = ctx.RIGHT().GetText()
	}

	return NaturalJoin{
		Join:            join,
		TableSourceItem: ctx.TableSourceItem().Accept(v).(TableSourceItem),
	}
}
