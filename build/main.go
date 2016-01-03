package main

import (
	client "../client"
	"os"
)

func main() {
	client.ParseCommands(os.Args).Run()
}
