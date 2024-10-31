package input

import (
	"fmt"
	"strings"

	f "github.com/basileb/custom_text_editor/files"
	r "github.com/basileb/custom_text_editor/renderer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NavigationData struct {
	SelectedLine        int // 0 indexed
	AbsoluteSelectedRow int // 0 indexed, number of characters depends on nothing
	SelectedRow         int // 0 indexed, number of characters depends on current line
}

type ProgramState struct {
	Nav            *NavigationData
	AcitveFile     string
	RenderUpdate   bool
	ActiveLanguage r.Language
	SavedFile      []string
	SaveState      bool
	ForceQuit      bool
}

func lastNonSpaceCharIndex(s string) int {
	lastIdx := -1
	for i, c := range s {
		if c >= 32 && c <= 126 {
			lastIdx = i
		}
	}
	return lastIdx
}

func InputManager(text *[]string, nav *NavigationData, state *ProgramState) bool {
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

	if rl.IsKeyPressedRepeat(rl.KeyBackspace) || rl.IsKeyPressed(rl.KeyBackspace) {
		state.RenderUpdate = true
		state.SaveState = false
		state.ForceQuit = false
		backSpace(text, nav)
	}

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
	}

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

	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft) {
		state.RenderUpdate = true
		state.ForceQuit = false
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

					if (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] == ' ' && (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] >= 32 && (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] <= 126 {
						for {
							if (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] == ' ' {
								nav.AbsoluteSelectedRow--
							} else {
								nav.AbsoluteSelectedRow++
								break
							}
						}
					}
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
		state.RenderUpdate = true
		state.ForceQuit = false
		if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
			if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
				nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			}
			if rl.IsKeyDown(rl.KeyLeftControl) {
				jumpTo := strings.Index((*text)[nav.SelectedLine][nav.AbsoluteSelectedRow+1:], " ")
				if jumpTo == -1 {
					nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
				} else {
					nav.AbsoluteSelectedRow = jumpTo + nav.AbsoluteSelectedRow + 1
					for {
						if (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] == ' ' {
							if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine])-1 {
								nav.AbsoluteSelectedRow++
							} else {
								nav.AbsoluteSelectedRow++
								break
							}
						} else {
							break
						}
					}
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
		state.RenderUpdate = true
		state.ForceQuit = false
		nav.SelectedLine--
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
			nav.SelectedRow = len((*text)[nav.SelectedLine])
		} else {
			nav.SelectedRow = nav.AbsoluteSelectedRow
		}
	}

	if (rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressedRepeat(rl.KeyDown)) && nav.SelectedLine < len(*text)-1 {
		state.RenderUpdate = true
		state.ForceQuit = false
		nav.SelectedLine++
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
			nav.SelectedRow = len((*text)[nav.SelectedLine])
		} else {
			nav.SelectedRow = nav.AbsoluteSelectedRow
		}
	}
	return false
}

func backSpace(text *[]string, nav *NavigationData) {
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
