package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pborman/getopt"
	//rsapi "github.com/pebble-dev/rebblestore-api"
)

func main() {
	var version bool
	getopt.BoolVarLong(&version, "version", 'V', "Get the current version info")
	getopt.Parse()
	if version {
		//fmt.Fprintf(os.Stderr, "Version %s\nBuild Host: %s\nBuild Date: %s\nBuild Hash: %s\n", rsapi.Buildversionstring, rsapi.Buildhost, rsapi.Buildstamp, rsapi.Buildgithash)
		fmt.Fprintf(os.Stderr, "Version %s\nBuild Host: %s\nBuild Date: %s\nBuild Hash: %s\n", Buildversionstring, Buildhost, Buildstamp, Buildgithash)
		return
	}
	r := Handlers()
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.Handle("/", r)
	http.ListenAndServe(":8080", loggedRouter)
}
