package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type searchScreen struct {
	model *model
}

func (q *searchScreen) Reset() {
}

func (q *searchScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:

			q.model.selectedQuery = q.model.queryTextInput.Value()
			res := SearchSong(q.model.selectedQuery)
			q.model.response = res

			var items []list.Item
			for _, song := range res.Results {
				items = append(items, item{
					title: song.Name,
					desc:  fmt.Sprintf("%s - %s", song.Artist, song.Album.Name),
				})
			}

			lm := list.New(items, list.NewDefaultDelegate(), 0, 0)
			q.model.list = lm
			q.model.list.Title = "Search Results"

			setCurrentScreen(&listScreen{model: q.model})
			return q.model, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return q.model, tea.Quit
		}

	case errMsg:
		q.model.err = msg
		return q.model, nil
	}

	var cmd tea.Cmd
	q.model.queryTextInput, cmd = q.model.queryTextInput.Update(msg)
	return q.model, cmd
}

func (q searchScreen) View() string {
	return fmt.Sprintf("Type your query\n\n%s\n\n(esc to quit)\n", q.model.queryTextInput.View())
}
