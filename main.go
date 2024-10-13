package main

import (
	"log"

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
	PaddingTop    float64
	PaddingRight  float64
	PaddingBottom float64
	PaddingLeft   float64
	ColorTheme    Theme
}

var compact WindowStyle = WindowStyle{
	PaddingTop:    13.0,
	PaddingRight:  13.0,
	PaddingBottom: 13.0,
	PaddingLeft:   13.0,
	ColorTheme:    darkTheme,
}

func inputManager(text *string) {
	char := rl.GetCharPressed()
	for char > 0 {
		// refuse non ascii and non printable chars
		if char >= 32 && char <= 126 {
			*text += string(rune(char))
		}
		char = rl.GetCharPressed()
	}

	if rl.IsKeyPressedRepeat(rl.KeyBackspace) && len(*text) > 0 {
		*text = (*text)[:len(*text)-1]
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		*text += "\n"
	}
}

func main() {
	rl.InitWindow(1000, 800, "My custom text editor")
	if !rl.IsWindowReady() {
		log.Panic("Window didn't open correctly ???")
	}
	defer rl.CloseWindow()
	userFont := rl.LoadFontEx("/usr/share/fonts/GeistMono/GeistMonoNerdFont-Regular.otf", 100, nil)
	rl.SetTextureFilter(userFont.Texture, rl.FilterBilinear)
	rl.SetTargetFPS(144)

	userStyle := compact
	var userText string

	for !rl.WindowShouldClose() {

		inputManager(&userText)

		rl.BeginDrawing()

		rl.ClearBackground(userStyle.ColorTheme.Background)
		textPos := rl.NewVector2(float32(userStyle.PaddingLeft), float32(userStyle.PaddingTop))
		rl.DrawTextEx(userFont, userText, textPos, 30, 1, userStyle.ColorTheme.Text)

		rl.EndDrawing()
	}
	rl.UnloadFont(userFont)
}
