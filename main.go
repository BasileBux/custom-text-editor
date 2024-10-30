package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/basileb/custom_text_editor/input"
	r "github.com/basileb/custom_text_editor/renderer"
	st "github.com/basileb/custom_text_editor/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var renderUpdate bool = true

type ProgramState struct {
	Nav            input.NavigationData
	AcitveFile     string
	RenderUpdate   bool
	ActiveLanguage r.Language
}

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

func DrawCursor(userText []string, nav *input.NavigationData, userStyle *st.WindowStyle) {
	textSize := rl.MeasureTextEx(userStyle.Font, userText[nav.SelectedLine], userStyle.FontSize, userStyle.FontSpacing)
	charSize := textSize.X / float32(len(userText[nav.SelectedLine]))

	var cursorHorizontalPos int32
	if len(userText[nav.SelectedLine]) <= 0 {
		cursorHorizontalPos = int32(userStyle.PaddingLeft)
	} else {
		cursorHorizontalPos = int32(charSize*float32(nav.SelectedRow)+charSize) + userStyle.CursorOffset
	}

	rl.DrawRectangle(cursorHorizontalPos, int32(userStyle.PaddingTop)+int32(nav.SelectedLine)*int32(textSize.Y)+int32(nav.SelectedLine+int(userStyle.FontSpacing)), int32(userStyle.CursorWidth), int32(textSize.Y*userStyle.CursorRatio), userStyle.ColorTheme.Editor.Fg)
}

func main() {
	RedirectLogs()
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(800, 800, "My custom text editor")
	if !rl.IsWindowReady() {
		log.Panic("Window didn't open correctly ???")
	}
	defer rl.CloseWindow()

	userStyle := st.Compact
	themeName := "ayu-light"
	var err error
	userStyle.ColorTheme, err = st.GetColorThemeFromFileName(&themeName)
	if err != nil {
		fmt.Println("Error could not open color theme")
		return
	}
	// userStyle.Font = rl.LoadFontEx("/usr/share/fonts/GeistMono/GeistMonoNerdFont-Regular.otf", 100, nil)
	userStyle.Font = rl.LoadFontEx("/home/basileb/.local/share/fonts/GeistMono/GeistMonoNerdFont-Regular.otf", 100, nil)

	rl.SetTextLineSpacing(int(userStyle.FontSpacing))
	rl.SetTextureFilter(userStyle.Font.Texture, rl.FilterBilinear)
	rl.SetExitKey(0)
	rl.SetTargetFPS(144)

	var userText []string
	userText = append(userText, "")
	nav := input.NavigationData{
		SelectedLine: 0,
		SelectedRow:  0,
	}

	state := ProgramState{
		Nav:            nav,
		RenderUpdate:   true,
		AcitveFile:     "",
		ActiveLanguage: r.C,
	}

	for !rl.WindowShouldClose() {

		renderUpdate = input.InputManager(&userText, &nav)
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyQ) {
			break
		}

		rl.BeginDrawing()
		rl.ClearBackground(userStyle.ColorTheme.Editor.Bg)

		var textToRender string
		for _, l := range userText {
			textToRender += l
			textToRender += "\n"
		}
		textToRender = strings.TrimRight(textToRender, "\n")
		r.RenderText(state.ActiveLanguage, &textToRender, &userStyle)

		DrawCursor(userText, &nav, &userStyle)
		rl.EndDrawing()

	}
	rl.UnloadFont(userStyle.Font)
}
