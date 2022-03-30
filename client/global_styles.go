package main

import "github.com/charmbracelet/lipgloss"

var dialogBoxStyle = lipgloss.NewStyle().
Border(lipgloss.RoundedBorder()).
BorderForeground(lipgloss.Color("#874BFD")).
Padding(1, 0).
BorderTop(true).
BorderLeft(true).
BorderRight(true).
BorderBottom(true)

var buttonStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#888B7E")).
		Padding(0, 3).
		MarginRight(2).
		MarginTop(1)

var activeButtonStyle = buttonStyle.Copy().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#F25D94")).
			MarginRight(2).
			Underline(true)
var subtle  = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
var errorColor = lipgloss.Color("#FF5252")