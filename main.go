package main

import (
	bf "bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	hfm "./huffman"
)

func main() {

	function := flag.String("f", "encode", "Defines huffman function")
	flag.Parse()

	ioreader := bf.NewReader(os.Stdin)

	if *function == "build-tree" {
		fmt.Print("Enter base text: ")
		text, _ := ioreader.ReadString('\n')
		text = text[:len(text)-1]
		tree := hfm.BuildTreeFromText(text)
		b, _ := json.Marshal(&tree)
		ioutil.WriteFile("./tree.json", b, 0644)
		fmt.Println("tree created")
	} else if *function == "encode" {
		fmt.Print("Enter text to encode: ")
		text, _ := ioreader.ReadString('\n')
		text = text[:len(text)-1]
		b, _ := ioutil.ReadFile("./tree.json")
		var tree hfm.Tree
		json.Unmarshal(b, &tree)
		fmt.Println(hfm.Encode(tree, text))
	} else if *function == "decode" {
		fmt.Print("Enter text to encode: ")
		text, _ := ioreader.ReadString('\n')
		text = text[:len(text)-1]
		b, _ := ioutil.ReadFile("./tree.json")
		var tree hfm.Tree
		json.Unmarshal(b, &tree)
		fmt.Println(hfm.Decode(tree, text))
	}
}
