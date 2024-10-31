package renderer

import (
	st "github.com/basileb/custom_text_editor/settings"
	t "github.com/basileb/custom_text_editor/types"
	rl "github.com/gen2brain/raylib-go/raylib"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

func RenderText(lang t.Language, text *string, state *t.ProgramState, scrollOffset *rl.Vector2, userStyle *st.WindowStyle) {
	if lang != t.NONE {

		code := ([]byte)(*text)
		parser := tree_sitter.NewParser()
		defer parser.Close()

		switch lang {
		case t.C:
			parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_c.Language()))

		default:
			break
		}

		tree := parser.Parse(code, nil)
		defer tree.Close()

		root := tree.RootNode()

		switch lang {
		case t.C:
			syntaxHighlightingC(root, code, state, scrollOffset, userStyle)
		}
	} else {
		noSyntaxHighlight(text, userStyle, scrollOffset, userStyle)
	}
}
