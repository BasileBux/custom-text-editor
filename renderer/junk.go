package renderer

import (
	"fmt"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
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
		cursor.GotoDescendant(i)
		start, finish := cursor.Node().ByteRange()
		text := code[start:finish]
		text = []byte(strings.TrimSuffix(string(text), "\n"))
		fmt.Println(string(text), " -> ", cursor.Node().Kind(), " -> ", cursor.Node().KindId())
		// if i > 100 {
		// 	return
		// }
	}

}
