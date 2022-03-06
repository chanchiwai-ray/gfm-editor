package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chanchiwai-ray/github-markdown/editor"
)

var file *editor.File
var path *string
var port *string
var err error

func init() {
	path = flag.String("f", "", "The path to the markdown file (Required).")
	port = flag.String("port", "8080", "Start the webserver at this port.")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
		fmt.Println(`A simple web-based editor that allows you to preview the
content of a markdown file just as what you will see in GitHub.`)
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *path == "" {
		log.Fatal("-f is required.")
	}

	file, err = editor.ReadMD(*path)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	editor.Serve(file, *path, *port)
}
