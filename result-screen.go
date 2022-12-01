package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	headerHeight = 3
	footerHeight = 3
)

type resultsScreen struct {
	model *model
}

func (r *resultsScreen) Reset() {
	r.model.viewport.SetContent(r.model.resultContent)
}

func (r *resultsScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape:
			r.model.selectedTopic = nil
			setCurrentScreen(&searchScreen{model: r.model})
			return r.model, nil
		case tea.KeyRunes:
			if string(msg.Runes) == "q" {
				return r.model, tea.Quit
			}
		case tea.KeyCtrlC:
			return r.model, tea.Quit
		}

	case tea.WindowSizeMsg:
		verticalMargins := headerHeight + footerHeight
		r.model.viewport.Width = msg.Width
		r.model.viewport.Height = msg.Height - verticalMargins
	}

	var cmd tea.Cmd
	r.model.viewport, cmd = r.model.viewport.Update(msg)
	return r.model, cmd
}

func (r resultsScreen) View() string {
	if !r.model.ready {
		return "\n  Initializing..."
	}

	return "not handled yet"
}
