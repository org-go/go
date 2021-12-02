package ast

import (
	"reflect"
	"strconv"
	"strings"
)

const (
	Number   = 0
	Operator = 1
	A        = `+`
	S        = `-`
)

type Node struct {
	Type  int
	Value string
	Left  *Node
	Right *Node
}

func getAst(expr string) *Node {

	operator := make(map[string]int)
	operator[A], operator[S] = Operator, Operator
	nodes := make([]Node, 0)
	var root *Node
	words := strings.Split(strings.Trim(expr, " "), " ")
	for _, word := range words {
		var tp int
		if _, ok := operator[word]; ok {
			tp = Operator
		} else {
			tp = Number
		}
		nodes = append(nodes, Node{
			Type:  tp,
			Value: word,
		})
	}
	for i := 0; i < len(nodes); i++ {
		if root == nil {
			root = &nodes[i]
			continue
		}
		switch nodes[i].Type {
		case Operator:
			nodes[i].Left = root
			root = &nodes[i]
		case Number:
			root.Right = &nodes[i]
		}
	}
	return root
}

func getResult(node *Node) string {
	switch node.Type {
	case Operator:
		l := getResult(node.Left)
		r := getResult(node.Right)
		return cale(l, r, node.Value)
	case Number:
		return node.Value
	}
	return ""
}

func cale(left, right string, operator string) string {

	lv, _ := transToInt(left)
	rv, _ := transToInt(right)
	val := 0
	switch operator {
	case A:
		val = lv + rv
	case S:
		val = lv - rv
	}
	return transToString(val)

}

func transToString(data interface{}) string {

	return strconv.FormatInt(reflect.ValueOf(data).Int(), 10)
}

func transToInt(data interface{}) (int, error) {

	return strconv.Atoi(strings.TrimSpace(data.(string)))
}
