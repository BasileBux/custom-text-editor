package main

import (
	"fmt"
	"os"

	// ts "github.com/basileb/custom_text_editor/treesitter"
	ts "github.com/basileb/custom_text_editor/treesitter"
	t "github.com/basileb/custom_text_editor/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RedirectLogs() {
	logFile, err := os.OpenFile("raylib.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logFile.Close()
	rl.SetTraceLogCallback(func(msgType int, text string) {
		fmt.Fprintf(logFile, "%s\n", text)
	})
}

func DrawCursor(userText []string, nav *t.NavigationData, userStyle *t.WindowStyle) {
	textSize := rl.MeasureTextEx(userStyle.Font, userText[nav.SelectedLine], userStyle.FontSize, 1)
	charSize := textSize.X / float32(len(userText[nav.SelectedLine]))

	var cursorHorizontalPos int32
	if len(userText[nav.SelectedLine]) <= 0 {
		cursorHorizontalPos = int32(userStyle.PaddingLeft)
	} else {
		cursorHorizontalPos = int32(charSize*float32(nav.SelectedRow)+charSize) + userStyle.CursorOffset
	}

	rl.DrawRectangle(cursorHorizontalPos, int32(userStyle.PaddingTop)+int32(nav.SelectedLine)*int32(textSize.Y)+int32(nav.SelectedLine+1), int32(userStyle.CursorWidth), int32(textSize.Y*userStyle.CursorRatio), userStyle.ColorTheme.Text)
}

// func main() {
// 	RedirectLogs()

// 	rl.InitWindow(800, 800, "My custom text editor")
// 	if !rl.IsWindowReady() {
// 		log.Panic("Window didn't open correctly ???")
// 	}
// 	defer rl.CloseWindow()

// 	userStyle := t.Compact
// 	// userStyle.Font = rl.LoadFontEx("/usr/share/fonts/GeistMono/GeistMonoNerdFont-Regular.otf", 100, nil)
// 	userStyle.Font = rl.LoadFontEx("/home/basileb/.local/share/fonts/GeistMono/GeistMonoNerdFont-Regular.otf", 100, nil)
// 	rl.SetTextLineSpacing(1)
// 	rl.SetTextureFilter(userStyle.Font.Texture, rl.FilterBilinear)
// 	rl.SetTargetFPS(144)

// 	var userText []string
// 	userText = append(userText, "")
// 	nav := t.NavigationData{
// 		SelectedLine: 0,
// 		SelectedRow:  0,
// 	}
// 	textPos := rl.NewVector2(userStyle.PaddingLeft, userStyle.PaddingTop)

// 	for !rl.WindowShouldClose() {

// 		input.InputManager(&userText, &nav)

// 		rl.BeginDrawing()
// 		rl.ClearBackground(userStyle.ColorTheme.Background)

// 		var textToRender string
// 		for _, l := range userText {
// 			textToRender += l
// 			textToRender += "\n"
// 		}
// 		rl.DrawTextEx(userStyle.Font, textToRender, textPos, userStyle.FontSize, 1, userStyle.ColorTheme.Text)

// 		DrawCursor(userText, &nav, &userStyle)

// 		rl.EndDrawing()
// 	}
// 	rl.UnloadFont(userStyle.Font)
// }

func main() {
	// myCode := "int main(int a, int b)\n{\n\treturn a + b;\n}"
	myCode := "func main() {\n\tvar oue int\n\tnon := a.b(c)\n\treturn true}"
	ts.ParseText(ts.GO, &myCode)
}
