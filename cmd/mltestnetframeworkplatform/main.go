// cmd/mltestnetframeworkplatform/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"mltestnetframeworkplatform/internal/mltestnetframeworkplatform"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := mltestnetframeworkplatform.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
