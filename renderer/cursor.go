package renderer

import (
	"math"

	st "github.com/basileb/custom_text_editor/settings"
	t "github.com/basileb/custom_text_editor/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawCursor(userText []string, nav *t.NavigationData, userStyle *st.WindowStyle) {
	textSize := rl.MeasureTextEx(userStyle.Font, userText[nav.SelectedLine], userStyle.FontSize, userStyle.FontSpacing)
	charSize := textSize.X / float32(len(userText[nav.SelectedLine]))

	var cursorHorizontalPos int32
	if len(userText[nav.SelectedLine]) <= 0 {
		cursorHorizontalPos = int32(userStyle.PaddingLeft)
	} else {
		cursorHorizontalPos = int32(math.Floor(float64(charSize)*float64(nav.SelectedRow)+float64(charSize))) + userStyle.Cursor.CursorOffset - int32(math.Floor(float64(nav.ScrollOffset.X)*float64(charSize)))
	}
	cursorVerticalPos := int32(userStyle.PaddingTop) + int32(nav.SelectedLine)*int32(textSize.Y) + int32(nav.SelectedLine+int(userStyle.FontSpacing)) - int32(nav.ScrollOffset.Y)*int32(userStyle.FontSpacing) - int32(nav.ScrollOffset.Y)*int32(textSize.Y)

	rl.DrawRectangle(cursorHorizontalPos, cursorVerticalPos, int32(userStyle.Cursor.CursorWidth), int32(textSize.Y*userStyle.Cursor.CursorRatio), userStyle.ColorTheme.Editor.Fg)
}
