# Hexagon
![Logo](https://i.ibb.co/MMVFhfq/logo.png)
## What is Hexagon?
Hexagon is a terminal based Homelab dashboard so you can monitor all the apps you run locally and remember them 😉.
## Features
- 🔥 Blazing fast
- ✅ Status indicators
- 🔐 Secure authentication
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
      "checkStatus": true
    }
  ]
}
```
## Installing the server
To run the server natively just run the below command to compile the binary then just run it!
```
go build .
```
Docker container coming soon!
## Installing the client
Just run the below command to download the binary and then add it to your path
```
wget https://github.com/maxall41/hexagon/releases/download/PUBLIC/client
```
Or run this command in the source directory to build it yourself
```
go build .
```
## Contributing
If you want to contribute to this project feel free to!
## Roadmap:
- [ ] Docker file
- [ ] Helm config
- [ ] Docker-Compose config
- [ ] Automated Install of both server & client
- [ ] Icons
- [ ] Keycloak authentication
- [ ] Alternate views
- [ ] In-client configuration
- [ ] Localization
- [ ] Custom status code resolution with custom code
- [ ] Custom acceptable status codes returned from URL endpoint
