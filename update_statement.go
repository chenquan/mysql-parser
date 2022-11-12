package parser

type (
	UpdateStatement interface {
		IsUpdateStatement()
	}
	SingleUpdateStatement struct {
	}
)
