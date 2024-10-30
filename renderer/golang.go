package renderer

import (
	"fmt"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func syntaxHighlightingGo(node *tree_sitter.Node, code []byte) {
	cursor := node.Walk()
	defer cursor.Close()

	lastFinish := uint(0)

	for i := uint32(0); !(cursor.Node().KindId() == 94 && i > 1); i++ {
		cursor.GotoDescendant(i)
		start, finish := cursor.Node().ByteRange()

		if start > lastFinish {
			fmt.Print(string(code[lastFinish:start]))
		}

		text := code[start:finish]

		switch cursor.Node().KindId() {
		case 216: // Primitive type
			printBlue(string(text))

		case 1: // Identifier
			if cursor.Node().Parent().KindId() == 106 { // Function name
				printMagenta(string(text))
			} else {
				printGreen(string(text))
			}

		case 9, 10:
			printMagenta(string(text))

		case 15: // func keyword
			printBlue(string(text))

		case 7: // dot operator
			printGreen(string(text))

		case 188: // String literal
			printYellow(string(text))

		default: // Other elements
			if cursor.Node().KindId() < 70 {
				printWhite(string(text))
			}
		}

		lastFinish = finish
	}

	if lastFinish < uint(len(code)) {
		fmt.Print(string(code[lastFinish:]))
	}
}
