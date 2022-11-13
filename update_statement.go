package parser

type (
	UpdateStatement interface {
		isUpdateStatement()
	}
	SingleUpdateStatement struct {
	}
)
