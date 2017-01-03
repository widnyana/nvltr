# nvltr

**N** - NICE
**V** - VIVACIOUS
**L** - LUMINOUS
**T** - TELEGRAM
**R** - ROBOT

*--okay, that's sounds so cheesy*

a telegram bot boilerplate using golang.

# Adding Command

I've provide some basic command, please take a look at `command` directory.

# Deploying

- open `config.yml` and put your config there, such as telegram bot token, google maps api key, etc
- extend yout bot
- place your TLS cert and key at `./cert/your.domain.name/{fullchain,privkey}.pem`
- setup reverse proxy on nginx to bot's bind address, default: '127.0.0.1:8088`
- execute `go generate`
- run it like usual `go run main.go` or `go build -o nvltr main.go && ./nvltr`
- ...
- profit!

# License

[MIT License](https://widnyana.mit-license.org/2016)
