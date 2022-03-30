package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func HandlePasswordInputScreen(m model,msg tea.KeyMsg) (tea.Model,tea.Cmd) {
	if msg.String() == "left" && m.setupPointer != 0 {
		m.setupPointer--;
	}
	if msg.String() == "right" && m.setupPointer != 1 {
		m.setupPointer++;
	}
	if msg.String() == "enter" && m.setupPointer == 0 {
		f, err := os.Create(os.Getenv("HOME") + "/.hexagon")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		sent_hash := sha256.Sum256([]byte(m.passwordInput.Value()))
		config := &Config{ Url: m.urlInput.Value(), Hash: strings.ToUpper(strings.TrimRight(fmt.Sprintf("%x\n", sent_hash),"\n")) }
		b, err := json.Marshal(config)
		if err != nil {
			fmt.Println(err)
		}

		_, err2 := f.WriteString(string(b))

		if err2 != nil {
			log.Fatal(err2)
		}
		items := []list.Item{
		}
	
		itemURLs := []string{
		}
	
		if _, err := os.Stat(os.Getenv("HOME") + "/.hexagon"); err == nil {
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
				items = append(items,item{ title: app.Name, desc: app.Description })
				itemURLs = append(itemURLs, app.Url)
			}
			m.urls = itemURLs
			m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
			m.list.Title = "Apps"
		}
		m.screen = MainScreen
		return m, nil
	}
	return m, nil
}

func RenderPasswordInputScreen(m model) string {
	okButton := buttonStyle.Render("Failed")
		cancelButton := buttonStyle.Render("Failed")
		if (m.setupPointer == 0) {
			okButton = activeButtonStyle.Render("Confirm")
			cancelButton = buttonStyle.Render("Cancel")
		} else {
			okButton = buttonStyle.Render("Confirm")
			cancelButton = activeButtonStyle.Render("Cancel")
		}

		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("This Hexagon server requires a password. Enter it below")
		errorText := lipgloss.NewStyle().Foreground(errorColor).Width(50).Align(lipgloss.Center).Render(m.setupModeError)
		buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
		ui := lipgloss.JoinVertical(lipgloss.Center, question, m.passwordInput.View(), errorText, buttons)

		dialog := lipgloss.Place(m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars("/"),
			lipgloss.WithWhitespaceForeground(subtle),
		)
		style := lipgloss.NewStyle()
		
		return style.Render(dialog)
}