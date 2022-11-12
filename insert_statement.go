package parser

var (
	_ DmlStatement = (*InsertStatement)(nil)
)

type InsertStatement struct {
	Priority string
	Ignore   bool
}

func (i InsertStatement) IsSqlStatement() {
}

func (i InsertStatement) IsDmlStatement() {
}
