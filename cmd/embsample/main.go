package main

import (
	"embed"

	"github.com/rwxrob/emb"
)

//go:embed files/*
var files embed.FS

func init() {
	emb.FS = files
	emb.Top = "files"
}

func main() { emb.Cmd.Run() }
