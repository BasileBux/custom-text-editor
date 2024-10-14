package main

import (
	"fmt"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Theme struct {
	Background rl.Color
	Text       rl.Color
}

var darkTheme Theme = Theme{
	Background: rl.Black,
	Text:       rl.RayWhite,
}

var lightTheme Theme = Theme{
	Background: rl.RayWhite,
	Text:       rl.Black,
}

type WindowStyle struct {
	PaddingTop    float32
	PaddingRight  float32
	PaddingBottom float32
	PaddingLeft   float32
	Font          rl.Font
	FontSize      float32
	CursorOffset  int32 // horizontal distance to text
	CursorWidth   int32
	cursorRatio   float32 // ratio with the text height
	ColorTheme    Theme
}

var compact WindowStyle = WindowStyle{
	PaddingTop:    13.0,
	PaddingRight:  13.0,
	PaddingBottom: 13.0,
	PaddingLeft:   13.0,
	Font:          rl.Font{},
	FontSize:      30,
	CursorOffset:  -2,
	CursorWidth:   1,
	cursorRatio:   1,
	ColorTheme:    darkTheme,
}

type NavigationData struct {
	SelectedLine        int // 0 indexed
	AbsoluteSelectedRow int // 0 indexed, number of characters depends on nothing
	SelectedRow         int // 0 indexed, number of characters depends on current line
}

func inputManager(text *[]string, nav *NavigationData) {
	char := rl.GetCharPressed()
	for char > 0 {
		// refuse non ascii and non printable chars
		if char >= 32 && char <= 126 {
			if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
				nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			}
			if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
				(*text)[nav.SelectedLine] = (*text)[nav.SelectedLine][:nav.AbsoluteSelectedRow] + string(rune(char)) + (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow:]
				nav.AbsoluteSelectedRow++
				nav.SelectedRow = nav.AbsoluteSelectedRow
			} else {
				(*text)[nav.SelectedLine] += string(rune(char))
				nav.AbsoluteSelectedRow++
				nav.SelectedRow = nav.AbsoluteSelectedRow
			}
		}
		char = rl.GetCharPressed()
	}

	if rl.IsKeyPressedRepeat(rl.KeyBackspace) || rl.IsKeyPressed(rl.KeyBackspace) {
		if len((*text)[nav.SelectedLine]) <= 0 && nav.SelectedLine > 0 {
			nav.SelectedLine--
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			nav.SelectedRow = nav.AbsoluteSelectedRow
			return
		}

		if len((*text)[nav.SelectedLine]) >= 1 && nav.AbsoluteSelectedRow > 0 {
			if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
				nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			}
			if nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
				(*text)[nav.SelectedLine] = (*text)[nav.SelectedLine][:nav.AbsoluteSelectedRow-1] + (*text)[nav.SelectedLine][nav.AbsoluteSelectedRow:]
				nav.AbsoluteSelectedRow--
				nav.SelectedRow = nav.AbsoluteSelectedRow
			} else {
				(*text)[nav.SelectedLine] = (*text)[nav.SelectedLine][:len((*text)[nav.SelectedLine])-1]
				nav.AbsoluteSelectedRow--
				nav.SelectedRow = nav.AbsoluteSelectedRow
			}

		} else if nav.SelectedLine > 0 {
			remaining := (*text)[nav.SelectedLine]
			newText := make([]string, len(*text)-1)
			copy(newText, (*text)[:nav.SelectedLine-1])
			copy(newText[nav.SelectedLine-1:], (*text)[nav.SelectedLine-1:])
			*text = newText
			nav.SelectedLine--
			(*text)[nav.SelectedLine] += remaining
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine]) - len(remaining)
			nav.SelectedRow = nav.AbsoluteSelectedRow
		}
	}

	if rl.IsKeyPressedRepeat(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyEnter) {
		nav.SelectedLine++
		if len((*text)) <= nav.SelectedLine {
			*text = append(*text, "")
			nav.AbsoluteSelectedRow = 0
			nav.SelectedRow = nav.AbsoluteSelectedRow
		} else {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
			nav.SelectedRow = nav.AbsoluteSelectedRow

		}
	}

	if (rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft)) && nav.AbsoluteSelectedRow >= 1 {
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		}
		nav.AbsoluteSelectedRow--
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	if (rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight)) && nav.AbsoluteSelectedRow < len((*text)[nav.SelectedLine]) {
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine]) {
			nav.AbsoluteSelectedRow = len((*text)[nav.SelectedLine])
		}
		nav.AbsoluteSelectedRow++
		nav.SelectedRow = nav.AbsoluteSelectedRow
	}

	if (rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressedRepeat(rl.KeyUp)) && nav.SelectedLine >= 1 {
		nav.SelectedLine--
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
			nav.SelectedRow = len((*text)[nav.SelectedLine])
		} else {
			nav.SelectedRow = nav.AbsoluteSelectedRow
		}
	}

	if (rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressedRepeat(rl.KeyDown)) && nav.SelectedLine < len(*text)-1 {
		nav.SelectedLine++
		if nav.AbsoluteSelectedRow > len((*text)[nav.SelectedLine])-1 {
			nav.SelectedRow = len((*text)[nav.SelectedLine])
		} else {
			nav.SelectedRow = nav.AbsoluteSelectedRow
		}
	}

}

func main() {
	logFile, err := os.OpenFile("raylib.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close()
	rl.SetTraceLogCallback(func(msgType int, text string) {
		fmt.Fprintf(logFile, "%s\n", text)
	})

	rl.InitWindow(800, 800, "My custom text editor")
	if !rl.IsWindowReady() {
		log.Panic("Window didn't open correctly ???")
	}
	defer rl.CloseWindow()

	userStyle := compact
	// userStyle.Font = rl.LoadFontEx("/usr/share/fonts/GeistMono/GeistMonoNerdFont-Regular.otf", 100, nil)
	userStyle.Font = rl.LoadFontEx("/home/basileb/.local/share/fonts/GeistMono/GeistMonoNerdFont-Regular.otf", 100, nil)
	rl.SetTextLineSpacing(1)
	rl.SetTextureFilter(userStyle.Font.Texture, rl.FilterBilinear)
	rl.SetTargetFPS(144)

	var userText []string
	userText = append(userText, "")
	nav := NavigationData{
		SelectedLine: 0,
	}

	for !rl.WindowShouldClose() {

		inputManager(&userText, &nav)

		rl.BeginDrawing()

		rl.ClearBackground(userStyle.ColorTheme.Background)

		textPos := rl.NewVector2(userStyle.PaddingLeft, userStyle.PaddingTop)
		var textToRender string

		for _, l := range userText {
			textToRender += l
			textToRender += "\n"
		}
		rl.DrawTextEx(userStyle.Font, textToRender, textPos, userStyle.FontSize, 1, userStyle.ColorTheme.Text)

		// Draw cursor rectangle
		textSize := rl.MeasureTextEx(userStyle.Font, userText[nav.SelectedLine], userStyle.FontSize, 1)
		charSize := textSize.X / float32(len(userText[nav.SelectedLine]))

		var cursorHorizontalPos int32
		if len(userText[nav.SelectedLine]) <= 0 {
			cursorHorizontalPos = int32(userStyle.PaddingLeft)
		} else {
			cursorHorizontalPos = int32(charSize*float32(nav.SelectedRow)+charSize) + userStyle.CursorOffset
		}

		// Previous x: int32(userStyle.PaddingLeft)+int32(textSize.X)+int32(userStyle.CursorOffset)
		rl.DrawRectangle(cursorHorizontalPos, int32(userStyle.PaddingTop)+int32(nav.SelectedLine)*int32(textSize.Y)+int32(nav.SelectedLine+1), int32(userStyle.CursorWidth), int32(textSize.Y*userStyle.cursorRatio), rl.RayWhite)

		rl.EndDrawing()
	}
	rl.UnloadFont(userStyle.Font)
}
