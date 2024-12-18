package settings

import (
	_ "embed"
	"encoding/json"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var VALID_CONFIG_FILES = [3]string{"settings", "user", "kenzan"}

type Cursor struct {
	Width             int32
	Ratio             float32 // ratio with the text height
	HorizontalPadding int32   // number of chars to show when scrolling
	VerticalPadding   int32
}

type LineNumbers struct {
	PaddingLeft   int
	PaddingRight  int
	LineWidth     int
	OffsetCurrent bool
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
	LineNumbers   LineNumbers
}

type Settings struct {
	Padding struct {
		Top    *int `json:"top,omitempty"`
		Right  *int `json:"right,omitempty"`
		Bottom *int `json:"bottom,omitempty"`
		Left   *int `json:"left,omitempty"`
	} `json:"padding,omitempty"`
	FontFamily    *string `json:"font_family,omitempty"`
	FontSize      *int    `json:"font_size,omitempty"`
	FontSpacing   *int    `json:"font_spacing,omitempty"`
	ScrollPadding *int    `json:"scroll_padding,omitempty"`
	CursorRatio   *int    `json:"cursor_ratio,omitempty"`
	Theme         *string `json:"theme,omitempty"`

	LineNumbers struct {
		Show          *bool `json:"show,omitempty"`
		Relative      *bool `json:"relative,omitempty"`
		PaddingLeft   *int  `json:"padding_left,omitempty"`
		PaddingRight  *int  `json:"padding_right,omitempty"`
		LineWidth     *int  `json:"line_width,omitempty"`
		OffsetCurrent *bool `json:"offset_current,omitempty"`
	} `json:"line_numbers,omitempty"`

	LineHighlight      *bool `json:"line_highlight,omitempty"`
	HighDpi            *bool `json:"high_dpi,omitempty"`
	PerformanceDisplay *bool `json:"performance_display,omitempty"`
}

//go:embed default.json
var defaultData []byte

func loadDefaultSettings() (*Settings, error) {
	var settings Settings
	if err := json.Unmarshal(defaultData, &settings); err != nil {
		return nil, err
	}

	return &settings, nil
}

func loadSettings(path string) (*Settings, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var settings Settings
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, err
	}

	return &settings, nil
}

func loadUserSettings(path string) (*Settings, error) {
	var err error
	for _, p := range VALID_CONFIG_FILES {
		file := path + "/" + p + ".json"
		data, err := loadSettings(file)
		if err == nil {
			return data, nil
		}
	}

	return nil, err
}

func LoadAllSettings() (*Settings, error) {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir = os.ExpandEnv("$HOME/.config")
	}
	configDir += "/kenzan"

	defaults, err := loadDefaultSettings()
	if err != nil {
		return nil, err
	}

	user, err := loadUserSettings(configDir)
	if err != nil {
		return defaults, nil
	}

	merged := MergeSettings(defaults, user)
	return merged, nil
}

func MergeSettings(defaults *Settings, user *Settings) *Settings {
	if user == nil {
		return defaults
	}

	merged := *defaults // Create a copy of defaults

	// Merge padding settings
	if user.Padding.Top != nil {
		merged.Padding.Top = user.Padding.Top
	}
	if user.Padding.Right != nil {
		merged.Padding.Right = user.Padding.Right
	}
	if user.Padding.Bottom != nil {
		merged.Padding.Bottom = user.Padding.Bottom
	}
	if user.Padding.Left != nil {
		merged.Padding.Left = user.Padding.Left
	}

	// Merge other settings
	if user.FontFamily != nil {
		merged.FontFamily = user.FontFamily
	}
	if user.FontSize != nil {
		merged.FontSize = user.FontSize
	}
	if user.FontSpacing != nil {
		merged.FontSpacing = user.FontSpacing
	}
	if user.ScrollPadding != nil {
		merged.ScrollPadding = user.ScrollPadding
	}
	if user.CursorRatio != nil {
		merged.CursorRatio = user.CursorRatio
	}
	if user.Theme != nil {
		merged.Theme = user.Theme
	}

	// Merge line number settings
	if user.LineNumbers.Show != nil {
		merged.LineNumbers.Show = user.LineNumbers.Show
	}
	if user.LineNumbers.Relative != nil {
		merged.LineNumbers.Relative = user.LineNumbers.Relative
	}
	if user.LineNumbers.PaddingLeft != nil {
		merged.LineNumbers.PaddingLeft = user.LineNumbers.PaddingLeft
	}
	if user.LineNumbers.PaddingRight != nil {
		merged.LineNumbers.PaddingRight = user.LineNumbers.PaddingRight
	}
	if user.LineNumbers.LineWidth != nil {
		merged.LineNumbers.LineWidth = user.LineNumbers.LineWidth
	}
	if user.LineNumbers.OffsetCurrent != nil {
		merged.LineNumbers.OffsetCurrent = user.LineNumbers.OffsetCurrent
	}

	if user.LineHighlight != nil {
		merged.LineHighlight = user.LineHighlight
	}

	// Merge system settings
	if user.HighDpi != nil {
		merged.HighDpi = user.HighDpi
	}

	if user.PerformanceDisplay != nil {
		merged.PerformanceDisplay = user.PerformanceDisplay
	}

	return &merged
}
