package main

import (
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

type listScreen struct {
	model *model
}

func (t *listScreen) Reset() {
	t.model.list.ResetFilter()
	t.model.queryTextInput.Reset()
}

func (t *listScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "space", " ":
			t.model.selectedTopic = append(t.model.selectedTopic, t.model.list.SelectedItem())
			t.model.list.RemoveItem(t.model.list.Index())
			return t.model, nil

		}
		switch msg.Type {
		case tea.KeyEsc:
			tm := textinput.NewModel()
			tm.Placeholder = "Your query"
			tm.Focus()
			tm.CharLimit = 156
			tm.Width = 20
			m := model{
				queryTextInput: tm,
			}
			setCurrentScreen(&searchScreen{model: &m})

		case tea.KeyCtrlC:
			return t.model, tea.Quit
		case tea.KeyEnter:
			setCurrentScreen(&resultsScreen{model: t.model})
			return t.model, nil
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		t.model.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	t.model.list, cmd = t.model.list.Update(msg)
	return t.model, cmd
}

func (t listScreen) View() string {
	return docStyle.Render(t.model.list.View())
}
