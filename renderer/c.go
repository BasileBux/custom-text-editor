package renderer

import (
	"strings"

	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func syntaxHighlightingC(node *tree_sitter.Node, code []byte, state *t.ProgramState, style *st.WindowStyle) {
	cursor := node.Walk()
	defer cursor.Close()

	// Reset slice
	state.Cache.Syntax = state.Cache.Syntax[:0]

	textRenderCursor := &t.TextRenderCursor{
		Line: style.PaddingTop,
		Row:  style.PaddingLeft,
		Stop: false,
	}

	lastFinish := uint(0)

	for i := uint32(0); !(cursor.Node().KindId() == 161 && i > 1); i++ {
		cursor.GotoDescendant(i)
		start, finish := cursor.Node().ByteRange()

		if start > lastFinish {
			stringText := string(code[lastFinish:start])
			highlight(&stringText, &style.ColorTheme.Editor.Fg, textRenderCursor, state, style)
		}

		text := code[start:finish]
		stringText := string(text)

		switch cursor.Node().KindId() {
		case 93: // Primitive type
			highlight(&stringText, &style.ColorTheme.Syntax.Tag, textRenderCursor, state, style)

		case 1, 362: // Identifier
			if cursor.Node().Parent().KindId() == 230 { // Function name
				highlight(&stringText, &style.ColorTheme.Syntax.Func, textRenderCursor, state, style)
			} else if cursor.Node().Parent().KindId() == 165 { // Defined word
				highlight(&stringText, &style.ColorTheme.Syntax.Func, textRenderCursor, state, style)
			} else if cursor.Node().Parent().KindId() == 299 { // Called function -> can also be 266
				highlight(&stringText, &style.ColorTheme.Syntax.Func, textRenderCursor, state, style)
			} else if cursor.Node().Parent().KindId() == 199 {
				highlight(&stringText, &style.ColorTheme.Syntax.Tag, textRenderCursor, state, style)
			} else if cursor.Node().Parent().KindId() == 198 {
				highlight(&stringText, &style.ColorTheme.Syntax.Tag, textRenderCursor, state, style)
			} else { // variable name
				highlight(&stringText, &style.ColorTheme.Editor.Fg, textRenderCursor, state, style)
			}

		case 205, 44, 96, 106, 98, 99, 100, 102, 101, 103, 104, 105, 107, 108, 109, 127: // Keywords
			highlight(&stringText, &style.ColorTheme.Syntax.Keyword, textRenderCursor, state, style)

		case 146, 147, 153: // String content
			if cursor.Node().Parent().KindId() == 164 { // include string (local lib)
				highlight(&stringText, &style.ColorTheme.Syntax.String, textRenderCursor, state, style)
			} else {
				highlight(&stringText, &style.ColorTheme.Syntax.String, textRenderCursor, state, style)
			}

		case 7, 95: // Comma, column
			highlight(&stringText, &style.ColorTheme.Syntax.Comment, textRenderCursor, state, style)

		case 154: // Escape sequence
			highlight(&stringText, &style.ColorTheme.Syntax.Escape, textRenderCursor, state, style)

		case 152: // Quotes
			highlight(&stringText, &style.ColorTheme.Syntax.String, textRenderCursor, state, style)

		case 5, 8: // Parenthesis
			highlight(&stringText, &style.ColorTheme.Syntax.Entity, textRenderCursor, state, style)

		case 64, 65: // Curly brackets
			highlight(&stringText, &style.ColorTheme.Syntax.Entity, textRenderCursor, state, style)

		case 70, 72: // Angle brackets
			highlight(&stringText, &style.ColorTheme.Syntax.Entity, textRenderCursor, state, style)

		case 26, 33: // Pointers operators
			highlight(&stringText, &style.ColorTheme.Syntax.Operator, textRenderCursor, state, style)

		case 140: // Pointers arrow (struct->item)
			highlight(&stringText, &style.ColorTheme.Syntax.Comment, textRenderCursor, state, style)

		case 42: // ;
			highlight(&stringText, &style.ColorTheme.Syntax.Comment, textRenderCursor, state, style)

		case 24, 25, 27, 28, 73, 118, 119, 126, 125: // Operators (arithmetic)
			highlight(&stringText, &style.ColorTheme.Syntax.Operator, textRenderCursor, state, style)

		case 31, 32: // Operators bitwise
			highlight(&stringText, &style.ColorTheme.Syntax.Operator, textRenderCursor, state, style)

		case 22, 35, 36, 37, 38, 39: // Operators (comparison)
			highlight(&stringText, &style.ColorTheme.Syntax.Operator, textRenderCursor, state, style)

		case 30, 29: // and, or
			highlight(&stringText, &style.ColorTheme.Syntax.Entity, textRenderCursor, state, style)

		case 141: // number literals
			highlight(&stringText, &style.ColorTheme.Syntax.Constant, textRenderCursor, state, style)

		case 156, 157: // true / false constant
			highlight(&stringText, &style.ColorTheme.Syntax.Constant, textRenderCursor, state, style)

		case 2, 4: // include, define
			highlight(&stringText, &style.ColorTheme.Syntax.Keyword, textRenderCursor, state, style)

		case 155: // system lib string
			highlight(&stringText, &style.ColorTheme.Syntax.String, textRenderCursor, state, style)

		case 18: // preproc_arg = defined keyword
			highlight(&stringText, &style.ColorTheme.Syntax.Constant, textRenderCursor, state, style)

		case 160: // !comments -> block comments can be on multiple lines. This is really bad
			if stringText[1] == '*' {
				lines := strings.Split(stringText, "\n")
				for j, s := range lines {
					if j != 0 {
						s += "\n"
					}
					highlight(&s, &style.ColorTheme.Syntax.Comment, textRenderCursor, state, style)
				}
			} else {
				highlight(&stringText, &style.ColorTheme.Syntax.Comment, textRenderCursor, state, style)
			}

		default: // Other elements which are leaf
			if cursor.Node().ChildCount() == 0 {
				stringText := string(text)
				highlight(&stringText, &style.ColorTheme.Editor.Fg, textRenderCursor, state, style)
			}
		}

		lastFinish = finish
	}
}
