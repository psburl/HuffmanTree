package huffman

import (
	"sort"
	"strings"
)

// Tree is
type Tree struct {
	Root *Node
}

// Node is
type Node struct {
	Left      *Node
	Right     *Node
	Value     string
	Frequency int
}

// Encode is
func Encode(tree Tree, text string) string {
	coding := ""

	for _, char := range text {
		encondeChar := ""
		curNode := tree.Root
		if strings.Contains(tree.Root.Value, string(char)) == false {
			return "Texto contém caracteres não exitentes na arvore"
		}
		for curNode.Value != string(char) {
			if curNode.Left == nil && curNode.Right == nil {
				encondeChar = encondeChar + "?"
				break
			}
			if curNode.Left != nil {
				if strings.Contains(curNode.Left.Value, string(char)) {
					curNode = curNode.Left
					encondeChar = encondeChar + "0"
				}
			}
			if curNode.Right != nil {
				if strings.Contains(curNode.Right.Value, string(char)) {
					curNode = curNode.Right
					encondeChar = encondeChar + "1"
				}
			}
		}
		coding += encondeChar
	}

	return coding
}

// Decode is
func Decode(tree Tree, encodedText string) string {

	text := ""
	curText := ""
	curNode := tree.Root
	for len(encodedText) > 0 {
		curText = string(encodedText[0])
		encodedText = encodedText[1:]
		if curText == "0" {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
		if curNode.Left == nil && curNode.Right == nil {
			text += curNode.Value
			curNode = tree.Root
		}
	}
	return text
}

// BuildTreeFromText is
func BuildTreeFromText(text string) Tree {

	nodes := buildCharFrequency(text)

	for len(nodes) > 1 {
		nodes = sortNodes(nodes)
		newNode := newNode(
			nodeCopy(&nodes[0]),
			nodeCopy(&nodes[1]),
			nodes[0].Value+nodes[1].Value,
			nodes[0].Frequency+nodes[1].Frequency)
		nodes = remove(nodes, 0)
		nodes = remove(nodes, 0)
		nodes = append(nodes, newNode)
	}

	tree := Tree{Root: &nodes[0]}
	return tree
}

func buildCharFrequency(text string) []Node {
	frequencies := map[string]int{}
	for _, char := range text {
		frequency, exists := frequencies[string(char)]
		if exists {
			frequencies[string(char)] = frequency + 1
		} else {
			frequencies[string(char)] = 1
		}
	}
	nodes := []Node{}
	for data, frequency := range frequencies {
		nodes = append(nodes, Node{Value: data, Frequency: frequency})
	}
	return nodes
}

func sortNodes(nodes []Node) []Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Frequency < nodes[j].Frequency
	})
	return nodes
}

func remove(slice []Node, s int) []Node {
	if len(slice) == 1 {
		return []Node{}
	}
	return append(slice[:s], slice[s+1:]...)
}

func newNode(left, right *Node, value string, freq int) Node {
	return Node{
		Left:      left,
		Right:     right,
		Value:     value,
		Frequency: freq}
}

func nodeCopy(node *Node) *Node {
	copy := newNode(node.Left, node.Right, node.Value, node.Frequency)
	return &copy
}
