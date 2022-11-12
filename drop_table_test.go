package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitDropTable(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("DROP TABLE IF EXISTS A")
		result := mySqlParser.DropTable().Accept(visitor)

		assert.EqualValues(t, DropTable{
			IfExists:  true,
			Temporary: false,
			TableNames: []TableName{{
				Uid:   "A",
				DotId: "",
			}},
			DropType: "",
		}, result)
	})

	t.Run("2", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("DROP TABLE IF EXISTS A,B")
		result := mySqlParser.DropTable().Accept(visitor)

		assert.EqualValues(t, DropTable{
			IfExists:  true,
			Temporary: false,
			TableNames: []TableName{
				{
					Uid:   "A",
					DotId: "",
				},
				{
					Uid:   "B",
					DotId: "",
				},
			},
			DropType: "",
		}, result)
	})

	t.Run("3", func(t *testing.T) {
		mySqlParser, visitor := createMySqlParser("DROP TEMPORARY TABLE IF EXISTS A,B RESTRICT")
		result := mySqlParser.DropTable().Accept(visitor)

		assert.EqualValues(t, DropTable{
			IfExists:  true,
			Temporary: true,
			TableNames: []TableName{
				{
					Uid:   "A",
					DotId: "",
				},
				{
					Uid:   "B",
					DotId: "",
				},
			},
			DropType: "RESTRICT",
		}, result)
	})
}
