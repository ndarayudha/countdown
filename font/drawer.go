package font

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// DrawChar returns the argument character in bigger size.
func DrawChar(c rune) string {
	b := strings.Builder{}
	for _, row := range smallFonts[c] {
		for _, char := range row {
			s := lipgloss.NewStyle().SetString(" ")
			if char == rune('#') {
				s = s.Background(lipgloss.Color("#ffd670"))
			}
			b.WriteString(s.String())
		}
		b.WriteRune('\n')
	}

	return b.String()
}
