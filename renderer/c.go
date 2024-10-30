package renderer

import (
	"fmt"
	"strings"

	st "github.com/basileb/custom_text_editor/settings"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func syntaxHighlightingC(node *tree_sitter.Node, code []byte, userStyle *st.WindowStyle) {
	cursor := node.Walk()
	defer cursor.Close()

	lastFinish := uint(0)
	textRenderCursor := &TextRenderCursor{
		line: userStyle.PaddingTop,
		row:  userStyle.PaddingLeft,
	}

	for i := uint32(0); !(cursor.Node().KindId() == 161 && i > 1); i++ {
		cursor.GotoDescendant(i)
		start, finish := cursor.Node().ByteRange()

		if start > lastFinish {
			stringText := string(code[lastFinish:start])
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Editor.Fg, userStyle)
		}

		text := code[start:finish]
		stringText := string(text)

		switch cursor.Node().KindId() {
		case 93: // Primitive type
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Tag, userStyle)

		case 1, 362: // Identifier
			if cursor.Node().Parent().KindId() == 230 { // Function name
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Func, userStyle)
			} else if cursor.Node().Parent().KindId() == 165 { // Defined word
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Func, userStyle)
			} else if cursor.Node().Parent().KindId() == 299 { // Called function -> can also be 266
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Func, userStyle)
			} else if cursor.Node().Parent().KindId() == 199 {
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Tag, userStyle)
			} else if cursor.Node().Parent().KindId() == 198 {
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Tag, userStyle)
			} else { // variable name
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Editor.Fg, userStyle)
			}

		case 205, 44, 96, 106, 98, 99, 100, 102, 101, 103, 104, 105, 107, 108, 109, 127: // Keywords
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Keyword, userStyle)

		case 146, 147, 153: // String content
			if cursor.Node().Parent().KindId() == 164 { // include string (local lib)
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.String, userStyle)
			} else {
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.String, userStyle)
			}

		case 7, 95: // Comma, column
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Comment, userStyle)

		case 154: // Escape sequence
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Regexp, userStyle)

		case 152: // Quotes
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.String, userStyle)

		case 5, 8: // Parenthesis
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Entity, userStyle)

		case 64, 65: // Curly brackets
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Entity, userStyle)

		case 70, 72: // Angle brackets
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Entity, userStyle)

		case 26, 33: // Pointers operators
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Operator, userStyle)

		case 140: // Pointers arrow (struct->item)
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Comment, userStyle)

		case 42: // ;
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Comment, userStyle)

		case 24, 25, 27, 28, 73, 118, 119, 126, 125: // Operators (arithmetic)
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Operator, userStyle)

		case 31, 32: // Operators bitwise
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Operator, userStyle)

		case 22, 35, 36, 37, 38, 39: // Operators (comparison)
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Operator, userStyle)

		case 30, 29: // and, or
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Entity, userStyle)

		case 141: // number literals
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Constant, userStyle)

		case 156, 157: // true / false constant
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Constant, userStyle)

		case 2, 4: // include, define
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Keyword, userStyle)

		case 155: // system lib string
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.String, userStyle)

		case 18: // preproc_arg = defined keyword
			textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Constant, userStyle)

		case 160: // !comments -> block comments can be on multiple lines. This is really bad
			if stringText[1] == '*' {
				lines := strings.Split(stringText, "\n")
				for i, s := range lines {
					if i != 0 {
						s += "\n"
					}
					textRenderCursor.DrawTextPart(&s, userStyle.ColorTheme.Syntax.Comment, userStyle)
				}
			} else {
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Syntax.Comment, userStyle)
			}

		default: // Other elements which are leaf
			if cursor.Node().ChildCount() == 0 {
				stringText := string(text)
				textRenderCursor.DrawTextPart(&stringText, userStyle.ColorTheme.Editor.Fg, userStyle)
			}
		}

		lastFinish = finish
	}

	if lastFinish < uint(len(code)) {
		fmt.Print(string(code[lastFinish:]))
	}
}
