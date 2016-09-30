package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	portPtr := flag.Int("p", 8080, "port to run server on")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Serve static files on the web:\nweb [flags] [directory]")
		flag.PrintDefaults()
	}

	flag.Parse()

	port := fmt.Sprintf(":%d", *portPtr)

	dir := flag.Arg(0)
	if dir == "" {
		dir, _ = os.Getwd()
	}

	dir = strings.Replace(dir, " ", "\\ ", -1)

	fmt.Printf("Serving %s on localhost%s\n", dir, port)
	panic(http.ListenAndServe(port, http.FileServer(http.Dir(dir))))
}
