package settings

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Theme struct {
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

	VCS struct {
		Added    string `yaml:"added"`
		Modified string `yaml:"modified"`
		Removed  string `yaml:"removed"`
	} `yaml:"vcs"`

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

func getColorThemeFromFileName(themeName *string) (Theme, error) {
	filename := "./themes/" + (*themeName) + ".yaml"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return Theme{}, err
	}

	var theme Theme
	err = yaml.Unmarshal(data, &theme)
	if err != nil {
		fmt.Println(err)
		return Theme{}, err
	}

	return theme, nil
}
