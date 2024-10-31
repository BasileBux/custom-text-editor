package renderer

import (
	t "github.com/basileb/custom_text_editor/types"
)

func ScrollUp(nav *t.NavigationData) {
	if nav.ScrollOffset.Y > 0 && nav.SelectedLine < int(nav.ScrollOffset.Y) {
		nav.ScrollOffset.Y--
	}
}

func ScrollDown(nav *t.NavigationData, state *t.ProgramState) {
	if nav.SelectedLine > int(nav.ScrollOffset.Y)+state.ViewPortSteps.Y-2 {
		nav.ScrollOffset.Y++
	}
}
