package huffman

import (
	"errors"
	"sort"
	"strings"
)

// Tree is a structure that represents a huffman tree and stores the tree root node
type Tree struct {
	Root *Node
}

// Node is a structure that represents a huffman tree node and stores the node value and frenquency
type Node struct {
	Left      *Node
	Right     *Node
	Value     string
	Frequency int
}

// Encode is a method that encodes a string to huffman base using a given huffman tree
func Encode(tree Tree, text string) (string, error) {
	coding := ""

	for _, char := range text {
		encodeChar := ""
		curNode := tree.Root
		if strings.Contains(tree.Root.Value, string(char)) == false {
			return "", errors.New("Text contains invalid chars, char \"" + string(char) + "\" not found on huffman tree")
		}
		for curNode.Value != string(char) {
			if curNode.Left == nil && curNode.Right == nil {
				encodeChar = encodeChar + "?"
				break
			}
			if curNode.Left != nil {
				if strings.Contains(curNode.Left.Value, string(char)) {
					curNode = curNode.Left
					encodeChar = encodeChar + "0"
				}
			}
			if curNode.Right != nil {
				if strings.Contains(curNode.Right.Value, string(char)) {
					curNode = curNode.Right
					encodeChar = encodeChar + "1"
				}
			}
		}
		coding += encodeChar
	}

	return coding, nil
}

// Decode is method that decodes a huffman base to string using a given huffman tree
func Decode(tree Tree, encodedText string) (string, error) {

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
	return text, nil
}

// BuildTreeFromText is a methdo that build a huffman tree based on characters that exists on a given string
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
