package treesitter

import (
	"fmt"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
	tree_sitter_go "github.com/tree-sitter/tree-sitter-go/bindings/go"
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
	NONE
)

// ANSI color codes
const (
	Reset   = "\033[0m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	White   = "\033[37m"
)

// Helper functions for color-coding
func printBlue(text string) {
	fmt.Printf("%s%s%s", Blue, text, Reset)
}

func printMagenta(text string) {
	fmt.Printf("%s%s%s", Magenta, text, Reset)
}

func printGreen(text string) {
	fmt.Printf("%s%s%s", Green, text, Reset)
}

func printYellow(text string) {
	fmt.Printf("%s%s%s", Yellow, text, Reset)
}

func printWhite(text string) {
	fmt.Printf("%s%s%s", White, text, Reset)
}

func traverseAST(node *tree_sitter.Node, code []byte) {
	cursor := node.Walk()
	defer cursor.Close()

	for i := uint32(0); !(cursor.Node().KindId() == 161 && i > 1); i++ {
		fmt.Println(i)
		cursor.GotoDescendant(i)
		start, finish := cursor.Node().ByteRange()
		text := code[start:finish]
		fmt.Println(string(text), " -> ", cursor.Node().Kind(), " -> ", cursor.Node().KindId())
		if i > 100 {
			return
		}
	}

}

func ParseText(lang Language, text *string) {
	code := ([]byte)(*text)
	parser := tree_sitter.NewParser()
	defer parser.Close()

	switch lang {
	case C:
		parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_c.Language()))

	case GO:
		parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_go.Language()))

	default:
		break
	}

	tree := parser.Parse(code, nil)
	defer tree.Close()

	root := tree.RootNode()
	traverseAST(root, code)
	fmt.Println()

	switch lang {
	case C:
		syntaxHighlightingC(root, code)

	case GO:
		syntaxHighlightingGo(root, code)

	}
	fmt.Println()
}
