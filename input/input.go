package input

import (
	"strings"

	t "github.com/basileb/custom_text_editor/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func InputManager(text *[]string, nav *t.NavigationData) {
	char := rl.GetCharPressed()
	for char > 0 {
		// refuse non ascii and non printable chars
		if char >= 32 && char <= 126 {
			if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
				nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			}
			if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
				(*text)[nav.SelectedLine] = (*text)[nav.SelectedLine][:nav.AbsoluteSelectedRow] + string(rune(char)) + (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow:]
				nav.AbsoluteSelectedRow++
				nav.SelectedRow = nav.AbsoluteSelectedRow
			} else {
				(*text)[nav.SelectedLine] += string(rune(char))
				nav.AbsoluteSelectedRow++
				nav.SelectedRow = nav.AbsoluteSelectedRow
			}
		}
		char = rl.GetCharPressed()
	}

	if rl.IsKeyPressedRepeat(rl.KeyBackspace) || rl.IsKeyPressed(rl.KeyBackspace) {
		backSpace(text, nav)
	}

	if rl.IsKeyPressedRepeat(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyEnter) {
		nav.SelectedLine++
		if len((*text)) <= nav.SelectedLine {
			*text = append(*text, "")
			nav.AbsoluteSelectedRow = 0
			nav.SelectedRow = nav.AbsoluteSelectedRow
		} else {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			nav.SelectedRow = nav.AbsoluteSelectedRow

		}
	}

	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft) {
		if nav.AbsoluteSelectedRow >= 1 {
			if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
				nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			}
			if rl.IsKeyDown(rl.KeyLeftControl) {
				jumpTo := strings.LastIndex((*text)[nav.SelectedLine][:nav.AbsoluteSelectedRow-1], " ")
				if jumpTo == -1 {
					nav.AbsoluteSelectedRow = 0
				} else {
					nav.AbsoluteSelectedRow = jumpTo + 1
				}
			} else {
				nav.AbsoluteSelectedRow--
			}
		} else if nav.SelectedLine >= 1 {
			// when on left line end, go up end
			nav.SelectedLine--
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		}
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight) {
		if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
			if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
				nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			}
			if rl.IsKeyDown(rl.KeyLeftControl) {
				jumpTo := strings.Index((*text)[nav.SelectedLine][nav.AbsoluteSelectedRow+1:], " ")
				if jumpTo == -1 {
					nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
				} else {
					nav.AbsoluteSelectedRow = jumpTo + nav.AbsoluteSelectedRow + 2
				}
			} else {
				nav.AbsoluteSelectedRow++
			}
		} else if nav.SelectedLine < len((*text))-1 {
			// when on right line end, go down and 0
			nav.SelectedLine++
			nav.AbsoluteSelectedRow = 0
		}
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	if (rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressedRepeat(rl.KeyUp)) && nav.SelectedLine >= 1 {
		nav.SelectedLine--
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
			nav.SelectedRow = len((*text)[nav.SelectedLine])
		} else {
			nav.SelectedRow = nav.AbsoluteSelectedRow
		}
	}

	if (rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressedRepeat(rl.KeyDown)) && nav.SelectedLine < len(*text)-1 {
		nav.SelectedLine++
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
			nav.SelectedRow = len((*text)[nav.SelectedLine])
		} else {
			nav.SelectedRow = nav.AbsoluteSelectedRow
		}
	}

}

func backSpace(text *[]string, nav *t.NavigationData) {
	// SelectedLine is not index 0 and deleting last char so going one up
	if len((*text)[nav.SelectedLine]) <= 0 && nav.SelectedLine > 0 {
		// remove line
		newText := make([]string, len(*text)-1)
		copy(newText, (*text)[:nav.SelectedLine])
		copy(newText[nav.SelectedLine:], (*text)[1+nav.SelectedLine:])
		*text = newText

		// move one up
		nav.SelectedLine--
		nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		nav.SelectedRow = nav.AbsoluteSelectedRow
		return
	}

	// Deleting inside and at the end of a non empty line anywhere
	if len((*text)[nav.SelectedLine]) >= 1 && nav.AbsoluteSelectedRow > 0 {
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		}

		// At the end
		if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
			(*text)[nav.SelectedLine] = (*text)[nav.SelectedLine][:nav.AbsoluteSelectedRow-1] + (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow:]
			nav.AbsoluteSelectedRow--
			nav.SelectedRow = nav.AbsoluteSelectedRow

		} else { // inside string
			(*text)[nav.SelectedLine] = (*text)[nav.SelectedLine][:len((*text)[nav.SelectedLine])-1]
			nav.AbsoluteSelectedRow--
			nav.SelectedRow = nav.AbsoluteSelectedRow
		}

		// inside and erasing last char
	} else if nav.SelectedLine > 0 {
		remaining := (*text)[nav.SelectedLine]
		// remove line
		newText := make([]string, len(*text)-1)
		copy(newText, (*text)[:nav.SelectedLine])
		copy(newText[nav.SelectedLine:], (*text)[1+nav.SelectedLine:])
		*text = newText
		// move and append remaining text to line up one
		nav.SelectedLine--
		(*text)[nav.SelectedLine] += remaining
		nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine]) - len(remaining)
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}
}
