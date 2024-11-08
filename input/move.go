package input

import (
	"strings"

	r "github.com/basileb/custom_text_editor/renderer"
	st "github.com/basileb/custom_text_editor/settings"
	t "github.com/basileb/custom_text_editor/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func arrowLeft(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav
	if nav.AbsoluteSelectedRow >= 1 {
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		}

		// control + right moves whole words
		if rl.IsKeyDown(rl.KeyLeftControl) {
			jumpTo := strings.LastIndex((*text)[nav.SelectedLine][:nav.AbsoluteSelectedRow-1], " ")
			if jumpTo == -1 {
				nav.AbsoluteSelectedRow = 0
				nav.SelectedRow = nav.AbsoluteSelectedRow
				nav.ScrollOffset.X = 0
			} else {
				offset := nav.AbsoluteSelectedRow - (jumpTo + 1)
				nav.AbsoluteSelectedRow = jumpTo + 1
				nav.SelectedRow = nav.AbsoluteSelectedRow
				r.ScrollLeft(offset, nav, style)

				if (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] == ' ' && (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] >= 32 && (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] <= 126 {
					for {
						if (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] == ' ' {
							nav.AbsoluteSelectedRow--
							nav.SelectedRow = nav.AbsoluteSelectedRow
							r.ScrollLeft(1, nav, style)
						} else {
							nav.AbsoluteSelectedRow++
							nav.SelectedRow = nav.AbsoluteSelectedRow
							r.ScrollRight(1, state, style)
							break
						}
					}
				}
			}
		} else {
			nav.AbsoluteSelectedRow--
			nav.SelectedRow = nav.AbsoluteSelectedRow
			r.ScrollLeft(1, nav, style)
		}
	} else if nav.SelectedLine >= 1 {
		// when on left line end, go up end
		nav.SelectedLine--
		nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		nav.SelectedRow = nav.AbsoluteSelectedRow
		r.ResetHorizontalScrollRight(float32(nav.AbsoluteSelectedRow), state, style)
		r.ScrollUp(1, nav, style)
	}
}

func arrowRight(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav
	if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		}

		// control + right moves whole words
		if rl.IsKeyDown(rl.KeyLeftControl) {
			jumpTo := strings.Index((*text)[nav.SelectedLine][nav.AbsoluteSelectedRow+1:], " ")
			if jumpTo == -1 {
				offset := len((*text)[nav.SelectedLine]) - nav.AbsoluteSelectedRow
				nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
				r.ScrollRight(offset, state, style)
			} else {
				offset := jumpTo + nav.AbsoluteSelectedRow + 1 - nav.AbsoluteSelectedRow
				nav.AbsoluteSelectedRow = jumpTo + nav.AbsoluteSelectedRow + 1
				r.ScrollRight(offset, state, style)
				for {
					if (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow] == ' ' {
						if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine])-1 {
							nav.AbsoluteSelectedRow++
							r.ScrollRight(1, state, style)
						} else {
							nav.AbsoluteSelectedRow++
							r.ScrollRight(1, state, style)
							break
						}
					} else {
						break
					}
				}
			}
		} else {
			nav.AbsoluteSelectedRow++
			nav.SelectedRow = nav.AbsoluteSelectedRow
			r.ScrollRight(1, state, style)
		}
	} else if nav.SelectedLine < len((*text))-1 {
		// when on right line end, go down and 0
		nav.SelectedLine++
		nav.AbsoluteSelectedRow = 0
		nav.SelectedRow = nav.AbsoluteSelectedRow

		// going to begining of next line so reset X scroll offset and scroll down
		nav.ScrollOffset.X = 0
		r.ScrollDown(1, state, style)
	}
	nav.SelectedRow = nav.AbsoluteSelectedRow
}

func arrowUp(text *[]string, nav *t.NavigationData, style *st.WindowStyle) {
	nav.SelectedLine--
	if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
		nav.SelectedRow = len((*text)[nav.SelectedLine])
	} else {
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	r.ScrollUp(1, nav, style)
}

func arrowDown(text *[]string, nav *t.NavigationData, state *t.ProgramState, style *st.WindowStyle) {
	nav.SelectedLine++
	if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
		nav.SelectedRow = len((*text)[nav.SelectedLine])
	} else {
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	r.ScrollDown(1, state, style)
}
