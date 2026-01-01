package main

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type state int

const (
	listState state = iota
	inputState
)

type model struct {
	// App state
	state state

	// Components
	textInput textinput.Model
	list      list.Model

	err error
}

func initialModel() model {
	items := []list.Item{
		item{title: "Get flight information.", desc: "Display real-time information of a flight."},
		item{title: "Get flight schedule from airport.", desc: "Display flight schedule of an airport."},
	}

	ti := textinput.New()
	ti.Placeholder = "MEA212"
	ti.CharLimit = 156
	ti.Width = 20

	m := model{
		state:     listState,
		textInput: ti,
		list:      list.New(items, list.NewDefaultDelegate(), 0, 0),
		err:       nil,
	}
	m.list.Title = "Go ATC"
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == tea.KeyCtrlC.String() {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	switch m.state {
	case listState:
		// On laisse la list gérer la navigation (up/down, etc.)
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)

		// On intercepte Enter pour décider si on bascule en inputState
		if key, ok := msg.(tea.KeyMsg); ok && key.Type == tea.KeyEnter {
			if it, ok := m.list.SelectedItem().(item); ok {
				if it.title == "Get flight information." {
					m.state = inputState
					m.textInput.SetValue("")
					m.textInput.Focus()
					return m, textinput.Blink
				}
			}
		}

		return m, cmd

	case inputState:
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.Type {
			case tea.KeyEsc:
				m.state = listState
				m.textInput.Blur()
				return m, nil

			case tea.KeyEnter:
				m.state = listState
				m.textInput.Blur()
				return m, nil
			}
		}

		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {
	case listState:
		header := ""
		return docStyle.Render(header + m.list.View())

	case inputState:
		return docStyle.Render(
			"Enter flight callsign (Esc to cancel)\n\n" +
				m.textInput.View(),
		)
	}

	return docStyle.Render(m.list.View())
}
