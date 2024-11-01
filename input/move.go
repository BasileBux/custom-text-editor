package input

import (
	"strings"

	r "github.com/basileb/custom_text_editor/renderer"
	t "github.com/basileb/custom_text_editor/types"
	st "github.com/basileb/custom_text_editor/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func arrowLeft(text *[]string, nav *t.NavigationData, style *st.WindowStyle) {

	if nav.AbsoluteSelectedRow >= 1 {
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		}

		// control + right moves whole words
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

			if nav.ScrollOffset.X > 0 && nav.SelectedRow < int(nav.ScrollOffset.X+1+float32(style.Cursor.CursorHorizontalPadding)) {
				nav.ScrollOffset.X--
				// fmt.Println("Scroll left -> Scroll offset = ", nav.ScrollOffset.X)
			}
		}
	} else if nav.SelectedLine >= 1 {
		// when on left line end, go up end
		nav.SelectedLine--
		nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])

		r.ScrollUp(nav)
	}
	nav.SelectedRow = nav.AbsoluteSelectedRow
}

func arrowRight(text *[]string, nav *t.NavigationData, state *t.ProgramState, style *st.WindowStyle) {
	if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		}

		// control + right moves whole words
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

			if nav.AbsoluteSelectedRow > int(nav.ScrollOffset.X)+state.ViewPortSteps.X-4-int(style.Cursor.CursorHorizontalPadding) {
				nav.ScrollOffset.X++
			}

		}
	} else if nav.SelectedLine < len((*text))-1 {
		// when on right line end, go down and 0
		nav.SelectedLine++
		nav.AbsoluteSelectedRow = 0

		r.ScrollDown(nav, state)
	}
	nav.SelectedRow = nav.AbsoluteSelectedRow

}

func arrowUp(text *[]string, nav *t.NavigationData) {
	nav.SelectedLine--
	if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
		nav.SelectedRow = len((*text)[nav.SelectedLine])
	} else {
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	r.ScrollUp(nav)
}

func arrowDown(text *[]string, nav *t.NavigationData, state *t.ProgramState) {
	nav.SelectedLine++
	if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
		nav.SelectedRow = len((*text)[nav.SelectedLine])
	} else {
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	r.ScrollDown(nav, state)
}
