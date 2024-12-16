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
		Tag      string `json:"tag"`
		Func     string `json:"func"`
		Entity   string `json:"entity"`
		String   string `json:"string"`
		Escape   string `json:"escape"`
		Keyword  string `json:"keyword"`
		Comment  string `json:"comment"`
		Constant string `json:"constant"`
		Operator string `json:"operator"`
	} `json:"syntax"`

	Editor struct {
		Fg     string `json:"fg"`
		Bg     string `json:"bg"`
		Gutter struct {
			Active string `json:"active"`
			Normal string `json:"normal"`
		} `json:"gutter"`
	} `json:"editor"`
}

type Theme struct {
	Syntax struct {
		Tag      rl.Color
		Func     rl.Color
		Entity   rl.Color
		String   rl.Color
		Escape   rl.Color
		Keyword  rl.Color
		Comment  rl.Color
		Constant rl.Color
		Operator rl.Color
	}

	Editor struct {
		Fg     rl.Color
		Bg     rl.Color
		Gutter struct {
			Active rl.Color
			Normal rl.Color
		}
	}
}

func HexToRayColorTheme(hexTheme ThemeHex) Theme {
	var rayTheme Theme

	// Syntax
	rayTheme.Syntax.Tag = GetRayColor(hexTheme.Syntax.Tag)
	rayTheme.Syntax.Func = GetRayColor(hexTheme.Syntax.Func)
	rayTheme.Syntax.Entity = GetRayColor(hexTheme.Syntax.Entity)
	rayTheme.Syntax.String = GetRayColor(hexTheme.Syntax.String)
	rayTheme.Syntax.Escape = GetRayColor(hexTheme.Syntax.Escape)
	rayTheme.Syntax.Keyword = GetRayColor(hexTheme.Syntax.Keyword)
	rayTheme.Syntax.Comment = GetRayColor(hexTheme.Syntax.Comment)
	rayTheme.Syntax.Constant = GetRayColor(hexTheme.Syntax.Constant)
	rayTheme.Syntax.Operator = GetRayColor(hexTheme.Syntax.Operator)

	// Editor
	rayTheme.Editor.Fg = GetRayColor(hexTheme.Editor.Fg)
	rayTheme.Editor.Bg = GetRayColor(hexTheme.Editor.Bg)
	rayTheme.Editor.Gutter.Active = GetRayColor(hexTheme.Editor.Gutter.Active)
	rayTheme.Editor.Gutter.Normal = GetRayColor(hexTheme.Editor.Gutter.Normal)
	return rayTheme
}

func GetColorThemeFromFileName(themeName *string) (Theme, error) {
	filename := "./themes/" + (*themeName) + ".json"
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
