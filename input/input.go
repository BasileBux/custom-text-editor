package input

import (
	"fmt"
	"strings"

	f "github.com/basileb/kenzan/files"
	r "github.com/basileb/kenzan/renderer"
	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func lastNonSpaceCharIndex(s string) int {
	lastIdx := -1
	for i, c := range s {
		if c >= 32 && c <= 126 {
			lastIdx = i
		}
	}
	return lastIdx
}

func InputManager(text *[]string, state *t.ProgramState, style *st.WindowStyle) bool {
	nav := state.Nav
	char := rl.GetCharPressed()
	for char > 0 {
		// refuse non ascii and non printable chars
		if char >= 32 && char <= 126 {
			state.SaveState = false
			state.ForceQuit = false
			state.Update.SyntaxHighlight = true
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

				// Scroll right if needed
				r.ScrollRight(1, state, style)
			}
		}
		char = rl.GetCharPressed()
	}

	// Save
	if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyS) {
		err := f.WriteFile(state.AcitveFile, *text)
		if err != nil {
			fmt.Println("Couldn't save file")
		} else {
			state.SaveState = true
			state.SavedFile = make([]string, len(*text))
			copy(state.SavedFile, *text)
		}
	}

	// Backspace
	if rl.IsKeyPressedRepeat(rl.KeyBackspace) || rl.IsKeyPressed(rl.KeyBackspace) {
		state.Update.SyntaxHighlight = true
		state.SaveState = false
		state.ForceQuit = false
		backSpace(text, state, style)
	}

	// Enter
	if rl.IsKeyPressedRepeat(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyEnter) {
		if state.ForceQuit {
			return true
		}
		state.Update.SyntaxHighlight = true
		state.SaveState = false
		newText := make([]string, len(*text)+1)
		copy(newText, (*text)[:nav.SelectedLine+1])
		copy(newText[nav.SelectedLine+2:], (*text)[nav.SelectedLine+1:])
		*text = newText

		remainingString := (*text)[nav.SelectedLine][nav.SelectedRow:]
		if len(remainingString) != 0 {
			(*text)[nav.SelectedLine] = (*text)[nav.SelectedLine][:nav.SelectedRow]
			(*text)[nav.SelectedLine+1] += remainingString
		}

		nav.SelectedLine++
		nav.AbsoluteSelectedRow = 0
		nav.SelectedRow = nav.AbsoluteSelectedRow

		// Scroll down and reset horizontal scroll
		r.ScrollDown(1, state, style)
		nav.ScrollOffset.X = 0

		// Check and change width of line numbers gutter if max nb changes
		newFileNb := fmt.Sprintf("%d", len(*text)-1)
		if len(newFileNb) > int(state.Cache.LineNumbers.Width) {
			newNbSize := rl.MeasureTextEx(style.Font, newFileNb, style.FontSize, style.FontSpacing)
			style.PaddingLeft -= float32(state.Cache.LineNumbers.Width)
			style.PaddingLeft += newNbSize.X
		}
	}

	// Tab
	if rl.IsKeyPressed(rl.KeyTab) {
		state.Update.SyntaxHighlight = true
		state.SaveState = false
		state.ForceQuit = false
		begin := (*text)[nav.SelectedLine][:nav.SelectedRow]
		end := (*text)[nav.SelectedLine][nav.SelectedRow:]
		(*text)[nav.SelectedLine] = begin + strings.Repeat(" ", 4) + end
		nav.AbsoluteSelectedRow += 4
		nav.SelectedRow = nav.AbsoluteSelectedRow
		r.ScrollRight(4, state, style)
	}

	// Left
	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft) {
		state.Update.Cursor = true
		state.ForceQuit = false
		arrowLeft(text, state, style)
	}

	// Right
	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight) {
		state.Update.Cursor = true
		state.ForceQuit = false
		arrowRight(text, state, style)
	}

	// Up
	if (rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressedRepeat(rl.KeyUp)) && nav.SelectedLine >= 1 {
		state.Update.Cursor = true
		state.ForceQuit = false
		arrowUp(text, state, style)
	}

	// Down
	if (rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressedRepeat(rl.KeyDown)) && nav.SelectedLine < len(*text)-1 {
		state.Update.Cursor = true
		state.ForceQuit = false
		arrowDown(text, state, style)
	}
	return false
}
