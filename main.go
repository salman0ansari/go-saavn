package main

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type screen interface {
	Reset()
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

var currentScreen screen

type model struct {
	err error

	list           list.Model
	queryTextInput textinput.Model
	viewport       viewport.Model

	ready         bool
	selectedTopic []list.Item
	response      SearchResponse
	selectedQuery string
	resultContent string
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		if !m.ready {
			m.ready = true
			m.viewport = viewport.Model{Width: msg.Width, Height: msg.Height - (headerHeight + footerHeight)}
			m.viewport.YPosition = headerHeight + 1
		}
	}
	return currentScreen.Update(msg)
}

func (m model) View() string {
	return currentScreen.View()
}

func InitialModel() *model {

	tm := textinput.NewModel()
	tm.Placeholder = "Your query"
	tm.Focus()
	tm.CharLimit = 156
	tm.Width = 20

	m := model{
		queryTextInput: tm,
	}
	m.list.Title = "Topics"
	return &m
}

func setCurrentScreen(scr screen) {
	currentScreen = scr
	currentScreen.Reset()
}

func main() {

	m := InitialModel()
	setCurrentScreen(&searchScreen{model: m})

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}
