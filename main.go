package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	CursorOffset  int // horizontal distance to text
	CursorWidth   int
	ColorTheme    Theme
}

var compact WindowStyle = WindowStyle{
	PaddingTop:    13.0,
	PaddingRight:  13.0,
	PaddingBottom: 13.0,
	PaddingLeft:   13.0,
	Font:          rl.Font{},
	FontSize:      30,
	CursorOffset:  3,
	CursorWidth:   8,
	ColorTheme:    darkTheme,
}

type NavigationData struct {
	SelectedLine int // 0 indexed
}

func inputManager(text *string, nav *NavigationData) {
	char := rl.GetCharPressed()
	for char > 0 {
		// refuse non ascii and non printable chars
		if char >= 32 && char <= 126 {
			*text += string(rune(char))
		}
		char = rl.GetCharPressed()
	}

	if (rl.IsKeyPressedRepeat(rl.KeyBackspace) || rl.IsKeyPressed(rl.KeyBackspace)) && len(*text) > 0 {
		if (*text)[len(*text)-1] == '\n' {
			nav.SelectedLine--
		}
		*text = (*text)[:len(*text)-1]
	}

	if rl.IsKeyPressedRepeat(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyEnter) {
		*text += "\n"
		nav.SelectedLine++
	}
}

func getTextSize(text *string, nav *NavigationData, style *WindowStyle) (rl.Vector2, error) {
	lines := strings.Split(*text, "\n")
	if nav.SelectedLine > len(*text)-1 {
		fmt.Println("Error, navigation index is too big!")
		return rl.Vector2{}, fmt.Errorf("Navigation index to big\n")
	}
	currentLine := strings.Split(lines[nav.SelectedLine], "\n")
	textSize := rl.MeasureTextEx(style.Font, currentLine[0], style.FontSize, 1)
	return textSize, nil
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

	rl.InitWindow(1000, 800, "My custom text editor")
	if !rl.IsWindowReady() {
		log.Panic("Window didn't open correctly ???")
	}
	defer rl.CloseWindow()

	userStyle := compact
	userStyle.Font = rl.LoadFontEx("/usr/share/fonts/GeistMono/GeistMonoNerdFont-Regular.otf", 100, nil)
	rl.SetTextureFilter(userStyle.Font.Texture, rl.FilterBilinear)
	rl.SetTargetFPS(144)

	var userText string
	nav := NavigationData{
		SelectedLine: 0,
	}

	for !rl.WindowShouldClose() {

		inputManager(&userText, &nav)

		rl.BeginDrawing()

		rl.ClearBackground(userStyle.ColorTheme.Background)

		textPos := rl.NewVector2(userStyle.PaddingLeft, userStyle.PaddingTop)
		rl.DrawTextEx(userStyle.Font, userText, textPos, userStyle.FontSize, 1, userStyle.ColorTheme.Text)

		// Draw cursor rectangle
		// TODO: This gives error of nav index being wrong. Check this out
		textSize, err := getTextSize(&userText, &nav, &userStyle)
		if err != nil {
			fmt.Println("Resetting nav index")
			nav.SelectedLine = 0
		}
		// TODO: Show cursor when there is no text
		rl.DrawRectangle(int32(userStyle.PaddingLeft)+int32(textSize.X)+int32(userStyle.CursorOffset), int32(userStyle.PaddingTop)+int32(nav.SelectedLine)*int32(textSize.Y)+int32(nav.SelectedLine+1), int32(userStyle.CursorWidth), int32(textSize.Y), rl.RayWhite)

		rl.EndDrawing()
	}
	rl.UnloadFont(userStyle.Font)
}
