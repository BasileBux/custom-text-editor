package renderer

import (
	"strings"

	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderHighlight(state *t.ProgramState, style *st.WindowStyle) {
	for _, c := range state.Cache.Syntax {
		scrollHeight := (state.Nav.ScrollOffset.Y * style.CharSize.Y) + (state.Nav.ScrollOffset.Y * style.FontSpacing)
		scrollWidth := (state.Nav.ScrollOffset.X * style.CharSize.X) + (state.Nav.ScrollOffset.X * style.FontSpacing)
		textPos := rl.NewVector2(c.Cursor.Row-scrollWidth, c.Cursor.Line-scrollHeight)

		if textPos.Y > 0 || textPos.Y < -(style.CharSize.Y/2) { // small optimization
			rl.DrawTextEx(style.Font, c.Text, textPos, style.FontSize, 1, *c.Color)
		}
		if c.Cursor.Stop {
			return
		}
	}
}

// Trailing spaces are spaces blocks followed by a '\n'
func removeTrailingSpaces(input string) string {
	var result []rune
	spaceBlock := []rune{}

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			spaceBlock = append(spaceBlock, ' ')
		} else {
			if input[i] == '\n' {
				spaceBlock = nil
			} else {
				result = append(result, spaceBlock...)
				spaceBlock = nil
			}
			result = append(result, rune(input[i]))
		}
	}
	result = append(result, spaceBlock...)
	return string(result)
}

func calculateOffset(cursor *t.TextRenderCursor, text *string, state *t.ProgramState, style *st.WindowStyle) t.TextRenderCursor {
	textSize := rl.MeasureTextEx(style.Font, *text, style.FontSize, style.FontSpacing)

	if strings.Contains(*text, "\n") {
		// Never render trailing spaces as their offset will transfer on new lines
		*text = removeTrailingSpaces(*text)
		// Remove last new line to have correct height
		lastNewline := strings.LastIndex(*text, "\n")
		begin := (*text)[:lastNewline]
		end := (*text)[lastNewline+1:]
		*text = begin + end

		textSize = rl.MeasureTextEx(style.Font, *text, style.FontSize, style.FontSpacing)
		cursor.Line += textSize.Y + style.FontSpacing
		cursor.Row = style.PaddingLeft
	}
	result := *cursor
	cursor.Row += textSize.X + style.FontSpacing
	if cursor.Line > (state.ViewPortSize.Y*style.CharSize.Y + state.Nav.ScrollOffset.Y) {
		result.Stop = true
	}
	return result
}

func noSyntaxHighlight(text *string, scrollOffset *rl.Vector2, style *st.WindowStyle) {
	scrollHeight := (scrollOffset.Y * style.CharSize.Y) + (scrollOffset.Y * style.FontSpacing)
	scrollWidth := (scrollOffset.X * style.CharSize.X) + (scrollOffset.X * style.FontSpacing)
	textPos := rl.NewVector2(style.PaddingLeft-scrollWidth, style.PaddingTop-scrollHeight)
	rl.DrawTextEx(style.Font, *text, textPos, style.FontSize, 1, style.ColorTheme.Editor.Fg)
}
