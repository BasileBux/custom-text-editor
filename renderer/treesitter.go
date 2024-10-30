package renderer

import (
	// "os"
	// "fmt"
	st "github.com/basileb/custom_text_editor/settings"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
	tree_sitter_go "github.com/tree-sitter/tree-sitter-go/bindings/go"
)

type Language = uint8

const (
	C Language = iota
	GO
	// CPP
	// PYTHON
	// RUST
	// JAVA
	// JAVASCRIPT
	NONE
)

func RenderText(lang Language, text *string, userStyle *st.WindowStyle) {
	if lang != NONE {

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

		// fmt.Println()
		// traverseAST(root, code)
		// fmt.Println()
		// os.Exit(1)

		switch lang {
		case C:
			syntaxHighlightingC(root, code, userStyle)

		case GO:
			syntaxHighlightingGo(root, code)
		}
	} else {
		noSyntaxHighlight(text, userStyle)
	}
}
