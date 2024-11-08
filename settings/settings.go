package settings

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cursor struct {
	Offset            int32 // horizontal distance to text
	Width             int32
	Ratio             float32 // ratio with the text height
	HorizontalPadding int32   // number of chars to show when scrolling
	VerticalPadding   int32
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
	FontSize:      40,
	FontSpacing:   1,
	Cursor: Cursor{
		Offset:            -2,
		Width:             1,
		Ratio:             1,
		HorizontalPadding: 5,
		VerticalPadding:   5,
	},
}
