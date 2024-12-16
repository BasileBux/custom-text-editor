package settings

import (
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

type Settings struct {
	UI struct {
		Padding struct {
			Top    *int `json:"top,omitempty"`
			Right  *int `json:"right,omitempty"`
			Bottom *int `json:"bottom,omitempty"`
			Left   *int `json:"left,omitempty"`
		} `json:"padding,omitempty"`
		FontFamily    *string `json:"font_familly,omitempty"`
		FontSize      *int    `json:"font_size,omitempty"`
		FontSpacing   *int    `json:"font_spacing,omitempty"`
		ScrollPadding *int    `json:"scroll_padding,omitempty"`
		CursorRatio   *int    `json:"cursor_ratio,omitempty"`
		Theme         *string `json:"theme,omitempty"`

		LineNumbers struct {
			Show         *bool `json:"show,omitempty"`
			Relative     *bool `json:"relative,omitempty"`
			PaddingLeft  *int  `json:"padding_left,omitempty"`
			PaddingRight *int  `json:"padding_right,omitempty"`
		} `json:"line_numbers,omitempty"`
	} `json:"ui,omitempty"`
	System struct {
		HighDpi *bool `json:"high_dpi,omitempty"`
	} `json:"system,omitempty"`
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
	defaultPath := configDir + "/default.json"

	defaults, err := loadSettings(defaultPath)
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
	if user.UI.Padding.Top != nil {
		merged.UI.Padding.Top = user.UI.Padding.Top
	}
	if user.UI.Padding.Right != nil {
		merged.UI.Padding.Right = user.UI.Padding.Right
	}
	if user.UI.Padding.Bottom != nil {
		merged.UI.Padding.Bottom = user.UI.Padding.Bottom
	}
	if user.UI.Padding.Left != nil {
		merged.UI.Padding.Left = user.UI.Padding.Left
	}

	// Merge other UI settings
	if user.UI.FontFamily != nil {
		merged.UI.FontFamily = user.UI.FontFamily
	}
	if user.UI.FontSize != nil {
		merged.UI.FontSize = user.UI.FontSize
	}
	if user.UI.FontSpacing != nil {
		merged.UI.FontSpacing = user.UI.FontSpacing
	}
	if user.UI.ScrollPadding != nil {
		merged.UI.ScrollPadding = user.UI.ScrollPadding
	}
	if user.UI.CursorRatio != nil {
		merged.UI.CursorRatio = user.UI.CursorRatio
	}
	if user.UI.Theme != nil {
		merged.UI.Theme = user.UI.Theme
	}

	// Merge line number settings
	if user.UI.LineNumbers.Show != nil {
		merged.UI.LineNumbers.Show = user.UI.LineNumbers.Show
	}
	if user.UI.LineNumbers.Relative != nil {
		merged.UI.LineNumbers.Relative = user.UI.LineNumbers.Relative
	}
	if user.UI.LineNumbers.PaddingLeft != nil {
		merged.UI.LineNumbers.PaddingLeft = user.UI.LineNumbers.PaddingLeft
	}
	if user.UI.LineNumbers.PaddingRight != nil {
		merged.UI.LineNumbers.PaddingRight = user.UI.LineNumbers.PaddingRight
	}

	// Merge system settings
	if user.System.HighDpi != nil {
		merged.System.HighDpi = user.System.HighDpi
	}

	return &merged
}
