package renderer

import (
	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
)

// Arbitrary padding to allign scroll to side
const UP_PADDING int = -1
const DOWN_PADDING int = 0
const LEFT_PADDING int = -1
const RIGHT_PADDING int = 0

func ResetHorizontalScrollRight(lineSize float32, state *t.ProgramState, style *st.WindowStyle) {
	if lineSize > float32(state.ViewPortSteps.X)-float32(RIGHT_PADDING) {
		state.Nav.ScrollOffset.X = lineSize - float32(state.ViewPortSteps.X) + float32(RIGHT_PADDING) + float32(style.Cursor.HorizontalPadding)
		state.Update.Highlight = true
	}
}

func ScrollLeft(size int, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav
	if nav.ScrollOffset.X > float32(size-1) {
		if nav.SelectedRow < int(nav.ScrollOffset.X+float32(LEFT_PADDING)+float32(style.Cursor.HorizontalPadding)) {
			nav.ScrollOffset.X -= float32(size)
			state.Update.Highlight = true
		}
	} else {
		nav.ScrollOffset.X = 0
		state.Update.Highlight = true
	}
}

func ScrollRight(size int, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav

	if nav.AbsoluteSelectedRow > int(nav.ScrollOffset.X)+state.ViewPortSteps.X-RIGHT_PADDING-int(style.Cursor.HorizontalPadding) {
		nav.ScrollOffset.X += float32(size)
		state.Update.Highlight = true
	}
}

func ScrollUp(size int, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav
	if int(nav.ScrollOffset.Y) > (size) {
		if nav.SelectedLine < int(nav.ScrollOffset.Y)+int(style.Cursor.VerticalPadding)+UP_PADDING {
			nav.ScrollOffset.Y -= float32(size)
			state.Update.Highlight = true
		}
	} else {
		nav.ScrollOffset.Y = 0
		state.Update.Highlight = true
	}
}

func ScrollDown(size int, state *t.ProgramState, style *st.WindowStyle) {
	nav := state.Nav
	if nav.SelectedLine > int(nav.ScrollOffset.Y)+state.ViewPortSteps.Y-DOWN_PADDING-int(style.Cursor.VerticalPadding) {
		nav.ScrollOffset.Y += float32(size)
		state.Update.Highlight = true
	}
}
