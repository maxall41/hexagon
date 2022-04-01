package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	list list.Model
	listPointer int
	urls []string
	setupModeError string
	screen int
	width int
	height int
	setupPointer int
	urlInput textinput.Model
	passwordInput textinput.Model
}