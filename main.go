package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello, World!")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops, problem!", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops, problem!"))
			return
		}
		// log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye, World!")
	})
	http.ListenAndServe(":9090", nil)
}
