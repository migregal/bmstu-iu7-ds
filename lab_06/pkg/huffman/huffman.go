package huffman

import (
	"errors"
)

func NewHuffmanTree(data []byte) *Tree {
	huffmanTable := NewTable(data)
	var j *Node
	i := 0

	for len(huffmanTable.Table) > 1 {
		firstSmallestNode := huffmanTable.GetSmallestNode()
		delete(huffmanTable.Table, firstSmallestNode.value)

		secondSmallestNode := huffmanTable.GetSmallestNode()
		delete(huffmanTable.Table, secondSmallestNode.value)

		j = JoinNodes(firstSmallestNode, secondSmallestNode, byte(i))
		i++
		huffmanTable.Table[j.value] = j
	}
	return &Tree{Root: j, Data: data}
}

type Node struct {
	value  byte
	weight int
	Left   *Node
	Right  *Node
}

type Tree struct {
	Data []byte
	Root *Node
}

func (t *Tree) Compress(res *[]byte) error {
	if t.Root == nil {
		return errors.New("root cannot be null")
	}
	for _, c := range t.Data {
		stack := Stack{}
		t.encodeByte(c, stack.New(), t.Root, res)
	}
	return nil
}

func (t *Tree) Decompress(encoded []byte) ([]byte, error) {
	if t.Root == nil {
		return nil, errors.New("root cannot be null")
	}

	n := t.Root
	var decoded []byte
	for _, b := range encoded {
		if b == '0' {
			n = n.Left
		}
		if b == '1' {
			n = n.Right
		}
		if n.Left == nil && n.Right == nil {
			decoded = append(decoded, n.value)
			n = t.Root
		}
	}

	return decoded, nil
}

func (t *Tree) encodeByte(b byte, s *Stack, rootNode *Node, res *[]byte) {
	if rootNode.Left == nil && rootNode.Right == nil {
		if rootNode.value == b {
			for _, i := range s.items {
				*res = append(*res, byte(i))
			}
		}
		return
	}
	s.Push('0')
	t.encodeByte(b, s, rootNode.Left, res)
	s.Pop()

	s.Push('1')
	t.encodeByte(b, s, rootNode.Right, res)
	s.Pop()
}
