package input

import (
	r "github.com/basileb/kenzan/renderer"
	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
)

func backSpace(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav
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

		r.ResetHorizontalScrollRight(float32(nav.AbsoluteSelectedRow), state, style)
		r.ScrollUp(1, nav, style)
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
		r.ScrollLeft(1, nav, style)

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

		// Scroll one up and go at end of line
		r.ResetHorizontalScrollRight(float32(nav.AbsoluteSelectedRow), state, style)
		r.ScrollUp(1, nav, style)
	}
}
