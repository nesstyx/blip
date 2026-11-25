package main

import (
	"fmt"
	"net/http"
)

var btn_count int

func click(w http.ResponseWriter, r *http.Request) {
	btn_count += 1
	fmt.Fprint(w, btn_count)
}

func restore(w http.ResponseWriter, r *http.Request) {
	btn_count = 0
	fmt.Fprint(w, btn_count)
}

func addclicks(w http.ResponseWriter, r *http.Request) {
	btn_count = btn_count + 100 // or btn_count += 100
	fmt.Fprint(w, btn_count)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/click", click)
	http.HandleFunc("/restore", restore)
	http.HandleFunc("/addclicks", addclicks)

	http.ListenAndServe(":8080", nil)
}
