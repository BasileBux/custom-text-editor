package treesitter

import (
	"fmt"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

type Language = uint8

const (
	C Language = iota
	CPP
	GO
	PYTHON
	RUST
	JAVA
	JAVASCRIPT
)

func traverseAST(node *tree_sitter.Node, code []byte) {
	cursor := node.Walk()
	defer cursor.Close()

	for i := uint32(0); !(cursor.Node().KindId() == 161 && i > 1); i++ {
		cursor.GotoDescendant(i)
		start, finish := cursor.Node().ByteRange()
		text := code[start:finish]
		fmt.Println(string(text), " -> ", cursor.Node().Kind(), " -> ", cursor.Node().KindId())
	}

}

func ParseText(lang Language, text *string) {
	code := ([]byte)(*text)
	parser := tree_sitter.NewParser()
	defer parser.Close()
	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_c.Language()))

	tree := parser.Parse(code, nil)
	defer tree.Close()

	root := tree.RootNode()
	traverseAST(root, code)
}
