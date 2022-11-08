package parser

import (
	"strings"

	"github.com/chenquan/mysql-parser/internal/parser"
)

type (
	CreateDatabaseOption interface {
		IsCreateDatabaseOption()
	}
	CreateDatabaseOptionCharSet struct {
		CharSet     string
		CharsetName string
	}

	CreateDatabaseOptionCollate struct {
		CollationName string
	}
	CreateDatabaseOptionEncryption struct {
		Encryption string
	}
	CreateDatabaseOptionReadonly struct {
		Default bool
		Num     int
	}
)

func (c CreateDatabaseOptionReadonly) IsCreateDatabaseOption() {
}

func (c CreateDatabaseOptionCollate) IsCreateDatabaseOption() {
}

func (c CreateDatabaseOptionEncryption) IsCreateDatabaseOption() {
}

func (c CreateDatabaseOptionCharSet) IsCreateDatabaseOption() {
}

func (v *parseTreeVisitor) VisitCreateDatabaseOption(ctx *parser.CreateDatabaseOptionContext) interface{} {
	charSetCtx := ctx.CharSet()
	if charSetCtx != nil {
		var charsetName string
		charsetNameCtx := ctx.CharsetName()
		if charsetNameCtx != nil {
			charsetName = charsetNameCtx.GetText()
		} else {
			allDEFAULT := ctx.AllDEFAULT()
			charsetName = allDEFAULT[len(allDEFAULT)-1].GetText()
		}
		return CreateDatabaseOptionCharSet{
			CharSet:     charSetCtx.Accept(v).(string),
			CharsetName: charsetName,
		}
	}

	if ctx.COLLATE() != nil {
		return CreateDatabaseOptionCollate{
			CollationName: ctx.CollationName().GetText(),
		}
	}

	if ctx.ENCRYPTION() != nil {
		return CreateDatabaseOptionEncryption{Encryption: strings.Trim(ctx.STRING_LITERAL().GetText(), "'`\"")}
	}
	if ctx.ONLY() != nil && ctx.READ() != nil {
		if ctx.DEFAULT(0) != nil {
			return CreateDatabaseOptionReadonly{
				Default: true,
			}
		}

		if ctx.ZERO_DECIMAL() != nil {
			return CreateDatabaseOptionReadonly{
				Num: 0,
			}
		}

		if ctx.ONE_DECIMAL() != nil {
			return CreateDatabaseOptionReadonly{
				Num: 1,
			}
		}
	}

	return nil
}

func (v *parseTreeVisitor) VisitCharSet(ctx *parser.CharSetContext) interface{} {
	if ctx.CHARACTER() != nil {
		return "CHARACTER SET"
	}

	if ctx.CHARSET() != nil {
		return "CHARSET"
	}

	if ctx.CHAR() != nil {
		return "CHAR SET"
	}

	return ""
}
