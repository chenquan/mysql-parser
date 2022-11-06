package parser

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTreeVisitor_VisitCurrentTimestamp(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		strs := []string{"CURRENT_TIMESTAMP", "LOCALTIME", "LOCALTIMESTAMP", "CURDATE", "CURTIME"}
		for _, str := range strs {
			for i := 0; i < 7; i++ {
				mySqlParser, visitor := createMySqlParser(str + "(" + strconv.Itoa(i) + ")")
				result := mySqlParser.CurrentTimestamp().Accept(visitor)

				assert.EqualValues(t, CurrentTimestamp{
					Current:   str,
					Precision: i,
				}, result)
			}

		}
	})
}
