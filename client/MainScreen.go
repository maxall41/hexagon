package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pkg/browser"
)

func HandleMainScreen(m model,msg tea.KeyMsg) (tea.Model,tea.Cmd) {
	if msg.String() == "enter" {
		browser.OpenURL(m.urls[m.list.Cursor()])
	}
	if msg.String() == "down" {
		m.listPointer--;
	}
	if msg.String() == "up" {
		m.listPointer++;
	}
	fmt.Println(m.listPointer)
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func RenderMainScreen(m model) string {
	return docStyle.Render(m.list.View())
}