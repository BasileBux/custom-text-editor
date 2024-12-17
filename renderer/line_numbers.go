package renderer

import (
	"fmt"

	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderLineNumbers(state *t.ProgramState, style *st.WindowStyle) {
	width := state.Cache.LineNumbers.Width +
		int32(style.LineNumbers.PaddingLeft) + int32(style.LineNumbers.PaddingRight)
	rl.DrawRectangle(0, 0, width, int32(state.ViewPortSize.Y), style.ColorTheme.Editor.Bg)
	for i, pos := range state.Cache.LineNumbers.Positions {
		rl.DrawTextEx(style.Font, state.Cache.LineNumbers.Numbers[i], pos,
			style.FontSize, style.FontSpacing, state.Cache.LineNumbers.Colors[i])
	}
	rl.DrawRectangle(width, 0, int32(style.LineNumbers.LineWidth),
		int32(state.ViewPortSize.Y), style.ColorTheme.Editor.Gutter.Normal)
}

func CalculateLineNbPositions(relative bool, state *t.ProgramState, style *st.WindowStyle) {
	if !relative {
		calculateAbsLineNbPositions(style.LineNumbers.PaddingLeft,
			style.LineNumbers.PaddingRight, state, style)
	} else {
		calculateRelLineNbPositions(style.LineNumbers.PaddingLeft,
			style.LineNumbers.PaddingRight, state, style)
	}
}

func calculateAbsLineNbPositions(paddingL int, paddingR int, state *t.ProgramState, style *st.WindowStyle) {
	state.Cache.LineNumbers.Positions = make([]rl.Vector2, 0) // empty slice
	state.Cache.LineNumbers.Colors = make([]rl.Color, 0)      // empty slice
	state.Cache.LineNumbers.Numbers = make([]string, 0)       // empty slice

	width := state.Cache.LineNumbers.Width + int32(paddingL) + int32(paddingR)
	Ypos := style.PaddingTop

	offset := 0
	if state.Nav.ScrollOffset.Y > 0 {
		nbSize := rl.MeasureTextEx(style.Font, "0", style.FontSize, style.FontSpacing)
		Ypos -= nbSize.Y + style.FontSpacing
		offset = 1
	}

	for i := range state.ViewPortSteps.Y + 2 {
		index := int(state.Nav.ScrollOffset.Y) - offset + 1
		lineNb := fmt.Sprintf("%d", i+index)
		nbSize := rl.MeasureTextEx(style.Font, lineNb, style.FontSize, style.FontSpacing)
		Xpos := width - int32(paddingR) - int32(nbSize.X)
		pos := rl.Vector2{
			X: float32(Xpos),
			Y: Ypos,
		}

		currentColor := style.ColorTheme.Editor.Gutter.Normal
		if i+index == state.Nav.SelectedLine+1 {
			currentColor = style.ColorTheme.Editor.Gutter.Active

			if style.LineNumbers.OffsetCurrent && i+index < 100 {
				pos.X -= nbSize.X / float32(len(lineNb))
			}
		}

		state.Cache.LineNumbers.Positions = append(state.Cache.LineNumbers.Positions, pos)
		state.Cache.LineNumbers.Colors = append(state.Cache.LineNumbers.Colors, currentColor)

		if i+index >= len(state.SavedFile) {
			lineNb = "~"
		}
		state.Cache.LineNumbers.Numbers = append(state.Cache.LineNumbers.Numbers, lineNb)
		Ypos += nbSize.Y + style.FontSpacing
	}
}

func calculateRelLineNbPositions(paddingL int, paddingR int, state *t.ProgramState, style *st.WindowStyle) {
	state.Cache.LineNumbers.Positions = make([]rl.Vector2, 0) // empty slice
	state.Cache.LineNumbers.Colors = make([]rl.Color, 0)      // empty slice
	state.Cache.LineNumbers.Numbers = make([]string, 0)       // empty slice

	width := state.Cache.LineNumbers.Width + int32(paddingL) + int32(paddingR)
	Ypos := style.PaddingTop

	offset := 0
	if state.Nav.ScrollOffset.Y > 0 {
		nbSize := rl.MeasureTextEx(style.Font, "0", style.FontSize, style.FontSpacing)
		Ypos -= nbSize.Y + style.FontSpacing
		offset = 1
	}

	index := (int(state.Nav.ScrollOffset.Y)-state.Nav.SelectedLine)*-1 + offset
	for range state.ViewPortSteps.Y + 2 {
		var lineNb string
		if index > 0 {
			lineNb = fmt.Sprintf("%d", index)
		} else if index == 0 {
			lineNb = fmt.Sprintf("%d", state.Nav.SelectedLine+1)
		} else {
			lineNb = fmt.Sprintf("%d", (index)*-1)
		}
		index--
		nbSize := rl.MeasureTextEx(style.Font, lineNb, style.FontSize, style.FontSpacing)
		Xpos := width - int32(paddingR) - int32(nbSize.X)
		pos := rl.Vector2{
			X: float32(Xpos),
			Y: Ypos,
		}

		currentColor := style.ColorTheme.Editor.Gutter.Normal

		if index == -1 {
			if style.LineNumbers.OffsetCurrent && state.Nav.SelectedLine < 100 {
				pos.X -= nbSize.X / float32(len(lineNb))
			}
			currentColor = style.ColorTheme.Editor.Gutter.Active
		}

		state.Cache.LineNumbers.Positions = append(state.Cache.LineNumbers.Positions, pos)
		state.Cache.LineNumbers.Colors = append(state.Cache.LineNumbers.Colors, currentColor)
		state.Cache.LineNumbers.Numbers = append(state.Cache.LineNumbers.Numbers, lineNb)
		Ypos += nbSize.Y + style.FontSpacing
	}
}
