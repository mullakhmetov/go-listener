package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Echo: "+r.URL.String())
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: listener [port ...]")
	}
	flag.Parse()
	wg := &sync.WaitGroup{}
	ports := flag.Args()
	if len(ports) == 0 {
		fmt.Println("No ports passed. See 'listener --help'.")
	}
	for _, p := range flag.Args() {
		wg.Add(1)
		port := fmt.Sprintf(":%s", p)
		handler := &handler{}
		go func() {
			defer wg.Done()
			fmt.Println("Listening", port)
			log.Fatal(http.ListenAndServe(port, handler))
		}()
	}
	wg.Wait()
}
