package input

import (
	"fmt"
	"strings"

	f "github.com/basileb/custom_text_editor/files"
	st "github.com/basileb/custom_text_editor/settings"
	t "github.com/basileb/custom_text_editor/types"
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

func InputManager(text *[]string, nav *t.NavigationData, state *t.ProgramState, style *st.WindowStyle) bool {
	char := rl.GetCharPressed()
	for char > 0 {
		// refuse non ascii and non printable chars
		if char >= 32 && char <= 126 {
			state.SaveState = false
			state.ForceQuit = false
			state.RenderUpdate = true
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
		state.RenderUpdate = true
		state.SaveState = false
		state.ForceQuit = false
		backSpace(text, nav)
	}

	// Enter
	if rl.IsKeyPressedRepeat(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyEnter) {
		if state.ForceQuit {
			return true
		}
		state.RenderUpdate = true
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

		// Scroll down
		if nav.SelectedLine > int(nav.ScrollOffset.Y)+state.ViewPortSteps.Y-2 {
			nav.ScrollOffset.Y++
		}
	}

	// Tab
	if rl.IsKeyPressed(rl.KeyTab) {
		state.RenderUpdate = true
		state.SaveState = false
		state.ForceQuit = false
		begin := (*text)[nav.SelectedLine][:nav.SelectedRow]
		end := (*text)[nav.SelectedLine][nav.SelectedRow:]
		(*text)[nav.SelectedLine] = begin + strings.Repeat(" ", 4) + end
		nav.AbsoluteSelectedRow += 4
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	// Left
	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft) {
		state.RenderUpdate = true
		state.ForceQuit = false
		arrowLeft(text, nav, style)
	}

	// Right
	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight) {
		state.RenderUpdate = true
		state.ForceQuit = false
		arrowRight(text, nav, state, style)
	}

	// Up
	if (rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressedRepeat(rl.KeyUp)) && nav.SelectedLine >= 1 {
		state.RenderUpdate = true
		state.ForceQuit = false
		arrowUp(text, nav)
	}

	// Down
	if (rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressedRepeat(rl.KeyDown)) && nav.SelectedLine < len(*text)-1 {
		state.RenderUpdate = true
		state.ForceQuit = false
		arrowDown(text, nav, state)
	}
	return false
}
