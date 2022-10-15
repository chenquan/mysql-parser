package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitConstantExpressionAtom(t *testing.T) {
	// FIXME
	mySqlParser, visitor := createMySqlParser("-132.3")
	fmt.Println(mySqlParser.ExpressionAtom().GetText())
	result := mySqlParser.ExpressionAtom().Accept(visitor)
	assert.EqualValues(t, ConstantExpressionAtom{Constant: Constant{
		Constant: "-132.3",
	}}, result)
	_ = result
}
