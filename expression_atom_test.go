package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/mysql-parser/internal/parser"
	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitConstantExpressionAtom(t *testing.T) {
	// FIXME
	sql := "-132.3"
	sql = strings.ToUpper(sql)
	inputStream := antlr.NewInputStream(sql)
	lexer := parser.NewMySqlLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	mySqlParser := parser.NewMySqlParser(tokens)
	visitor := &parseTreeVisitor{
		Result: &Result{},
	}
	fmt.Println(mySqlParser.ExpressionAtom().GetText())
	result := mySqlParser.ExpressionAtom().Accept(visitor)
	_ = result
}

func Test_parseTreeVisitor_VisitFullColumnNameExpressionAtom(t *testing.T) {
	sql := "a.c.name"
	sql = strings.ToUpper(sql)
	inputStream := antlr.NewInputStream(sql)
	lexer := parser.NewMySqlLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	mySqlParser := parser.NewMySqlParser(tokens)
	visitor := &parseTreeVisitor{
		Result: &Result{},
	}

	result := mySqlParser.ExpressionAtom().Accept(visitor)
	assert.EqualValues(t, result, FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
		Uid: "A",
		DottedIds: []DottedId{
			{Uid: "C"},
			{Uid: "NAME"},
		},
	}})
}

func Test_parseTreeVisitor_VisitFunctionCallExpressionAtom(t *testing.T) {
	mySqlParser, visitor := createMySqlParser("CURRENT_DATE()")
	result := mySqlParser.ExpressionAtom().Accept(visitor)
	assert.EqualValues(t, result, FunctionCallExpressionAtom{FunctionCall: SpecificFunctionCall{SpecificFunction: SimpleFunctionCall{Function: "CURRENT_DATE"}}})
}

func Test_parseTreeVisitor_VisitMathExpressionAtom(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("a+2")
		result := mySqlParser.ExpressionAtom().Accept(visitor)
		assert.EqualValues(t,
			MathExpressionAtom{
				LeftExpressionAtom: FullColumnNameExpressionAtom{FullColumnName: FullColumnName{
					Uid: "a",
				}},
				MathOperator:        "+",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: "2"}},
			},
			result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1+2")
		result := mySqlParser.ExpressionAtom().Accept(visitor)
		assert.EqualValues(t,
			MathExpressionAtom{
				LeftExpressionAtom:  ConstantExpressionAtom{Constant: ConstantDecimal{Val: "1"}},
				MathOperator:        "+",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: "2"}},
			},
			result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("1+2+3")
		result := mySqlParser.ExpressionAtom().Accept(visitor)
		assert.EqualValues(t,
			MathExpressionAtom{
				LeftExpressionAtom: MathExpressionAtom{
					LeftExpressionAtom:  ConstantExpressionAtom{Constant: ConstantDecimal{Val: "1"}},
					MathOperator:        "+",
					RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: "2"}},
				},
				MathOperator:        "+",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: "3"}},
			},
			result)

	})

	t.Run("4", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("(2+3)*1")
		result := mySqlParser.ExpressionAtom().Accept(visitor)
		assert.EqualValues(t,
			MathExpressionAtom{
				LeftExpressionAtom: NestedExpressionAtom{Expressions: []Expression{
					ExpressionAtomPredicate{ExpressionAtom: MathExpressionAtom{
						LeftExpressionAtom:  ConstantExpressionAtom{Constant: ConstantDecimal{Val: "2"}},
						MathOperator:        "+",
						RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: "3"}},
					}},
				}},
				MathOperator:        "*",
				RightExpressionAtom: ConstantExpressionAtom{Constant: ConstantDecimal{Val: "1"}},
			},
			result)

	})

}

func Test_parseTreeVisitor_VisitExistsExpressionAtom(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("EXISTS (select count(*) from u)")
		result := mySqlParser.ExpressionAtom().Accept(visitor)
		assert.EqualValues(t,
			ExistsExpressionAtom{
				SelectStatement: SimpleSelect{QuerySpecification: QuerySpecification{
					SelectSpecs: nil,
					SelectElements: SelectElements{
						All: false,
						SelectElements: []SelectElement{
							SelectFunctionElement{FunctionCall: AggregateWindowedFunction{
								Function: "count",
								StarArg:  true,
							}},
						},
					},
					FromClause: &FromClause{TableSources: &TableSources{TableSources: []TableSource{
						TableSourceBase{TableSourceItem: AtomTableItem{TableName: "u"}},
					}}},
				}},
			},
			result)
	})
}
