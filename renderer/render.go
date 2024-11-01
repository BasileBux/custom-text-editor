package renderer

import (
	"strings"

	st "github.com/basileb/custom_text_editor/settings"
	t "github.com/basileb/custom_text_editor/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextRenderCursor struct {
	line         float32 // pixels
	row          float32 // pixels
	scrollOffset rl.Vector2
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

func (t *TextRenderCursor) DrawTextPart(text *string, color rl.Color, state *t.ProgramState, style *st.WindowStyle) bool {

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
		t.line += textSize.Y + style.FontSpacing
		t.row = style.PaddingLeft
	}

	scrollHeight := (t.scrollOffset.Y * style.CharSize.Y) + (t.scrollOffset.Y * style.FontSpacing)
	scrollWidth := (t.scrollOffset.X * style.CharSize.X) + (t.scrollOffset.X * style.FontSpacing)
	textPos := rl.NewVector2(t.row-scrollHeight, t.line-scrollWidth)

	if textPos.Y > 0 || textPos.Y < -(textSize.Y/2) { // small optimization
		rl.DrawTextEx(style.Font, *text, textPos, style.FontSize, 1, color)
	}
	t.row += textSize.X + style.FontSpacing

	if t.line > (state.ViewPortSize.Y*style.CharSize.Y + t.scrollOffset.Y) {
		return true
	}

	return false
}

func noSyntaxHighlight(text *string, userStyle *st.WindowStyle, scrollOffset *rl.Vector2, style *st.WindowStyle) {
	scrollHeight := (scrollOffset.Y * style.CharSize.Y) + (scrollOffset.Y * style.FontSpacing)
	scrollWidth := (scrollOffset.X * style.CharSize.X) + (scrollOffset.X * style.FontSpacing)
	textPos := rl.NewVector2(userStyle.PaddingLeft-scrollWidth, userStyle.PaddingTop-scrollHeight)
	rl.DrawTextEx(userStyle.Font, *text, textPos, userStyle.FontSize, 1, userStyle.ColorTheme.Editor.Fg)
}
