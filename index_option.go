package parser

import (
	"strings"

	"github.com/chenquan/mysql-parser/internal/parser"
)

var (
	_ IndexOption = (*IndexOptionKeyBlockSize)(nil)
	_ IndexOption = (*IndexOptionWithParser)(nil)
	_ IndexOption = (*IndexOptionComment)(nil)
	_ IndexOption = (*IndexOptionVisible)(nil)
	_ IndexOption = (*IndexOptionEngineAttribute)(nil)
	_ IndexOption = (*IndexOptionSecondaryEngineAttribute)(nil)
)

type (
	IndexOption interface {
		IsIndexOption()
	}
	IndexOptionKeyBlockSize struct {
		BlockSize string
	}
	IndexOptionWithParser struct {
		Uid string
	}
	IndexOptionComment struct {
		Comment string
	}

	IndexOptionVisible struct {
		Visible bool
	}
	IndexOptionEngineAttribute struct {
		EngineAttribute string
	}
	IndexOptionSecondaryEngineAttribute struct {
		EngineAttribute string
	}
)

func (i IndexOptionSecondaryEngineAttribute) IsIndexOption() {
}

func (i IndexOptionEngineAttribute) IsIndexOption() {
}

func (i IndexOptionVisible) IsIndexOption() {}

func (i IndexOptionComment) IsIndexOption() {
}

func (i IndexOptionWithParser) IsIndexOption() {
}

func (i IndexOptionKeyBlockSize) IsIndexOption() {
}

func (v *parseTreeVisitor) VisitIndexOption(ctx *parser.IndexOptionContext) interface{} {
	if ctx.KEY_BLOCK_SIZE() != nil {
		return IndexOptionKeyBlockSize{BlockSize: ctx.FileSizeLiteral().GetText()}
	}

	indexTypeContext := ctx.IndexType()
	if indexTypeContext != nil {
		return indexTypeContext.Accept(v).(IndexType)
	}

	if ctx.WITH() != nil && ctx.PARSER() != nil {
		return IndexOptionWithParser{Uid: ctx.Uid().GetText()}
	}

	if ctx.COMMENT() != nil {
		return IndexOptionComment{Comment: strings.Trim(ctx.STRING_LITERAL().GetText(), "'`\"")}
	}

	if ctx.VISIBLE() != nil {
		return IndexOptionVisible{Visible: true}
	}

	if ctx.INVISIBLE() != nil {
		return IndexOptionVisible{}
	}

	if ctx.ENGINE_ATTRIBUTE() != nil {
		return IndexOptionEngineAttribute{EngineAttribute: strings.Trim(ctx.STRING_LITERAL().GetText(), "'`\"")}
	}

	if ctx.SECONDARY_ENGINE_ATTRIBUTE() != nil {
		return IndexOptionSecondaryEngineAttribute{EngineAttribute: strings.Trim(ctx.STRING_LITERAL().GetText(), "'`\"")}
	}

	return nil
}
