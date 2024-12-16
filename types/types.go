package types

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Language = uint8

const (
	C Language = iota
	// GO
	// CPP
	// PYTHON
	// RUST
	// JAVA
	// JAVASCRIPT
	NONE
)

type NavigationData struct {
	SelectedLine        int        // 0 indexed
	AbsoluteSelectedRow int        // 0 indexed, number of characters depends on nothing
	SelectedRow         int        // 0 indexed, number of characters depends on current line
	ScrollOffset        rl.Vector2 // offset from top and left of window in number of chars
}

type Vec2 struct {
	X int
	Y int
}

type Update struct {
	Cursor    bool
	Highlight bool
}

func (u *Update) Reset() {
	u.Cursor = false
	u.Highlight = false
}

type ProgramState struct {
	Nav            *NavigationData
	AcitveFile     string
	Update         Update
	ActiveLanguage Language
	SavedFile      []string
	SaveState      bool
	ForceQuit      bool
	ViewPortSize   rl.Vector2
	ViewPortSteps  Vec2
	Cache          Cache
}

type Cache struct {
	Syntax      []SyntaxCache
	Cursor      Vec2
	LineNumbers LineNumbersCache
}

type SyntaxCache struct {
	Text   string
	Color  *rl.Color
	Cursor TextRenderCursor
}

type TextRenderCursor struct {
	Line float32 // pixels
	Row  float32 // pixels
	Stop bool
}

type LineNumbersCache struct {
	Width     int32
	Positions []rl.Vector2
	Colors    []rl.Color
	Numbers   []string
}
