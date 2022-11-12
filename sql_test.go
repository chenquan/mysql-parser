package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	result := Parse("DROP TABLE if exists A")
	fmt.Println(result)
}
