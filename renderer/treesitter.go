package renderer

import (
	st "github.com/basileb/custom_text_editor/settings"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

type Language = uint8

const (
	C Language = iota
	// GO
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

		default:
			break
		}

		tree := parser.Parse(code, nil)
		defer tree.Close()

		root := tree.RootNode()

		switch lang {
		case C:
			syntaxHighlightingC(root, code, userStyle)
		}
	} else {
		noSyntaxHighlight(text, userStyle)
	}
}
