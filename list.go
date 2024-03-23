package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		choices: getWatchList(),

		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "s":
			result := getRemainChoices(m.choices, m.selected)
			res := updateWatchList(result)

			if res {
				fmt.Println("Watchlist has been updated!")
			} else {
				fmt.Println("Failed to update watchlist!")
			}
			return m, tea.Quit

		// exit the program.
		case "ctrl+c", "q":

			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	s := "Watchlist\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := "x" // not selected
		if _, ok := m.selected[i]; ok {
			checked = " " // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress s to save.\nPress q to quit.\n"

	return s
}

func getRemainChoices(choices []string, selected map[int]struct{}) string {
	var remain []string

	for i, choice := range choices {
		if _, ok := selected[i]; !ok {
			remain = append(remain, choice)
		}
	}

	result := strings.Join(remain, ",")
	return result
}
