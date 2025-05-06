package setup

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel() model {
	return model{
		choices:  []string{OllamaBackend, ChatGPTBackend},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	s := "Select LLM:\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		selected := " "
		if _, ok := m.selected[i]; ok {
			selected = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, selected, choice)
	}

	s += "\nPress q to quit.\n"
	s += "Press space to select.\n"
	s += "Press up/down to navigate.\n"
	s += "Press enter to confirm.\n"

	return s

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "enter":
			if len(m.selected) != 1 {
				return m, tea.Quit
			}
			for k := range m.selected {
				return InitialConfigModel(m.choices[k]), nil
			}

		}

	}

	return m, nil
}
