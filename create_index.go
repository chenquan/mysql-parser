package parser

type (
	CreateIndex struct {
		Replace       bool
		InTimeAction  string
		IndexCategory string
		IfNotExists   bool
		IndexName     string
		IndexType     string
		TableName     string
	}
)
