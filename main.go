package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	f "github.com/basileb/custom_text_editor/files"
	"github.com/basileb/custom_text_editor/input"
	r "github.com/basileb/custom_text_editor/renderer"
	st "github.com/basileb/custom_text_editor/settings"
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

func main() {

	if len(os.Args) > 2 {
		fmt.Println("You need to provide one or no filepath")
		return
	}

	var userText []string
	filename := ""
	fileLanguage := t.NONE
	userText = append(userText, "")
	var err error
	if len(os.Args) == 2 {
		userText, err = f.OpenFile(os.Args[1])
		if err != nil {
			fmt.Println("Error: Couldn't open specified file")
			return
		}
		if len(userText) == 0 {
			userText = append(userText, "")
		}
		filename = os.Args[1]
		fileLanguage = f.GetFileExtension(os.Args[1])
	}

	RedirectLogs()

	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(800, 800, "My custom text editor")
	if !rl.IsWindowReady() {
		log.Panic("Window didn't open correctly ???")
	}
	defer rl.CloseWindow()

	userStyle := st.Compact
	themeName := "ayu-light"
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

	charSize := rl.MeasureTextEx(userStyle.Font, "a", userStyle.FontSize, userStyle.FontSpacing)
	userStyle.CharSize = charSize

	nav := t.NavigationData{
		SelectedLine:        0,
		SelectedRow:         0,
		AbsoluteSelectedRow: 0,
		ScrollOffset: rl.Vector2{
			X: 0,
			Y: 0,
		},
	}

	state := t.ProgramState{
		Nav:            &nav,
		RenderUpdate:   true,
		AcitveFile:     filename,
		ActiveLanguage: fileLanguage,
		SavedFile:      make([]string, len(userText)),
		SaveState:      true,
		ForceQuit:      false,
		ViewPortSize: rl.Vector2{
			X: float32(rl.GetRenderWidth()),
			Y: float32(rl.GetRenderHeight())},
	}
	state.ViewPortSteps.X = int(state.ViewPortSize.X / userStyle.CharSize.X)
	state.ViewPortSteps.Y = int(state.ViewPortSize.Y / userStyle.CharSize.Y)

	copy(state.SavedFile, userText)

	for !rl.WindowShouldClose() {

		terminate := input.InputManager(&userText, &state, &userStyle)
		if terminate {
			break
		}
		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyQ) {
			if state.SaveState {
				break
			} else {
				noChanges := f.DiffText(userText, state.SavedFile)
				if noChanges {
					break
				}
				fmt.Println("The file wasn't saved. Are you sure you want to close the editor ?")
				fmt.Println("press enter to confirm")
				state.ForceQuit = true
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(userStyle.ColorTheme.Editor.Bg)

		if rl.IsWindowResized() {
			state.ViewPortSize.X = float32(rl.GetRenderWidth())
			state.ViewPortSize.Y = float32(rl.GetRenderHeight())

			state.ViewPortSteps.X = int(state.ViewPortSize.X / userStyle.CharSize.X)
			state.ViewPortSteps.Y = int(state.ViewPortSize.Y / userStyle.CharSize.Y)
		}

		var textToRender string
		for _, l := range userText {
			textToRender += l
			textToRender += "\n"
		}
		textToRender = strings.TrimRight(textToRender, "\n")
		r.RenderText(state.ActiveLanguage, &textToRender, &state, &userStyle)

		r.DrawCursor(userText, &nav, &userStyle)
		rl.EndDrawing()

	}
	rl.UnloadFont(userStyle.Font)
}
