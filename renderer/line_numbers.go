package renderer

import (
	"fmt"

	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
Modifications needed:
	- Deprecate width and calculate regarding the max width
		- Or maybe make it min width ?
	- Calculate padding on both sides
	- Color current line
	- Implement relative line numbers
*/

func RenderLineNumbers(relative bool, width int, padding int, state *t.ProgramState, style *st.WindowStyle) {
	if relative {
		relativeNumbering(width, padding, state, style)
	} else {
		absoluteNumbering(width, padding, state, style)
	}
}

func absoluteNumbering(width int, padding int, state *t.ProgramState, style *st.WindowStyle) {
	// Clear background
	rl.DrawRectangle(0, 0, int32(width), int32(state.ViewPortSize.Y), style.ColorTheme.Editor.Bg)

	previousY := style.PaddingTop
	for i := range state.ViewPortSteps.Y + 1 {
		lineNb := fmt.Sprintf("%d", i+int(state.Nav.ScrollOffset.Y))
		nbHeight := rl.MeasureTextEx(style.Font, lineNb, style.FontSize, style.FontSpacing)
		pos := rl.Vector2{
			X: float32(padding),
			Y: previousY,
		}
		rl.DrawTextEx(style.Font, lineNb, pos, style.FontSize, style.FontSpacing, style.ColorTheme.Syntax.Comment)
		previousY += nbHeight.Y + style.FontSpacing
	}

	// Draw nice delimitation line
	rl.DrawLine(int32(width), 0, int32(width), int32(state.ViewPortSize.Y), style.ColorTheme.Syntax.Comment)
}

func relativeNumbering(width int, padding int, state *t.ProgramState, style *st.WindowStyle) {
}
