# Websocket chat written on golang

In this project I try to understand, how to create web apps using golang.

There are so many examples about creating simple chats, but unfortunately to little about creating complex web solutions with autentication and config management, etc.

So in this project I try to discribe as many different web problem solving, as I can.

To start server you need to copy `example.json` to `config.json` and change values, configure database connection
```bash
cp example.json config.json
```

Also you need to install all dependencies using [Govendor](https://github.com/kardianos/govendor) tool

```bash
go get -u github.com/kardianos/govendor
govendor sync
```

Than you need to build binary file
```bash
go build -o output_binary .
```

And run binary file

```bash
./output_binary
```
