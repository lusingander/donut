package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	*pages
	current int
	err     string
}

func newModel(pages *pages) model {
	return model{
		pages:   pages,
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
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		if msg.Width < m.pages.width || msg.Height < m.pages.height {
			m.err = fmt.Sprintf("not enough size [terminal: %dx%d, require: %dx%d]", msg.Width, msg.Height, m.pages.width, m.pages.height)
		} else {
			m.err = ""
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != "" {
		return m.err
	}

	p := m.pages.ps[m.current-1]
	s := strings.Join(p.lines, "\n")
	return s
}

func start(pages *pages) error {
	m := newModel(pages)
	p := tea.NewProgram(m, tea.WithAltScreen())
	return p.Start()
}
