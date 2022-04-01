package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)


type Authentication struct {
    Enabled   bool `json:"enabled"`
    Password   string `json:"sha256_password_hash"`
}

type App struct {
    Name  string `json:"name"`
    Description   string `json:"description"`
    Url string  `json:"url"`
    CheckStatus bool `string:"checkStatus"`
    ResolvedStatus bool `string:"resolvedStatus"`
    AcceptableStatusCodes []int `string:"acceptableStatusCodes"`
}

type Authenticate struct {
    Password string
}


type Config struct {
    Authentication   Authentication `json:"authentication"`;
    Apps   []App `json:"apps"`
}


func load_config() Config {
    // Open our jsonFile
    jsonFile, err := os.Open("config.json")
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

func check_password(hash string) bool {
    config := load_config()

    correct := strings.EqualFold(strings.TrimRight(hash, "\n"),config.Authentication.Password)
    return correct
}

func intArrayContains(arr []int,element int) bool {
    result := false
    for _, arrElement := range arr {
        if (element == arrElement) {
            result = true
        }
    }
    return result
}

func main() {
    app := fiber.New()

    app.Static("/", "./public")

    app.Get("/verify", func (c *fiber.Ctx) error {
        return c.SendString("This is a valid Hexagon server")
    })

    app.Get("/check-pass/:password", func (c *fiber.Ctx) error {
        config := load_config()
        if (check_password(c.Params("password")) == false && config.Authentication.Enabled == true) {
            return c.SendString("false")
        } else {
            return c.SendString("true")
        }
    })

    app.Get("/get-apps/:password?", func (c *fiber.Ctx) error {
        config := load_config()
        // Check Authentication
        if (check_password(c.Params("password")) == false && config.Authentication.Enabled == true) {
            return c.SendString("Invalid password")
        }
        // Inject status
        for i := 0; i < len(config.Apps); i++ {
            if config.Apps[i].CheckStatus == true {
                client := http.Client{
                    Timeout: 5 * time.Second,
                }
                resp, err := client.Get(config.Apps[i].Url)
                if err != nil {
                    if (strings.Contains(err.Error(),"no such host") || strings.Contains(err.Error(),"context deadline exceeded")) {
                        config.Apps[i].ResolvedStatus = false
                    } else {
                        log.Fatal(err)
                    }
                } else {
                    if intArrayContains(config.Apps[i].AcceptableStatusCodes,resp.StatusCode) || resp.StatusCode == 200 {
                        config.Apps[i].ResolvedStatus = true
                    } else {
                        config.Apps[i].ResolvedStatus = false
                    }
                }
            }
        }
        // Return resolved config
        j, _ := json.Marshal(config.Apps)
        return c.SendString(string(j))
    })
    
    app.Get("/isAuthenticationRequired", func (c *fiber.Ctx) error {
        config := load_config()
        return c.SendString(strconv.FormatBool(config.Authentication.Enabled))
    })

    log.Fatal(app.Listen(":3000"))
}