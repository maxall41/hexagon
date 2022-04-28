# Hexagon
![Logo](https://i.ibb.co/MMVFhfq/logo.png)
---
![Badge](https://img.shields.io/github/release-date/maxall41/hexagon)
![BadgeStars](https://img.shields.io/github/stars/maxall41/hexagon?style=social)
## What is Hexagon?
Hexagon is a terminal based Homelab dashboard so you can monitor all the apps you run locally and remember them 😉.
## Features
- 🔥 Blazing fast
- ✅ Status indicators
- 🔐 Secure authentication
- 🎉 Terminal UI
## Why use a terminal?
Using a terminal application instead of a browser makes viewing your dashboard a more streamlined experience. Because if you are like me you prefer using your terminal over your browser.
## Usage
After you setup the server just put the server info into the client (You will be prompted) when it first starts. Once you get into the main screen of the client you should be able to see all of your apps. Press enter to goto it if you configured a URL. Emojis indicate status checking if you have it enabled.
## Configuring
The server loads a config when it starts which stores all of your apps and authentication details. It is pretty simple here is an example config:
```
{
  "authentication": {
    "enabled": false,
    "sha256_password_hash": "AEC8084845B41A6952D46CBAA1C9B798659487FFD133796D95D05BA45D9096C2"
  },
  "apps": [
    {
      "name": "App 1",
      "description": "This is an app (1)",
      "url": "https://example.com/",
      "checkStatus": true
    },
    {
      "name": "App 2",
      "description": "This is an app (2)",
      "url": "https://example.com/",
      "checkStatus": true
    },
    {
      "name": "App 3",
      "description": "This is an app (3)",
      "url": "https://fdsfsdfsdfs.com/3",
      "checkStatus": true,
      "_comment": "This is optional defaults to 200 as accepted code:",
      "acceptableStatusCode": [200,201]
    }
  ]
}
```
## Installing the server (NATIVE)
To run the server natively just run the below command to compile the binary then just run it!
```
go build .
```
## Installing the server (Docker-Compose)
To run the server inside Docker compose just download the repo and then run the below command inside of the server/ directory to use docker-compose file
```
docker-compose up
```
Or use the built docker image https://hub.docker.com/r/maxall4/hexagon
## Installing the client
Just run the below command to download the binary and then add it to your path
```
wget https://github.com/maxall41/hexagon/releases/latest/download/client
```
Or run this command in the source directory to build it yourself
```
go build .
```
## Contributing
If you want to contribute to this project I'm happy to review your PRs!
## Roadmap:
- [x] Docker file
- [x] Docker-Compose config
- [x] Custom acceptable status codes returned from URL endpoint
- [ ] Homebrew formula
- [ ] Helm config
- [ ] Automated install for client
- [ ] Widgets
- [ ] In-client configuration
- [ ] Localization
- [ ] Custom status code resolvers
