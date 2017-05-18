// this is the documentation found using godoc
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

var re = regexp.MustCompile("^(.+)@golang.org$")

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1)

	if match != nil {
		fmt.Fprintf(w, "Hello, gopher %s\n", match[0][1])
		return
	}
	fmt.Fprintf(w, "Hello %s on the web!\n", path)
}
