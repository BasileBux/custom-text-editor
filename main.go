package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	f "github.com/basileb/kenzan/files"
	"github.com/basileb/kenzan/input"
	r "github.com/basileb/kenzan/renderer"
	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const FPS int32 = 60

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

	settings, err := st.LoadAllSettings()
	if err != nil {
		panic("Settings couldn't load")
	}

	// Config flags
	rl.SetConfigFlags(rl.FlagWindowResizable)
	if *settings.HighDpi {
		rl.SetConfigFlags(rl.FlagWindowHighdpi)
	}

	rl.InitWindow(800, 800, "kenzan")
	if !rl.IsWindowReady() {
		log.Panic("Window didn't open correctly ???")
	}
	defer rl.CloseWindow()

	userStyle := st.WindowStyle{
		PaddingTop:    float32(*settings.Padding.Top),
		PaddingRight:  float32(*settings.Padding.Right),
		PaddingBottom: float32(*settings.Padding.Bottom),
		PaddingLeft:   float32(*settings.Padding.Left),
		Font:          rl.LoadFontEx(*settings.FontFamily, 100, nil),
		FontSize:      float32(*settings.FontSize),
		FontSpacing:   float32(*settings.FontSpacing),
		Cursor: st.Cursor{
			Width:             1,
			Ratio:             float32(*settings.CursorRatio),
			HorizontalPadding: int32(*settings.ScrollPadding),
			VerticalPadding:   int32(*settings.ScrollPadding),
		},
		LineNumbers: st.LineNumbers{
			PaddingLeft:   *settings.LineNumbers.PaddingLeft,
			PaddingRight:  *settings.LineNumbers.PaddingRight,
			LineWidth:     *settings.LineNumbers.LineWidth,
			OffsetCurrent: *settings.LineNumbers.OffsetCurrent,
		},
	}

	userStyle.ColorTheme, err = st.GetColorThemeFromFileName(settings.Theme)
	if err != nil {
		fmt.Println("Error could not open color theme")
		return
	}

	rl.SetTextLineSpacing(int(userStyle.FontSpacing))
	rl.SetTextureFilter(userStyle.Font.Texture, rl.FilterBilinear)
	rl.SetExitKey(0)
	rl.SetTargetFPS(FPS)

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
		Update:         t.Update{Cursor: true, SyntaxHighlight: true},
		AcitveFile:     filename,
		ActiveLanguage: fileLanguage,
		SavedFile:      make([]string, len(userText)),
		SaveState:      true,
		ForceQuit:      false,
		ViewPortSize: rl.Vector2{
			X: float32(rl.GetRenderWidth()),
			Y: float32(rl.GetRenderHeight())},
	}
	state.ViewPortSteps.X = int(state.ViewPortSize.X / (userStyle.CharSize.X + userStyle.FontSpacing))
	state.ViewPortSteps.Y = int(state.ViewPortSize.Y / (userStyle.CharSize.Y + userStyle.FontSpacing))

	copy(state.SavedFile, userText)

	// Prepare line number cache
	if *settings.LineNumbers.Show {
		lastLineNb := len(userText) - 1
		if userStyle.LineNumbers.OffsetCurrent && len(userText)-1 < 100 {
			lastLineNb *= 10
		}
		largestLineNb := fmt.Sprintf("%d", lastLineNb)
		largestNbSize := rl.MeasureTextEx(userStyle.Font, largestLineNb, userStyle.FontSize, userStyle.FontSpacing)
		state.Cache.LineNumbers.Width = int32(largestNbSize.X)

		userStyle.PaddingLeft += float32(state.Cache.LineNumbers.Width) +
			float32(*settings.LineNumbers.PaddingLeft) + float32(*settings.LineNumbers.PaddingRight)
	}

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
				fmt.Println("The file wasn't saved. Are you sure you want to close the editor ?\npress enter to confirm")
				state.ForceQuit = true
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(userStyle.ColorTheme.Editor.Bg)

		if rl.IsWindowResized() {
			state.ViewPortSize.X = float32(rl.GetRenderWidth())
			state.ViewPortSize.Y = float32(rl.GetRenderHeight())

			state.ViewPortSteps.X = int(state.ViewPortSize.X / (userStyle.CharSize.X + userStyle.FontSpacing))
			state.ViewPortSteps.Y = int(state.ViewPortSize.Y / (userStyle.CharSize.Y + userStyle.FontSpacing))
		}

		if *settings.LineHighlight {
			rl.DrawRectangle(0, int32(state.Cache.Cursor.Y), int32(state.ViewPortSize.X),
				int32(userStyle.CharSize.Y), userStyle.ColorTheme.Editor.Highlight)
		}

		var textToRender string
		for _, l := range userText {
			textToRender += l
			textToRender += "\n"
		}
		textToRender = strings.TrimRight(textToRender, "\n")
		r.RenderText(state.ActiveLanguage, &textToRender, &state, &userStyle)

		if state.Update.Cursor || state.Update.SyntaxHighlight {
			r.CalculateCursorPos(userText, &nav, &state.Cache, &userStyle)
		}
		rl.DrawRectangle(
			int32(state.Cache.Cursor.X),
			int32(state.Cache.Cursor.Y),
			int32(userStyle.Cursor.Width),
			int32(userStyle.FontSize*userStyle.Cursor.Ratio),
			userStyle.ColorTheme.Editor.Fg,
		)

		if state.Update.Cursor {
			r.CalculateLineNbPositions(*settings.LineNumbers.Relative, &state, &userStyle)
		}

		if *settings.LineNumbers.Show {
			r.RenderLineNumbers(&state, &userStyle)
		}

		state.Update.Reset()
		rl.EndDrawing()

	}
	rl.UnloadFont(userStyle.Font)
}
