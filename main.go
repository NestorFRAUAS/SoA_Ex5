package main

		import (
			"log"
			"net/http"
	"time"
)

func main() {

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", home)

	srv.ListenAndServe()
	mainHTTP()
}

func home(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1><center> Hello from Go using HTTP/1.1 ! </center></h1>"))
	handler(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
}