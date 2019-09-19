package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	hfm "./huffman"
)

func main() {

	function := flag.String("m", "encode", "Defines huffman mode."+
		" Values: [\"build-tree\", \"encode\", \"decode\"]")
	text := flag.String("t", "", "text to be used on selected mode")
	flag.Parse()
	treePath := "./tree.json"

	if *function == "build-tree" {
		tree := hfm.BuildTreeFromText(*text)
		b, _ := json.Marshal(&tree)
		ioutil.WriteFile("./tree.json", b, 0644)
		fmt.Println("tree created on " + treePath)
	} else if *function == "encode" {
		b, _ := ioutil.ReadFile(treePath)
		var tree hfm.Tree
		json.Unmarshal(b, &tree)

		if encodedText, err := hfm.Encode(tree, *text); err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("Encoded text: " + encodedText)
		}

	} else if *function == "decode" {
		b, _ := ioutil.ReadFile(treePath)
		var tree hfm.Tree
		json.Unmarshal(b, &tree)

		if decodedText, err := hfm.Decode(tree, *text); err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("Decoded text: " + decodedText)
		}
	}
}
