package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type mode int

const (
	modeNormal mode = iota
	modeSeekBar
)

type model struct {
	*pages
	mode
	tw, th  int
	current int
	err     string
}

func newModel(pages *pages) model {
	return model{
		pages:   pages,
		mode:    modeNormal,
		current: 1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if m.current < len(m.pages.ps) {
				m.current += 1
			}
		case "k":
			if m.current > 1 {
				m.current -= 1
			}
		case "tab":
			switch m.mode {
			case modeNormal:
				m.mode = modeSeekBar
			case modeSeekBar:
				m.mode = modeNormal
			}
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.tw = msg.Width
		m.th = msg.Height
		if m.tw < m.pages.width || m.th < m.pages.height {
			m.err = fmt.Sprintf("not enough size [terminal: %dx%d, require: %dx%d]", m.tw, m.th, m.pages.width, m.pages.height)
		} else {
			m.err = ""
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != "" {
		return lipgloss.Place(m.tw, m.th, lipgloss.Center, lipgloss.Center, m.err)
	}

	p := m.pages.ps[m.current-1]

	switch m.mode {
	case modeNormal:
		s := strings.Join(p.lines, "\n")
		return s
	case modeSeekBar:
		current := m.current
		total := len(m.pages.ps)

		sb := seekBar(m.tw, current, total)
		pn := pageNumber(m.tw, current, total)

		ft := lipgloss.JoinVertical(lipgloss.Left, sb, pn)
		ft = lipgloss.NewStyle().
			Width(m.tw).
			Border(lipgloss.DoubleBorder(), true, false, false, false).
			BorderForeground(lipgloss.Color("238")).
			Render(ft)
		_, fh := lipgloss.Size(ft)
		fts := strings.Split(ft, "\n")

		lines := make([]string, 0, m.th)
		for i, l := range p.lines {
			if i >= m.th-fh {
				break
			}
			s := lipgloss.NewStyle().
				Foreground(lipgloss.Color("245")).
				Render(l)
			lines = append(lines, s)
		}
		n := m.th - len(p.lines) - fh
		for i := 0; i < n; i++ {
			lines = append(lines, "")
		}
		for i := 0; i < fh; i++ {
			lines = append(lines, fts[i])
		}

		s := strings.Join(lines, "\n")
		return s
	default:
		return ""
	}
}

func seekBar(w, current, total int) string {
	barLength := w - 4
	pos := int((float64(current-1) / float64(total-1)) * float64(barLength))
	bar := strings.Repeat("━", pos) + strings.Repeat("─", barLength-pos)
	return lipgloss.PlaceHorizontal(w, lipgloss.Center, bar)
}

func pageNumber(w, current, total int) string {
	d := digit(uint(total))
	ps := fmt.Sprintf("%*d / %d", d, current, total)
	return lipgloss.PlaceHorizontal(w, lipgloss.Center, ps)
}

func digit(n uint) uint {
	if n == 0 {
		return 1
	}
	var c uint
	for n > 0 {
		n /= 10
		c++
	}
	return c
}

func start(pages *pages) error {
	m := newModel(pages)
	p := tea.NewProgram(m, tea.WithAltScreen())
	return p.Start()
}
