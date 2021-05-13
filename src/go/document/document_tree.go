package document

import (
	"encoding/json"
	"fmt"
)

type Kind uint32

const (
	ListNode Kind = 1 << iota
	MapNode
	ScalarNode
	RootNode
)

//TODO I should probably stringify this
func (k Kind) String() string {
	if k == ListNode {
		return "ListNode"
	}
	if k == MapNode {
		return "MapNode"
	}
	if k == ScalarNode {
		return "ScalarNode"
	}
	return "RootNode"
}

type Node struct {
	Begin    int
	End      int
	Kind     Kind
	Children []*Node
}

func newNode(begin, end int, kind Kind) *Node {
	return &Node{
		Begin:    begin,
		End:      end,
		Kind:     kind,
		Children: make([]*Node, 0),
	}
}

func (n *Node) addChild(c *Node) {
	n.Children = append(n.Children, c)
}

type NodeAsJSON struct {
	Begin    int      `json:"begin"`
	End      int      `json:"end"`
	Kind     string   `json:"kind"`
	Children []string `json:"children"`
}

func (n *Node) ToJSON() *NodeAsJSON {
	childStrings := make([]string, len(n.Children))
	for i, c := range n.Children {
		childStrings[i] = c.Kind.String()
	}
	return &NodeAsJSON{
		Begin:    n.Begin,
		End:      n.End,
		Kind:     n.Kind.String(),
		Children: childStrings,
	}
}

func (n *Node) String() string {
	jsonStruct := n.ToJSON()
	json, err := json.Marshal(jsonStruct)
	if err != nil {
		return fmt.Sprintf(`{"%s"}`, err)
	}
	return string(json)
}

type Tree Node

func (t *Tree) PreOrder(cb func(*Node)) {
	cb((*Node)(t))
	for _, child := range t.Children {
		(*Tree)(child).PreOrder(cb)
	}
}

func (t *Tree) Len() int {
	count := 0
	t.PreOrder(func(n *Node) {
		count++
	})
	return count
}

func (t *Tree) ToJSON() ([]byte, error) {
	var jsonTree = make([]*NodeAsJSON, t.Len())
	count := 0
	t.PreOrder(func(n *Node) {
		jsonTree[count] = n.ToJSON()
		count++
	})
	jData, err := json.Marshal(jsonTree)
	if err != nil {
		return nil, err
	}
	return jData, nil
}
