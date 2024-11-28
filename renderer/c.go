package renderer

import (
	"strings"

	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func syntaxHighlightingC(node *tree_sitter.Node, code []byte, state *t.ProgramState, userStyle *st.WindowStyle) {
	cursor := node.Walk()
	defer cursor.Close()

	lastFinish := uint(0)
	textRenderCursor := &TextRenderCursor{
		line:         userStyle.PaddingTop,
		row:          userStyle.PaddingLeft,
		scrollOffset: state.Nav.ScrollOffset,
	}

	terminateRender := false

	for i := uint32(0); !(cursor.Node().KindId() == 161 && i > 1); i++ {
		cursor.GotoDescendant(i)
		start, finish := cursor.Node().ByteRange()

		if start > lastFinish {
			stringText := string(code[lastFinish:start])
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Editor.Fg, state, userStyle)
		}

		text := code[start:finish]
		stringText := string(text)

		switch cursor.Node().KindId() {
		case 93: // Primitive type
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Tag, state, userStyle)

		case 1, 362: // Identifier
			if cursor.Node().Parent().KindId() == 230 { // Function name
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Func, state, userStyle)
			} else if cursor.Node().Parent().KindId() == 165 { // Defined word
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Func, state, userStyle)
			} else if cursor.Node().Parent().KindId() == 299 { // Called function -> can also be 266
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Func, state, userStyle)
			} else if cursor.Node().Parent().KindId() == 199 {
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Tag, state, userStyle)
			} else if cursor.Node().Parent().KindId() == 198 {
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Tag, state, userStyle)
			} else { // variable name
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Editor.Fg, state, userStyle)
			}

		case 205, 44, 96, 106, 98, 99, 100, 102, 101, 103, 104, 105, 107, 108, 109, 127: // Keywords
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Keyword, state, userStyle)

		case 146, 147, 153: // String content
			if cursor.Node().Parent().KindId() == 164 { // include string (local lib)
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.String, state, userStyle)
			} else {
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.String, state, userStyle)
			}

		case 7, 95: // Comma, column
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Comment, state, userStyle)

		case 154: // Escape sequence
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Escape, state, userStyle)

		case 152: // Quotes
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.String, state, userStyle)

		case 5, 8: // Parenthesis
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Entity, state, userStyle)

		case 64, 65: // Curly brackets
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Entity, state, userStyle)

		case 70, 72: // Angle brackets
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Entity, state, userStyle)

		case 26, 33: // Pointers operators
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Operator, state, userStyle)

		case 140: // Pointers arrow (struct->item)
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Comment, state, userStyle)

		case 42: // ;
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Comment, state, userStyle)

		case 24, 25, 27, 28, 73, 118, 119, 126, 125: // Operators (arithmetic)
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Operator, state, userStyle)

		case 31, 32: // Operators bitwise
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Operator, state, userStyle)

		case 22, 35, 36, 37, 38, 39: // Operators (comparison)
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Operator, state, userStyle)

		case 30, 29: // and, or
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Entity, state, userStyle)

		case 141: // number literals
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Constant, state, userStyle)

		case 156, 157: // true / false constant
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Constant, state, userStyle)

		case 2, 4: // include, define
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Keyword, state, userStyle)

		case 155: // system lib string
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.String, state, userStyle)

		case 18: // preproc_arg = defined keyword
			terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Constant, state, userStyle)

		case 160: // !comments -> block comments can be on multiple lines. This is really bad
			if stringText[1] == '*' {
				lines := strings.Split(stringText, "\n")
				for i, s := range lines {
					if i != 0 {
						s += "\n"
					}
					terminateRender = textRenderCursor.DrawTextPart(&s, userStyle.ColorTheme.Syntax.Comment, state, userStyle)
				}
			} else {
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Comment, state, userStyle)
			}

		default: // Other elements which are leaf
			if cursor.Node().ChildCount() == 0 {
				stringText := string(text)
				terminateRender = textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Editor.Fg, state, userStyle)
			}
		}

		lastFinish = finish
		if terminateRender {
			break
		}
	}
}
