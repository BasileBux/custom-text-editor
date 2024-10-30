package settings

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"gopkg.in/yaml.v3"
)

type ThemeHex struct {
	Syntax struct {
		Tag      string `yaml:"tag"`
		Func     string `yaml:"func"`
		Entity   string `yaml:"entity"`
		String   string `yaml:"string"`
		Regexp   string `yaml:"regexp"`
		Markup   string `yaml:"markup"`
		Keyword  string `yaml:"keyword"`
		Special  string `yaml:"special"`
		Comment  string `yaml:"comment"`
		Constant string `yaml:"constant"`
		Operator string `yaml:"operator"`
	} `yaml:"syntax"`

	Editor struct {
		Fg          string `yaml:"fg"`
		Bg          string `yaml:"bg"`
		Line        string `yaml:"line"`
		Selection   string `yaml:"selection"`
		FindMatch   string `yaml:"findMatch"`
		Gutter      string `yaml:"gutter"`
		IndentGuide string `yaml:"indentGuide"`
	} `yaml:"editor"`

	UI struct {
		Fg        string `yaml:"fg"`
		Bg        string `yaml:"bg"`
		Line      string `yaml:"line"`
		Selection string `yaml:"selection"`
		Panel     struct {
			Bg     string `yaml:"bg"`
			Shadow string `yaml:"shadow"`
		} `yaml:"panel"`
	} `yaml:"ui"`

	Common struct {
		Accent string `yaml:"accent"`
		Error  string `yaml:"error"`
	} `yaml:"common"`
}

type Theme struct {
	Syntax struct {
		Tag      rl.Color
		Func     rl.Color
		Entity   rl.Color
		String   rl.Color
		Regexp   rl.Color
		Markup   rl.Color
		Keyword  rl.Color
		Special  rl.Color
		Comment  rl.Color
		Constant rl.Color
		Operator rl.Color
	}

	Editor struct {
		Fg          rl.Color
		Bg          rl.Color
		Line        rl.Color
		Selection   rl.Color
		FindMatch   rl.Color
		Gutter      rl.Color
		IndentGuide rl.Color
	}

	UI struct {
		Fg        rl.Color
		Bg        rl.Color
		Line      rl.Color
		Selection rl.Color
		Panel     struct {
			Bg     rl.Color
			Shadow rl.Color
		}
	}

	Common struct {
		Accent rl.Color
		Error  rl.Color
	}
}

func HexToRayColorTheme(hexTheme ThemeHex) Theme {
	var rayTheme Theme

	// Syntax
	rayTheme.Syntax.Tag = GetRayColor(hexTheme.Syntax.Tag)
	rayTheme.Syntax.Func = GetRayColor(hexTheme.Syntax.Func)
	rayTheme.Syntax.Entity = GetRayColor(hexTheme.Syntax.Entity)
	rayTheme.Syntax.String = GetRayColor(hexTheme.Syntax.String)
	rayTheme.Syntax.Regexp = GetRayColor(hexTheme.Syntax.Regexp)
	rayTheme.Syntax.Markup = GetRayColor(hexTheme.Syntax.Markup)
	rayTheme.Syntax.Keyword = GetRayColor(hexTheme.Syntax.Keyword)
	rayTheme.Syntax.Special = GetRayColor(hexTheme.Syntax.Special)
	rayTheme.Syntax.Comment = GetRayColor(hexTheme.Syntax.Comment)
	rayTheme.Syntax.Constant = GetRayColor(hexTheme.Syntax.Constant)
	rayTheme.Syntax.Operator = GetRayColor(hexTheme.Syntax.Operator)

	// Editor
	rayTheme.Editor.Fg = GetRayColor(hexTheme.Editor.Fg)
	rayTheme.Editor.Bg = GetRayColor(hexTheme.Editor.Bg)
	rayTheme.Editor.Line = GetRayColor(hexTheme.Editor.Line)
	rayTheme.Editor.Selection = GetRayColor(hexTheme.Editor.Selection)
	rayTheme.Editor.FindMatch = GetRayColor(hexTheme.Editor.FindMatch)
	rayTheme.Editor.Gutter = GetRayColor(hexTheme.Editor.Gutter)
	rayTheme.Editor.IndentGuide = GetRayColor(hexTheme.Editor.IndentGuide)

	// UI
	rayTheme.UI.Fg = GetRayColor(hexTheme.UI.Fg)
	rayTheme.UI.Bg = GetRayColor(hexTheme.UI.Bg)
	rayTheme.UI.Line = GetRayColor(hexTheme.UI.Line)
	rayTheme.UI.Selection = GetRayColor(hexTheme.UI.Selection)
	rayTheme.UI.Panel.Bg = GetRayColor(hexTheme.UI.Panel.Bg)
	rayTheme.UI.Panel.Shadow = GetRayColor(hexTheme.UI.Panel.Shadow)

	// Common
	rayTheme.Common.Accent = GetRayColor(hexTheme.Common.Accent)
	rayTheme.Common.Error = GetRayColor(hexTheme.Common.Error)

	return rayTheme
}

func GetColorThemeFromFileName(themeName *string) (Theme, error) {
	filename := "./themes/" + (*themeName) + ".yaml"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return Theme{}, err
	}

	var themeHex ThemeHex
	err = yaml.Unmarshal(data, &themeHex)
	if err != nil {
		fmt.Println(err)
		return Theme{}, err
	}

	theme := HexToRayColorTheme(themeHex)

	return theme, nil
}

// solution from: https://www.cloudhadoop.com/2018/12/golang-example-convertcast-hexa-to20.html
func GetRayColor(hex string) rl.Color {
	if len(hex) != 9 {
		log.Panic("The theme you try to use is wrong. One of the colors isn't the right format")
		return rl.Color{}
	}

	var outputColor [4]uint8 // R8 G8 B8 A8

	numberStr := strings.Replace(hex, "#", "", -1)

	for i := 0; i < 4; i++ {
		hexPart := numberStr[i*2 : i*2+2]
		intValue, err := strconv.ParseUint(hexPart, 16, 64)
		if err != nil {
			log.Panic(err)
			return rl.Color{}
		}
		outputColor[i] = uint8(intValue)
	}

	return rl.Color{R: outputColor[0],
		G: outputColor[1],
		B: outputColor[2],
		A: outputColor[3]}
}
