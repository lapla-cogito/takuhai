/*
Copyright © 2023 lapla <lapla.cogito@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"log"
	"os"
	"strings"
)

const (
	compmode     = iota
	tracknummode = iota
)

const listHeight = 14

var selectedCompany, selectedtn string

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string
type itemDelegate struct{}

func (i item) FilterValue() string                             { return "" }
func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list      list.Model
	choice    string
	quitting  bool
	textInput textinput.Model
	mode      int
}

type (
	errMsg error
)

func (m model) Init() tea.Cmd {
	if m.mode == compmode {
		return nil
	}
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.mode == compmode {
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.list.SetWidth(msg.Width)
			return m, nil

		case tea.KeyMsg:
			switch keypress := msg.String(); keypress {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.choice = string(i)
				}
				return m, tea.Quit
			}
		}
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.choice = m.textInput.Value()
			return m, tea.Quit
		}
	case errMsg:
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.mode == compmode {
		selectedCompany = m.choice
		if m.choice != "" {
			return quitTextStyle.Render("Selected company: " + m.choice)
		}
		if m.quitting {
			return quitTextStyle.Render("No company selected. Quitting..")
		}
		return "\n" + m.list.View()
	}

	selectedtn = m.textInput.Value()
	return fmt.Sprintf(
		"Input tracking number\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func selectcompany() string {
	items := []list.Item{
		item("sagawa"),
		item("yamato"),
		item("jpost"),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select company delivers your package"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{
		list: l,
		mode: compmode,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return selectedCompany
}

func inputtn() string {
	ti := textinput.New()
	ti.Placeholder = "1234567890"
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 20

	m := model{
		textInput: ti,
		mode:      tracknummode,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, ti.Value())
	return selectedtn
}
