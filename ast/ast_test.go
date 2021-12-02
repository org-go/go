package ast

import (
	"fmt"
	"testing"
)

func Test_cale(t *testing.T) {
	expr := `1 + 4 - 2 + 100 - 20 + 12`
	ast := getAst(expr)

	result := getResult(ast)
	fmt.Println(result)
}
