package settings

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type WindowStyle struct {
	PaddingTop    float32
	PaddingRight  float32
	PaddingBottom float32
	PaddingLeft   float32
	Font          rl.Font
	FontSize      float32
	FontSpacing   float32
	CursorOffset  int32 // horizontal distance to text
	CursorWidth   int32
	CursorRatio   float32 // ratio with the text height
	ColorTheme    Theme
}

var Compact WindowStyle = WindowStyle{
	PaddingTop:    13.0,
	PaddingRight:  13.0,
	PaddingBottom: 13.0,
	PaddingLeft:   13.0,
	Font:          rl.Font{},
	FontSize:      30,
	FontSpacing:   1,
	CursorOffset:  -2,
	CursorWidth:   1,
	CursorRatio:   1,
}
