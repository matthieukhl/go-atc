package main

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/matthieukhl/go-atc/internal"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	option int
	title  string
	desc   string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type appState int

const (
	menu appState = iota
	checkFlight
	flightDepartures
	flightArrivals
)

type model struct {
	// App state
	appState appState

	// Components
	textInput textinput.Model
	list      list.Model

	// HTTP Client
	client internal.Client

	err error
}

func initialModel() model {
	items := []list.Item{
		item{option: 1, title: "Get flight information.", desc: "Display real-time information of a flight."},
		item{option: 2, title: "Get flight schedule from airport.", desc: "Display flight schedule of an airport."},
	}

	m := model{
		appState:  menu,
		textInput: textinput.New(),
		list:      list.New(items, list.NewDefaultDelegate(), 0, 0),
		client:    internal.NewClient(),
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

	switch m.appState {
	case menu:
		// On laisse la list gérer la navigation (up/down, etc.)
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)

		// On intercepte Enter pour décider si on bascule en inputState
		if key, ok := msg.(tea.KeyMsg); ok && key.Type == tea.KeyEnter {
			switch m.list.SelectedItem().(item).option {
			case 1:
				m.appState = checkFlight
				m.textInput.Placeholder = "MEA212"
				m.textInput.CharLimit = 156
				m.textInput.Width = 20
				m.textInput.SetValue("")
				m.textInput.Focus()
				return m, textinput.Blink
			case 2:
				m.appState = flightDepartures
				m.textInput.Placeholder = "LFPG"
				m.textInput.CharLimit = 156
				m.textInput.Width = 20
				m.textInput.SetValue("")
				m.textInput.Focus()
				return m, textinput.Blink
			}
		}

		return m, cmd

	case checkFlight:
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.Type {
			case tea.KeyEsc:
				m.appState = menu
				m.textInput.Blur()
				return m, nil

			case tea.KeyEnter:
				m.appState = menu
				m.textInput.Blur()
				return m, nil
			}
		}

		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd

	case flightDepartures:
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.Type {
			case tea.KeyEsc:
				m.appState = menu
				m.textInput.Blur()
				return m, nil
			case tea.KeyEnter:
				err := internal.GetFlightDepartures(m.client, m.textInput.Value())
				if err != nil {
					log.Fatal(err)
				}
				return m, nil
			}

			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.appState {
	case menu:
		header := ""
		return docStyle.Render(header + m.list.View())

	case checkFlight:
		return docStyle.Render(
			"Enter flight callsign (Esc to cancel)\n\n" +
				m.textInput.View(),
		)

	case flightDepartures:
		return docStyle.Render(
			"Enter airport ICAO (Esc to cancel)\n\n" + m.textInput.View(),
		)
	}

	return docStyle.Render(m.list.View())
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
