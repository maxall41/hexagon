package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pkg/browser"
)

func HandleMainScreen(m model,msg tea.KeyMsg) (tea.Model,tea.Cmd) {
	if msg.String() == "enter" {
		browser.OpenURL(m.urls[m.list.Cursor()])
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func RenderMainScreen(m model) string {
	return docStyle.Render(m.list.View())
}