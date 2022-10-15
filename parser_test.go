package parser

//func TestParseTreeVisitor_VisitCreateDatabase(t *testing.T) {
//
//	inputStream := antlr.NewInputStream("CREATE TABLE IOT( A INT);CREATE TABLE IOT1( A INT);CREATE DATABASE `name`;")
//	lexer := parser.NewMySqlLexer(inputStream)
//	lexer.RemoveErrorListeners()
//	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
//	mysqlParser := parser.NewMySqlParser(tokens)
//	visitor := &parseTreeVisitor{}
//	mysqlParser.Root().Accept(visitor)
//}
