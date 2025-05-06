package setup

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type configModel struct {
	focusIndex int
	inputs     []textinput.Model
	backendIdx int
	done       bool
	choice     string
}

func InitialConfigModel(choice string) configModel {
	m := configModel{}
	inputLabels := []string{}
	switch choice {
	case OllamaBackend:
		inputLabels = []string{"Model", "Host", "Stream (true/false)"}
		m.choice = OllamaBackend
	case ChatGPTBackend:
		inputLabels = []string{"Model", "Host", "Stream (true/false)", "API Key"}
		m.choice = ChatGPTBackend
	default:
		inputLabels = []string{"Model", "Host", "Stream (true/false)"}
	}

	for _, label := range inputLabels {
		ti := textinput.New()
		ti.Placeholder = label
		ti.Focus()
		m.inputs = append(m.inputs, ti)
	}
	return m
}

func (m configModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m configModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.focusIndex == len(m.inputs) {
				m.done = true
				return m, tea.Quit
			}
			m.focusIndex++
			if m.focusIndex >= len(m.inputs) {
				m.focusIndex = len(m.inputs)
			}
			return m, nil
		case "up":
			m.focusIndex--
		case "down":
			m.focusIndex++
			if m.focusIndex > len(m.inputs) {
				m.focusIndex = len(m.inputs)
			}
		}
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		if i == m.focusIndex {
			m.inputs[i].Focus()
		} else {
			m.inputs[i].Blur()
		}
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m configModel) View() string {
	if m.done {
		//TODO: Write to config file

		return m.String()
	}
	var b strings.Builder
	b.WriteString("DevMind Config Editor (use Tab to switch backend, ↑/↓ to navigate)\n\n")
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View() + "\n")
	}
	b.WriteString("\n[Press Enter to Save, Esc to Cancel]")
	return b.String()
}

func (m configModel) String() string {
	switch m.choice {
	case OllamaBackend:
		return fmt.Sprintf(
			"Saved config:\nModel: %s\nHost: %s\nStream: %s\n",
			m.inputs[0].Value(),
			m.inputs[1].Value(),
			m.inputs[2].Value(),
		)
	case ChatGPTBackend:
		return fmt.Sprintf(
			"Saved config:\nModel: %s\nHost: %s\nStream: %s\nAPI Key: %s\n",
			m.inputs[0].Value(),
			m.inputs[1].Value(),
			m.inputs[2].Value(),
			m.inputs[3].Value(),
		)
	default:
		return "invalid choice"
	}
}
