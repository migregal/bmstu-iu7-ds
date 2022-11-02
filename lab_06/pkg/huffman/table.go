package huffman

import (
	"math"
)

func NewTable(data []byte) *Table {
	table := make(map[byte]*Node)
	for _, b := range data {
		if _, ok := table[b]; !ok {
			table[b] = &Node{value: b, weight: 1}
		} else {
			table[b].weight++
		}
	}
	return &Table{Table: table}
}

func JoinNodes(firstNode, secondNode *Node, label byte) *Node {
	return &Node{
		Left:   firstNode,
		Right:  secondNode,
		weight: firstNode.weight + secondNode.weight,
		value:  label,
	}
}

type Table struct {
	Table map[byte]*Node
}

func (hft *Table) GetSmallestNode() *Node {
	smallest := &Node{weight: math.MaxInt32}
	for _, v := range hft.Table {
		if v.weight < smallest.weight {
			smallest = v
		}
	}
	return smallest
}
