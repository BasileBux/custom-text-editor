package renderer

import (
	st "github.com/basileb/custom_text_editor/settings"
	t "github.com/basileb/custom_text_editor/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawCursor(userText []string, nav *t.NavigationData, userStyle *st.WindowStyle) {
	textSize := rl.MeasureTextEx(userStyle.Font, userText[nav.SelectedLine], userStyle.FontSize, userStyle.FontSpacing)
	charSize := textSize.X / float32(len(userText[nav.SelectedLine]))

	var cursorHorizontalPos int32
	if len(userText[nav.SelectedLine]) <= 0 {
		cursorHorizontalPos = int32(userStyle.PaddingLeft) - int32(nav.ScrollOffset.X)*int32(textSize.X)
	} else {
		cursorHorizontalPos = int32(charSize*float32(nav.SelectedRow)+charSize) + userStyle.CursorOffset - int32(nav.ScrollOffset.X)*int32(textSize.X)
	}
	cursorVerticalPos := int32(userStyle.PaddingTop) + int32(nav.SelectedLine)*int32(textSize.Y) + int32(nav.SelectedLine+int(userStyle.FontSpacing)) - int32(nav.ScrollOffset.Y) - int32(nav.ScrollOffset.Y)*int32(textSize.Y)

	rl.DrawRectangle(cursorHorizontalPos, cursorVerticalPos, int32(userStyle.CursorWidth), int32(textSize.Y*userStyle.CursorRatio), userStyle.ColorTheme.Editor.Fg)
}
