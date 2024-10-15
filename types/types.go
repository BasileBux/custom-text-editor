package types

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ThemeOld struct {
	Background rl.Color
	Text       rl.Color
}

var darkTheme ThemeOld = ThemeOld{
	Background: rl.Black,
	Text:       rl.RayWhite,
}

var lightTheme ThemeOld = ThemeOld{
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
	CursorRatio   float32 // ratio with the text height
	ColorTheme    ThemeOld
}

var Compact WindowStyle = WindowStyle{
	PaddingTop:    13.0,
	PaddingRight:  13.0,
	PaddingBottom: 13.0,
	PaddingLeft:   13.0,
	Font:          rl.Font{},
	FontSize:      30,
	CursorOffset:  -2,
	CursorWidth:   1,
	CursorRatio:   1,
	ColorTheme:    lightTheme,
}

type NavigationData struct {
	SelectedLine        int // 0 indexed
	AbsoluteSelectedRow int // 0 indexed, number of characters depends on nothing
	SelectedRow         int // 0 indexed, number of characters depends on current line
}
