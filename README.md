# Garden Minder

A garden monitoring system. Automatically water, plan, and provide data for further actions needed for a living, green garden!

Built with **Go** using **CoAP** for system communication. **Gorm/sqlite** for storage and **client side web stack** for its user interface. Including an extended api for interfacing with other automation.

For the other half of this project see [Garden-Device](https://github.com/Volution-Labs/garden-device)

**A work in progress.**

## Installation & Running

```bash
# Install With:

go get github.com/volution-labs/garden-server

```

Before starting up the server, you should create an '.env' file inside the config folder for you specific environment. See [.envExample](config/exampledotenv)

```bash
# rename and use as .env | should not be commited

APP_ENV=dev

HTTP_SERVER_PORT=:8080
```

```bash
# Build and Run

cd garden-server

go build

./garden-server

```

## Documentation

System Overview: [Diagram](docs/diagram.jpg)
Full Documentation: Still in development

## CLI Tool

A tool to aid in device setup and debugging is provided in /tools/garden-cli. See the [garden-cli README](tools/garden-cli/README.md) for more info.

## To do

- [ ] Logging system
- [x] Routes for CoAP
- [ ] Routes for Http
- [ ] Add auth to CoAP and HTTP
- [ ] Command system for water timer
- [ ] Front end
- [x] CLI for device setup/control/debug
- [ ] Failsafe systems
- [ ] Scheduling System
