package renderer

import (
	"strings"

	st "github.com/basileb/custom_text_editor/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextRenderCursor struct {
	line       float32 // pixels
	row        float32 // pixels
	widthReset bool
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

func (t *TextRenderCursor) DrawTextPart(text *string, color rl.Color, style *st.WindowStyle) {

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

	textPos := rl.NewVector2(t.row, t.line)
	rl.DrawTextEx(style.Font, *text, textPos, style.FontSize, 1, color)
	t.row += textSize.X + style.FontSpacing
}

func noSyntaxHighlight(text *string, userStyle *st.WindowStyle) {
	textPos := rl.NewVector2(userStyle.PaddingLeft, userStyle.PaddingTop)
	rl.DrawTextEx(userStyle.Font, *text, textPos, userStyle.FontSize, 1, userStyle.ColorTheme.Editor.Fg)
}
