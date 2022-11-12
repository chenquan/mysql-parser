package parser

import (
	"github.com/chenquan/mysql-parser/internal/parser"
)

func (v *parseTreeVisitor) VisitUidList(ctx *parser.UidListContext) interface{} {
	allUid := ctx.AllUid()
	uids := make([]string, 0, len(allUid))

	for _, uidContext := range allUid {
		uids = append(uids, uidContext.GetText())
	}
	return uids
}
