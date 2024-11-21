package settings

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cursor struct {
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
	FontSize:      30,
	FontSpacing:   1,
	Cursor: Cursor{
		Width:             1,
		Ratio:             1,
		HorizontalPadding: 5,
		VerticalPadding:   5,
	},
}

type Settings struct {
    UI struct {
        Padding struct {
            Top    int    `json:"top"`
            Right  int    `json:"right"`
            Bottom int    `json:"bottom"`
            Left   int    `json:"left"`
        } `json:"padding"`
        FontFamily    string `json:"font_familly"`
        FontSpacing   int    `json:"font_spacing"`
        ScrollPadding int    `json:"scroll_padding"`
        CursorRatio   int    `json:"cursor_ratio"`
        Theme         string `json:"theme"`
    } `json:"ui"`
}

// func loadSettings(path string) (*Settings, error) {
//     data, err := os.ReadFile(path)
//     if err != nil {
//         return nil, err
//     }
//
//     var settings Settings
//     if err := json.Unmarshal(data, &settings); err != nil {
//         return nil, err
//     }
//
//     return &settings, nil
// }
//
// func LoadAllSettings() (*Settings, error) {
// 	default, err := loadSettings("default.toml"); err != nil {
// 		return nil, err
// 	}
//
// 	return &settings, nil
// }
