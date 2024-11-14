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

type ProgramState struct {
	Nav            *NavigationData
	AcitveFile     string
	RenderUpdate   bool
	ActiveLanguage Language
	SavedFile      []string
	SaveState      bool
	ForceQuit      bool
	ViewPortSize   rl.Vector2
	ViewPortSteps  Vec2
}
