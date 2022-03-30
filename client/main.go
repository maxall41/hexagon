package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nathan-fiscaletti/consolesize-go"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc, url string
}

type App struct {
    Name  string `json:"name"`
    Description   string `json:"description"`
    Url string  `json:"url"`
	ResolvedStatus bool `string:"resolvedStatus"`
}

type Config struct {
    Url string `json:"url"`
    Hash   string `json:"hash"`
}


func (i item) Title() string       { return i.title }
func (i item) URL() string       { return i.url }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func (m model) Init() tea.Cmd {
	return nil
}

func load_config() Config {
    // Open our jsonFile
    jsonFile, err := os.Open(".hexagon")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    // read our opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var config Config

    // we unmarshal our byteArray which contains our
    json.Unmarshal(byteValue, &config)
    defer jsonFile.Close()
    return config
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if m.screen == URLInputScreen {
			return HandleURLInputScreen(m,msg)
		}
		if m.screen == PasswordInputScreen {
			return HandlePasswordInputScreen(m,msg)
		}
		if m.screen == MainScreen {
			return HandleMainScreen(m,msg)
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}
	cols, rows := consolesize.GetConsoleSize()
	m.width = cols
	m.height = rows
	var cmd tea.Cmd
	m.passwordInput, cmd = m.passwordInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if (m.screen == MainScreen) {
		return RenderMainScreen(m)
	} else if (m.screen == PasswordInputScreen) {
		return RenderPasswordInputScreen(m)
	} else if (m.screen == URLInputScreen) {
		return RenderURLInputPage(m)
	} else {
		style := lipgloss.NewStyle()
		return style.Render("Failed to find state " + string(m.screen))
	}
}

func main() {

	items := []list.Item{
	}

	itemURLs := []string{
	}

	urlInput := textinput.New()
	urlInput.Placeholder = "http://192.168.1.239:3000"
	urlInput.Focus()
	urlInput.CharLimit = 30
	urlInput.Width = 10

	passwordInput := textinput.New()
	passwordInput.Placeholder = "***********"
	passwordInput.CharLimit = 30
	passwordInput.Width = 10

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}

	if _, err := os.Stat(".hexagon"); err == nil {
		config := load_config()

		if err != nil {
			 log.Fatal(err)
		}
		url := string(config.Url) + "/get-apps/" + config.Hash
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// we initialize our Users array
		var apps []App

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(body, &apps)

		for _, app := range apps {
			var statusIcon = " 🔥" // 🔥 = Error resolving status
			if app.ResolvedStatus == true {
				statusIcon = " ✅"
			} else if app.ResolvedStatus == false {
				statusIcon = " ❌"
			}
			items = append(items,item{ title: app.Name + statusIcon, desc: app.Description })
			itemURLs = append(itemURLs, app.Url)
		}
		m.urls = itemURLs
		m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
		m.list.Title = "Apps"
		m.screen = MainScreen
	}

	m.list.Title = "Apps"
	m.urlInput = urlInput
	m.passwordInput = passwordInput
	m.setupPointer = 0;
	m.urls = itemURLs

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}