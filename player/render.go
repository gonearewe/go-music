package player

import (
	"bufio"
	"strings"

	. "github.com/fatih/color"
)

type ColorTheme [2]Attribute

var (
	Spring    = [2]Attribute{FgHiGreen, FgGreen}
	Autumn    = [2]Attribute{FgHiYellow, FgYellow}
	Winter    = [2]Attribute{FgHiBlue, FgBlue}
	Rose      = [2]Attribute{FgHiRed, FgHiMagenta}
	Valentine = [2]Attribute{FgHiMagenta, FgMagenta}
)

func RenderText(text string, theme ColorTheme) string {
	var flag bool
	s := make([]string, 5)
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		if flag {
			s = append(s, New(theme[1]).Sprint(scanner.Text()))
		} else {
			s = append(s, New(theme[0]).Sprint(scanner.Text()))
		}

		flag = !flag
	}

	return strings.Join(s, "")
}
