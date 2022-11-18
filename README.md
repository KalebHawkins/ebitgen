# Ebitgen

An Ebiten project generator.

## Installation

```bash
go install github.com/KalebHawkins/ebitgen@latest
```

## Usage

Start a new project like you would any Go project. 

```bash
mkdir goapp
go mod init example.com/goapp
```

Now generate your `game.go` file and you are ready to *Go*. Yea, I said it. Pun intended. 😐🖐️🎤.

```bash
ebitgen 
```

You can customize some parameters, for example, say you want a smaller window and a different file name and title window.

```bash
ebiten -w 340 -H 240 -t "Awesome Game Name" -o main.go
```

Notice the `-H` is capital because it conflicts with the `--help, -h`. I couldn't think of another letter to use that made sense.