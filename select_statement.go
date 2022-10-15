package parser

type (
	SelectStatement interface {
		IsSelectStatement()
	}
	SimpleSelect struct {
	}
)

func (s SimpleSelect) IsSelectStatement() {
}
