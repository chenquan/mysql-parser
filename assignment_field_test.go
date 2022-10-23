package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitAssignmentField(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("xx")
		result := mySqlParser.AssignmentField().Accept(visitor)
		assert.EqualValues(t, AssignmentField{Val: "xx"}, result)
	})
}
