package renderer

import (
	st "github.com/basileb/custom_text_editor/settings"
	t "github.com/basileb/custom_text_editor/types"
)

func ResetHorizontalScrollRight(lineSize float32, state *t.ProgramState, style *st.WindowStyle) {
	// +4 is a magic number. This won't adapt to highDPI settings which I want later
	if lineSize > float32(state.ViewPortSteps.X)-4 {
		state.Nav.ScrollOffset.X = lineSize - float32(state.ViewPortSteps.X) + 4 + float32(style.Cursor.HorizontalPadding)
	}
}

func ScrollLeft(size int, nav *t.NavigationData, style *st.WindowStyle) {
	if nav.ScrollOffset.X > float32(size-1) {
		// +1 is a magic number. This won't adapt to highDPI settings which I want later
		if nav.SelectedRow < int(nav.ScrollOffset.X+1+float32(style.Cursor.HorizontalPadding)) {
			nav.ScrollOffset.X -= float32(size)
		}
	} else {
		nav.ScrollOffset.X = 0
	}
}

func ScrollRight(size int, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav
	// -4 is a magic number. This won't adapt to highDPI settings which I want later
	if nav.AbsoluteSelectedRow > int(nav.ScrollOffset.X)+state.ViewPortSteps.X-4-int(style.Cursor.HorizontalPadding) {
		nav.ScrollOffset.X += float32(size)
	}
}

func ScrollUp(size int, nav *t.NavigationData, style *st.WindowStyle) {
	if int(nav.ScrollOffset.Y) > (size - 1) {
		if nav.SelectedLine < int(nav.ScrollOffset.Y)+int(style.Cursor.VerticalPadding) {
			nav.ScrollOffset.Y -= float32(size)
		}
	} else {
		nav.ScrollOffset.Y = 0
	}
}

func ScrollDown(size int, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav
	// -2 is a magic number. This won't adapt to highDPI settings which I want later
	if nav.SelectedLine > int(nav.ScrollOffset.Y)+state.ViewPortSteps.Y-2-int(style.Cursor.VerticalPadding) {
		nav.ScrollOffset.Y += float32(size)
	}
}
