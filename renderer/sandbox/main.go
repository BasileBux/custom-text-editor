package main

import (
	"fmt"
	"strings"

	f "github.com/basileb/kenzan/files"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

const END_NODE_KIND_ID uint16 = 161

// Function to print the AST and allow for new language syntax highlighting
func traverseAST(node *tree_sitter.Node, code []byte) {
	cursor := node.Walk()
	defer cursor.Close()

	for i := uint32(0); !(cursor.Node().KindId() == END_NODE_KIND_ID && i > 1); i++ {
		cursor.GotoDescendant(i)
		start, finish := cursor.Node().ByteRange()
		text := code[start:finish]
		text = []byte(strings.TrimSuffix(string(text), "\n"))

		// Only show leaf nodes which will be highlighted
		if cursor.Node().ChildCount() == 0 {
			fmt.Println(string(text), " -> ", cursor.Node().Kind(), " -> ", cursor.Node().KindId())
		}
	}

}

// Test new languages. Change the treesitter language and the input file
func main() {

	userText, err := f.OpenFile("../../highlight_samples/main.c")
	if err != nil {
		fmt.Println(err)
		return
	}

	var textToRender string
	for _, l := range userText {
		textToRender += l
		textToRender += "\n"
	}
	textToRender = strings.TrimRight(textToRender, "\n")

	code := ([]byte)(textToRender)
	parser := tree_sitter.NewParser()
	defer parser.Close()

	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_c.Language()))

	tree := parser.Parse(code, nil)
	defer tree.Close()

	root := tree.RootNode()
	traverseAST(root, code)
}
