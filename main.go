package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yofan2408/countdown/font"
)

type Window struct {
	Width  int
	Height int
}

type Countdown struct {
	Window Window
	Hour   int
	Min    int
	Sec    int
	Done   bool
}

func initialCountdown(hour, min, sec int) Countdown {
	return Countdown{Hour: hour, Min: min, Sec: sec, Done: false}
}

type tickMsg time.Time

func Tick() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (c Countdown) Init() tea.Cmd {
	return Tick()
}

func (c Countdown) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.Window.Width, c.Window.Height = msg.Width, msg.Height
		return c, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+x":
			return c, tea.Quit
		}
	case tickMsg:
		if c.Hour == 0 && c.Min == 0 && c.Sec == 0 {
			c.Done = true
		}

		if c.Done {
			break
		}

		hour, min, sec := c.Hour, c.Min, c.Sec-1
		if sec < 0 {
			min, sec = min-1, 59
		}

		if min < 0 {
			hour, min = hour-1, 59
		}

		c.Hour, c.Min, c.Sec = hour, min, sec

		return c, Tick()
	}

	return c, nil
}

func (c Countdown) View() string {
	timerStr := ""
	if c.Hour > 0 {
		timerStr = fmt.Sprintf("%02d:", c.Hour)
	}
	timerStr = timerStr + fmt.Sprintf("%02d:%02d", c.Min, c.Sec)
	timer := ""
	for _, c := range timerStr {
		timer = lipgloss.JoinHorizontal(lipgloss.Center, timer, font.DrawChar(c))
	}

	ui := lipgloss.Place(
		c.Window.Width, c.Window.Height,
		lipgloss.Center, lipgloss.Center,
		timer,
	)

	return ui
}

func main() {
	var d = flag.Duration("d", 5*time.Second, "Duration of timer.")

	flag.Parse()

	seconds := int(d.Seconds())
	if seconds < 0 {
		fmt.Print("No time left\n")
		os.Exit(1)
	}

	minutes := seconds / 60
	seconds = seconds % 60
	hours := minutes / 60
	minutes = minutes % 60

	p := tea.NewProgram(initialCountdown(hours, minutes, seconds), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
