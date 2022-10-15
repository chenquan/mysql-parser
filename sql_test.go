package parser

//func TestParser(t *testing.T) {
//	t.Run("", func(t *testing.T) {
//		Parser("CREATE TABLE IOT( A INT);CREATE TABLE IOT1( A INT);CREATE DATABASE TT;DROP DATABASE TT;")
//
//	})
//
//	t.Run("ALTER DATABASE", func(t *testing.T) {
//		Parser("ALTER DATABASE A DEFAULT CHARACTER SET gb2312")
//	})
//
//	t.Run("ALTER TABLE", func(t *testing.T) {
//		t.Run("ADD COLUMN", func(t *testing.T) {
//			//Parser("ALTER TABLE user ADD (age INT , class int ,user varchar(2) )")
//			Parser("ALTER TABLE user ADD IF NOT EXISTS age INT ,Add IF NOT EXISTS class varchar(20);ALTER TABLE user ADD IF NOT EXISTS age INT ,Add IF NOT EXISTS class varchar(1001);")
//		})
//	})
//}
