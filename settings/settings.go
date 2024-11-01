package settings

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cursor struct {
	CursorOffset            int32 // horizontal distance to text
	CursorWidth             int32
	CursorRatio             float32 // ratio with the text height
	CursorHorizontalPadding int32   // number of chars to show when scrolling
}

type WindowStyle struct {
	PaddingTop    float32
	PaddingRight  float32
	PaddingBottom float32
	PaddingLeft   float32
	Font          rl.Font
	FontSize      float32
	FontSpacing   float32
	Cursor        Cursor
	ColorTheme    Theme
	CharSize      rl.Vector2
}

var Compact WindowStyle = WindowStyle{
	PaddingTop:    13.0,
	PaddingRight:  13.0,
	PaddingBottom: 13.0,
	PaddingLeft:   13.0,
	Font:          rl.Font{},
	FontSize:      30,
	FontSpacing:   1,
	Cursor: Cursor{
		CursorOffset: -2,
		CursorWidth:  1,
		CursorRatio:  1,
		CursorHorizontalPadding: 5,
	},
}
