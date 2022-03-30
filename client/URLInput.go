package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func HandleURLInputScreen(m model,msg tea.KeyMsg) (tea.Model,tea.Cmd) {
	if msg.String() == "enter" && m.setupPointer == 0 {
		_, err := url.ParseRequestURI(m.urlInput.Value())
		if err != nil {
			m.setupModeError = "Invalid URL";
		} else {
			resp, err := http.Get(m.urlInput.Value() + "/verify")
			if err != nil {
				log.Fatalln(err)
			}
			//We Read the response body on the line below.
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			if (string(body) == "This is a valid Hexagon server") {
				resp, err := http.Get(m.urlInput.Value() + "/isAuthenticationRequired")
				if err != nil {
					log.Fatalln(err)
				}
				//We Read the response body on the line below.
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				if (string(body) == "true") {
					m.urlInput.Blur()
					m.passwordInput.Focus()
					m.screen = PasswordInputScreen;
				} else {
					f, err := os.Create(os.Getenv("HOME") + "/.hexagon")
					if err != nil {
						log.Fatal(err)
					}

					defer f.Close()

					config := &Config{ Url: m.urlInput.Value(), Hash: "" }
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
				}

			} else {
				m.setupModeError = "Failed to validate that URL coresponds to valid Hexagon server"
			}
		}
		return m, nil
	}
	if msg.String() == "left" && m.setupPointer != 0 {
		m.setupPointer--;
	}
	if msg.String() == "right" && m.setupPointer != 1 {
		m.setupPointer++;
	}
	var cmd tea.Cmd
	m.urlInput, cmd = m.urlInput.Update(msg)
	return m, cmd
}

func RenderURLInputPage(m model) string {
	okButton := buttonStyle.Render("Failed")
	cancelButton := buttonStyle.Render("Failed")
	if (m.setupPointer == 0) {
		okButton = activeButtonStyle.Render("Confirm")
		cancelButton = buttonStyle.Render("Cancel")
	} else {
		okButton = buttonStyle.Render("Confirm")
		cancelButton = activeButtonStyle.Render("Cancel")
	}

	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render("Enter the URL for the Hexagon server:")
	errorText := lipgloss.NewStyle().Foreground(errorColor).Width(50).Align(lipgloss.Center).Render(m.setupModeError)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, m.urlInput.View(), errorText, buttons)

	dialog := lipgloss.Place(m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	style := lipgloss.NewStyle()
	
	return style.Render(dialog)
}