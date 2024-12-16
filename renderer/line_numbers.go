package renderer

import (
	"fmt"

	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"

	// u "github.com/basileb/kenzan/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO:
// - Make line width an setting (can be 0)

func RenderLineNumbers(paddingL int, paddingR int, state *t.ProgramState, style *st.WindowStyle) {
	drawLineNumbers(paddingL, paddingR, state, style)
}

func drawLineNumbers(paddingL int, paddingR int, state *t.ProgramState, style *st.WindowStyle) {
	width := state.Cache.LineNumbers.Width + int32(paddingL) + int32(paddingR)
	rl.DrawRectangle(0, 0, width, int32(state.ViewPortSize.Y), style.ColorTheme.Editor.Bg)
	for i, pos := range state.Cache.LineNumbers.Positions {
		rl.DrawTextEx(style.Font, state.Cache.LineNumbers.Numbers[i], pos,
			style.FontSize, style.FontSpacing, state.Cache.LineNumbers.Colors[i])
	}
	// 2 == line width
	rl.DrawRectangle(width, 0, 2, int32(state.ViewPortSize.Y), style.ColorTheme.Editor.Gutter.Normal)
}

func CalculateLineNbPositions(relative bool, paddingL int, paddingR int, state *t.ProgramState, style *st.WindowStyle) {
	if !relative {
		calculateAbsLineNbPositions(paddingL, paddingR, state, style)
	} else {
		calculateRelLineNbPositions(paddingL, paddingR, state, style)
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
		lineNb := fmt.Sprintf("%d", i+int(state.Nav.ScrollOffset.Y)-offset)
		nbSize := rl.MeasureTextEx(style.Font, lineNb, style.FontSize, style.FontSpacing)
		Xpos := width - int32(paddingR) - int32(nbSize.X)
		pos := rl.Vector2{
			X: float32(Xpos),
			Y: Ypos,
		}
		state.Cache.LineNumbers.Positions = append(state.Cache.LineNumbers.Positions, pos)

		currentColor := style.ColorTheme.Editor.Gutter.Normal
		if i+int(state.Nav.ScrollOffset.Y)-offset == state.Nav.SelectedLine {
			currentColor = style.ColorTheme.Editor.Gutter.Active
		}
		state.Cache.LineNumbers.Colors = append(state.Cache.LineNumbers.Colors, currentColor)

		if i+int(state.Nav.ScrollOffset.Y)-offset >= len(state.SavedFile) {
			lineNb = "~"
		}
		state.Cache.LineNumbers.Numbers = append(state.Cache.LineNumbers.Numbers, lineNb)
		Ypos += nbSize.Y + style.FontSpacing
	}
}

// Code can be reafactored easily. The only thing changing is the number to display
// Right now, there is a ton of copy paste but it works tho
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

	var index int
	for index = (int(state.Nav.ScrollOffset.Y)-state.Nav.SelectedLine)*-1 + offset; index > 0; index-- {
		lineNb := fmt.Sprintf("%d", index)
		nbSize := rl.MeasureTextEx(style.Font, lineNb, style.FontSize, style.FontSpacing)
		Xpos := width - int32(paddingR) - int32(nbSize.X)
		pos := rl.Vector2{
			X: float32(Xpos),
			Y: Ypos,
		}
		state.Cache.LineNumbers.Positions = append(state.Cache.LineNumbers.Positions, pos)
		currentColor := style.ColorTheme.Editor.Gutter.Normal
		state.Cache.LineNumbers.Colors = append(state.Cache.LineNumbers.Colors, currentColor)
		state.Cache.LineNumbers.Numbers = append(state.Cache.LineNumbers.Numbers, lineNb)
		Ypos += nbSize.Y + style.FontSpacing
	}

	lineNb := fmt.Sprintf("%d", state.Nav.SelectedLine)
	nbSize := rl.MeasureTextEx(style.Font, lineNb, style.FontSize, style.FontSpacing)
	Xpos := width - int32(paddingR) - int32(nbSize.X)
	pos := rl.Vector2{
		X: float32(Xpos),
		Y: Ypos,
	}
	state.Cache.LineNumbers.Positions = append(state.Cache.LineNumbers.Positions, pos)
	currentColor := style.ColorTheme.Editor.Gutter.Active
	state.Cache.LineNumbers.Colors = append(state.Cache.LineNumbers.Colors, currentColor)
	state.Cache.LineNumbers.Numbers = append(state.Cache.LineNumbers.Numbers, lineNb)
	Ypos += nbSize.Y + style.FontSpacing

	for i := 1; i < state.ViewPortSteps.Y+1-index; i++ {
		lineNb := fmt.Sprintf("%d", i)
		nbSize := rl.MeasureTextEx(style.Font, lineNb, style.FontSize, style.FontSpacing)
		Xpos := width - int32(paddingR) - int32(nbSize.X)
		pos := rl.Vector2{
			X: float32(Xpos),
			Y: Ypos,
		}
		state.Cache.LineNumbers.Positions = append(state.Cache.LineNumbers.Positions, pos)
		currentColor := style.ColorTheme.Editor.Gutter.Normal
		state.Cache.LineNumbers.Colors = append(state.Cache.LineNumbers.Colors, currentColor)
		state.Cache.LineNumbers.Numbers = append(state.Cache.LineNumbers.Numbers, lineNb)
		Ypos += nbSize.Y + style.FontSpacing
	}
}
